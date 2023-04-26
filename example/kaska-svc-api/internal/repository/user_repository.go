package repository

import (
	"errors"

	pgdb "github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/db/pgsql"
	"go.uber.org/zap"
)

//"github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/entities"

type (
	UserRepository struct {
		//AppDBSession *pgdb.AppDBSession
		UserDbTable string
		Logger      *zap.Logger
	}

	UserData struct {
		Username  string
		UserFname string
		UserLname string
		UserEmail string
		UserUUID  string
		//UserAppId    string
		UserChannels []string
	}
)

func NewUserRepository(appDBSession *pgdb.AppDBSession,
	userDbTable string,
	logger *zap.Logger,
) *UserRepository {

	return &UserRepository{
		//AppDBSession: appDBSession,
		UserDbTable: userDbTable,
		Logger:      logger,
	}
}

func (r *UserRepository) GetUserByUsername(username string) (*UserData, error) {

	if username == "janedoe" || username == "johncena" {

		if username == "janedoe" {

			UsernameFromDB := "janedoe"
			UserFnameFromDB := "Jane"
			UserLnameFromDB := "Doe"
			UserEmailFromDB := "jane@doe.org"
			UserUUIDFromDB := "92303f7f-bcbb-428f-9314-21a6e2d70616"
			//UserAppIdFromDB := "72303f7f-bcbb-428f-9314-21a6e2d7061c"
			UserChannelsFromDB := []string{"internalTeam", "cx1"}

			return &UserData{
				Username:     UsernameFromDB,
				UserFname:    UserFnameFromDB,
				UserLname:    UserLnameFromDB,
				UserEmail:    UserEmailFromDB,
				UserUUID:     UserUUIDFromDB,
				UserChannels: UserChannelsFromDB,
			}, nil

		} else if username == "johncena" {

			UsernameFromDB := "johncena"
			UserFnameFromDB := "John"
			UserLnameFromDB := "Cena"
			UserEmailFromDB := "john@wwe.org"
			UserUUIDFromDB := "12303f7f-bcbb-428f-9314-21a6e2d70619"
			UserChannelsFromDB := []string{"internalTeam", "cx2"}

			return &UserData{
				Username:     UsernameFromDB,
				UserFname:    UserFnameFromDB,
				UserLname:    UserLnameFromDB,
				UserEmail:    UserEmailFromDB,
				UserUUID:     UserUUIDFromDB,
				UserChannels: UserChannelsFromDB,
			}, nil
		}

		/*
			return &UserData{
				Username:     UsernameFromDB,
				UserFname:    UserFnameFromDB,
				UserLname:    UserLnameFromDB,
				UserEmail:    UserEmailFromDB,
				UserUUID:     UserUUIDFromDB,
				//UserAppId: UserAppIdFromDB,
				UserChannels: UserChannelsFromDB,
			}, nil
		*/

	}

	errorMessage := "invalid username"
	return &UserData{}, errors.New(errorMessage)
}

func (r *UserRepository) GetUserByUUID(userUUID string) (*UserData, error) {

	errorMessage := "invalid user id"
	return &UserData{}, errors.New(errorMessage)
}
