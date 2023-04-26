package entities

type (
	UserLoginDataUnP struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	AuthenticatedUser struct {
		Username    string `json:"username"`
		UserUUID    string `json:"useruuid"`
		FirstName   string `json:"firstname"`
		Lastname    string `json:"lastname"`
		ChatIdToken string `json:"chattoken"`
		ChatAppId   string `json:"chatappid"`
	}
)
