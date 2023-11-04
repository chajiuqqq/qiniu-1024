package service

import (
	"github.com/stretchr/testify/assert"
	"qiniu-1024-server/model"
	"qiniu-1024-server/types"
	"testing"
)

func actionInit(t *testing.T) {
	//videos
	initVideos(t)
	//users
	initialUser(t)
}
func TestDefaultActionService_PlayVideo(t *testing.T) {
	actionInit(t)
	_, err := srv.ActionService.PlayVideo(ctx, 100001, 1)
	assert.NoError(t, err)
	v, err := srv.VideoDetailDB(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), v.PlayCount)
}
func TestDefaultActionService_LikeVideo(t *testing.T) {
	actionInit(t)
	// like 1,2
	_, err := srv.ActionService.LikeVideo(ctx, 100001, 1)
	assert.NoError(t, err)
	_, err = srv.ActionService.LikeVideo(ctx, 100001, 2)
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
	_, err = srv.ActionService.UnLikeVideo(ctx, 100001, 1)
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
	actionInit(t)
	// collect 1,2
	_, err := srv.ActionService.CollectVideo(ctx, 100001, 1)
	assert.NoError(t, err)
	_, err = srv.ActionService.CollectVideo(ctx, 100001, 2)
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
	_, err = srv.ActionService.UnCollectVideo(ctx, 100001, 1)
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
	actionInit(t)
	// collect 1,2
	_, err := srv.ActionService.CommentVideo(ctx, 100001, 1, "test comment")
	assert.NoError(t, err)
	v, err := srv.VideoDetailDB(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(v.Comments))
	assert.Equal(t, "test comment", v.Comments[0].Content)
}
func TestDefaultActionService_QueryAction(t *testing.T) {
	actionInit(t)
	_, err := srv.ActionService.LikeVideo(ctx, 100001, 1)
	assert.NoError(t, err)
	_, err = srv.ActionService.CollectVideo(ctx, 100001, 1)
	assert.NoError(t, err)
	vs := []types.MainVideoItem{
		{
			Video: model.Video{
				ID: 1,
			},
		},
		{
			Video: model.Video{
				ID: 2,
			},
		},
		{
			Video: model.Video{
				ID: 3,
			},
		},
	}
	res, err := srv.ActionService.QueryAction(ctx, 100001, vs)
	assert.NoError(t, err)
	assert.True(t, res[0].Liked, "video 1 should be liked")
	assert.True(t, res[0].Collected, "video 1 should be collected")
	assert.False(t, res[1].Liked)
	assert.False(t, res[1].Collected)
	assert.False(t, res[2].Liked)
	assert.False(t, res[2].Collected)
}
