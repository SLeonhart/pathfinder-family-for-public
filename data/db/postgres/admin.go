package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) AddDonate(ctx context.Context, request model.AddDonateRequest) error {
	/*	span := jaeger.GetSpan(ctx, "AddDonate")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `do $$
	begin
		insert into helper ("name") values ($1) on conflict("name") do nothing;
		insert into donate (dt, sum, helper_id)
			select $2, $3, id from helper where "name" = $1;
	end $$;`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "AddDonate")

	_, err := db.GetClient().ExecContext(ctx, query, request.Helper, request.Dt, request.Sum)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "request": request}, "AddDonate error", err)
		return err
	}

	return nil
}

func (db *Postgres) GetPushTokens(ctx context.Context) ([]model.PushToken, error) {
	/*	span := jaeger.GetSpan(ctx, "GetPushTokens")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select id, token from push_token`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetPushTokens")

	res := make([]model.PushToken, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetPushTokens error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) DeletePushToken(ctx context.Context, id int) error {
	/*	span := jaeger.GetSpan(ctx, "DeletePushToken")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `delete from push_token where id = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "id": id}, "DeletePushToken")

	_, err := db.GetClient().ExecContext(ctx, query, id)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "id": id}, "DeletePushToken error", err)
		return err
	}

	return nil
}
