package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetGods(ctx context.Context) ([]model.God, error) {
	/*	span := jaeger.GetSpan(ctx, "GetGods")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select g.alias, g."name", g.eng_name, g.aligment, g.title, g.portfolios, g.favored_weapon, g.symbol, g.sacred_animal, g.sacred_colors,
	jsonb_build_object('alias', gt.alias, 'name', gt."name", 'description', gt.description) as god_type,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book,
	(select jsonb_agg(jsonb_build_object('type', dt.alias, 'name', d."name", 'alias', d.alias)) 
		from god_domain gd
			join "domain" d on d.id = gd.domain_id
			join domain_type dt on dt.id = d.type_id
		where gd.god_id = g.id) as domains
	from god g 
		join book b on b.id = g.book_id
		join god_type gt on gt.id = g.type_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetGods")

	res := make([]model.God, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetGods error", err)
		return nil, err
	}

	return res, nil
}
