package service

import (
	"github.com/stretchr/testify/assert"
	"qiniu-1024-server/model"
	"testing"
)

func TestService_SaveMainCategories(t *testing.T) {
	err := srv.Mongo.Collection(model.Category{}.Collection()).Drop(ctx)
	assert.NoError(t, err)

	err = srv.SaveMainCategories(ctx, []model.Category{
		{
			Name:  "娱乐",
			Order: 2,
		},
		{
			Name:  "旅游",
			Order: 1,
		},
	})
	assert.NoError(t, err)

	data, err := srv.MainCategories(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(data))
	assert.Equal(t, "旅游", data[0].Name)
	assert.Equal(t, "娱乐", data[1].Name)
}
