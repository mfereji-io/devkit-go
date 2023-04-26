package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/repository"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/spin"

	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/config"
	pgdb "github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/db/pgsql"
	applogger "github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/pkg/logger"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	appConfig := config.InitAppConfig()

	AppLogger := applogger.InitAppLogger(appConfig)

	defer AppLogger.Sync()
	appConfig.AppLogger = AppLogger

	//################################################################

	pgDbConfig := &pgdb.PgDbConfig{
		PgHost:     appConfig.PgDbHost,
		PgPort:     appConfig.PgDbPort,
		PgDb:       appConfig.PgDb,
		PgUsername: appConfig.PgDbUsername,
		PgPassword: appConfig.PgDbPassword,
		Logger:     appConfig.AppLogger,
	}

	if pgDBSession, err := pgdb.InitPgDbSession(pgDbConfig); err != nil {
		AppLogger.Sugar().Panicf("Could not connect to pg db session : %s", err.Error())
	} else {

		appConfig.UserRepository = repository.NewUserRepository(pgDBSession,
			appConfig.UserDBTable,
			appConfig.AppLogger,
		)

	}

	//################################################################

	healthCheckerReady := health.NewChecker(health.WithCacheDuration(1*time.Second),
		health.WithTimeout(10*time.Second))

	healthCheckerReady.Start()

	healthCheckerLive := health.NewChecker()

	appConfig.HealthCheckerLive = healthCheckerLive
	appConfig.HealthCheckerReady = healthCheckerReady

	s := &spin.Servers{
		AppConfig: appConfig,
	}

	//################################
	//Handle Signals
	ec := make(chan error, 2)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt, syscall.SIGTERM,
	)

	go func() {
		ec <- s.Run(context.Background())
	}()

	//################################
	var err error

	select {

	case err = <-ec:
	case <-ctx.Done():

		haltCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		s.Shutdown(haltCtx)
		stop()

		err = <-ec
	}

	if err != nil {
		AppLogger.Sugar().Infof("kaskazini service %s shutdown : %s", appConfig.SvcName, err)
	}

}
