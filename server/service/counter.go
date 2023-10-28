package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"qiniu-1024-server/model"
)

func (s *Service) GetUserSeq(ctx context.Context) (int64, error) {
	return s.getOneSeq(ctx, model.User{}.Collection())
}
func (s *Service) GetVideoSeq(ctx context.Context) (int64, error) {
	return s.getOneSeq(ctx, model.Video{}.Collection())
}
func (s *Service) GetCategorySeq(ctx context.Context) (int64, error) {
	return s.getOneSeq(ctx, model.Category{}.Collection())
}

// 返回新的seq
func (s *Service) getOneSeq(ctx context.Context, name string) (int64, error) {
	var cnt model.Counter
	r := options.After
	upsert := true
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &r,
		Upsert:         &upsert,
	}
	err := s.Mongo.Collection(model.Counter{}.Collection()).
		FindOneAndUpdate(ctx, bson.M{"_id": name},
			bson.M{"$inc": bson.M{"seq": 1}},
			&opt).Decode(&cnt)
	if err != nil {
		return 0, fmt.Errorf("get one counter[%s] failed: %w", name, err)
	}
	return cnt.Seq, nil
}
