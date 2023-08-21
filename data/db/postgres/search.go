package postgres

import (
	"context"
	"strings"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetSearchInfo(ctx context.Context) ([]model.SearchInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetSearchInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	queryArray := []string{
		`select 'ability_'||alias as id, 'ability' as "type", '/ability/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description as "content"
			from ability`,
		`select 'afflictions_'||alias as id, 'afflictions' as "type", '/afflictions?menu=alphabet_'||upper(left("name",1))||'&scroll='||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description as "content"
			from affliction`,
		`select 'animal_companion_'||act.alias||'_'||ac.alias as id, 'animal_companion_'||act.alias as "type", '/bestiary/'||act.alias||'?menu='||upper(left(ac."name",1))||'&scroll='||ac.alias as url, case when ac.eng_name is null then ac."name" else ac."name"||' ('||ac.eng_name||')' end as h1, null as "content"
			from animal_companion ac
				join animal_companion_type act on act.id = ac.type_id`,
		`select 'armor_'||alias as id, 'armor' as "type", '/armors?menu=alphabet_'||upper(left("name",1))||'&scroll='||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description as "content"
			from armor`,
		`select 'beast_'||alias as id, 'beast' as "type", '/bestiary/beast/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, full_description as "content"
			from beast`,
		`select 'bloodline_'||c.alias||'_'||b.alias as id, 'bloodline' as "type", '/class/'||c.alias||'/bloodline/'||b.alias as url, case when b.eng_name is null then b."name" else b."name"||' ('||b.eng_name||')' end as h1, b.description as "content"
			from bloodline b
				join "class" c on c.id = b.class_id`,
		`select 'class_'||alias as id, 'class' as "type", '/class/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description||'
'||full_description||'
'||features as "content"
			from "class"
			where coalesce(is_archetype,false) = false and coalesce(is_npc,false) = false`,
		`select 'archetype_'||pc.alias||'_'||c.alias as id, 'archetype' as "type", '/class/archetype/'||pc.alias||'/'||c.alias as url, case when c.eng_name is null then c."name" else c."name"||' ('||c.eng_name||')' end as h1, c.description||'
'||c.full_description||'
'||c.features as "content"
			from "class" c
				join class_class cc on cc.child_class_id = c.id
				join "class" pc on pc.id = cc.parent_class_id
			where coalesce(c.is_archetype,false) = true and coalesce(c.is_npc,false) = false`,
		`select 'npc_'||alias as id, 'npc' as "type", '/npc/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description||'
'||full_description||'
'||features as "content"
			from "class"
			where coalesce(is_archetype,false) = false and coalesce(is_npc,false) = true`,
		`select 'domain_'||dt.alias||'_'||d.alias as id, 'domain' as "type", '/god/'||dt.alias||'/'||d.alias as url, case when d.eng_name is null then d."name" else d."name"||' ('||d.eng_name||')' end as h1, d.description as "content"
			from "domain" d
				join domain_type dt on dt.id = d.type_id `,
		`select 'equipment_'||alias as id, 'equipment' as "type", '/goodsAndServices?menu=alphabet_'||upper(left("name",1))||'&scroll='||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description as "content"
			from equipment`,
		`select 'feat_'||alias as id, 'feat' as "type", '/feat/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description||'
'||full_description||'
'||benefit||'
'||normal||'
'||special as "content"
			from feat`,
		`select 'god_'||g.alias as id, 'god' as "type", '/god?menu=type_'||gt.alias||'&scroll='||g.alias as url, case when g.eng_name is null then g."name" else g."name"||' ('||g.eng_name||')' end as h1, g.title||'
'||portfolios as "content"
			from god g
				join god_type gt on gt.id = g.type_id `,
		`select 'haunt_'||alias as id, 'haunt' as "type", '/haunt?menu=cr_'||cr||'&scroll='||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, effect as "content"
			from haunt`,
		`select 'magic_item_'||alias as id, 'magic_item' as "type", '/magicItem/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description as "content"
			from magic_item`,
		`select 'magic_item_ability_armor_shield_'||mia.alias as id, 'magic_item_ability_armor_shield' as "type", '/magicItems/armor?scroll='||mia.alias as url, case when mia.eng_name is null then mia."name" else mia."name"||' ('||mia.eng_name||')' end as h1, mia.description as "content"
			from magic_item_ability mia
				join magic_item_ability_magic_item_type miamit on miamit.magicitem_ability_id = mia.id
				join magic_item_type mit on mit.id = miamit.magicitem_type_id
			where mit.alias in ('armor','shield')
			group by mia.alias, mia.eng_name, mia."name", mia.description`,
		`select 'magic_item_ability_rangeWeapon_meeleWeapon_'||mia.alias as id, 'magic_item_ability_rangeWeapon_meeleWeapon' as "type", '/magicItems/weapons?scroll='||mia.alias as url, case when mia.eng_name is null then mia."name" else mia."name"||' ('||mia.eng_name||')' end as h1, mia.description as "content"
			from magic_item_ability mia
				join magic_item_ability_magic_item_type miamit on miamit.magicitem_ability_id = mia.id
				join magic_item_type mit on mit.id = miamit.magicitem_type_id
			where mit.alias in ('rangeWeapon','meeleWeapon')
			group by mia.alias, mia.eng_name, mia."name", mia.description`,
		`select 'monster_ability_'||alias as id, 'monster_ability' as "type", '/bestiary/appendix/universalMonsterRules?scroll='||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description as "content"
			from monster_ability`,
		`select 'creature_type_'||alias as id, 'creature_type' as "type", '/bestiary/appendix/creatureTypes?scroll='||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description as "content"
			from creature_type`,
		`select 'orders_'||alias as id, 'orders' as "type", '/class/samurai_cavalier/order/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description||'
'||full_description as "content"
			from orders`,
		`select 'prestige_class_'||alias as id, 'prestige_class' as "type", '/class/prestige/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description||'
'||full_description||'
'||"role"||'
'||requirements||'
'||features as "content"
			from prestige_class`,
		`select 'race_'||alias as id, 'race' as "type", '/race/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description||'
'||adventurers||'
'||alignment_and_religion||'
'||physical_description||'
'||relations||'
'||society as "content"
			from race`,
		`select 'school_'||alias as id, 'school' as "type", '/spell/school/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description||'
'||full_description as "content"
			from school`,
		`select 'skill_'||alias as id, 'skill' as "type", '/skill/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description||'
'||full_description as "content"
			from skill
			where alias is not null`,
		`select 'spell_'||alias as id, 'spell' as "type", '/spell/'||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description||'
'||short_description as "content"
			from spell`,
		`select 'trait_'||alias as id, 'trait' as "type", '/traits?menu=alphabet_'||upper(left("name",1))||'&scroll='||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, benefit as "content"
			from trait`,
		`select 'trap_'||alias as id, 'trap' as "type", '/traps?menu=cr_'||cr||'&scroll='||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, effect as "content"
			from trap`,
		`select 'weapon_'||alias as id, 'weapon' as "type", '/weapons?menu=alphabet_'||upper(left("name",1))||'&scroll='||alias as url, case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as h1, description as "content"
			from weapon`,
		`select 'wild_talent_'||wt.alias as id, 'wild_talent' as "type", '/class/kineticist/wildTalent?menu=type_'||wtt.alias||'&scroll='||wt.alias as url, case when wt.eng_name is null then wt."name" else wt."name"||' ('||wt.eng_name||')' end as h1, wt.description as "content"
			from wild_talent wt
				join wild_talent_type wtt on wtt.id =  wt.type_id`,
	}

	query := strings.Join(queryArray, " union all ")

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetSearchInfo")

	classes := make([]model.SearchInfo, 0)
	err := db.GetClient().SelectContext(ctx, &classes, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetSearchInfo error", err)
		return nil, err
	}

	return classes, nil
}
