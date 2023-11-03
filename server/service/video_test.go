package service

import (
	"github.com/stretchr/testify/assert"
	"qiniu-1024-server/model"
	"testing"
)

func TestService_VideoStatusUpdate(t *testing.T) {
	err := srv.Mongo.Collection(model.Video{}.Collection()).Drop(ctx)
	assert.NoError(t, err)
	vid := int64(10000001)
	// presave
	v, err := srv.PreSaveVideo(ctx, 1, vid)
	assert.NoError(t, err)
	assert.Equal(t, model.VideoStatusUploading, v.Status)

	err = srv.VideoStatusUpdate(ctx, vid, model.VideoStatusNew)
	assert.NoError(t, err)
	v, err = srv.VideoDetailDB(ctx, vid)
	assert.NoError(t, err)
	assert.Equal(t, model.VideoStatusNew, v.Status)
}
func TestService_VideoCoverStatusUpdate(t *testing.T) {
	err := srv.Mongo.Collection(model.Video{}.Collection()).Drop(ctx)
	assert.NoError(t, err)

	vid := int64(10000001)
	// presave
	v, err := srv.PreSaveVideo(ctx, 1, vid)
	assert.NoError(t, err)
	assert.Equal(t, model.CoverStatusUploading, v.CoverStatus)

	err = srv.VideoCoverStatusUpdate(ctx, vid, model.CoverStatusSuccess)
	assert.NoError(t, err)
	v, err = srv.VideoDetailDB(ctx, vid)
	assert.NoError(t, err)
	assert.Equal(t, model.CoverStatusSuccess, v.CoverStatus)
}
