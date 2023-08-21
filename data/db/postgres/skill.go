package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetSkills(ctx context.Context) ([]model.SkillWithClasses, error) {
	/*	span := jaeger.GetSpan(ctx, "GetSkills")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `with cs as (
		select s."name", coalesce(s.alias, (select alias from skill where id = s.parent_skill_id)) as alias, jsonb_agg(jsonb_build_object('alias', c.alias, 'name', c."name", 'shortName', c.short_name, 'isClassSkill', cs.class_id is not null, 'book', jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order")) order by c.short_name) as classes 
			from skill s 
				cross join "class" c
				join book b on b.id = c.book_id 
				left join class_skill cs on cs.skill_id = s.id and cs.class_id = c.id
			where not coalesce(c.is_archetype, false) and not coalesce(c.is_npc, false) and coalesce(s.alias, '') != 'knowledge'
			group by s."name", s.alias, s.parent_skill_id
	), pcs as (
		select s."name", coalesce(s.alias, (select alias from skill where id = s.parent_skill_id)) as alias, jsonb_agg(jsonb_build_object('alias', pc.alias, 'name', pc."name", 'shortName', pc.short_name, 'isClassSkill', pcs.prestigeclass_id is not null, 'book', jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order")) order by pc.short_name) as classes 
			from skill s 
				cross join prestige_class pc 
				join book b on b.id = pc.book_id 
				left join prestige_class_skill pcs on pcs.skill_id = s.id and pcs.prestigeclass_id = pc.id
			where coalesce(s.alias, '') != 'knowledge'
			group by s."name", s.alias, s.parent_skill_id
	)
	select coalesce(cs.name, pcs.name) as name, coalesce(cs.alias, pcs.alias) as alias, cs.classes as classes, pcs.classes as prestige_classes
		from cs cs
			full join pcs pcs on pcs.name = cs.name`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetSkills")

	res := make([]model.SkillWithClasses, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetSkills error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetSkillsPerLvl(ctx context.Context) ([]model.SkillsPerLvlInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetSkillsPerLvl")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select alias, name, skill_ranks_per_level, is_prestige from (
		select c.alias, c."name", c.skill_ranks_per_level, false as is_prestige from "class" c where not coalesce(c.is_archetype, false) and not coalesce(c.is_npc, false)
		union all
		select pc.alias, pc."name", pc.skill_ranks_per_level, true as is_prestige from prestige_class pc
	) a
	order by is_prestige, name`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetSkillsPerLvl")

	res := make([]model.SkillsPerLvlInfo, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetSkillsPerLvl error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetSkillInfo(ctx context.Context, skillAlias string) (*model.SkillInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetSkillInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select s.description, s.full_description, jsonb_build_object('name', a.name, 'alias', a.alias, 'shortName', a.short_name, 'isSkillArmorPenalty', a.is_skill_armor_penalty) as ability 
	from skill s 
		join ability a on a.id = s.ability_id 
	where s.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "skillAlias": skillAlias}, "GetSkillInfo")

	var res model.SkillInfo
	err := db.GetClient().GetContext(ctx, &res, query, skillAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "skillAlias": skillAlias}, "GetSkillInfo error", err)
		}
		return nil, err
	}

	return &res, nil
}
