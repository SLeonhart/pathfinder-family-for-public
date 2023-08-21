package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetNews(ctx context.Context, offset int, limit int, onlyActual bool) ([]model.News, error) {
	/*	span := jaeger.GetSpan(ctx, "GetNews")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select id, push_title, push_body, push_image_link, push_target_link, dt, body
		from news `

	if onlyActual {
		query += ` where now() >= dt `
	}

	query += ` order by dt desc
	offset $1 limit $2`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetNews")

	news := make([]model.News, 0)
	err := db.GetClient().SelectContext(ctx, &news, query, offset, limit)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetNews error", err)
		return nil, err
	}

	return news, nil
}

func (db *Postgres) AddNews(ctx context.Context, request model.AddNewsRequest) error {
	/*	span := jaeger.GetSpan(ctx, "AddNews")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `insert into news (push_title, push_body, push_image_link, push_target_link, dt, body) values ($1,$2,$3,$4,now(),$5)`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "AddNews")

	_, err := db.GetClient().ExecContext(ctx, query, request.PushTitle, request.PushBody, request.PushImageLink, request.PushTargetLink, request.Body)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "AddNews error", err)
		return err
	}

	return nil
}
