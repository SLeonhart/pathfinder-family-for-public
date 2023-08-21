package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetTraps(ctx context.Context) ([]model.Trap, error) {
	/*	span := jaeger.GetSpan(ctx, "GetTraps")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select t."name", t.eng_name, t.alias, t.cr, t."type", t.dc_perception, t.dc_disable_device, t."trigger", t.duration, t."reset", t.effect,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from trap t
		join book b on b.id = t.book_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetTraps")

	res := make([]model.Trap, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetTraps error", err)
		return nil, err
	}

	return res, nil
}
