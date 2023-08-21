package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetTraits(ctx context.Context) ([]model.Trait, error) {
	/*	span := jaeger.GetSpan(ctx, "GetTraits")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select t.alias, t."name", t.eng_name, t.benefit, t.prerequisites,
	jsonb_build_object('alias', tt.alias, 'name', tt."name", 'description', tt.description, 'parentType', (case when ptt.id is null then null else jsonb_build_object('alias', ptt.alias, 'name', ptt."name", 'description', ptt.description) end)) as trait_type,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from trait t
		join book b on b.id = t.book_id
		join trait_type tt on tt.id = t.type_id 
		left join trait_type ptt on ptt.id = tt.parent_type_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetTraits")

	res := make([]model.Trait, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetTraits error", err)
		return nil, err
	}

	return res, nil
}
