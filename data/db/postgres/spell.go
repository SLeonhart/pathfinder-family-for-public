package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pathfinder-family/model"
	"pathfinder-family/utils"
	"strings"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetSpellSchools(ctx context.Context) ([]model.School, error) {
	cacheKey := "GetSpellSchools"
	if res := db.cache.Get(cacheKey); res != nil {
		return res.([]model.School), nil
	}

	/*	span := jaeger.GetSpan(ctx, "GetSpellSchools")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select s."name", s.alias, s.description,
	jsonb_build_object('alias', st.alias, 'name', st."name", 'description', st.description) as school_type,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from school s
		join school_type st on st.id = s.type_id
		join book b on b.id = s.book_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetSpellSchools")

	res := make([]model.School, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetSpellSchools error", err)
		return nil, err
	}

	go db.cache.Set(cacheKey, res)

	return res, nil
}

func (db *Postgres) GetSpellSchoolInfo(ctx context.Context, schoolAlias string) (*model.SchoolInfo, error) {
	cacheKey := "GetSpellSchoolInfo:" + schoolAlias
	if res := db.cache.Get(cacheKey); res != nil {
		return res.(*model.SchoolInfo), nil
	}

	/*	span := jaeger.GetSpan(ctx, "GetSpellSchoolInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select s."name", s.alias, s.description, s.full_description,
	jsonb_build_object('alias', st.alias, 'name', st."name", 'description', st.description) as school_type,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from school s
		join school_type st on st.id = s.type_id
		join book b on b.id = s.book_id
	where s.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "schoolAlias": schoolAlias}, "GetSpellSchoolInfo")

	var res model.SchoolInfo
	err := db.GetClient().GetContext(ctx, &res, query, schoolAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "schoolAlias": schoolAlias}, "GetSpellSchoolInfo error", err)
		}
		return nil, err
	}

	go db.cache.Set(cacheKey, &res)

	return &res, nil
}

func (db *Postgres) GetSpells(ctx context.Context, classAlias *string) ([]model.Spell, error) {
	cacheKey := "GetSpells:" + utils.PtrToString(classAlias)

	if res := db.cache.Get(cacheKey); res != nil {
		return res.([]model.Spell), nil
	}

	/*	span := jaeger.GetSpan(ctx, "GetSpells")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/
	args := make([]interface{}, 0)
	query := `with cs as (
		select distinct c.alias, c."name", cs.spell_id, min(cs."level") as "level"
			from "class" c
			  left join class_class cc on cc.child_class_id = c.id
			  left join "class" pc on pc.id = cc.parent_class_id
			  left join class_spell cs on cs.class_id = c.id or (not coalesce(c.is_own_spell_list, false) and cs.class_id = pc.id and cs."level" <= coalesce(c.max_spell_lvl, 99))
			where not coalesce(c.is_archetype,false) and cs.id is not null
			group by c.alias, c."name", cs.spell_id
	)
	select s."name", s.alias, s.short_description,
		s.short_description_components||case when exists(select * from race_spell rs where rs.spell_id = s.id) then 'Н' else '' end as short_description_components,
	  (select jsonb_agg(jsonb_build_object('alias', sc.alias, 'name', sc."name", 'type', jsonb_build_object('alias', st.alias, 'name', st."name")))
		from school sc
		  join spell_school ss on ss.school_id = sc.id
		  join school_type st on st.id = sc.type_id
		where ss.spell_id = s.id) as schools,
	  (select jsonb_agg(distinct jsonb_build_object('alias', cs.alias, 'name', cs."name", 'level', cs."level"))
		from cs
		where cs.spell_id = s.id) as classes,
	  jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	  from spell s
		join book b on b.id = s.book_id `
	if classAlias != nil {
		query += ` where exists(select * from cs where spell_id = s.id and alias = $1)`
		args = append(args, classAlias)
	}

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetSpells")

	res := make([]model.Spell, 0)

	err := db.GetClient().SelectContext(ctx, &res, query, args...)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetSpells error", err)
		return nil, err
	}

	go db.cache.Set(cacheKey, res)

	return res, nil
}

func (db *Postgres) GetSpellInfo(ctx context.Context, spellAlias string) (*model.SpellInfo, error) {
	cacheKey := "GetSpellInfo:" + spellAlias
	if res := db.cache.Get(cacheKey); res != nil {
		return res.(*model.SpellInfo), nil
	}

	/*	span := jaeger.GetSpan(ctx, "GetSpellInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `with cs as (
		select c.alias, c."name", cs.spell_id, min(cs."level") as "level"
			from "class" c
			  left join class_class cc on cc.child_class_id = c.id
			  left join "class" pc on pc.id = cc.parent_class_id
			  left join class_spell cs on cs.class_id = c.id or (not coalesce(c.is_own_spell_list, false) and cs.class_id = pc.id and cs."level" <= coalesce(c.max_spell_lvl, 99))
			where not coalesce(c.is_archetype,false) and cs.id is not null
			group by c.alias, c."name", cs.spell_id
	)
	select s."name", s.eng_name, s.casting_time, s.components, s."range", s.target, s.area, s.effect, s.duration, s.saving_throw, s.spell_resistance, s.description, s.sub_school, 
	  jsonb_build_object('alias', sc.alias, 'name', sc."name", 'type', jsonb_build_object('alias', st.alias, 'name', st."name")) as school,
	  (select jsonb_agg(distinct jsonb_build_object('alias', cs.alias, 'name', cs."name", 'level', cs."level"))
		from cs
		where cs.spell_id = s.id) as classes,
	  (select jsonb_agg(jsonb_build_object('alias', r.alias, 'name', r."name"))
		from race_spell rs
			join race r on r.id = rs.race_id
		where rs.spell_id = s.id) as races,
	  jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book,
	  case when g.id is null then null else jsonb_build_object('alias', g.alias, 'name', g."name") end as god,
	  (select jsonb_agg(jsonb_build_object('alias', h.alias, 'name', h."name", 'isMain', hs.is_main_helper) order by h."name") from helper_spell hs join helper h on h.id = hs.helper_id where hs.spell_id = s.id) as helpers
	  from spell s
		join book b on b.id = s.book_id
		join spell_school ss on ss.spell_id = s.id 
		join school sc on sc.id = ss.school_id
		join school_type st on st.id = sc.type_id and st.alias = 'standart'
		left join god g on g.id = s.god_id
	  where s.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "spellAlias": spellAlias}, "GetSpellInfo")

	var res model.SpellInfo
	err := db.GetClient().GetContext(ctx, &res, query, spellAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "spellAlias": spellAlias}, "GetSpellInfo error", err)
		}
		return nil, err
	}

	go db.cache.Set(cacheKey, &res)

	return &res, nil
}

func (db *Postgres) GetBotSpells(ctx context.Context, id *int, name *string, engName *string, classId *int, alias *string, level *int, rulebookIds []int) ([]model.BotSpellInfo, error) {
	// cacheKey := "GetSpellInfo:" + spellAlias
	// if res := db.cache.Get(cacheKey); res != nil {
	// 	return res.(*model.SpellInfo), nil
	// }

	/*span := jaeger.GetSpan(ctx, "GetBotSpells")
	defer span.Finish()
	ctx = jaeger.SetParentSpan(ctx, span)*/

	i := 0
	where := ``
	args := []interface{}{}
	if id != nil {
		i++
		args = append(args, id)
		where += fmt.Sprintf(` and s.id = $%v `, i)
	}
	if name != nil {
		i++
		args = append(args, fmt.Sprintf("%%%v%%", *name))
		where += fmt.Sprintf(` and lower(s."name") like lower($%v) `, i)
	}
	if engName != nil {
		i++
		args = append(args, fmt.Sprintf("%%%v%%", *engName))
		where += fmt.Sprintf(` and lower(s.eng_name) like lower($%v) `, i)
	}
	if classId != nil {
		i++
		args = append(args, classId)
		where += fmt.Sprintf(` and exists(select * from cs where id = $%v and spell_id = s.id) `, i)
	}
	if alias != nil {
		i++
		args = append(args, alias)
		where += fmt.Sprintf(` and s.alias = $%v `, i)
	}
	if level != nil {
		i++
		args = append(args, level)
		where += fmt.Sprintf(` and exists(select * from cs where "level" = $%v and spell_id = s.id) `, i)
	}
	if len(rulebookIds) > 0 {
		where += fmt.Sprintf(` and b.id in (%v) `, strings.Trim(strings.Replace(fmt.Sprint(rulebookIds), " ", ",", -1), "[]"))
	}

	query := `with cs as (
		select c.id, c.alias, c."name", cs.spell_id, min(cs."level") as "level"
			from "class" c
			  left join class_class cc on cc.child_class_id = c.id
			  left join "class" pc on pc.id = cc.parent_class_id
			  left join class_spell cs on cs.class_id = c.id or (not coalesce(c.is_own_spell_list, false) and cs.class_id = pc.id and cs."level" <= coalesce(c.max_spell_lvl, 99))
			where not coalesce(c.is_archetype,false) and cs.id is not null
			group by c.id, c.alias, c."name", cs.spell_id
	)
	select s.id, s.alias, s."name", s.eng_name, s.short_description, s.short_description_components, s.casting_time, s.components, s."range", s.target, s.area, s.effect, s.duration, s.saving_throw, s.spell_resistance, s.description, s.sub_school,
	  jsonb_build_object('alias', sc.alias, 'name', sc."name", 'type', jsonb_build_object('alias', st.alias, 'name', st."name")) as school,
	  (select jsonb_agg(distinct jsonb_build_object('id', cs.id, 'alias', cs.alias, 'name', cs."name", 'level', cs."level"))
		from cs
		where cs.spell_id = s.id) as classes,
	  (select jsonb_agg(jsonb_build_object('alias', r.alias, 'name', r."name"))
		from race_spell rs
			join race r on r.id = rs.race_id
		where rs.spell_id = s.id) as races,
	  jsonb_build_object('id', b.id, 'alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book,
	  case when g.id is null then null else jsonb_build_object('alias', g.alias, 'name', g."name") end as god,
	  (select jsonb_agg(jsonb_build_object('alias', h.alias, 'name', h."name", 'isMain', hs.is_main_helper) order by h."name") from helper_spell hs join helper h on h.id = hs.helper_id where hs.spell_id = s.id) as helpers
	  from spell s
		join book b on b.id = s.book_id
		join spell_school ss on ss.spell_id = s.id 
		join school sc on sc.id = ss.school_id
		join school_type st on st.id = sc.type_id and st.alias = 'standart'
		left join god g on g.id = s.god_id
	  where 1=1 ` + where

	// logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "name": name, "engName": engName, "classId": classId, "level": level, "rulebookIds": rulebookIds}, "GetBotSpells")

	res := make([]model.BotSpellInfo, 0)

	err := db.GetClient().SelectContext(ctx, &res, query, args...)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "name": name, "engName": engName, "classId": classId, "level": level, "rulebookIds": rulebookIds}, "GetBotSpells error", err)
		}
		return nil, err
	}

	// go db.cache.Set(cacheKey, &res)

	return res, nil
}
