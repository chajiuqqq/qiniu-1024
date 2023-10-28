package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"qiniu-1024-server/model"
)

func (s *Service) MainCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	opts := &options.FindOptions{
		Sort: bson.D{{"order", 1}},
	}
	cur, err := s.Mongo.Collection(model.Category{}.Collection()).Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("find categories db error: %w", err)
	}
	err = cur.All(ctx, &categories)
	if err != nil {
		return nil, fmt.Errorf("find categories db cursor error: %w", err)
	}
	return categories, nil
}

func (s *Service) SaveMainCategories(ctx context.Context, categories []model.Category) error {
	col := s.Mongo.Collection(model.Category{}.Collection())
	for _, category := range categories {
		id, err := s.GetCategorySeq(ctx)
		if err != nil {
			return err
		}
		category.ID = id
		if category.Order == 0 {
			category.Order = 100
		}
		_, err = col.InsertOne(ctx, category)
		if err != nil {
			return fmt.Errorf("insert categories db error: %w", err)
		}
	}
	return nil
}
func (s *Service) CategoryDetail(ctx context.Context, id int64) (model.Category, error) {
	var category model.Category
	err := s.Mongo.Collection(model.Category{}.Collection()).FindOne(ctx, bson.M{"id": id}).Decode(&category)
	if err != nil {
		return model.Category{}, fmt.Errorf("find category db error: %w", err)
	}
	return category, nil
}
