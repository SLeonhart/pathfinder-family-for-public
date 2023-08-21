package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetNpcs(ctx context.Context) ([]model.Class, error) {
	/*	span := jaeger.GetSpan(ctx, "GetNpcs")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select c.alias, c."name", c.description, jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation) as book_json
	from class c 
		join book b on b.id = c.book_id 
	where not coalesce(c.is_archetype, false) and coalesce(c.is_npc, false)
	order by c."name"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetNpcs")

	npcs := make([]model.Class, 0)
	err := db.GetClient().SelectContext(ctx, &npcs, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetNpcs error", err)
		return nil, err
	}

	return npcs, nil
}

func (db *Postgres) GetNpcInfo(ctx context.Context, classAlias string) (*model.NpcInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetNpcInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select c."name", c.alignment, c.hit_die, c.skill_ranks_per_level, c.table_features, c.table_spell_count, c.features,
	(select jsonb_agg(jsonb_build_object('alias', s.alias, 'name', s."name", 'ability', jsonb_build_object('alias', a.alias, 'name', a.name, 'shortName', a.short_name, 'isSkillArmorPenalty', a.is_skill_armor_penalty)) order by s."name") 
		from skill s
			join ability a on a.id = s.ability_id
			join class_skill cs on cs.skill_id = s.id
		where cs.class_id = c.id) as skills,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation) as book_json,
	(select jsonb_agg(jsonb_build_object('alias', h.alias, 'name', h."name", 'isMain', hc.is_main_helper) order by h."name") from helper_class hc join helper h on h.id = hc.helper_id where hc.class_id = c.id) as helpers_json
	from "class" c
		join book b on b.id = c.book_id 
	where not coalesce(c.is_archetype, false) and coalesce(c.is_npc, false) and c.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetNpcInfo")

	var npcInfo model.NpcInfo
	err := db.GetClient().GetContext(ctx, &npcInfo, query, classAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetNpcInfo error", err)
		}
		return nil, err
	}

	return &npcInfo, nil
}
