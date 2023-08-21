package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetGoods(ctx context.Context, token *string) ([]model.Good, error) {
	/*	span := jaeger.GetSpan(ctx, "GetGoods")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `with u as (
		select su.shop_id
			from "user" u
				join shop_user su on su.user_id = u.id
			where u."token" = $1
	)
	select s.id, s.cnt, s.priority, s."name", s.url, s.image_urls, s.price, exists(select * from u where shop_id = s.id) as in_waiting_list
		from shop s
		where (s.dt_from is null or s.dt_from < now()) and (s.dt_to is null or s.dt_to > now())
		order by s.cnt = 0, s.priority desc, s.name`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "token": token}, "GetGoods")

	res := make([]model.Good, 0)
	err := db.GetClient().SelectContext(ctx, &res, query, token)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "token": token}, "GetGoods error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) AddGoodInWaitingList(ctx context.Context, userId int, goodId int) error {
	/*	span := jaeger.GetSpan(ctx, "AddGoodInWaitingList")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `insert into shop_user (user_id,shop_id) values ($1,$2)`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId, "goodId": goodId}, "AddGoodInWaitingList")

	_, err := db.GetClient().ExecContext(ctx, query, userId, goodId)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "userId": userId, "goodId": goodId}, "AddGoodInWaitingList error", err)
		return err
	}

	return nil
}
