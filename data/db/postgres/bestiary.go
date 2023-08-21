package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetBeasts(ctx context.Context) ([]model.Beast, error) {
	cacheKey := "GetBeasts"
	if res := db.cache.Get(cacheKey); res != nil {
		return res.([]model.Beast), nil
	}

	/*	span := jaeger.GetSpan(ctx, "GetBeasts")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `with b as (
		select b.id, b.alias, b."name", b.eng_name, b.parent_id, b.cr, b.description, b."exp", b.full_creature_type, b.names, b.is_unique,
				case when c.alias is null then null else jsonb_build_object('alias', c.alias, 'name', c."name") end as climate,
				case when ct.alias is null then null else jsonb_build_object('alias', ct.alias, 'name', ct."name") end as creature_type,
				case when t.alias is null then null else jsonb_build_object('alias', t.alias, 'name', t."name") end as terrain,
				(select jsonb_agg(jsonb_build_object('name', cr."name", 'alias', cr.alias))
					from creature_role cr
						join beast_role br on br.creature_role_id = cr.id
					where br.beast_id = b.id) as roles,
				jsonb_build_object('alias', bk.alias, 'name', bk."name", 'order', bk."order", 'abbreviation', bk.abbreviation) as book
			from beast b
				join book bk on bk.id = b.book_id
				left join climate c on c.id = b.climate_id
				left join creature_type ct on ct.id = b.creature_type_id
				left join terrain t on t.id = b.terrain_id
	)
	select b.id, b.alias, b."name", b.eng_name, b.cr, b.description, b."exp", b.full_creature_type, b.names, b.is_unique, b.climate, b.creature_type, b.terrain, b.roles, b.book,
		(select array_agg(cb.id)
			from b cb
			where cb.parent_id = b.id) as childs
		from b`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetBeasts")

	res := make([]model.Beast, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetBeasts error", err)
		return nil, err
	}

	go db.cache.Set(cacheKey, res)

	return res, nil
}

func (db *Postgres) GetBeastInfo(ctx context.Context, beastAlias string) (*model.BeastInfo, error) {
	cacheKey := "GetBeastInfo:" + beastAlias
	if res := db.cache.Get(cacheKey); res != nil {
		return res.(*model.BeastInfo), nil
	}

	/*	span := jaeger.GetSpan(ctx, "GetBeastInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select b.alias, b."name", b.eng_name, b.description, b.names, b.root_page_description, b.cr, b."exp", b.full_creature_type, b.initiative, b.senses, b.perception, b.perception_comment, b.aura, b.strength, b.dexterity, b.constitution, b.intelligence, b.wisdom, b.charisma, b.max_ac_dexterity, b.ac_natural, b.ac_armor, b.ac_shield, b.ac_dodge, b.ac_deflection, b.ac_insight, b.ac_rage, b.ac_wisdom, b.ac_monk, b.ac_description, b.ac_string, b.hit_points, b.hit_points_description, b.hit_points_comment, b.fast_healing, b.regeneration, b.fortitude, b.reflex, b.will, b.will_comment, b.defensive_abilities, b.damage_resist, b.immune, b.resist, b.spell_resist, b.spell_resist_comment, b.weaknesses, b.speed, b.melee_attacks, b.ranged_attacks, b."space", b.reach, b.special_attacks, b.spell_like_abilities, b.spells_prepared, b.spells_known, b.domains, b.patron, b.bloodline, b.school, b.opposition_schools, b.before_combat, b.during_combat, b.morale, b.base_parameters, b.base_attack, b.combat_maneuver_bonus, b.combat_maneuver_bonus_comment, b.combat_maneuver_defense, b.combat_maneuver_defense_comment, b.feats, b.skills, b.skills_racial_modifiers, b.languages, b.combat_gear, b.other_gear, b.gear, b.spellbook, b.special_qualities, b.environment, b.organization, b.treasure, b.special_abilities, b.full_description, b.construction_description, b.construction_caster_level, b.construction_price, b.construction_price_comment, b.construction_requirements, b.construction_skill, b.cost_price, b.cost_price_comment,
	case when c.alias is null then null else jsonb_build_object('alias', c.alias, 'name', c."name") end as climate,
	case when ct.alias is null then null else jsonb_build_object('alias', ct.alias, 'name', ct."name") end as creature_type,
	case when t.alias is null then null else jsonb_build_object('alias', t.alias, 'name', t."name") end as terrain,
	case when st.alias is null then null else jsonb_build_object('alias', st.alias, 'name', st."name") end as size_type,
	case when pb.alias is null then null else jsonb_build_object('alias', pb.alias, 'name', pb."name") end as parent,
	(select jsonb_agg(jsonb_build_object('alias', ac.alias, 'name', ac."name", 'typeName', act."name", 'typeAlias', act.alias))
		from animal_companion ac
			join animal_companion_type act on act.id = ac.type_id
		where ac.beast_id = b.id) as animal_companion,
	(select jsonb_agg(jsonb_build_object('name', cr."name", 'alias', cr.alias))
		from creature_role cr
			join beast_role br on br.creature_role_id = cr.id
		where br.beast_id = b.id) as roles,
	(select jsonb_agg(jsonb_build_object('alias', cb.alias, 'name', cb."name"))
		from beast cb
		where cb.parent_id = b.id) as childs,
	jsonb_build_object('alias', bk.alias, 'name', bk."name", 'order', bk."order", 'abbreviation', bk.abbreviation) as book,
	(select jsonb_agg(jsonb_build_object('alias', h.alias, 'name', h."name", 'isMain', hb.is_main_helper) order by h."name") from helper_beast hb join helper h on h.id = hb.helper_id where hb.beast_id = b.id) as helpers
	from beast b
		join book bk on bk.id = b.book_id
		left join climate c on c.id = b.climate_id
		left join creature_type ct on ct.id = b.creature_type_id
		left join terrain t on t.id = b.terrain_id
		left join size_type st on st.id = b.size_id
		left join beast pb on pb.id = b.parent_id
	where b.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "beastAlias": beastAlias}, "GetBeastInfo")

	var res model.BeastInfo
	err := db.GetClient().GetContext(ctx, &res, query, beastAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "beastAlias": beastAlias}, "GetBeastInfo error", err)
		}
		return nil, err
	}

	go db.cache.Set(cacheKey, &res)

	return &res, nil
}

func (db *Postgres) GetMonsterAbilities(ctx context.Context) ([]model.MonsterAbility, error) {
	cacheKey := "GetMonsterAbilities"
	if res := db.cache.Get(cacheKey); res != nil {
		return res.([]model.MonsterAbility), nil
	}

	/*	span := jaeger.GetSpan(ctx, "GetMonsterAbilities")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select ma.alias, ma."name", ma.eng_name, ma.description, ma."type", ma.format, ma."location",
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from monster_ability ma
		join book b on b.id = ma.book_id
	order by ma."name"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetMonsterAbilities")

	res := make([]model.MonsterAbility, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetMonsterAbilities error", err)
		return nil, err
	}

	go db.cache.Set(cacheKey, res)

	return res, nil
}

func (db *Postgres) GetCreatureTypes(ctx context.Context) ([]model.CreatureType, error) {
	cacheKey := "GetCreatureTypes"
	if res := db.cache.Get(cacheKey); res != nil {
		return res.([]model.CreatureType), nil
	}

	/*	span := jaeger.GetSpan(ctx, "GetCreatureTypes")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select ct.alias, ct."name", ct.eng_name, ct.description, ct.features, ct.traits, ct.is_subtype
	from creature_type ct
	order by ct."name"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetCreatureTypes")

	res := make([]model.CreatureType, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetCreatureTypes error", err)
		return nil, err
	}

	go db.cache.Set(cacheKey, res)

	return res, nil
}

func (db *Postgres) GetAnimalCompanions(ctx context.Context) ([]model.AnimalCompanion, error) {
	cacheKey := "GetAnimalCompanions"
	if res := db.cache.Get(cacheKey); res != nil {
		return res.([]model.AnimalCompanion), nil
	}

	/*	span := jaeger.GetSpan(ctx, "GetAnimalCompanions")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select ac.alias, ac."name", ac.eng_name, ac.prerequisites, sst.description as start_size, ust.description as update_size, ac.start_speed, ac.update_speed, ac.start_ac, ac.update_ac, ac.start_attack, ac.update_attack, ac.start_strength, ac.update_strength, ac.start_dexterity, ac.update_dexterity, ac.start_constitution, ac.update_constitution, ac.start_intelligence, ac.update_intelligence, ac.start_wisdom, ac.update_wisdom, ac.start_charisma, ac.update_charisma, ac.start_languages, ac.update_languages, ac.start_special_attacks, ac.update_special_attacks, ac.start_special_qualities, ac.update_special_qualities, ac.start_combat_maneuver_defense, ac.update_combat_maneuver_defense, ac.update_level, ac.mastery_level, ac.mastery_description, ac.description, ac.start_racial_skill_modifiers, ac.update_racial_skill_modifiers, start_bonus_feat, update_bonus_feat, act.alias as "type",
	case when b.alias is null then null else jsonb_build_object('alias', b.alias, 'name', b."name") end as beast,
	jsonb_build_object('alias', bk.alias, 'name', bk."name", 'order', bk."order", 'abbreviation', bk.abbreviation) as book
	from animal_companion ac
		join animal_companion_type act on act.id = ac.type_id
		join book bk on bk.id = ac.book_id
		left join size_type sst on sst.id = ac.start_size_id
		left join size_type ust on ust.id = ac.update_size_id
		left join beast b on b.id = ac.beast_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetAnimalCompanions")

	res := make([]model.AnimalCompanion, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetAnimalCompanions error", err)
		return nil, err
	}

	go db.cache.Set(cacheKey, res)

	return res, nil
}
