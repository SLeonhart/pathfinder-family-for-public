package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetWeapons(ctx context.Context) ([]model.Weapon, error) {
	/*	span := jaeger.GetSpan(ctx, "GetWeapons")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `with w as (
		select w.id, w.alias, w."name", w.eng_name, w."type", w."cost", w.damage_s, w.damage_m, w.critical_roll, w.critical_damage, w."range", w.misfire, w.capacity, w.weight, w.special, w.description,
			case when wpc.alias is null then null else jsonb_build_object('alias', wpc.alias, 'name', wpc."name") end as proficient_category,
			case when wrc.alias is null then null else jsonb_build_object('alias', wrc.alias, 'name', wrc."name") end as range_category,
			case when wec.alias is null then null else jsonb_build_object('alias', wec.alias, 'name', wec."name") end as encumbrance_category,
			(select array_agg(pw."name")
				from weapon pw
					join weapon_weapon ww on ww.parent_weapon_id = pw.id
				where ww.child_weapon_id = w.id) as parents,
			jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
			from weapon w
				join book b on b.id = w.book_id
				left join weapon_proficient_category wpc on wpc.id = w.proficient_category_id
				left join weapon_range_category wrc on wrc.id = w.range_category_id
				left join weapon_encumbrance_category wec on wec.id = w.encumbrance_category_id
	)
	select w.alias, w."name", w.eng_name, w."type", w."cost", w.damage_s, w.damage_m, w.critical_roll, w.critical_damage, w."range", w.misfire, w.capacity, w.weight, w.special, w.description, w.proficient_category, w.range_category, w.encumbrance_category, w.parents, w.book,
		(select jsonb_agg(jsonb_build_object('alias', cw.alias, 'name', cw."name", 'engName', cw.eng_name, 'type', cw."type", 'cost', cw."cost", 'damageS', cw.damage_s, 'damageM', cw.damage_m, 'criticalRoll', cw.critical_roll, 'criticalDamage', cw.critical_damage, 'range', cw."range", 'misfire', cw.misfire, 'capacity', cw.capacity, 'weight', cw.weight, 'special', cw.special, 'description', cw.description, 'proficientCategory', cw.proficient_category, 'rangeCategory', cw.range_category, 'encumbranceCategory', cw.encumbrance_category, 'parents', cw.parents, 'book', cw.book))
			from w cw
				join weapon_weapon ww on ww.child_weapon_id = cw.id
			where ww.parent_weapon_id = w.id) as childs
		from w`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetWeapons")

	res := make([]model.Weapon, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetWeapons error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetArmors(ctx context.Context) ([]model.Armor, error) {
	/*	span := jaeger.GetSpan(ctx, "GetArmors")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select a.alias, a."name", a.eng_name, a.armor_bonus, a.max_dex_bonus, a.armor_check_penalty, a."cost", a.arcane_spell_failure_chance, a.speed30, a.speed20, a.weight, a.description,
	jsonb_build_object('alias', atp.alias, 'name', atp."name") as armor_type,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from armor a
		join book b on b.id = a.book_id
		join armor_type atp on atp.id = a.armor_type_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetArmors")

	res := make([]model.Armor, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetArmors error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetGoodsAndServices(ctx context.Context) ([]model.GoodAndService, error) {
	/*	span := jaeger.GetSpan(ctx, "GetGoodsAndServices")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `with e as (
		select e.id, e.alias, e."name", e.eng_name, e.equipment_sub_type, e."cost", e.cost_special, e.cost_description, e.cost_of_passage, e.weight, e.weight_special, e.weight_description, e.craft_dc, e.description,
			jsonb_build_object('alias', et.alias, 'name', et."name", 'description', et.description) as equipment_type,
			(select array_agg(pe."name")
				from equipment pe
					join equipment_equipment ee on ee.parent_equipment_id = pe.id
				where ee.child_equipment_id = e.id) as parents,
			jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
			from equipment e
				join book b on b.id = e.book_id
				join equipment_type et on et.id = e.equipment_type_id
	)
	select e.alias, e."name", e.eng_name, e.equipment_sub_type, e."cost", e.cost_special, e.cost_description, e.cost_of_passage, e.weight, e.weight_special, e.weight_description, e.craft_dc, e.description, e.equipment_type, e.parents, e.book,
			(select jsonb_agg(jsonb_build_object('alias', ce.alias, 'name', ce."name", 'engName', ce.eng_name, 'equipmentSubType', ce.equipment_sub_type, 'cost', ce."cost", 'costSpecial', ce.cost_special, 'costDescription', ce.cost_description, 'costOfPassage', ce.cost_of_passage, 'weight', ce.weight, 'weightSpecial', ce.weight_special, 'weightDescription', ce.weight_description, 'craftDc', ce.craft_dc, 'description', ce.description, 'type', ce.equipment_type, 'parents', ce.parents, 'book', ce.book))
				from e ce
					join equipment_equipment ee on ee.child_equipment_id = ce.id
				where ee.parent_equipment_id = e.id) as childs
		from e e`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetGoodsAndServices")

	res := make([]model.GoodAndService, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetGoodsAndServices error", err)
		return nil, err
	}

	return res, nil
}
