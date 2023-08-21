package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetAllMagicItems(ctx context.Context) ([]model.MagicItemForList, error) {
	/*	span := jaeger.GetSpan(ctx, "GetAllMagicItems")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select mi.alias, mi."name", mi.eng_name,
	case when mis.id is null then null else jsonb_build_object('alias', mis.alias, 'name', mis."name") end as slot,
	jsonb_build_object('alias', mit.alias, 'name', mit."name") as magic_item_type,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from magic_item mi
		join book b on b.id = mi.book_id
		join magic_item_type mit on mit.id = mi.type_id
		left join magic_item_slot mis on mis.id = mi.slot_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetAllMagicItems")

	res := make([]model.MagicItemForList, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetAllMagicItems error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetMagicItemInfo(ctx context.Context, magicItemAlias string) (*model.MagicItemInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetMagicItemInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select mi.alias, mi."name", mi.eng_name, mi.aura, mi.cl, mi.slot_comment, mi.price, mi.price_comment, mi.weight, mi.weight_comment, mi.description, mi.construction_requirements, mi.construction_cost, mi.creation_magic_items, mi."statistics", mi.destruction,
	case when mis.id is null then null else jsonb_build_object('alias', mis.alias, 'name', mis."name") end as slot,
	jsonb_build_object('alias', mit.alias, 'name', mit."name") as magic_item_type,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book,
	(select jsonb_agg(jsonb_build_object('alias', h.alias, 'name', h."name", 'isMain', hmi.is_main_helper) order by h."name") from helper_magic_item hmi join helper h on h.id = hmi.helper_id where hmi.magic_item_id = mi.id) as helpers
	from magic_item mi
		join book b on b.id = mi.book_id
		join magic_item_type mit on mit.id = mi.type_id
		left join magic_item_slot mis on mis.id = mi.slot_id
	where mi.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "magicItemAlias": magicItemAlias}, "GetMagicItemInfo")

	var res model.MagicItemInfo
	err := db.GetClient().GetContext(ctx, &res, query, magicItemAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "magicItemAlias": magicItemAlias}, "GetMagicItemInfo error", err)
		}
		return nil, err
	}

	return &res, nil
}

func (db *Postgres) GetMagicItemsByTypes(ctx context.Context, types []string) ([]model.MagicItemInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetMagicItemsByTypes")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	args := make([]interface{}, 0, len(types))
	typeQuerySlice := make([]string, 0, len(types))
	for i, v := range types {
		args = append(args, v)
		typeQuerySlice = append(typeQuerySlice, fmt.Sprintf("$%v", i+1))
	}
	query := fmt.Sprintf(`select mi.alias, mi."name", mi.eng_name, mi.aura, mi.cl, mi.slot_comment, mi.price, mi.price_comment, mi.weight, mi.weight_comment, mi.description, mi.construction_requirements, mi.construction_cost, mi.creation_magic_items, mi."statistics", mi.destruction,
	case when mis.id is null then null else jsonb_build_object('alias', mis.alias, 'name', mis."name") end as slot,
	jsonb_build_object('alias', mit.alias, 'name', mit."name") as magic_item_type,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from magic_item mi
		join book b on b.id = mi.book_id
		join magic_item_type mit on mit.id = mi.type_id
		left join magic_item_slot mis on mis.id = mi.slot_id
	where mit.alias in (%v)`, strings.Join(typeQuerySlice, ","))

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "types": types}, "GetMagicItemsByTypes")

	res := make([]model.MagicItemInfo, 0)
	err := db.GetClient().SelectContext(ctx, &res, query, args...)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "types": types}, "GetMagicItemsByTypes error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetMagicItemAbilitiesByTypes(ctx context.Context, types []string) ([]model.MagicItemAbility, error) {
	/*	span := jaeger.GetSpan(ctx, "GetMagicItemAbilitiesByTypes")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	args := make([]interface{}, 0, len(types))
	typeQuerySlice := make([]string, 0, len(types))
	for i, v := range types {
		args = append(args, v)
		typeQuerySlice = append(typeQuerySlice, fmt.Sprintf("$%v", i+1))
	}
	query := fmt.Sprintf(`select mia.alias, mia."name", mia.eng_name, mia.aura, mia.cl, mia.bonus_price, mia.money_price, mia.description, mia.construction_requirements,
    jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from magic_item_ability mia
		join book b on b.id = mia.book_id
		where exists(select *
			from magic_item_ability_magic_item_type miamit
				join magic_item_type mit on mit.id = miamit.magicitem_type_id
			where miamit.magicitem_ability_id = mia.id and mit.alias in (%v))`, strings.Join(typeQuerySlice, ","))

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "types": types}, "GetMagicItemAbilitiesByTypes")

	res := make([]model.MagicItemAbility, 0)
	err := db.GetClient().SelectContext(ctx, &res, query, args...)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "types": types}, "GetMagicItemAbilitiesByTypes error", err)
		return nil, err
	}

	return res, nil
}
