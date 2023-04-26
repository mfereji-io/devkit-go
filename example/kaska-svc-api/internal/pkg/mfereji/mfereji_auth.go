package mfereji

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type (
	UserInfoPayload struct {
		Anonymous bool

		UserUUId string
		Username string
		Email    string

		Issuer   string
		Channels []string
	}

	MferejiJwtChatTokenClaims struct {
		UserId   string `json:"session_user_id"`
		Username string `json:"session_username"`
		Email    string `json:"session_email"`

		ConnExpireAt int64    `json:"expire_at"`
		Info         string   `json:"info"`
		Meta         string   `json:"meta"`
		Channels     []string `json:"channels"`

		jwt.StandardClaims
	}

	MferejiAuth struct {
		MferejiAppId         string
		MferejiAppSigningKey string
	}
)

func NewMferejiAuth(mferejiAppId string, mferejiAppSigningKey string) *MferejiAuth {
	return &MferejiAuth{
		MferejiAppId:         mferejiAppId,
		MferejiAppSigningKey: mferejiAppSigningKey,
	}
}

func (ma *MferejiAuth) GenerateMferejiJwtChatToken(userInfoPayload *UserInfoPayload) (string, error) {

	if ma.MferejiAppId == "" || ma.MferejiAppSigningKey == "" {
		return "", errors.New("invalid mfereji appId or app signing key")
	}

	expirationTime := time.Now().Add(1 * 24 * 60 * time.Minute) //1d

	mferejiJwtChatTokenClaims := &MferejiJwtChatTokenClaims{
		UserId:   userInfoPayload.UserUUId,
		Username: userInfoPayload.Username,
		Email:    userInfoPayload.Email,
		Channels: userInfoPayload.Channels,
		StandardClaims: jwt.StandardClaims{
			Audience:  "mfereji.io",
			ExpiresAt: expirationTime.Unix(),
			Id:        uuid.New().String(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    userInfoPayload.Issuer,
			//NotBefore:
			Subject: userInfoPayload.UserUUId,
		},
	}
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, mferejiJwtChatTokenClaims)
	if signedToken, err := unsignedToken.SignedString([]byte(ma.MferejiAppSigningKey)); err != nil {

		//log.Printf(" could not gen signedToken %s", err.Error())
		return "", err

	} else {
		//log.Printf(" successfully signed token %s", signedToken)
		return signedToken, nil

	}

}

func (ma *MferejiAuth) VerifyMferejiJwtChatToken(strToken string) (*MferejiJwtChatTokenClaims, error) {

	if ma.MferejiAppId == "" || ma.MferejiAppSigningKey == "" {
		return nil, errors.New("invalid mfereji appId or app signing key")
	}

	claims := &MferejiJwtChatTokenClaims{}

	token, err := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ma.MferejiAppSigningKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, fmt.Errorf("invalid token signature")
		}
	}

	if !token.Valid {
		return claims, fmt.Errorf("invalid token")
	}

	return claims, nil

}
