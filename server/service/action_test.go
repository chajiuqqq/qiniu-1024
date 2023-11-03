package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"qiniu-1024-server/model"
	"testing"
)

func actionInit(ctx context.Context, t *testing.T) {
	err := srv.Mongo.Collection(model.Video{}.Collection()).Drop(ctx)
	assert.NoError(t, err)
	vs := []model.Video{
		{
			ID:       1,
			Comments: make([]model.Comment, 0),
		},
		{
			ID:       2,
			Comments: make([]model.Comment, 0),
		},
		{
			ID:       3,
			Comments: make([]model.Comment, 0),
		},
	}
	for _, v := range vs {
		_, err = srv.Mongo.Collection(model.Video{}.Collection()).InsertOne(ctx, v)
		if err != nil {
			panic(err)
		}
	}

	//users
	initialUser(ctx, t)
}
func TestDefaultActionService_PlayVideo(t *testing.T) {
	actionInit(ctx, t)
	err := srv.ActionService.PlayVideo(ctx, 100001, 1)
	assert.NoError(t, err)
	v, err := srv.VideoDetailDB(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), v.PlayCount)
}
