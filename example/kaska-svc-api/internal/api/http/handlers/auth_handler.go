package httpApi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/config"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/entities"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/pkg/mfereji"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/repository"
)

type (
	AuthHandler struct {
		AppConfig *config.AppConfig
	}
)

func NewAuthHandler(appConfig *config.AppConfig) *UserHandler {
	return &UserHandler{
		AppConfig: appConfig,
	}
}

func (h *UserHandler) AuthenticateUserWithUsernameAndPassword() func(c *gin.Context) {
	return func(c *gin.Context) {

		var incomingUserData entities.UserLoginDataUnP

		if err := c.ShouldBindJSON(&incomingUserData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser := &entities.UserLoginDataUnP{
			Username: incomingUserData.Username,
			Password: incomingUserData.Password,
		}

		mferejiAuth := mfereji.NewMferejiAuth(h.AppConfig.MferejiAppId, h.AppConfig.MferejiAppKey)
		authRepo := repository.NewAuthRepository(h.AppConfig.UserRepository,
			mferejiAuth,
			h.AppConfig.AppLogger,
		)

		if authenticatedUser, err := authRepo.AuthenticateWithUsernameAndPassword(newUser); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, authenticatedUser)
		}

	}
}

func (h *UserHandler) AuthenticateUserWithFireBase() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

func (h *UserHandler) RefreshUserAuthToken() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

func (h *UserHandler) RefreshUserChatToken() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
