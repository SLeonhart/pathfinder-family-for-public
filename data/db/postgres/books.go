package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetBooks(ctx context.Context) ([]model.Books, error) {
	/*	span := jaeger.GetSpan(ctx, "GetBooks")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select b.alias, b."name", b.abbreviation
	from book b
	order by b."name"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetBooks")

	books := make([]model.Books, 0)
	err := db.GetClient().SelectContext(ctx, &books, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetBooks error", err)
		return nil, err
	}

	return books, nil
}

func (db *Postgres) GetBookInfo(ctx context.Context, bookAlias string) (*model.BookInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetBookInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from race d
			join book b on b.id = d.book_id
		where b.alias = $1) as races_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from "class" d
			join book b on b.id = d.book_id
		where b.alias =$1 and not coalesce(d.is_archetype, false)) as classes_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name", 'classAlias', pc.alias) order by d."name")
		from "class" d
			join book b on b.id = d.book_id
			join class_class cc on cc.child_class_id = d.id
			join "class" pc on pc.id = cc.parent_class_id
		where b.alias =$1 and coalesce(d.is_archetype, false)) as archetypes_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from feat d
			join book b on b.id = d.book_id
		where b.alias = $1) as feats_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from prestige_class d
			join book b on b.id = d.book_id
		where b.alias = $1) as prestige_classes_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from trait d
			join book b on b.id = d.book_id
		where b.alias = $1) as traits_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from god d
			join book b on b.id = d.book_id
		where b.alias = $1) as gods_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from "domain" d
			join domain_type dt on dt.id = d.type_id
			join book b on b.id = d.book_id
		where b.alias =$1 and dt.alias = 'domain') as domains_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from "domain" d
			join domain_type dt on dt.id = d.type_id
			join book b on b.id = d.book_id
		where b.alias =$1 and dt.alias = 'subdomain') as subdomains_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from "domain" d
			join domain_type dt on dt.id = d.type_id
			join book b on b.id = d.book_id
		where b.alias =$1 and dt.alias = 'inquisition') as inquisitions_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name", 'classAlias', c.alias) order by d."name")
		from bloodline d
			join book b on b.id = d.book_id
			join "class" c on c.id = d.class_id
		where b.alias = $1) as bloodlines_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from school d
			join book b on b.id = d.book_id
		where b.alias = $1) as schools_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from spell d
			join book b on b.id = d.book_id
		where b.alias = $1) as spells_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from weapon d
			join book b on b.id = d.book_id
		where b.alias = $1) as weapons_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from armor d
			join book b on b.id = d.book_id
		where b.alias = $1) as armors_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from equipment d
			join book b on b.id = d.book_id
		where b.alias = $1) as equipments_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name", 'types', (select jsonb_agg(distinct mit.alias)
			from magic_item_type mit
				join magic_item_ability_magic_item_type miamit on mit.id = miamit.magicitem_type_id
			where miamit.magicitem_ability_id = d.id)) order by d."name")
		from magic_item_ability d
			join book b on b.id = d.book_id
		where b.alias = $1) as magic_item_abilities_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from magic_item d
			join book b on b.id = d.book_id
		where b.alias = $1) as magic_items_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from monster_ability d
			join book b on b.id = d.book_id
		where b.alias = $1) as monster_abilities_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from beast d
			join book b on b.id = d.book_id
		where b.alias = $1) as beasts_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from affliction d
			join book b on b.id = d.book_id
		where b.alias = $1) as afflictions_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from trap d
			join book b on b.id = d.book_id
		where b.alias = $1) as traps_json,
	(select jsonb_agg(jsonb_build_object('alias', d.alias, 'name', d."name") order by d."name")
		from haunt d
			join book b on b.id = d.book_id
		where b.alias = $1) as haunts_json`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "bookAlias": bookAlias}, "GetBookInfo")

	var bookInfo model.BookInfo
	err := db.GetClient().GetContext(ctx, &bookInfo, query, bookAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "bookAlias": bookAlias}, "GetBookInfo error", err)
		}
		return nil, err
	}

	return &bookInfo, nil
}

func (db *Postgres) GetBotBooks(ctx context.Context, withSpells bool) ([]model.BotBook, error) {
	/*	span := jaeger.GetSpan(ctx, "GetBooks")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select b.id, b."name"
	from book b 
	where 1=1 `

	if withSpells {
		query += ` and exists(select 1 from spell s where s.book_id = b.id)`
	}

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetBooks")

	books := make([]model.BotBook, 0)
	err := db.GetClient().SelectContext(ctx, &books, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetBooks error", err)
		return nil, err
	}

	return books, nil
}
