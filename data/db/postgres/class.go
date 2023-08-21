package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetClasses(ctx context.Context) ([]model.Class, error) {
	/*	span := jaeger.GetSpan(ctx, "GetClasses")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select c.alias, c."name", c.description, jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation) as book_json
	from class c 
		join book b on b.id = c.book_id 
	where not coalesce(c.is_archetype, false) and not coalesce(c.is_npc, false)
	order by c."name"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetClasses")

	classes := make([]model.Class, 0)
	err := db.GetClient().SelectContext(ctx, &classes, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetClasses error", err)
		return nil, err
	}

	return classes, nil
}

func (db *Postgres) GetClassInfo(ctx context.Context, classAlias string) (*model.ClassInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetClassInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select c."name", c.full_description, c."role", c.alignment, c.hit_die, c.starting_wealth, c.skill_ranks_per_level, c.table_features, c.table_spell_count, c.features, c.info_links,
	exists(select arc.id from "class" arc join class_class cc on cc.child_class_id = arc.id where cc.parent_class_id = c.id) as is_have_archetypes,
	(select jsonb_agg(jsonb_build_object('alias', s.alias, 'name', s."name", 'ability', jsonb_build_object('alias', a.alias, 'name', a.name, 'shortName', a.short_name, 'isSkillArmorPenalty', a.is_skill_armor_penalty)) order by s."name") 
		from skill s
			join ability a on a.id = s.ability_id
			join class_skill cs on cs.skill_id = s.id
		where cs.class_id = c.id) as skills,
	(select jsonb_agg(jsonb_build_object('alias', pc.alias, 'name', pc."name") order by pc."name") 
		from "class" pc
			join class_class cc on cc.parent_class_id = pc.id
		where not coalesce(pc.is_archetype, false) and not coalesce(pc.is_npc, false) and cc.child_class_id = c.id) as parent_classes,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation) as book_json,
	(select jsonb_agg(jsonb_build_object('alias', h.alias, 'name', h."name", 'isMain', hc.is_main_helper) order by h."name") from helper_class hc join helper h on h.id = hc.helper_id where hc.class_id = c.id) as helpers_json
	from "class" c
		join book b on b.id = c.book_id 
	where not coalesce(c.is_archetype, false) and not coalesce(c.is_npc, false) and c.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetClassInfo")

	var classInfo model.ClassInfo
	err := db.GetClient().GetContext(ctx, &classInfo, query, classAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetClassInfo error", err)
		}
		return nil, err
	}

	return &classInfo, nil
}

func (db *Postgres) GetBotClasses(ctx context.Context, id *int, alias *string, magicClass *bool) ([]model.BotClassInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetBotClasses")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	i := 0
	where := ``
	args := []interface{}{}
	if id != nil {
		i++
		args = append(args, id)
		where += fmt.Sprintf(` and c.id = $%v `, i)
	}
	if alias != nil {
		i++
		args = append(args, alias)
		where += fmt.Sprintf(` and c.alias = $%v `, i)
	}
	if magicClass != nil {
		if *magicClass {
			where += ` and cs.spell_levels is not null `
		} else {
			where += ` and cs.spell_levels is null `
		}
	}

	query := `with cs as (
		select c.id, array_agg(distinct cs."level" order by cs."level") as spell_levels
			from "class" c
			left join class_class cc on cc.child_class_id = c.id
			left join "class" pc on pc.id = cc.parent_class_id
			left join class_spell cs on cs.class_id = c.id or (not coalesce(c.is_own_spell_list, false) and cs.class_id = pc.id and cs."level" <= coalesce(c.max_spell_lvl, 99))
			where not coalesce(c.is_archetype,false) and cs.id is not null
			group by c.id
	)
	select c.id, c.alias, c."name", c.description, c.table_features, c.table_spell_count,
		cs.spell_levels as spell_levels
		from "class" c 
			left join cs on cs.id = c.id
		where not coalesce(c.is_archetype,false) and not coalesce(c.is_npc,false) ` + where

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetClasses")

	res := make([]model.BotClassInfo, 0)

	err := db.GetClient().SelectContext(ctx, &res, query, args...)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetClasses error", err)
		return nil, err
	}

	return res, nil
}
