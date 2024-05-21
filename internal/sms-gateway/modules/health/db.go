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
		p.counter.Store(-1)
		status = StatusFail
	}

	return Checks{
		"ping": {
			Description:   "Successful pings since startup or last failure",
			ObservedUnit:  "",
			ObservedValue: int(p.counter.Add(1)),
			Status:        status,
		},
	}, err
}

func NewDBProvider(params DBProviderParams) *DBProvider {
	return &DBProvider{
		db: params.DB,
	}
}
