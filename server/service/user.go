package service

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"qiniu-1024-server/model"
	"qiniu-1024-server/types"
	"qiniu-1024-server/utils/shared"
	"qiniu-1024-server/utils/xerr"
	"time"
)

func (s *Service) UserRegister(ctx context.Context, payload types.UserRegisterPayload) (*model.User, error) {
	col := s.Mongo.Collection(model.User{}.Collection())
	cnt, err := col.CountDocuments(ctx, bson.M{"username": payload.Username})
	if err != nil {
		return nil, fmt.Errorf("register user db failed: %w", err)
	}
	if cnt != 0 {
		return nil, xerr.New(400, "InvalidUsername", "username already exists")
	}

	seq, err := s.GetUserSeq(ctx)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		ID:            genUserID(seq),
		Username:      payload.Username,
		Password:      payload.Password,
		Phone:         payload.Phone,
		AvatarUrl:     payload.AvatarUrl,
		Description:   payload.Description,
		GithubAccount: "",
		WechatAccount: "",
		UserLikes:     nil,
		Follows:       nil,
		Followers:     nil,
		Likes:         nil,
		Collections:   nil,
	}
	res, err := col.InsertOne(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("register user db failed: %w", err)
	}
	var returnUser model.User
	err = col.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&returnUser)
	if err != nil {
		return nil, fmt.Errorf("register user db failed: %w", err)
	}
	returnUser.Password = ""
	return &returnUser, nil
}
func (s *Service) UserLogin(ctx context.Context, payload types.UserLoginPayload) (string, *model.User, error) {
	var user = new(model.User)
	err := s.Mongo.Collection(model.User{}.Collection()).FindOne(ctx, bson.M{"username": payload.Username}).Decode(user)
	if err != nil {
		return "", nil, fmt.Errorf("login user db failed: %w", err)
	}
	if user.Password != payload.Password {
		return "", nil, fmt.Errorf("password is wrong")
	}

	claims := &shared.JWTCustomClaims{
		UserID:   user.ID,
		Audience: shared.JWTAudienceUser,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(s.Conf.JWTSecret))
	if err != nil {
		return "", nil, fmt.Errorf("sign user jwt failed: %w", err)
	}
	user.Password = ""
	return t, user, nil
}

func genUserID(seq int64) int64 {
	return 1e5 + seq
}
