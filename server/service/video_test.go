package service

import (
	"github.com/stretchr/testify/assert"
	"qiniu-1024-server/model"
	"qiniu-1024-server/types"
	"testing"
)

func TestService_OssVideoCallback(t *testing.T) {
	err := srv.Mongo.Collection(model.Video{}.Collection()).Drop(ctx)
	assert.NoError(t, err)

	vid := int64(10000001)
	v, err := srv.SaveVideo(ctx, 1, types.MainVideoSubmit{
		VideoID:    vid,
		CategoryID: 1,
		Desc:       "test...",
	})
	assert.NoError(t, err)
	assert.Equal(t, model.VideoStatusUploading, v.Status)

	err = srv.VideoStatusUpdate(ctx, vid)
	assert.NoError(t, err)
	v, err = srv.VideoDetail(ctx, vid)
	assert.NoError(t, err)
	assert.Equal(t, model.VideoStatusOnShow, v.Status)
}

func TestService_VideoCoverStatusUpdate(t *testing.T) {
	err := srv.Mongo.Collection(model.Video{}.Collection()).Drop(ctx)
	assert.NoError(t, err)

	vid := int64(10000001)
	v, err := srv.SaveVideo(ctx, 1, types.MainVideoSubmit{
		VideoID:    vid,
		CategoryID: 1,
		Desc:       "test...",
	})
	assert.NoError(t, err)
	assert.Equal(t, model.CoverStatusUploading, v.CoverStatus)

	err = srv.VideoCoverStatusUpdate(ctx, vid, model.CoverStatusSuccess)
	assert.NoError(t, err)
	v, err = srv.VideoDetail(ctx, vid)
	assert.NoError(t, err)
	assert.Equal(t, model.CoverStatusSuccess, v.CoverStatus)
}
