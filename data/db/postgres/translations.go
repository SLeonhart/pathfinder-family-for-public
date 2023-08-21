package postgres

import (
	"context"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetTranslations(ctx context.Context, alias string) ([]model.Translation, error) {
	/*	span := jaeger.GetSpan(ctx, "GetTranslations")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select t."group", jsonb_agg(jsonb_build_object('forLevel', t.for_level, 'order', t."order", 'alias', t.alias, 'name', t."name", 'engName', t.name_eng, 'helpers', (
		select jsonb_agg(jsonb_build_object('name', h."name", 'alias', h.alias) order by h."name")
			from helper h  
				join helper_translation ht on ht.helper_id = h.id
			where t.id = ht.translation_id 
	)) order by t."order") as items
		from "translation" t
			join translation_type tt on tt.id = t.type_id
		where tt.alias = $1
		group by t."group"
		order by t."group"`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "alias": alias}, "GetTranslations")

	translations := make([]model.Translation, 0)
	err := db.GetClient().SelectContext(ctx, &translations, query, alias)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "alias": alias}, "GetTranslations error", err)
		return nil, err
	}

	return translations, nil
}
