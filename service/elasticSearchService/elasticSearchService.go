package elasticSearchService

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"pathfinder-family/config"
	"pathfinder-family/data/db/dbInterface"
	"pathfinder-family/infrastructure/logger"
	"pathfinder-family/model"

	goelastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticSearchService struct {
	cfg          *config.Config
	postgres     dbInterface.IPostgres
	client       *goelastic.Client
	indexName    string
	documentType string
}

func NewElasticSearchService(cfg *config.Config, postgres dbInterface.IPostgres) *ElasticSearchService {
	s := ElasticSearchService{
		cfg:          cfg,
		postgres:     postgres,
		indexName:    "pathfinder",
		documentType: "search",
	}

	s.createClient()

	go func() {
		timer := time.NewTicker(time.Duration(cfg.Elastic.ReconnectMsec) * time.Millisecond)
		defer timer.Stop()

		count := 5

		for range timer.C {
			if err := s.createClient(); err != nil {
				count--
				if count == 0 {
					logger.ErrorWithErr("", "Elastic reconnect error", err)
					count = 5
				}
			}
		}
	}()

	return &s
}

func (s *ElasticSearchService) createClient() error {
	client, err := goelastic.NewClient(goelastic.Config{
		Addresses: []string{
			s.cfg.Elastic.Url,
		},
	})
	if err != nil {
		logger.ErrorWithErr("", "Elastic client error", err)
		return err
	}
	res, err := client.Info()
	if err != nil {
		logger.ErrorWithErr("", "Elastic client info err", err)
		return err
	}
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		logger.ErrorWithErr("", "Elastic client info parsing err", err)
		return err
	}
	// logger.Infof("", "Client: %s", goelastic.Version)
	// logger.Infof("", "Server: %s", resMap["version"].(map[string]interface{})["number"])

	s.client = client

	ctx := context.Background()
	s.createIndex(ctx)

	return nil
}

func (s *ElasticSearchService) existsIndex(ctx context.Context) (bool, error) {
	req := esapi.IndicesGetRequest{
		Index: []string{s.indexName},
	}

	res, err := req.Do(ctx, s.client)
	if err != nil {
		logger.ErrorWithErr(logger.CreateRequestIDField(ctx), "Elastic request err", err)
		return false, err
	}
	defer res.Body.Close()

	if res.IsError() {
		if res.StatusCode == 404 {
			return false, nil
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		resStr := buf.String()
		logger.ErrorWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"status": res.Status(), "body": resStr}, "Elastic exists index document err")
		return false, errors.New("Elastic exists index document err")
	}
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"status": res.Status()}, "Elastic response parsing err", err)
		return false, err
	}
	return true, nil
}

func (s *ElasticSearchService) createIndex(ctx context.Context) error {
	exists, err := s.existsIndex(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	jsonQuery, _ := json.Marshal(map[string]interface{}{
		"settings": map[string]interface{}{
			"analysis": map[string]interface{}{
				"filter": map[string]interface{}{
					"russian_stop": map[string]interface{}{
						"type":      "stop",
						"stopwords": "_russian_",
					},
					"russian_stemmer": map[string]interface{}{
						"type":     "stemmer",
						"language": "russian",
					},
				},
				"analyzer": map[string]interface{}{
					"rebuilt_russian": map[string]interface{}{
						"tokenizer": "standard",
						"filter": []string{
							"lowercase",
							"russian_stop",
							"russian_stemmer",
						},
					},
				},
			},
		},
	})

	req := esapi.IndicesCreateRequest{
		Index: s.indexName,
		Body:  bytes.NewReader(jsonQuery),
	}

	res, err := req.Do(ctx, s.client)
	if err != nil {
		logger.ErrorWithErr(logger.CreateRequestIDField(ctx), "Elastic request err", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		resStr := buf.String()
		logger.ErrorWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"status": res.Status(), "body": resStr}, "Elastic create index document err")
		return errors.New("Elastic create index document err")
	}
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"status": res.Status()}, "Elastic response parsing err", err)
		return err
	}
	return nil
}

func (s *ElasticSearchService) UpdatePathfinderSearch(ctx context.Context) {
	for {
		err := s.Upsert(ctx)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	go func() {
		timer := time.NewTicker(15 * time.Minute)
		defer timer.Stop()

		for range timer.C {
			s.Upsert(context.Background())
		}
	}()
}

func (s *ElasticSearchService) Upsert(ctx context.Context) error {
	if s.client == nil {
		logger.Error(logger.CreateRequestIDField(ctx), "Elastic client is null")
		return errors.New("Elastic client is null")
	}

	updateTime := time.Now()

	err := s.ClearOld(ctx, updateTime)
	if err != nil {
		return err
	}

	err = s.Update(ctx, updateTime)
	if err != nil {
		return err
	}

	return nil
}

func (s *ElasticSearchService) Update(ctx context.Context, updateTime time.Time) error {
	if s.client == nil {
		logger.Error(logger.CreateRequestIDField(ctx), "Elastic client is null")
		return errors.New("Elastic client is null")
	}

	searchInfo, err := s.postgres.GetSearchInfo(ctx)
	if err != nil {
		return err
	}

	bulkArray := make([]string, 0)
	for _, item := range searchInfo {
		item.DtUpdate = updateTime
		jsonItem, _ := json.Marshal(item)
		bulkArray = append(bulkArray, fmt.Sprintf(`{"index":{"_id": "%v"}}`, item.Id))
		bulkArray = append(bulkArray, string(jsonItem))
	}
	bulkArray = append(bulkArray, "")
	req := esapi.BulkRequest{
		Index: s.indexName,
		//DocumentType: s.documentType,
		Body:    strings.NewReader(strings.Join(bulkArray, "\n")),
		Refresh: "true",
	}

	res, err := req.Do(ctx, s.client)
	if err != nil {
		logger.ErrorWithErr(logger.CreateRequestIDField(ctx), "Elastic request err", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		resStr := buf.String()
		logger.ErrorWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"status": res.Status(), "body": resStr}, "Elastic indexing document err")
		return errors.New("Elastic indexing document err")
	}
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"status": res.Status()}, "Elastic response parsing err", err)
		return err
	}
	return nil
}

