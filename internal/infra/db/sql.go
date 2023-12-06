package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	mysql "github.com/go-sql-driver/mysql"
)

func NewSQL(config Config) (*sql.DB, error) {
	dsn := makeDSN(config)

	mysql.SetLogger(log.Default())
	db, err := sql.Open(config.Dialect, dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(16)
	db.SetMaxIdleConns(16)

	return db, nil
}

func makeDSN(cfg Config) string {
	cfg = configDefault(cfg)
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&parseTime=true&loc=%s&tls=preferred",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, url.QueryEscape(cfg.Timezone),
	)
}
