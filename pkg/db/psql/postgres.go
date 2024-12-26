package psql

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/dlc-01/BackendCrypto/internal/conf"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migration/*.sql
var EmbedMigrations embed.FS

func NewSQLClient(cfg conf.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("projectError while opening conection: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("projectError while ping db: %s", err)
	}

	goose.SetBaseFS(EmbedMigrations)

	err = goose.SetDialect("postgres")
	if err != nil {
		return nil, fmt.Errorf("projectError while seting dialect: %s", err)
	}

	err = goose.Up(db, "migration")
	if err != nil {
		return nil, fmt.Errorf("projectError migarition: %s", err)
	}

	return db, nil
}
