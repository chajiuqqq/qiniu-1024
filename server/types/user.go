package types

type UserRegisterPayload struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Phone       string `bson:"phone" json:"phone"`
	AvatarUrl   string `bson:"avatar_url" json:"avatar_url"`
	Description string `bson:"description" json:"description"`
}

type UserLoginPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
