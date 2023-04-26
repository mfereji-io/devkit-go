package pgdb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type (
	PgDbConfig struct {
		PgHost     string
		PgPort     string
		PgDb       string
		PgUsername string
		PgPassword string
		Logger     *zap.Logger
	}

	AppDBSession struct {
		Pool *pgxpool.Pool
	}
)

func (p *AppDBSession) Close() {
	p.Pool.Close()
}

func InitPgDbSession(PgDbSessionConfig *PgDbConfig) (*AppDBSession, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", PgDbSessionConfig.PgUsername, PgDbSessionConfig.PgPassword, PgDbSessionConfig.PgHost, PgDbSessionConfig.PgPort, PgDbSessionConfig.PgDb)
	pool, err := pgxpool.Connect(context.Background(), connString)

	if err != nil {

		//PgDbSessionConfig.Logger.Sugar().Errorf("Could not connect to pgdb server %s with error:", PgDbSessionConfig.PgHost, err.Error())
		return &AppDBSession{}, err

	} else {

		//PgDbSessionConfig.Logger.Sugar().Infof("Connected to pgdb server %s", PgDbSessionConfig.PgHost)
		return &AppDBSession{Pool: pool}, nil

	}

}
