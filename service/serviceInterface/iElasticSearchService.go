package serviceInterface

import (
	"context"
	"pathfinder-family/model"
	"time"
)

//go:generate mockgen -source=iElasticSearchService.go -destination=../serviceMock/elasticSearchService.go -package=serviceMock

type IElasticSearchService interface {
	UpdatePathfinderSearch(ctx context.Context)
	Upsert(ctx context.Context) error
	Update(ctx context.Context, updateTime time.Time) error
	ClearOld(ctx context.Context, updateTime time.Time) error
	Get(ctx context.Context, searchString string) ([]model.ElasticResultHitsHits, error)
}
