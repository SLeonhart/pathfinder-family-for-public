package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetDomains(ctx context.Context) ([]model.Domain, error) {
	/*	span := jaeger.GetSpan(ctx, "GetDomains")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select d.alias, d."name", d.description, dt.alias as domain_type, 
	(select jsonb_agg(jsonb_build_object('alias', cd.alias, 'name', cd."name", 'description', cd.description, 'domain_type', cdt.alias, 'book', jsonb_build_object('alias', cb.alias, 'name', cb."name", 'order', cb."order", 'abbreviation', cb.abbreviation))) 
		from "domain" cd
			join domain_type cdt on cdt.id = cd.type_id
			join book cb on cb.id = cd.book_id
			join domain_domain dd on dd.child_domain_id = cd.id
		where dd.parent_domain_id = d.id) as child_domains,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from "domain" d 
		join domain_type dt on dt.id = d.type_id
		join book b on b.id = d.book_id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetDomains")

	res := make([]model.Domain, 0)
	err := db.GetClient().SelectContext(ctx, &res, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetDomains error", err)
		return nil, err
	}

	return res, nil
}

func (db *Postgres) GetDomainInfo(ctx context.Context, domainType string, domainAlias string) (*model.DomainInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetDomainInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select d."name", d.description, d.inner_description, d.power0_name, d.power0_description, d.power1_name, d.power1_description, d.power2_name, d.power2_description,
	(select array_agg(g."name" order by g."name") 
		from god g
			join god_domain gd on gd.god_id = g.id
		where gd.domain_id = d.id) as gods,
	(select jsonb_agg(jsonb_build_object('alias', cd.alias, 'name', cd."name") order by cd."name") 
		from "domain" cd
			join domain_domain dd on dd.child_domain_id = cd.id
		where dd.parent_domain_id = d.id) as childs,
	(select jsonb_agg(jsonb_build_object('alias', pd.alias, 'name', pd."name", 'power1Name', pd.power1_name, 'power1Description', pd.power1_description, 'power2Name', pd.power2_name, 'power2Description', pd.power2_description, 'spells', (select jsonb_agg(jsonb_build_object('level', pds."level", 'comment', pds."comment", 'alias', ps.alias, 'name', ps."name") order by pds."level") 
			from domain_spell pds
				join spell ps on ps.id = pds.spell_id
			where pds.domain_id = pd.id)) order by pd."name") 
		from "domain" pd
			join domain_domain dd on dd.parent_domain_id = pd.id
		where dd.child_domain_id = d.id) as parents,
	(select jsonb_agg(jsonb_build_object('level', ds."level", 'comment', ds."comment", 'alias', s.alias, 'name', s."name") order by ds."level") 
		from domain_spell ds
			join spell s on s.id = ds.spell_id
		where ds.domain_id = d.id) as spells,
	jsonb_build_object('alias', b.alias, 'name', b."name", 'order', b."order", 'abbreviation', b.abbreviation) as book
	from "domain" d 
		join domain_type dt on dt.id = d.type_id
		join book b on b.id = d.book_id
	where d.alias = $2 and dt.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "domainType": domainType, "domainAlias": domainAlias}, "GetDomainInfo")

	var res model.DomainInfo
	err := db.GetClient().GetContext(ctx, &res, query, domainType, domainAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "domainType": domainType, "domainAlias": domainAlias}, "GetDomainInfo error", err)
		}
		return nil, err
	}

	return &res, nil
}

func (db *Postgres) GetDomainName(ctx context.Context, domainType string, alias string) (*string, error) {
	/*	span := jaeger.GetSpan(ctx, "GetDomainName")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select case when d.eng_name is null then d."name" else d."name"||' ('||d.eng_name||')' end as "name" from "domain" d join domain_type dt on dt.id = d.type_id where d.alias = $1 and dt.alias = $2`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "alias": alias, "domainType": domainType}, "GetDomainName")

	var name string
	err := db.GetClient().GetContext(ctx, &name, query, alias, domainType)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "alias": alias, "domainType": domainType}, "GetDomainName error", err)
		}
		return nil, err
	}

	return &name, nil
}
