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

func (s *Service) MainVideosDB(ctx context.Context, q types.VideoQuery) ([]model.Video, error) {
	var col = s.Mongo.Collection(model.Video{}.Collection())

	// filter
	filter := bson.M{"status": model.VideoStatusOnShow}
	if q.UserID != 0 {
		filter["user_id"] = q.UserID
	}
	if q.CategoryID != 0 {
		filter["category_id"] = q.CategoryID
	}

	var data []model.Video
	opts := &options.FindOptions{
		Sort: bson.D{{"created_at", -1}},
	}
	cur, err := col.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("get main videos db failed: %w", err)
	}
	err = cur.All(ctx, &data)
	if err != nil {
		return nil, fmt.Errorf("get main videos db cursor failed: %w", err)
	}
	return data, nil
}

func (s *Service) MainVideos(ctx context.Context, q types.VideoQuery) ([]types.MainVideoItem, error) {
	data, err := s.MainVideosDB(ctx, q)
	if err != nil {
		return nil, err
	}
	// find all users
	var userIds []int64
	for _, v := range data {
		userIds = append(userIds, v.UserID)
	}

	usersMap, err := s.UsersMap(ctx, userIds)
	if err != nil {
		return nil, err
	}

	// compose
	var res = make([]types.MainVideoItem, len(data))
	for i := 0; i < len(data); i++ {
		v := data[i]
		u, ok := usersMap[data[i].UserID]
		if !ok {
			return nil, fmt.Errorf("get user [%d] of video [%d] failed, user not existed", v.UserID, v.ID)
		}
		publishedCnt, err := s.UserPublishedCnt(ctx, u.ID)
		if err != nil {
			return nil, err
		}
		res[i] = types.MainVideoItem{
			Video:        v,
			UserID:       u.ID,
			Nickname:     u.Name,
			AvatarUrl:    u.AvatarUrl,
			FollowerCnt:  len(u.Followers),
			PublishedCnt: publishedCnt,
			Score:        s.VideoScoreCal(v.PlayCount, v.LikesCount, v.CollectCount),
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
	prepared, err := s.VideoPrepared(ctx, req.VideoID)
	if err != nil {
		return nil, err
	}
	if !prepared {
		return nil, xerr.New(400, "VideoNotPrepared", "video is not prepared")
	}

	var col = s.Mongo.Collection(model.Video{}.Collection())
	var video = new(model.Video)
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	err = col.FindOneAndUpdate(ctx, bson.M{"id": req.VideoID, "user_id": uid},
		bson.M{"$set": bson.M{"category_id": req.CategoryID,
			"category":     category.Name,
			"description":  req.Desc,
			"status":       model.VideoStatusOnShow,
			"submitted_at": time.Now(),
		}}, &opt).Decode(video)
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

func (s *Service) VideoPrepared(ctx context.Context, vid int64) (bool, error) {
	existed, err := s.VideoExisted(ctx, vid)
	if err != nil || !existed {
		return false, err
	}
	v, err := s.VideoDetail(ctx, vid)
	if err != nil {
		return false, err
	}
	return v.Status == model.VideoStatusNew && v.CoverStatus == model.CoverStatusSuccess, nil
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
