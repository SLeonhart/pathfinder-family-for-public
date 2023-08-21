package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetAspects(ctx context.Context) ([]model.Aspect, error) {
	/*	span := jaeger.GetSpan(ctx, "GetAspects")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select a.alias, a.name, a.eng_name, a.description, a.minor_form, a.major_form, a.additional_description,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from aspect a
		join book b on b.id = a.book_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetAspects")

	res := make([]model.Aspect, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetAspects error", err)
		return nil, err
	}

	return res, nil
}
