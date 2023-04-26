package repository

import (
	"errors"

	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/entities"
	"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/pkg/mfereji"
	"go.uber.org/zap"
)

type (
	AuthRepository struct {
		UserRepository *UserRepository
		MferejiAuth    *mfereji.MferejiAuth
		Logger         *zap.Logger
	}
)

func NewAuthRepository(userRepo *UserRepository,
	mferejiAuth *mfereji.MferejiAuth,
	logger *zap.Logger,
) *AuthRepository {

	return &AuthRepository{
		UserRepository: userRepo,
		MferejiAuth:    mferejiAuth,
		Logger:         logger,
	}
}

func (r *AuthRepository) AuthenticateWithUsernameAndPassword(userLoginData *entities.UserLoginDataUnP) (*entities.AuthenticatedUser, error) {

	//Confirm id data for this username exists
	r.Logger.Sugar().Infof("authenticating user %s ", userLoginData.Username)
	//fmt.Printf("authenticating user %s ", userLoginData.Username)

	if userLoginData.Password != "123456" {
		return nil, errors.New("supplied password is incorrect")
	}

	if userData, err := r.UserRepository.GetUserByUsername(userLoginData.Username); err == nil {

		authenticatedUserEntity := &entities.AuthenticatedUser{

			Username:  userData.Username,
			UserUUID:  userData.UserUUID,
			FirstName: userData.UserFname,
			Lastname:  userData.UserLname,
		}

		//Add chat Id Token for use by your frontend to establish Connections to Mfereji API
		mferejiAuth := r.MferejiAuth
		userChatInfoPayload := &mfereji.UserInfoPayload{

			Anonymous: false,                    // bool
			UserUUId:  userData.UserUUID,        // string
			Username:  userData.Username,        // string
			Email:     userData.UserEmail,       // string
			Issuer:    mferejiAuth.MferejiAppId, // string
			Channels:  userData.UserChannels,    // []string
		}

		if chatIdToken, err := mferejiAuth.GenerateMferejiJwtChatToken(userChatInfoPayload); err == nil {
			authenticatedUserEntity.ChatIdToken = chatIdToken
			authenticatedUserEntity.ChatAppId = mferejiAuth.MferejiAppId

		}

		//Return value to be sent to frontend App
		return authenticatedUserEntity, nil
	}

	return nil, errors.New("supplied username was not found")
}
