package service

import (
	"github.com/stretchr/testify/assert"
	"qiniu-1024-server/model"
	"testing"
)

func TestService_GetUserSeq(t *testing.T) {
	err := srv.Mongo.Collection(model.Counter{}.Collection()).Drop(ctx)
	assert.NoError(t, err)

	seq, err := srv.GetUserSeq(ctx)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), seq)

	seq, err = srv.GetUserSeq(ctx)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), seq)
}
