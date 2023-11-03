package types

import "qiniu-1024-server/model"

const (
	UserActionLike         = "Like"
	UserActionFollow       = "Follow"
	UserActionLikeCancel   = "LikeCancel"
	UserActionFollowCancel = "FollowCancel"
)

type UserRegisterPayload struct {
	Name        string `json:"name" validate:"required"`
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

type UserDetail struct {
	model.User
	Videos []model.Video `json:"videos"`
}
