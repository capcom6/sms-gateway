package health

import (
	"context"
	"database/sql"
	"sync/atomic"

	"go.uber.org/fx"
)

type DBProviderParams struct {
	fx.In

	DB *sql.DB
}

type DBProvider struct {
	db *sql.DB

	counter atomic.Int32
}

func (p *DBProvider) Name() string {
	return "db"
}

func (p *DBProvider) HealthCheck(ctx context.Context) (Checks, error) {
	status := StatusPass

	err := p.db.PingContext(ctx)
	if err != nil {
		p.counter.Add(1)
		status = StatusFail
	} else {
		p.counter.Store(0)
	}

	return Checks{
		"ping": {
			Description:   "Failed sequential pings count",
			ObservedUnit:  "",
			ObservedValue: int(p.counter.Load()),
			Status:        status,
		},
	}, err
}

func NewDBProvider(params DBProviderParams) *DBProvider {
	return &DBProvider{
		db: params.DB,
	}
}
