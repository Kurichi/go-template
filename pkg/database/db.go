package database

import (
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	*bun.DB
}

type Config struct {
	Host     string
	Port     uint
	User     string
	Password string
	DBName   string
}

func New(cfg *Config) (*DB, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(addr),
		pgdriver.WithUser(cfg.User),
		pgdriver.WithPassword(cfg.Password),
		pgdriver.WithDatabase(cfg.DBName),
		pgdriver.WithInsecure(true),
	)

	sqldb := sql.OpenDB(pgconn)
	db := bun.NewDB(sqldb, pgdialect.New())

	// Check if the connection is working.
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
