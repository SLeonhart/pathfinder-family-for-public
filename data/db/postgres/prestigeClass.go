package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetPrestigeClasses(ctx context.Context) ([]model.Class, error) {
	/*	span := jaeger.GetSpan(ctx, "GetPrestigeClasses")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select pc.alias, pc."name", pc.description, jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation) as book_json
	from prestige_class pc 
		join book b on b.id = pc.book_id 
	order by pc."name"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetPrestigeClasses")

	prestigeClasses := make([]model.Class, 0)
	err := db.GetClient().SelectContext(ctx, &prestigeClasses, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetPrestigeClasses error", err)
		return nil, err
	}

	return prestigeClasses, nil
}

func (db *Postgres) GetPrestigeClassInfo(ctx context.Context, classAlias string) (*model.PrestigeClassInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetPrestigeClassInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select pc."name", pc.full_description, pc."role", pc.alignment, pc.hit_die, pc.requirements, pc.skill_ranks_per_level, pc.table_features, pc.table_spell_count, pc.features,
	(select jsonb_agg(jsonb_build_object('alias', s.alias, 'name', s."name", 'ability', jsonb_build_object('alias', a.alias, 'name', a.name, 'shortName', a.short_name, 'isSkillArmorPenalty', a.is_skill_armor_penalty)) order by s."name") 
		from skill s
			join ability a on a.id = s.ability_id
			join prestige_class_skill pcs on pcs.skill_id = s.id
		where pcs.prestigeclass_id = pc.id) as skills,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'abbreviation', b.abbreviation) as book_json,
	(select jsonb_agg(jsonb_build_object('alias', h.alias, 'name', h."name", 'isMain', hpc.is_main_helper) order by h."name") from helper_prestige_class hpc join helper h on h.id = hpc.helper_id where hpc.prestigeclass_id = pc.id) as helpers_json
	from prestige_class pc
		join book b on b.id = pc.book_id 
	where pc.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetPrestigeClassInfo")

	var prestigeClassInfo model.PrestigeClassInfo
	err := db.GetClient().GetContext(ctx, &prestigeClassInfo, query, classAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "classAlias": classAlias}, "GetPrestigeClassInfo error", err)
		}
		return nil, err
	}

	return &prestigeClassInfo, nil
}
