package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pathfinder-family/data/cache/cacheInterface"
	"pathfinder-family/data/cache/inmemory"
	"time"

	"pathfinder-family/config"

	"pathfinder-family/infrastructure/logger"

	"github.com/jmoiron/sqlx"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

type Postgres struct {
	dbClient *sqlx.DB
	cfg      *config.Config
	cache    cacheInterface.IInMemory
}

func NewPostgres(c *config.Config, memory *inmemory.InMemory) *Postgres {
	db := &Postgres{
		cfg:   c,
		cache: memory,
	}

	db.GetClient()
	if err := db.dbClient.Ping(); err != nil {
		logger.ErrorWithErr("", "Postgres ping error", err)
		panic(err)
	}
	db.runMigration()

	go func() {
		timer := time.NewTicker(time.Duration(db.cfg.DB.Postgres.ReconnectMsec) * time.Millisecond)
		defer timer.Stop()

		count := 5

		for range timer.C {
			if err := db.GetClient().Ping(); err != nil {
				count--
				if count == 0 {
					logger.ErrorWithErr("", "Postgres reconnect error", err)
					count = 5
				}
			}
		}
	}()

	return db
}

func (db *Postgres) GetClient() *sqlx.DB {
	if !db.isClientActive() {
		cfg, _ := pgx.ParseDSN(db.cfg.DB.Postgres.DataSource())
		cfg.RuntimeParams = map[string]string{
			"standard_conforming_strings": "on",
		}
		cfg.PreferSimpleProtocol = true
		sqlDb := stdlib.OpenDB(cfg)
		db.dbClient = sqlx.NewDb(sqlDb, "pgx")
		db.dbClient.SetMaxOpenConns(db.cfg.DB.Postgres.MaxConn)
	}
	return db.dbClient
}

func (db *Postgres) isClientActive() bool {
	if db.dbClient == nil {
		return false
	}

	// if db.dbClient.Ping() != nil {
	// 	return false
	// }

	return true
}

func (db *Postgres) Close() {
	if db.dbClient == nil {
		return
	}

	_ = db.dbClient.Close()
}

func (db *Postgres) GetNameByAlias(ctx context.Context, table string, alias string, withEng bool) (*string, error) {
	/*	span := jaeger.GetSpan(ctx, "GetNameByAlias")
		defer span.Finish()
		ctx = jaeger.SetParentSpan(ctx, span)*/

	nameField := `"name"`
	if withEng {
		nameField = `case when eng_name is null then "name" else "name"||' ('||eng_name||')' end as "name"`
	}
	query := fmt.Sprintf(`select %v from %v where alias = $1`, nameField, table)

	//logger.TraceWithFields(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "table": table, "alias": alias, "withEng": withEng}, "GetNameByAlias")

	var name string
	err := db.GetClient().GetContext(ctx, &name, query, alias)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.ErrorWithFieldsAndErr(logger.CreateRequestIDField(ctx), map[string]interface{}{"query": query, "table": table, "alias": alias, "withEng": withEng}, "GetNameByAlias error", err)
		}
		return nil, err
	}

	return &name, nil
}
