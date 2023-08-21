package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetRaces(ctx context.Context) ([]model.Race, error) {
	/*	span := jaeger.GetSpan(ctx, "GetRaces")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select r.alias, r."name", jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation) as book_json
	from race r 
		join book b on b.id = r.book_id 
	order by r."name"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetRaces")

	races := make([]model.Race, 0)
	err := db.GetClient().SelectContext(ctx, &races, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetRaces error", err)
		return nil, err
	}

	return races, nil
}

func (db *Postgres) GetRaceInfo(ctx context.Context, raceAlias string) (*model.RaceInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetRaceInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select r.description, r.physical_description, r.society, r.relations, r.alignment_and_religion, r.adventurers, r.names_description, r.additional_description,
	(select jsonb_agg(jsonb_build_object('name', rn."name", 'isMale', rn.is_male)) from race_name rn where rn.race_id = r.id) as names,
	(select jsonb_agg(jsonb_build_object('name', rt."name", 'description', rt.description) order by rt."order")
		from race_trait rt
			left join race_trait_race_trait rtrt on rtrt.additionalracetrait_id = rt.id 
		where rt.race_id = r.id and rtrt.baseracetrait_id is null) as base_race_traits_json,
	(select jsonb_agg(jsonb_build_object('name', rt."name", 'description', rt.description, 'baseRaceTraits', rt.base_race_traits) order by rt."order")
		from (select rt."name", rt.description, rt."order", array_agg(brt."name" order by brt."name") as base_race_traits
			from race_trait rt
				join race_trait_race_trait rtrt on rtrt.additionalracetrait_id = rt.id 
				join race_trait brt on brt.id = rtrt.baseracetrait_id 
			where rt.race_id = r.id and brt.race_id = r.id
			group by rt."name", rt.description, rt."order") rt) as alter_race_traits_json,
	(select jsonb_agg(jsonb_build_object('alias', c.alias, 'name', c."name", 'description', rc.favored_description) order by c."name") 
		from race_class rc 
			join "class" c on c.id = rc.class_id
		where (rc.race_id = r.id or rc.race_id is null) and coalesce(rc.favored_description,'') != '') as favored_class_json,
	(select jsonb_agg(jsonb_build_object('alias', c.alias, 'name', c."name", 'description', rc.description) order by c."name") 
		from race_class rc 
			join "class" c on c.id = rc.class_id
		where rc.race_id = r.id and coalesce(rc.description,'') != '') as adventurer_class_json,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation) as book_json,
	(select jsonb_agg(jsonb_build_object('alias', h.alias, 'name', h."name", 'isMain', hr.is_main_helper) order by h."name") from helper_race hr join helper h on h.id = hr.helper_id where hr.race_id = r.id) as helpers_json
	from race r
		join book b on b.id = r.book_id 
	where r.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "raceAlias": raceAlias}, "GetRaceInfo")

	var raceInfo model.RaceInfo
	err := db.GetClient().GetContext(ctx, &raceInfo, query, raceAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "raceAlias": raceAlias}, "GetRaceInfo error", err)
		}
		return nil, err
	}

	return &raceInfo, nil
}
