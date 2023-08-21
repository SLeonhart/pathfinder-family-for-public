package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetAfflictions(ctx context.Context) ([]model.Affliction, error) {
	/*	span := jaeger.GetSpan(ctx, "GetAfflictions")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select a."name", a.eng_name, a.alias, a.type_description, a.save, a.onset, a.frequency, a.effect, a.initial_effect, a.secondary_effect, a.cure, a."cost", a.description,
	jsonb_build_object('alias', mat.alias, 'name', mat."name") as main_type,
	(select jsonb_agg(jsonb_build_object('alias', sat.alias, 'name', sat."name") order by sat."name")
		from affliction_type sat
			join affliction_affliction_type aat on aat.affliction_type_id = sat.id
		where aat.affliction_id = a.id) as secondary_types,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from affliction a
		join book b on b.id = a.book_id
		join affliction_type mat on mat.id = a.main_type_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetAfflictions")

	res := make([]model.Affliction, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetAfflictions error", err)
		return nil, err
	}

	return res, nil
}
