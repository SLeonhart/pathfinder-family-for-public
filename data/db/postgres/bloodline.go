package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetBloodlines(ctx context.Context, classAlias string) ([]model.Bloodline, error) {
	/*	span := jaeger.GetSpan(ctx, "GetBloodlines")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select bl."name", bl.alias, bl.description,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from bloodline bl 
		join book b on b.id = bl.book_id
		join "class" c on c.id = bl.class_id 
	where c.alias = $1
	order by bl."name"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetBloodlines")

	res := make([]model.Bloodline, 0)
	err := db.GetClient().SelectContext(ctx, &res, query, classAlias)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetBloodlines error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetBloodlineInfo(ctx context.Context, classAlias string, bloodlineAlias string) (*model.BloodlineInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetBloodlineInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select bl."name", bl.description, bl.full_description,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from bloodline bl 
		join book b on b.id = bl.book_id
		join "class" c on c.id = bl.class_id 
	where c.alias = $1 and bl.alias = $2`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias, "bloodlineAlias": bloodlineAlias}, "GetBloodlineInfo")

	var res model.BloodlineInfo
	err := db.GetClient().GetContext(ctx, &res, query, classAlias, bloodlineAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias, "bloodlineAlias": bloodlineAlias}, "GetBloodlineInfo error", err)
		}
		return nil, err
	}

	return &res, nil
}

func (db *Postgres) GetBloodlineName(ctx context.Context, classAlias string, bloodlineAlias string) (*string, error) {
	/*	span := jaeger.GetSpan(ctx, "GetBloodlineName")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select case when bl.eng_name is null then bl."name" else bl."name"||' ('||bl.eng_name||')' end as "name"
	from bloodline bl 
		join "class" c on c.id = bl.class_id 
	where c.alias = $1 and bl.alias = $2`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "bloodlineAlias": bloodlineAlias, "classAlias": classAlias}, "GetBloodlineName")

	var name string
	err := db.GetClient().GetContext(ctx, &name, query, classAlias, bloodlineAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "bloodlineAlias": bloodlineAlias, "classAlias": classAlias}, "GetBloodlineName error", err)
		}
		return nil, err
	}

	return &name, nil
}
