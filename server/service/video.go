package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"qiniu-1024-server/model"
	"qiniu-1024-server/types"
	"qiniu-1024-server/utils/xerr"
	"slices"
	"strconv"
	"strings"
	"time"
)

func (s *Service) MainVideos(ctx context.Context, categoryID int64) ([]types.MainVideoItem, error) {
	var col = s.Mongo.Collection(model.Video{}.Collection())
	_, err := s.CategoryDetail(ctx, categoryID)
	if err != nil {
		return nil, xerr.New(400, "CategoryNotFound", "category id not exist")
	}

	var data []model.Video
	opts := &options.FindOptions{
		Sort: bson.D{{"created_at", 1}},
	}
	cur, err := col.Find(ctx, bson.M{"category_id": categoryID, "status": model.VideoStatusOnShow}, opts)
	if err != nil {
		return nil, fmt.Errorf("get main videos db failed: %w", err)
	}
	err = cur.All(ctx, &data)
	if err != nil {
		return nil, fmt.Errorf("get main videos db cursor failed: %w", err)
	}

	var res = make([]types.MainVideoItem, len(data))
	for i := 0; i < len(data); i++ {
		res[i] = types.MainVideoItem{
			Video: data[i],
			Score: s.VideoScoreCal(data[i].PlayCount, data[i].LikesCount, data[i].CollectCount),
		}
	}
	slices.SortFunc(res, func(a, b types.MainVideoItem) int {
		return int(a.Score - b.Score)
	})
	return res, nil
}

func (s *Service) VideoScoreCal(playCnt, likeCnt, collectCnt int64) int64 {
	// TODO: 添加用户是否浏览过参数
	return playCnt/1000 + likeCnt + collectCnt
}

func (s *Service) SaveVideo(ctx context.Context, uid int64, req types.MainVideoSubmit) (*model.Video, error) {
	category, err := s.CategoryDetail(ctx, req.CategoryID)
	if err != nil {
		return nil, xerr.New(400, "CategoryNotFound", "category id not exist")
	}
	existed, err := s.VideoExisted(ctx, req.VideoID)
	if err != nil {
		return nil, err
	}
	if existed {
		return nil, xerr.New(400, "VideoExisted", "video existed")
	}

	var col = s.Mongo.Collection(model.Video{}.Collection())
	var video = new(model.Video)
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err = col.FindOneAndUpdate(ctx, bson.M{"id": req.VideoID, "user_id": uid},
		bson.M{"$set": bson.M{"category_id": req.CategoryID, "category": category.Name, "description": req.Desc}}, &opt).Decode(video)
	if err != nil {
		return nil, fmt.Errorf("save video db failed: %w", err)
	}
	return video, nil
}

// PreSaveVideo 用于上传时预保存video
func (s *Service) PreSaveVideo(ctx context.Context, uid int64, vid int64) (*model.Video, error) {
	key := fmt.Sprintf("%d.mp4", vid)

	existed, err := s.VideoExisted(ctx, vid)
	if err != nil {
		return nil, err
	}
	if existed {
		return nil, xerr.New(400, "VideoExisted", "video existed")
	}

	video := &model.Video{
		ID:           vid,
		Number:       GetVideoNum(vid),
		UserID:       uid,
		PlayUrl:      s.Oss.ResourceUrl(key),
		CoverUrl:     s.Oss.CoverUrl(key),
		PlayCount:    0,
		LikesCount:   0,
		CollectCount: 0,
		Comments:     nil,
		Status:       model.VideoStatusUploading,
		CoverStatus:  model.CoverStatusUploading,
		IsDeleted:    false,
	}
	_, err = s.Mongo.Collection(model.Video{}.Collection()).InsertOne(ctx, video)
	if err != nil {
		return nil, fmt.Errorf("pre save video db failed: %w", err)
	}
	return video, nil
}

func (s *Service) VideoStatusUpdate(ctx context.Context, vid int64, status string) error {
	update := bson.M{"status": status}
	if status == model.VideoStatusNew {
		update["uploaded_at"] = time.Now()
	}
	res, err := s.Mongo.Collection(model.Video{}.Collection()).UpdateOne(ctx, bson.M{"id": vid},
		bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("update video db failed: %w", err)
	}
	if res.ModifiedCount == 0 {
		return xerr.New(400, "VideoNotFound", "video not found")
	}
	return nil
}
func (s *Service) VideoCoverStatusUpdate(ctx context.Context, vid int64, status string) error {
	update := bson.M{"cover_status": status}
	if status == model.CoverStatusSuccess {
		update["cover_uploaded_at"] = time.Now()
	}
	_, err := s.Mongo.Collection(model.Video{}.Collection()).UpdateOne(ctx, bson.M{"id": vid}, bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("update video cover db failed: %w", err)
	}
	return nil
}

func (s *Service) VideoDetail(ctx context.Context, vid int64) (*model.Video, error) {
	video := new(model.Video)
	err := s.Mongo.Collection(model.Video{}.Collection()).FindOne(ctx, bson.M{"id": vid}).Decode(video)
	if err != nil {
		return nil, fmt.Errorf("get video db failed: %w", err)
	}
	return video, nil
}
func (s *Service) VideoExisted(ctx context.Context, vid int64) (bool, error) {
	cnt, err := s.Mongo.Collection(model.Video{}.Collection()).CountDocuments(ctx, bson.M{"id": vid})
	if err != nil {
		return false, fmt.Errorf("get video db failed: %w", err)
	}
	return cnt != 0, nil
}

func GenVideoID(seq int64) int64 {
	return 1e8 + seq
}

func GetVideoNum(vid int64) int64 {
	return vid % 1e8
}
func GetVideoIDFromKey(key string) (int64, error) {
	ar := strings.Split(key, ".")
	return strconv.ParseInt(ar[0], 10, 64)
}
