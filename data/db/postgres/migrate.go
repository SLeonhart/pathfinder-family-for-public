package postgres

import (
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

func (db *Postgres) runMigration() error {
	migrations := &migrate.FileMigrationSource{
		Dir: db.cfg.DB.Postgres.MigrationDir,
	}

	migrate.SetTable(db.cfg.DB.Postgres.MigrationTable)

	_, err := migrate.Exec(db.GetClient().DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
