package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		Name:          payload.Name,
		Username:      payload.Username,
		Password:      payload.Password,
		Phone:         payload.Phone,
		AvatarUrl:     payload.AvatarUrl,
		Description:   payload.Description,
		GithubAccount: "",
		WechatAccount: "",
		UserLikes:     []model.UserLikeItem{},
		Follows:       []model.FollowItem{},
		Followers:     []model.FollowItem{},
		Likes:         []model.LikeItem{},
		Collections:   []model.CollectionItem{},
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

func (s *Service) Users(ctx context.Context, ids []int64) ([]model.User, error) {
	col := s.Mongo.Collection(model.User{}.Collection())
	var users []model.User
	cur, err := col.Find(ctx, bson.M{"id": bson.M{"$in": ids}})
	if err != nil {
		return nil, fmt.Errorf("get users db failed: %w", err)
	}
	err = cur.All(ctx, &users)
	if err != nil {
		return nil, fmt.Errorf("get users db cursor failed: %w", err)
	}
	return users, nil
}
func (s *Service) UsersMap(ctx context.Context, ids []int64) (map[int64]model.User, error) {
	users, err := s.Users(ctx, ids)
	if err != nil {
		return nil, err
	}
	var userMap = make(map[int64]model.User)
	for _, user := range users {
		userMap[user.ID] = user
	}
	return userMap, nil
}
func (s *Service) UserPublishedCnt(ctx context.Context, uid int64) (int, error) {
	cnt, err := s.Mongo.Collection(model.Video{}.Collection()).CountDocuments(ctx,
		bson.M{"user_id": uid, "status": model.VideoStatusOnShow})
	if err != nil {
		return 0, fmt.Errorf("get user[%d] published cnt failed: %w", uid, err)
	}
	return int(cnt), nil
}
func (s *Service) UserDetailDB(ctx context.Context, uid int64) (*model.User, error) {
	col := s.Mongo.Collection(model.User{}.Collection())
	var user model.User
	err := col.FindOne(ctx, bson.M{"id": uid}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, xerr.New(404, "NotFound", "user not found")
		}
		return nil, fmt.Errorf("get user[%d] db failed: %w", uid, err)
	}
	return &user, nil
}
func (s *Service) UserDetail(ctx context.Context, uid int64) (*types.UserDetail, error) {
	user, err := s.UserDetailDB(ctx, uid)
	if err != nil {
		return nil, err
	}
	user.Password = ""

	// videos
	videos, err := s.MainVideosDB(ctx, types.VideoQuery{
		UserID: uid,
	})
	if err != nil {
		return nil, err
	}

	var res = &types.UserDetail{
		User:   *user,
		Videos: videos,
	}
	return res, nil
}
func (s *Service) PostUserLike(ctx context.Context, from, to int64) (*model.User, error) {
	var col = s.Mongo.Collection(model.User{}.Collection())
	_, err := s.UserDetailDB(ctx, to)
	if err != nil {
		return nil, err
	}

	updateTo := bson.D{{"$push", bson.D{{"user_likes", model.UserLikeItem{
		UserID:    from,
		CreatedAt: time.Now(),
	}}}}}

	_, err = col.UpdateOne(ctx, bson.M{"id": to}, updateTo)
	if err != nil {
		return nil, fmt.Errorf("post user like db failed,from[%d] to[%d]: %w", from, to, err)
	}
	return s.UserDetailDB(ctx, from)
}
func (s *Service) PostUserFollow(ctx context.Context, from, to int64) (*model.User, error) {
	var col = s.Mongo.Collection(model.User{}.Collection())
	_, err := s.UserDetailDB(ctx, to)
	if err != nil {
		return nil, err
	}

	updateFrom := bson.D{{"$push", bson.D{{"follows", model.FollowItem{
		UserID:    to,
		CreatedAt: time.Now(),
	}}}}}
	_, err = col.UpdateOne(ctx, bson.M{"id": from}, updateFrom)
	if err != nil {
		return nil, fmt.Errorf("post user follow db failed,from[%d] to[%d]: %w", from, to, err)
	}

	updateTo := bson.D{{"$push", bson.D{{"followers", model.FollowItem{
		UserID:    from,
		CreatedAt: time.Now(),
	}}}}}

	_, err = col.UpdateOne(ctx, bson.M{"id": to}, updateTo)
	if err != nil {
		return nil, fmt.Errorf("post user follow db failed,from[%d] to[%d]: %w", from, to, err)
	}
	return s.UserDetailDB(ctx, from)
}
func (s *Service) PostUserLikeCancel(ctx context.Context, from, to int64) (*model.User, error) {
	var col = s.Mongo.Collection(model.User{}.Collection())
	_, err := s.UserDetailDB(ctx, to)
	if err != nil {
		return nil, err
	}

	updateTo := bson.D{{"$pull", bson.D{{"user_likes", bson.M{"user_id": from}}}}}
	_, err = col.UpdateOne(ctx, bson.M{"id": to}, updateTo)
	if err != nil {
		return nil, fmt.Errorf("cancel user like db failed,from[%d] to[%d]: %w", from, to, err)
	}
	return s.UserDetailDB(ctx, from)
}
func (s *Service) PostUserFollowCancel(ctx context.Context, from, to int64) (*model.User, error) {
	var col = s.Mongo.Collection(model.User{}.Collection())
	_, err := s.UserDetailDB(ctx, to)
	if err != nil {
		return nil, err
	}

	updateFrom := bson.D{{"$pull", bson.D{{"follows", bson.M{"user_id": to}}}}}
	_, err = col.UpdateOne(ctx, bson.M{"id": from}, updateFrom)
	if err != nil {
		return nil, fmt.Errorf("cancel user follow db failed,from[%d] to[%d]: %w", from, to, err)
	}

	updateTo := bson.D{{"$pull", bson.D{{"followers", bson.M{"user_id": from}}}}}

	_, err = col.UpdateOne(ctx, bson.M{"id": to}, updateTo)
	if err != nil {
		return nil, fmt.Errorf("cancel user follow db failed,from[%d] to[%d]: %w", from, to, err)
	}
	return s.UserDetailDB(ctx, from)
}
func (s *Service) PostUserAction(ctx context.Context, from, to int64, action string) (*model.User, error) {
	switch action {
	case types.UserActionLike:
		return s.PostUserLike(ctx, from, to)
	case types.UserActionFollow:
		return s.PostUserFollow(ctx, from, to)
	case types.UserActionLikeCancel:
		return s.PostUserLikeCancel(ctx, from, to)
	case types.UserActionFollowCancel:
		return s.PostUserFollowCancel(ctx, from, to)
	default:
		return nil, xerr.New(400, "InvalidAction", "invalid action")
	}
}

func genUserID(seq int64) int64 {
	return 1e5 + seq
}
