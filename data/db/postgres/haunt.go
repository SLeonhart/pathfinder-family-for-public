package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetHaunts(ctx context.Context) ([]model.Haunt, error) {
	/*	span := jaeger.GetSpan(ctx, "GetHaunts")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select h."name", h.eng_name, h.alias, h.cr, c.val_str as cr_str, c."exp" as "exp", h.alignment_area, h.caster_level, h.hp, h."notice", h.weakness, h."trigger", h."reset", h.destruction,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from haunt h
		join book b on b.id = h.book_id
		join cr c on c.val = h.cr`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetHaunts")

	res := make([]model.Haunt, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetHaunts error", err)
		return nil, err
	}

	return res, nil
}
