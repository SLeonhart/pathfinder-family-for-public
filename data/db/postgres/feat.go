package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetFeats(ctx context.Context) ([]model.Feat, error) {
	/*	span := jaeger.GetSpan(ctx, "GetFeats")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select f.id, f."name", f.alias, f.requirements, f.description, f.parent_feat_id,
	(select jsonb_agg(jsonb_build_object('name', ft."name", 'alias', ft.alias) order by ft."name") 
		from feat_type ft
			join feat_feat_type fft on fft.feat_type_id = ft.id
		where fft.feat_id = f.id) as feat_types,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from feat f  
		join book b on b.id = f.book_id
	order by f."name"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetFeats")

	res := make([]model.Feat, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetFeats error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetFeatInfo(ctx context.Context, featAlias string) (*model.FeatInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetFeatInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select f."name", f.full_description, f.prerequisites, f.benefit, f.normal, f.special,
	(select jsonb_agg(jsonb_build_object('name', ft."name", 'alias', ft.alias) order by ft."name") 
		from feat_type ft
			join feat_feat_type fft on fft.feat_type_id = ft.id
		where fft.feat_id = f.id) as feat_types,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation) as book,
	(select jsonb_agg(jsonb_build_object('alias', h.alias, 'name', h."name", 'isMain', hf.is_main_helper) order by h."name") from helper_feat hf join helper h on h.id = hf.helper_id where hf.feat_id = f.id) as helpers
	from feat f  
		join book b on b.id = f.book_id
	where f.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "featAlias": featAlias}, "GetFeatInfo")

	var res model.FeatInfo
	err := db.GetClient().GetContext(ctx, &res, query, featAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "featAlias": featAlias}, "GetFeatInfo error", err)
		}
		return nil, err
	}

	return &res, nil
}
