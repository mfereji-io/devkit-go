package grpcApi

import (
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/config"
)

type (
	KaskaApiSvcServerImpl struct {
		AppConfig *config.AppConfig
	}
)
