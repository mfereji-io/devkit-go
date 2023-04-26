package httpApi

import (
	"github.com/gin-gonic/gin"
	httpApiHandlers "github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/api/http/handlers"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/config"
)

func AddUserEndpoints(httpRoutingEngine *gin.Engine, c *config.AppConfig) *gin.Engine {

	v1RouterGroupUsers := httpRoutingEngine.Group("/api/v1/users").Use().(*gin.RouterGroup)

	userApiHandler := httpApiHandlers.NewUserHandler(c)

	v1RouterGroupUsers.POST("", userApiHandler.CreateUser())
	v1RouterGroupUsers.GET("/id/:userid", userApiHandler.GetUserById())

	return httpRoutingEngine

}
