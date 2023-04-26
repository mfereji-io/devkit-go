package httpApi

import (
	"github.com/gin-gonic/gin"
	httpApiHandlers "github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/api/http/handlers"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/config"
)

func AddAuthEndpoints(httpRoutingEngine *gin.Engine, c *config.AppConfig) *gin.Engine {

	v1RouterGroupAuth := httpRoutingEngine.Group("/api/v1/auth").Use().(*gin.RouterGroup)
	authApiHandler := httpApiHandlers.NewAuthHandler(c)

	v1RouterGroupAuth.POST("/login/u", authApiHandler.AuthenticateUserWithUsernameAndPassword())
	v1RouterGroupAuth.POST("/login/f", authApiHandler.AuthenticateUserWithFireBase())

	v1RouterGroupAuth.POST("/refresh/u", authApiHandler.RefreshUserAuthToken())
	v1RouterGroupAuth.POST("/refresh/c", authApiHandler.RefreshUserChatToken())

	return httpRoutingEngine

}
