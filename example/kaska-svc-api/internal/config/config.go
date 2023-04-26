package config

import (
	"log"

	"github.com/alexliesenfeld/health"
	"github.com/joho/godotenv"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/pkg/envo"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/repository"
	"go.uber.org/zap"
)

type (
	AppConfig struct {
		AppLogger *zap.Logger

		HealthCheckerLive  health.Checker
		HealthCheckerReady health.Checker

		SvcName string
		AppEnv  string
		GinMode string

		AppListenHostname string
		AppListenPortHttp string
		AppListenPortGrpc string

		HealthEndpointPrefix string
		HealthEndpointLive   string
		HealthEndpointReady  string

		PgDbHost     string
		PgDbPort     string
		PgDb         string
		PgDbUsername string
		PgDbPassword string

		UserRepository *repository.UserRepository
		UserDBTable    string

		MferejiAppId  string
		MferejiAppKey string
	}
)

func InitAppConfig() *AppConfig {

	envConfigFile := ".env"
	err := godotenv.Load(envConfigFile)

	if err != nil {
		log.Printf("could not load .env file %s", err)
	} else {
		log.Printf("loaded .env file at %s", envConfigFile)
	}

	//########################

	svcNameDefault := "kaska-svc-api"
	appEnvDefault := "dev"
	ginModeDefault := "debug"

	//appListenHostnameDefault := "127.0.0.1"
	appListenHostnameDefault := "0.0.0.0"
	appListenPortHttpDefault := "8081"
	appListenPortGrpcDefault := "8082"

	healthEndpointPrefixDefault := "/"
	healthEndpointLiveDefault := "live"
	healthEndpointReadyDefault := "ready"

	PgDbHostDefault := "127.0.0.1"
	PgDbPortDefault := "5432"
	PgDbDefault := "kaska_db"
	PgDbUsernameDefault := ""
	PgDbPasswordDefault := ""

	//#############################
	UserDBTableDefault := "kaska_users"

	//############################

	svcName := envo.EnvString("SVC_NAME", svcNameDefault)
	appEnv := envo.EnvString("APP_ENV", appEnvDefault)
	ginMode := envo.EnvString("GIN_MODE", ginModeDefault)

	appListenHostname := envo.EnvString("APP_LISTEN_HOSTNAME", appListenHostnameDefault)
	appListenPortHttp := envo.EnvString("APP_LISTEN_PORT_HTTP", appListenPortHttpDefault)
	appListenPortGrpc := envo.EnvString("APP_LISTEN_PORT_GRPC", appListenPortGrpcDefault)

	healthEndpointPrefix := envo.EnvString("HEALTH_ENDPOINT_PREFIX", healthEndpointPrefixDefault)
	healthEndpointLive := envo.EnvString("HEALTH_ENDPOINT_LIVE", healthEndpointLiveDefault)
	healthEndpointReady := envo.EnvString("HEALTH_ENDPOINT_READY", healthEndpointReadyDefault)

	pgDbHost := envo.EnvString("PG_DB_HOST", PgDbHostDefault)
	pgDbPort := envo.EnvString("PG_DB_PORT", PgDbPortDefault)
	pgDb := envo.EnvString("PG_DB", PgDbDefault)
	pgDbUsername := envo.EnvString("PG_DB_USERNAME", PgDbUsernameDefault)
	pgDbPassword := envo.EnvString("PG_DB_PASSWORD", PgDbPasswordDefault)

	UserDBTable := envo.EnvString("USERS_DB_TABLE", UserDBTableDefault)

	mferejiAppId := envo.EnvString("MFEREJI_APP_ID", "")
	mferejiAppKey := envo.EnvString("MFEREJI_APP_KEY", "")

	//#############################

	return &AppConfig{

		SvcName: svcName,
		AppEnv:  appEnv,
		GinMode: ginMode,

		AppListenHostname: appListenHostname,
		AppListenPortHttp: appListenPortHttp,
		AppListenPortGrpc: appListenPortGrpc,

		HealthEndpointPrefix: healthEndpointPrefix,
		HealthEndpointLive:   healthEndpointLive,
		HealthEndpointReady:  healthEndpointReady,

		PgDbHost:     pgDbHost,
		PgDbPort:     pgDbPort,
		PgDb:         pgDb,
		PgDbUsername: pgDbUsername,
		PgDbPassword: pgDbPassword,

		UserDBTable: UserDBTable,

		MferejiAppId:  mferejiAppId,
		MferejiAppKey: mferejiAppKey,
	}

}
