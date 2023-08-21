package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetThanks(ctx context.Context) ([]model.Thanks, error) {
	/*	span := jaeger.GetSpan(ctx, "GetThanks")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `with donate as (
		select h."name", h.alias, sum(d.sum) as cnt 
			from helper h
				join donate d on d.helper_id = h.id
			group by h."name", h.alias
			order by cnt desc, name
	), beast as (
		select h."name", h.alias, count(*) as cnt 
			from helper h
				join helper_beast x on x.helper_id = h.id
			group by h."name", h.alias
			order by cnt desc, name
	), race as (
		select h."name", h.alias, count(*) as cnt 
			from helper h
				join helper_race x on x.helper_id = h.id
			group by h."name", h.alias
			order by cnt desc, name
	), h_class as (
		select h."name", h.alias, count(*) as cnt 
			from helper h
				join helper_class x on x.helper_id = h.id
				join "class" c on c.id = x.class_id
			where coalesce(c.is_archetype, false) = false
			group by h."name", h.alias
			order by cnt desc, name
	), prestige_class as (
		select h."name", h.alias, count(*) as cnt 
			from helper h
				join helper_prestige_class x on x.helper_id = h.id
			group by h."name", h.alias
			order by cnt desc, name
	), feat as (
		select h."name", h.alias, count(*) as cnt 
			from helper h
				join helper_feat x on x.helper_id = h.id
			group by h."name", h.alias
			order by cnt desc, name
	), trait as (
		select h."name", h.alias, count(*) as cnt 
			from helper h
				join helper_trait x on x.helper_id = h.id
			group by h."name", h.alias
			order by cnt desc, name
	), spell as (
		select h."name", h.alias, count(*) as cnt 
			from helper h
				join helper_spell x on x.helper_id = h.id
			group by h."name", h.alias
			order by cnt desc, name
	), translation as (
		select h."name", h.alias, count(*) as cnt 
			from helper h
				join helper_translation x on x.helper_id = h.id
			group by h."name", h.alias
			order by cnt desc, name
	), archetype as (
		select h."name", h.alias, count(*) as cnt 
			from helper h
				join helper_class x on x.helper_id = h.id
				join "class" c on c.id = x.class_id
			where c.is_archetype
			group by h."name", h.alias
			order by cnt desc, name
	)
	select 'donate' as "type", (select jsonb_agg(jsonb_build_object('name', name, 'alias', alias, 'cnt', cnt)) from donate) as list
	union all
	select 'beast' as "type", (select jsonb_agg(jsonb_build_object('name', name, 'alias', alias, 'cnt', cnt)) from beast) as list
	union all
	select 'race' as "type", (select jsonb_agg(jsonb_build_object('name', name, 'alias', alias, 'cnt', cnt)) from race) as list
	union all
	select 'class' as "type", (select jsonb_agg(jsonb_build_object('name', name, 'alias', alias, 'cnt', cnt)) from h_class) as list
	union all
	select 'prestigeClass' as "type", (select jsonb_agg(jsonb_build_object('name', name, 'alias', alias, 'cnt', cnt)) from prestige_class) as list
	union all
	select 'archetype' as "type", (select jsonb_agg(jsonb_build_object('name', name, 'alias', alias, 'cnt', cnt)) from archetype) as list
	union all
	select 'feat' as "type", (select jsonb_agg(jsonb_build_object('name', name, 'alias', alias, 'cnt', cnt)) from feat) as list
	union all
	select 'trait' as "type", (select jsonb_agg(jsonb_build_object('name', name, 'alias', alias, 'cnt', cnt)) from trait) as list
	union all
	select 'spell' as "type", (select jsonb_agg(jsonb_build_object('name', name, 'alias', alias, 'cnt', cnt)) from spell) as list
	union all
	select 'translation' as "type", (select jsonb_agg(jsonb_build_object('name', name, 'alias', alias, 'cnt', cnt)) from translation) as list`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetThanks")

	thanks := make([]model.Thanks, 0)
	err := db.GetClient().SelectContext(ctx, &thanks, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetThanks error", err)
		return nil, err
	}

	return thanks, nil
}
