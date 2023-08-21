package postgres

import (
	"context"
	"database/sql"
	"errors"

	"pathfinder-family/model"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) GetAbilities(ctx context.Context) ([]model.NameAlias, error) {
	/*	span := jaeger.GetSpan(ctx, "GetAbilities")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select a."name", a.alias from ability a order by a.id`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetAbilities")

	classes := make([]model.NameAlias, 0)
	err := db.GetClient().SelectContext(ctx, &classes, query)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query}, "GetAbilities error", err)
		return nil, err
	}

	return classes, nil
}

func (db *Postgres) GetAbilityInfo(ctx context.Context, abilityAlias string) (*model.AbilityInfo, error) {
	/*	span := jaeger.GetSpan(ctx, "GetAbilityInfo")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `select a."name", a.alias, a.short_name, a.description, a.eng_name from ability a where a.alias = $1`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "abilityAlias": abilityAlias}, "GetAbilityInfo")

	var classInfo model.AbilityInfo
	err := db.GetClient().GetContext(ctx, &classInfo, query, abilityAlias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "abilityAlias": abilityAlias}, "GetAbilityInfo error", err)
		}
		return nil, err
	}

	return &classInfo, nil
}
