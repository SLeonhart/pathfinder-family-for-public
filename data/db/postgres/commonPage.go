package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"

	"database/sql"
	"errors"
)

func (db *Postgres) GetCommonPage(ctx context.Context, id int) (*model.StaticPage, error) {
	/*	span := jaeger.GetSpan(ctx, "GetCommonPage")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := "select p.id, p.name, p.slug, p.content from common_pages p where id = $1"

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "id": id}, "GetCommonPage")

	res := &model.StaticPage{}
	err := db.GetClient().GetContext(ctx, res, query, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "id": id}, "GetCommonPage error", err)
			return nil, err
		}
		return nil, nil
	}

	return res, nil
}

func (db *Postgres) AddCommonPage(ctx context.Context, page model.StaticPage) (*int, error) {
	/*	span := jaeger.GetSpan(ctx, "AddCommonPage")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := "insert into common_pages (name, slug, content) values ($1, $2, $3::jsonb) returning id"

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "page": page}, "AddCommonPage")

	var id int
	err := db.GetClient().GetContext(ctx, &id, query, page.Name, page.Slug, page.Content)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "page": page}, "AddCommonPage error", err)
		return nil, err
	}

	return &id, nil
}

func (db *Postgres) UpdateCommonPage(ctx context.Context, page model.StaticPage) error {
	/*	span := jaeger.GetSpan(ctx, "UpdateCommonPage")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := "update common_pages set name = $1, slug = $2, content = $3::jsonb where id = $4 returning id"

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "page": page}, "UpdateCommonPage")

	var id int
	err := db.GetClient().GetContext(ctx, &id, query, page.Name, page.Slug, page.Content, page.Id)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "page": page}, "UpdateCommonPage error", err)
		return err
	}

	return nil
}
