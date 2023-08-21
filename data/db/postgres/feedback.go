package postgres

import (
	"context"

	"pathfinder-family/infrastructure/logger"
)

func (db *Postgres) SendFeedback(ctx context.Context, theme string, email *string, message string) error {
	/*	span := jaeger.GetSpan(ctx, "SendFeedback")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	query := `insert into feedback (dt,email,is_readed,message,theme) values (now(),$1,false,$2,$3)`

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "theme": theme, "email": email, "message": message}, "SendFeedback")

	_, err := db.GetClient().ExecContext(ctx, query, theme, email, message)
	if err != nil {
		logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "theme": theme, "email": email, "message": message}, "SendFeedback error", err)
		return err
	}

	return nil
}
