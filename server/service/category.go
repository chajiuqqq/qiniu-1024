package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"qiniu-1024-server/model"
	"qiniu-1024-server/utils/xerr"
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
		if existed, err := s.CategoryExisted(ctx, category.Name); err != nil || existed {
			return xerr.New(400, "CategoryExisted", "category existed")
		}
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

func (s *Service) UpdateMainCategory(ctx context.Context, id int64, category model.Category) (*model.Category, error) {
	col := s.Mongo.Collection(model.Category{}.Collection())
	set := bson.M{
		"name":    category.Name,
		"order":   category.Order,
		"on_show": category.OnShow,
	}
	r, err := col.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": set})
	if err != nil {
		return nil, fmt.Errorf("update category db error: %w", err)
	}
	if r.MatchedCount == 0 {
		return nil, xerr.New(404, "CategoryNotFound", "category not found")
	}
	var newCategory model.Category
	err = col.FindOne(ctx, bson.M{"id": id}).Decode(&newCategory)
	if err != nil {
		return nil, fmt.Errorf("find category db error: %w", err)
	}
	return &newCategory, nil
}
func (s *Service) CategoryDetail(ctx context.Context, id int64) (model.Category, error) {
	var category model.Category
	err := s.Mongo.Collection(model.Category{}.Collection()).FindOne(ctx, bson.M{"id": id}).Decode(&category)
	if err != nil {
		return model.Category{}, fmt.Errorf("find category db error: %w", err)
	}
	return category, nil
}

func (s *Service) CategoryExisted(ctx context.Context, name string) (bool, error) {
	cnt, err := s.Mongo.Collection(model.Category{}.Collection()).CountDocuments(ctx, bson.M{"name": name})
	if err != nil {
		return false, fmt.Errorf("find category exist db error: %w", err)
	}
	if cnt != 0 && cnt != 1 {
		return false, fmt.Errorf("find multi category exist error: %w", err)
	}
	return cnt != 0, nil
}
