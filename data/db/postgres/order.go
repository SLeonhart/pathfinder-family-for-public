package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetOrders(ctx context.Context) ([]model.Order, error) {
	/*	span := jaeger.GetSpan(ctx, "GetOrders")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select o.alias, o."name", o.description, c."name" as class_name
	from orders o 
		join "class" c on c.id = o.class_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetOrders")

	res := make([]model.Order, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetOrders error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetOrderInfo(ctx context.Context, orderAlias string) (*model.OrderInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetOrderInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select o."name", o.description, o.full_description
	from orders o 
	where o.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "orderAlias": orderAlias}, "GetOrderInfo")

	var res model.OrderInfo
	err := db.GetClient().GetContext(ctx, &res, query, orderAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "orderAlias": orderAlias}, "GetOrderInfo error", err)
		}
		return nil, err
	}

	return &res, nil
}
