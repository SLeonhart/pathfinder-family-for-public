package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetWildTalents(ctx context.Context) ([]model.WildTalent, error) {
	/*	span := jaeger.GetSpan(ctx, "GetWildTalents")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select wt."name", wt.eng_name, wt.alias, wt.description, wt."element", wt.type_description, wt."level", wt.burn, wt.burn_description, wt.prerequisites, wt.blast_type, wt.damage, wt.saving_throw, wt.spell_resistance, wt.associated_blasts,
	jsonb_build_object('alias', wtt.alias, 'name', wtt."name", 'shortName', wtt.short_name, 'description', wtt.description, 'order', wtt."order") as "type",
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from wild_talent wt
		join wild_talent_type wtt on wtt.id = wt.type_id
		join book b on b.id = wt.book_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetWildTalents")

	res := make([]model.WildTalent, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetWildTalents error", err)
		return nil, err
	}

	return res, nil
}
