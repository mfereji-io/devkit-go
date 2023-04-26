package httpApi

import (
	"github.com/gin-gonic/gin"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/config"
)

type (
	UserHandler struct {
		AppConfig *config.AppConfig
	}
)

func NewUserHandler(appConfig *config.AppConfig) *UserHandler {

	return &UserHandler{
		AppConfig: appConfig,
	}

}

func (h *UserHandler) CreateUser() func(c *gin.Context) {

	return func(c *gin.Context) {

	}
}

func (h *UserHandler) GetUserById() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

func (h *UserHandler) UpdateUser() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

func (h *UserHandler) DeleteUser() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
