package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetArchetypes(ctx context.Context, classAlias *string) ([]model.ClassWithArchetypes, error) {
	/*	span := jaeger.GetSpan(ctx, "GetArchetypes")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	args := make([]interface{}, 0)

	query := `select pc."name", pc.alias, pc.archetypes_description, jsonb_agg(jsonb_build_object('name', c."name", 'alias', c.alias, 'description', c.description, 'book', jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation)) order by c."name") as archetypes
	from "class" c 
		join book b on b.id = c.book_id 
		join class_class cc on cc.child_class_id = c.id 
		join "class" pc on pc.id = cc.parent_class_id 
	where coalesce(c.is_archetype, false) and not coalesce(c.is_npc, false) `

	if classAlias != nil {
		query += ` and pc.alias = $1 `
		args = append(args, classAlias)
	}

	query += ` group by pc."name", pc.alias, pc.archetypes_description
	order by pc."name"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetArchetypes")

	classes := make([]model.ClassWithArchetypes, 0)
	err := db.GetClient().SelectContext(ctx, &classes, query, args...)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetArchetypes error", err)
		return nil, err
	}

	return classes, nil
}

func (db *Postgres) GetArchetypeInfo(ctx context.Context, archetypeAlias string) (*model.ArchetypeInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetArchetypeInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select c."name", c.alignment, c.hit_die, c.starting_wealth, c.skill_ranks_per_level, c.table_features, c.table_spell_count, c.features, c.archetype_features, c.description, c.info_links,
	(select jsonb_agg(jsonb_build_object('alias', s.alias, 'name', s."name", 'ability', jsonb_build_object('alias', a.alias, 'name', a.name, 'shortName', a.short_name, 'isSkillArmorPenalty', a.is_skill_armor_penalty)) order by s."name") 
		from skill s
			join ability a on a.id = s.ability_id
			join class_skill cs on cs.skill_id = s.id
		where cs.class_id = c.id) as skills,
	jsonb_build_object('alias', pc.alias, 'name', pc."name") as parent_class,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation) as book_json,
	(select jsonb_agg(jsonb_build_object('alias', h.alias, 'name', h."name", 'isMain', hc.is_main_helper) order by h."name") from helper_class hc join helper h on h.id = hc.helper_id where hc.class_id = c.id) as helpers_json
	from "class" c
		join book b on b.id = c.book_id 
		join class_class cc on cc.child_class_id = c.id 
		join "class" pc on pc.id = cc.parent_class_id 
	where coalesce(c.is_archetype, false) and not coalesce(c.is_npc, false) and c.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "archetypeAlias": archetypeAlias}, "GetArchetypeInfo")

	var archetypeInfo model.ArchetypeInfo
	err := db.GetClient().GetContext(ctx, &archetypeInfo, query, archetypeAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "archetypeAlias": archetypeAlias}, "GetArchetypeInfo error", err)
		}
		return nil, err
	}

	return &archetypeInfo, nil
}
