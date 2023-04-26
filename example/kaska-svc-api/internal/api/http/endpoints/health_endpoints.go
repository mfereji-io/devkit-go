package httpApi

import (
	"github.com/alexliesenfeld/health"
	"github.com/gin-gonic/gin"

	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/config"
)

func AddHealthEndpoints(httpRoutingEngine *gin.Engine, c *config.AppConfig) *gin.Engine {

	httpRoutingEngine.GET(c.HealthEndpointPrefix+c.HealthEndpointLive, gin.WrapH(health.NewHandler(c.HealthCheckerLive)))
	httpRoutingEngine.GET(c.HealthEndpointPrefix+c.HealthEndpointReady, gin.WrapH(health.NewHandler(c.HealthCheckerReady)))

	return httpRoutingEngine

}
