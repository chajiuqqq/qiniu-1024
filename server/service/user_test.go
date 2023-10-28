package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"qiniu-1024-server/model"
	"qiniu-1024-server/types"
	"qiniu-1024-server/utils/xerr"
	"testing"
)

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
		Username: "testname",
		Password: "test",
	})
	assert.NoError(t, err)
	assert.Equal(t, "testname", u.Username)

	token, err := srv.UserLogin(ctx, types.UserLoginPayload{
		Username: "testname",
		Password: "test",
	})
	assert.NoError(t, err)
	fmt.Println(token)
}
