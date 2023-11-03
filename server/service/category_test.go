package service

import (
	"github.com/stretchr/testify/assert"
	"qiniu-1024-server/model"
	"qiniu-1024-server/utils/xerr"
	"testing"
)

func TestService_SaveMainCategories(t *testing.T) {
	err := srv.Mongo.Collection(model.Category{}.Collection()).Drop(ctx)
	assert.NoError(t, err)

	err = srv.SaveMainCategories(ctx, []model.Category{
		{
			Name:   "娱乐",
			Order:  2,
			OnShow: true,
		},
		{
			Name:   "旅游",
			Order:  1,
			OnShow: true,
		},
	})
	assert.NoError(t, err)

	data, err := srv.MainCategories(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(data))
	assert.Equal(t, "旅游", data[0].Name)
	assert.Equal(t, "娱乐", data[1].Name)
	//  save same category
	err = srv.SaveMainCategories(ctx, []model.Category{
		{
			Name:   "娱乐",
			Order:  2,
			OnShow: true,
		},
	})
	assert.NotNil(t, err)
	if err != nil {
		e, ok := err.(*xerr.XError)
		assert.True(t, ok)
		if ok {
			assert.Equal(t, 400, e.StatusCode)
		}
	}

	// detail
	d, err := srv.CategoryDetail(ctx, data[0].ID)
	assert.NoError(t, err)
	assert.Equal(t, "旅游", d.Name)

	exist, err := srv.CategoryExisted(ctx, "旅游")
	assert.NoError(t, err)
	assert.True(t, exist)

	exist, err = srv.CategoryExisted(ctx, "旅游2")
	assert.NoError(t, err)
	assert.False(t, exist)

	// update
	m, err := srv.UpdateMainCategory(ctx, data[0].ID, model.Category{
		Name:   "旅游3",
		Order:  100,
		OnShow: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, "旅游3", m.Name)
	assert.Equal(t, int64(100), m.Order)
}
