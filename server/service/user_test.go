package service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"qiniu-1024-server/model"
	"qiniu-1024-server/types"
	"qiniu-1024-server/utils/xerr"
	"testing"
)

const (
	testUsername1 = "testname1"
	testUsername2 = "testname2"
)

func initialUser(ctx context.Context, t *testing.T) {
	err := srv.Mongo.Collection(model.User{}.Collection()).Drop(ctx)
	assert.NoError(t, err)
	err = srv.Mongo.Collection(model.Counter{}.Collection()).Drop(ctx)
	assert.NoError(t, err)

	_, err = srv.UserRegister(ctx, types.UserRegisterPayload{
		Name:        "myName1",
		Username:    testUsername1,
		Password:    "test",
		Description: "Love",
		Phone:       "1777777777777",
		AvatarUrl:   "https://test.com/",
	})
	assert.NoError(t, err)
	_, err = srv.UserRegister(ctx, types.UserRegisterPayload{
		Name:        "myName2",
		Username:    testUsername2,
		Password:    "test",
		Description: "Love",
		Phone:       "1777777777777",
		AvatarUrl:   "https://test.com/",
	})
	assert.NoError(t, err)
}
func TestService_UserRegister(t *testing.T) {
	err := srv.Mongo.Collection(model.User{}.Collection()).Drop(ctx)
	assert.NoError(t, err)

	err = srv.Mongo.Collection(model.Counter{}.Collection()).Drop(ctx)
	assert.NoError(t, err)

	u, err := srv.UserRegister(ctx, types.UserRegisterPayload{
		Username: "testname",
		Password: "test",
	})
	assert.NoError(t, err)
	assert.Equal(t, "testname", u.Username)
	assert.Equal(t, int64(100001), u.ID)

	u, err = srv.UserRegister(ctx, types.UserRegisterPayload{
		Username: "testname",
		Password: "test",
	})
	assert.NotNil(t, err)
	assert.Equal(t, 400, err.(*xerr.XError).StatusCode)

	u, err = srv.UserRegister(ctx, types.UserRegisterPayload{
		Username: "testname2",
		Password: "test",
	})
	assert.NoError(t, err)
	assert.Equal(t, "testname2", u.Username)
	assert.Equal(t, int64(100002), u.ID)
}

func TestService_UserLogin(t *testing.T) {
	err := srv.Mongo.Collection(model.User{}.Collection()).Drop(ctx)
	assert.NoError(t, err)

	u, err := srv.UserRegister(ctx, types.UserRegisterPayload{
		Name:        "myName",
		Username:    "testname",
		Password:    "test",
		Description: "Love",
		Phone:       "1777777777777",
		AvatarUrl:   "https://test.com/",
	})
	assert.NoError(t, err)
	assert.Equal(t, "testname", u.Username)

	token, u, err := srv.UserLogin(ctx, types.UserLoginPayload{
		Username: "testname",
		Password: "test",
	})
	assert.NoError(t, err)
	assert.Equal(t, "myName", u.Name)
	assert.Equal(t, "testname", u.Username)
	assert.Equal(t, "Love", u.Description)
	assert.Equal(t, "1777777777777", u.Phone)
	assert.Equal(t, "https://test.com/", u.AvatarUrl)

	fmt.Println(token)
}

func TestService_UsersMap(t *testing.T) {
	initialUser(ctx, t)
	mp, err := srv.UsersMap(ctx, []int64{100001, 100002})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(mp))
	assert.Equal(t, int64(100001), mp[100001].ID)
	assert.Equal(t, int64(100002), mp[100002].ID)
	assert.Equal(t, testUsername1, mp[100001].Username)
	assert.Equal(t, testUsername2, mp[100002].Username)
}
func TestService_PostUserAction(t *testing.T) {
	initialUser(ctx, t)
	// follow
	from, err := srv.PostUserAction(ctx, 100001, 100002, types.UserActionFollow)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(from.Follows))
	assert.Equal(t, int64(100002), from.Follows[0].UserID)

	to, err := srv.UserDetailDB(ctx, 100002)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(to.Followers))
	assert.Equal(t, int64(100001), to.Followers[0].UserID)

	// cancel follow
	from, err = srv.PostUserAction(ctx, 100001, 100002, types.UserActionFollowCancel)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(from.Follows))

	to, err = srv.UserDetailDB(ctx, 100002)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(to.Followers))

	// like
	_, err = srv.PostUserAction(ctx, 100001, 100002, types.UserActionLike)
	assert.NoError(t, err)

	to, err = srv.UserDetailDB(ctx, 100002)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(to.UserLikes))
	assert.Equal(t, int64(100001), to.UserLikes[0].UserID)

	// cancel like
	_, err = srv.PostUserAction(ctx, 100001, 100002, types.UserActionLikeCancel)
	assert.NoError(t, err)

	to, err = srv.UserDetailDB(ctx, 100002)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(to.UserLikes))
}
