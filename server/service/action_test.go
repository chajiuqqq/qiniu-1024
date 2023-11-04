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
func TestDefaultActionService_LikeVideo(t *testing.T) {
	actionInit(ctx, t)
	// like 1,2
	err := srv.ActionService.LikeVideo(ctx, 100001, 1)
	assert.NoError(t, err)
	err = srv.ActionService.LikeVideo(ctx, 100001, 2)
	assert.NoError(t, err)
	v, err := srv.VideoDetailDB(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), v.LikesCount)
	v, err = srv.VideoDetailDB(ctx, 2)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), v.LikesCount)

	u, err := srv.UserDetailDB(ctx, 100001)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(u.Likes))

	// unlike 1
	err = srv.ActionService.UnLikeVideo(ctx, 100001, 1)
	assert.NoError(t, err)
	v, err = srv.VideoDetailDB(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), v.LikesCount)

	u, err = srv.UserDetailDB(ctx, 100001)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(u.Likes))
	assert.Equal(t, int64(2), u.Likes[0].VideoID)

}
func TestDefaultActionService_CollectVideo(t *testing.T) {
	actionInit(ctx, t)
	// collect 1,2
	err := srv.ActionService.CollectVideo(ctx, 100001, 1)
	assert.NoError(t, err)
	err = srv.ActionService.CollectVideo(ctx, 100001, 2)
	assert.NoError(t, err)
	v, err := srv.VideoDetailDB(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), v.CollectCount)
	v, err = srv.VideoDetailDB(ctx, 2)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), v.CollectCount)

	u, err := srv.UserDetailDB(ctx, 100001)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(u.Collections))

	// uncollect 1
	err = srv.ActionService.UnCollectVideo(ctx, 100001, 1)
	assert.NoError(t, err)
	v, err = srv.VideoDetailDB(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), v.CollectCount)

	u, err = srv.UserDetailDB(ctx, 100001)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(u.Collections))
	assert.Equal(t, int64(2), u.Collections[0].VideoID)
}
func TestDefaultActionService_CommentVideo(t *testing.T) {
	actionInit(ctx, t)
	// collect 1,2
	err := srv.ActionService.CommentVideo(ctx, 100001, 1, "test comment")
	assert.NoError(t, err)
	v, err := srv.VideoDetailDB(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(v.Comments))
	assert.Equal(t, "test comment", v.Comments[0].Content)
}
