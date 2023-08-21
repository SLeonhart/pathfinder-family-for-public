package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetShamanSpirits(ctx context.Context) ([]model.ShamanSpiritInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetShamanSpirits")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select s.name, s.eng_name, s.alias, s.description, s.hexes, s.spirit_animal, s.spirit_ability, s.greater_spirit_ability, s.true_spirit_ability, s.manifestation,
	case when ps.alias is null then null else jsonb_build_object('alias', ps.alias, 'name', ps."name") end as parent,
	(select jsonb_agg(jsonb_build_object('level', ssp."level", 'comment', ssp."comment", 'alias', sp.alias, 'name', sp."name") order by ssp."level") 
		from spirit_spell ssp
			join spell sp on sp.id = ssp.spell_id
		where ssp.spirit_id = s.id) as spells,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from spirit s
		join book b on b.id = s.book_id
		left join spirit_spirit ss on ss.child_spirit_id = s.id
		left join spirit ps on ss.parent_spirit_id = ps.id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetShamanSpirits")

	res := make([]model.ShamanSpiritInfo, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetShamanSpirits error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetMediumSpirits(ctx context.Context, alias string) ([]model.MediumSpiritInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetMediumSpirits")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select ms.name, ms.eng_name, ms.alias, ms.description, ms."type", ms.spirit_bonus, ms.seance_boon, ms.favored_locations, ms.influence_penalty, ms.taboos, ms.spirit_power_base, ms.spirit_power_intermediate, ms.spirit_power_greater, ms.spirit_power_supreme,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from medium_spirit ms
		join book b on b.id = ms.book_id
		join "class" c on c.id = ms.class_id
	where c.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "alias": alias}, "GetMediumSpirits")

	res := make([]model.MediumSpiritInfo, 0)
	err := db.GetClient().SelectContext(ctx, &res, query, alias)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "alias": alias}, "GetMediumSpirits error", err)
		return nil, err
	}

	return res, nil
}