func (s *ElasticSearchService) ClearOld(ctx context.Context, updateTime time.Time) error {
	if s.client == nil {
		logger.Error(logger.CreateRequestIDField(ctx), "Elastic client is null")
		return errors.New("Elastic client is null")
	}

	jsonQuery, _ := json.Marshal(map[string]interface{}{
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"dtUpdate": map[string]interface{}{
					"lt": updateTime,
				},
			},
		},
	})
	req := esapi.DeleteByQueryRequest{
		Index: []string{s.indexName},
		//DocumentType: []string{s.documentType},
		Body: bytes.NewReader(jsonQuery),
	}

	res, err := req.Do(ctx, s.client)
	if err != nil {
		logger.ErrorWithErr(logger.CreateRequestIDField(ctx), "Elastic request err", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		resStr := buf.String()
		logger.ErrorWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"status": res.Status(), "body": resStr}, "Elastic indexing document err")
		return errors.New("Elastic indexing document err")
	}
	var resMap map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"status": res.Status()}, "Elastic response parsing err", err)
		return err
	}
	return nil
}

func (s *ElasticSearchService) Get(ctx context.Context, searchString string) ([]model.ElasticResultHitsHits, error) {
	if s.client == nil {
		logger.Error(logger.CreateRequestIDField(ctx), "Elastic client is null")
		return nil, errors.New("Elastic client is null")
	}

	reg, err := regexp.Compile("[^A-zА-яЁё0-9]")
	if err != nil {
		logger.ErrorWithErr("", "Elastic regex compile err", err)
		return nil, err
	}
	newSearchString := strings.ToLower(strings.TrimSpace(reg.ReplaceAllString(searchString, " ")))
	if len(newSearchString) == 0 {
		return nil, nil
	}

	searchWords := make([]string, 0)
	regExps := make([]string, 0)
	for _, k := range strings.Split(newSearchString, " ") {
		k = strings.TrimSpace(k)
		if len(k) > 0 {
			lower := k
			query := ".*" + lower + ".*"
			if len([]rune(k)) < 3 {
				query = lower + ".*"
			}
			regExps = append(regExps, query)
			searchWords = append(searchWords, lower)
		}
	}

	shouldQuery := make([]map[string]interface{}, 0)
	for _, r := range regExps {
		shouldQuery = append(shouldQuery, map[string]interface{}{"regexp": map[string]interface{}{"h1": r}})
		shouldQuery = append(shouldQuery, map[string]interface{}{"regexp": map[string]interface{}{"content": r}})
	}

	jsonQuery, _ := json.Marshal(map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": shouldQuery,
			},
		},
		"size": 1000,
	})

	req := esapi.SearchRequest{
		Index: []string{s.indexName},
		//DocumentType: []string{s.documentType},
		Body: bytes.NewReader(jsonQuery),
	}

	res, err := req.Do(ctx, s.client)
	if err != nil {
		logger.ErrorWithErr(logger.CreateRequestIDField(ctx), "Elastic request err", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		resStr := buf.String()
		logger.ErrorWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"status": res.Status(), "body": resStr}, "Elastic search document err")
		return nil, errors.New("Elastic search document err")
	}

	var result model.ElasticResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"status": res.Status()}, "Elastic response parsing err", err)
		return nil, err
	}

	for i := range result.Hits.Hits {
		lowerH1 := strings.ToLower(result.Hits.Hits[i].Source.H1)
		if strings.Contains(" "+lowerH1+" ", " "+newSearchString+" ") {
			result.Hits.Hits[i].Score += float64(10000)
		}
		for _, r := range searchWords {
			reg, _ := regexp.Compile(r)
			res := reg.FindAllString(lowerH1, -1)
			if len(res) > 0 {
				result.Hits.Hits[i].Score += float64(1000 * len(res))
			}

			if result.Hits.Hits[i].Source.Content != nil {
				res = reg.FindAllString(strings.ToLower(*result.Hits.Hits[i].Source.Content), -1)
				if len(res) > 0 {
					result.Hits.Hits[i].Score += float64(len(res))
				}
			}
		}
	}

	// sort.Slice(result.Hits.Hits, func(i, j int) bool {
	// 	if result.Hits.Hits[i].Score != result.Hits.Hits[j].Score {
	// 		return result.Hits.Hits[i].Score > result.Hits.Hits[j].Score
	// 	}
	// 	return result.Hits.Hits[i].Source.H1 < result.Hits.Hits[j].Source.H1
	// })

	return result.Hits.Hits, nil
}
