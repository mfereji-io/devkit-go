package httpApi

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/config"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/pkg/middleware"

	httpApiEndpoints "github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/api/http/endpoints"
)

func InitHTTPRoutingEngine(c *config.AppConfig) *gin.Engine {

	globalHTTPRoutingEngine := gin.New()

	if c.AppEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	myLogger := c.AppLogger

	//# Other middleware
	globalHTTPRoutingEngine.Use(ginzap.Ginzap(myLogger, time.RFC3339, true))
	//Gzip
	globalHTTPRoutingEngine.Use(middleware.GzipCompression())
	//TLS
	globalHTTPRoutingEngine.Use(middleware.Secure())
	//CORS/CSRF
	globalHTTPRoutingEngine.Use(middleware.CORSMiddleware())
	//Request Id
	globalHTTPRoutingEngine.Use(middleware.SetRequestId())

	globalHTTPRoutingEngine.Use(ginzap.RecoveryWithZap(myLogger, true))
	globalHTTPRoutingEngine.Use(gin.Recovery())

	return globalHTTPRoutingEngine

}

func AddHttpEndpoints(httpRoutingEngine *gin.Engine, c *config.AppConfig) *gin.Engine {

	httpApiEndpoints.AddHealthEndpoints(httpRoutingEngine, c)
	httpApiEndpoints.AddAuthEndpoints(httpRoutingEngine, c)
	httpApiEndpoints.AddUserEndpoints(httpRoutingEngine, c)

	return httpRoutingEngine
}
