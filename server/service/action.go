package service

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"qiniu-1024-server/model"
	"qiniu-1024-server/types"
	"qiniu-1024-server/utils"
	"qiniu-1024-server/utils/xerr"
	"qiniu-1024-server/utils/xmongo"
	"time"
)

// TODO: 校验
type ActionService interface {
	PlayVideo(ctx context.Context, uid, vid int64) (*model.Video, error)
	CollectVideo(ctx context.Context, uid, vid int64) (*model.Video, error)
	UnCollectVideo(ctx context.Context, uid, vid int64) (*model.Video, error)
	LikeVideo(ctx context.Context, uid, vid int64) (*model.Video, error)
	UnLikeVideo(ctx context.Context, uid, vid int64) (*model.Video, error)
	CommentVideo(ctx context.Context, uid, vid int64, content string) (*model.Video, error)
	QueryAction(ctx context.Context, uid int64, videos []types.MainVideoItem) ([]types.MainVideoItem, error)
}
type DefaultActionService struct {
	ActionService
	parent *Service
	mongo  *xmongo.Database
}

func NewDefaultActionService(parent *Service) *DefaultActionService {
	return &DefaultActionService{
		parent: parent,
		mongo:  parent.Mongo,
	}
}
func (s DefaultActionService) PlayVideo(ctx context.Context, uid, vid int64) (*model.Video, error) {
	res, err := s.mongo.Collection(model.Video{}.Collection()).
		UpdateOne(ctx, bson.M{"id": vid}, bson.M{"$inc": bson.M{"play_count": 1}})
	if err != nil {
		return nil, fmt.Errorf("update PlayVideo error,vid[%d]:%v", vid, err)
	}
	if res.MatchedCount == 0 {
		return nil, xerr.New(400, "VideoNotFound", "video not found")
	}
	v, err := s.parent.VideoDetailDB(ctx, vid)
	if err != nil {
		return nil, err
	}
	vlog := &model.VideoLog{
		Op:      model.ActionPlay,
		VideoID: vid,
		OwnerID: v.UserID,
		UserID:  uid,
	}
	_, err = s.mongo.Collection(model.VideoLog{}.Collection()).
		InsertOne(ctx, vlog)
	if err != nil {
		return nil, fmt.Errorf("save PlayVideo log error,vid[%d]:%v", vid, err)
	}
	return v, nil
}

func (s DefaultActionService) CollectVideo(ctx context.Context, uid, vid int64) (*model.Video, error) {
	// collect_count
	res, err := s.mongo.Collection(model.Video{}.Collection()).
		UpdateOne(ctx, bson.M{"id": vid}, bson.M{"$inc": bson.M{"collect_count": 1}})
	if err != nil {
		return nil, fmt.Errorf("CollectVideo inc collect_count error,vid[%d]:%v", vid, err)
	}

	if res.MatchedCount == 0 {
		return nil, xerr.New(400, "VideoNotFound", "video not found")
	}
	v, err := s.parent.VideoDetailDB(ctx, vid)
	if err != nil {
		return nil, err
	}

	// user
	_, err = s.mongo.Collection(model.User{}.Collection()).UpdateOne(ctx,
		bson.M{"id": uid}, bson.M{"$push": bson.M{"collections": &model.CollectionItem{
			VideoID:   vid,
			CreatedAt: time.Now(),
		}}})
	if err != nil {
		return nil, fmt.Errorf("CollectVideo update user error,vid[%d]:%v", vid, err)
	}

	// log
	vlog := &model.VideoLog{
		Op:      model.ActionCollect,
		VideoID: vid,
		OwnerID: v.UserID,
		UserID:  uid,
	}
	_, err = s.mongo.Collection(model.VideoLog{}.Collection()).
		InsertOne(ctx, vlog)
	if err != nil {
		return nil, fmt.Errorf("save CollectVideo log error,vid[%d]:%v", vid, err)
	}
	return v, nil
}
func (s DefaultActionService) UnCollectVideo(ctx context.Context, uid, vid int64) (*model.Video, error) {
	// collect_count
	res, err := s.mongo.Collection(model.Video{}.Collection()).
		UpdateOne(ctx, bson.M{"id": vid}, bson.M{"$inc": bson.M{"collect_count": -1}})
	if err != nil {
		return nil, fmt.Errorf("UnCollectVideo inc collect_count error,vid[%d]:%v", vid, err)
	}
	if res.MatchedCount == 0 {
		return nil, xerr.New(400, "VideoNotFound", "video not found")
	}
	v, err := s.parent.VideoDetailDB(ctx, vid)
	if err != nil {
		return nil, err
	}

	// user
	_, err = s.mongo.Collection(model.User{}.Collection()).UpdateOne(ctx,
		bson.M{"id": uid}, bson.M{"$pull": bson.M{"collections": bson.M{"video_id": vid}}})
	if err != nil {
		return nil, fmt.Errorf("UnCollectVideo update user error,vid[%d]:%v", vid, err)
	}

	// log
	vlog := &model.VideoLog{
		Op:      model.ActionUnCollect,
		VideoID: vid,
		OwnerID: v.UserID,
		UserID:  uid,
	}
	_, err = s.mongo.Collection(model.VideoLog{}.Collection()).
		InsertOne(ctx, vlog)
	if err != nil {
		return nil, fmt.Errorf("save UnCollectVideo log error,vid[%d]:%v", vid, err)
	}
	return v, nil
}
func (s DefaultActionService) LikeVideo(ctx context.Context, uid, vid int64) (*model.Video, error) {
	// collect_count
	res, err := s.mongo.Collection(model.Video{}.Collection()).
		UpdateOne(ctx, bson.M{"id": vid}, bson.M{"$inc": bson.M{"likes_count": 1}})
	if err != nil {
		return nil, fmt.Errorf("LikeVideo inc likes_count error,vid[%d]:%v", vid, err)
	}
	if res.MatchedCount == 0 {
		return nil, xerr.New(400, "VideoNotFound", "video not found")
	}
	v, err := s.parent.VideoDetailDB(ctx, vid)
	if err != nil {
		return nil, err
	}

	// user
	_, err = s.mongo.Collection(model.User{}.Collection()).UpdateOne(ctx,
		bson.M{"id": uid}, bson.M{"$push": bson.M{"likes": &model.LikeItem{
			VideoID:   vid,
			CreatedAt: time.Now(),
		}}})
	if err != nil {
		return nil, fmt.Errorf("LikeVideo update user error,vid[%d]:%v", vid, err)
	}

	// log
	vlog := &model.VideoLog{
		Op:      model.ActionLike,
		VideoID: vid,
		OwnerID: v.UserID,
		UserID:  uid,
	}
	_, err = s.mongo.Collection(model.VideoLog{}.Collection()).
		InsertOne(ctx, vlog)
	if err != nil {
		return nil, fmt.Errorf("save LikeVideo log error,vid[%d]:%v", vid, err)
	}
	return v, nil
}
func (s DefaultActionService) UnLikeVideo(ctx context.Context, uid, vid int64) (*model.Video, error) {
	// collect_count
	res, err := s.mongo.Collection(model.Video{}.Collection()).
		UpdateOne(ctx, bson.M{"id": vid}, bson.M{"$inc": bson.M{"likes_count": -1}})
	if err != nil {
		return nil, fmt.Errorf("UnLikeVideo inc likes_count error,vid[%d]:%v", vid, err)
	}
	if res.MatchedCount == 0 {
		return nil, xerr.New(400, "VideoNotFound", "video not found")
	}
	v, err := s.parent.VideoDetailDB(ctx, vid)
	if err != nil {
		return nil, err
	}

	// user
	_, err = s.mongo.Collection(model.User{}.Collection()).UpdateOne(ctx,
		bson.M{"id": uid}, bson.M{"$pull": bson.M{"likes": bson.M{"video_id": vid}}})
	if err != nil {
		return nil, fmt.Errorf("UnLikeVideo update user error,vid[%d]:%v", vid, err)
	}

	// log
	vlog := &model.VideoLog{
		Op:      model.ActionCancelLike,
		VideoID: vid,
		OwnerID: v.UserID,
		UserID:  uid,
	}
	_, err = s.mongo.Collection(model.VideoLog{}.Collection()).
		InsertOne(ctx, vlog)
	if err != nil {
		return nil, fmt.Errorf("save UnLikeVideo log error,vid[%d]:%v", vid, err)
	}
	return v, nil
}

func (s DefaultActionService) CommentVideo(ctx context.Context, uid, vid int64, content string) (*model.Video, error) {
	// comment
	id, err := s.parent.GetCommentSeq(ctx)
	if err != nil {
		return nil, err
	}
	comment := &model.Comment{
		ID:         id,
		Content:    content,
		UserID:     uid,
		VideoID:    vid,
		LikesCount: 0,
	}
	res, err := s.mongo.Collection(model.Video{}.Collection()).
		UpdateOne(ctx, bson.M{"id": vid}, bson.M{"$push": bson.M{"comments": comment}})
	if err != nil {
		return nil, fmt.Errorf("CommentVideo push comment error,vid[%d]:%v", vid, err)
	}
	if res.MatchedCount == 0 {
		return nil, xerr.New(400, "VideoNotFound", "video not found")
	}
	v, err := s.parent.VideoDetailDB(ctx, vid)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s DefaultActionService) QueryAction(ctx context.Context, uid int64, videos []types.MainVideoItem) ([]types.MainVideoItem, error) {
	u, err := s.parent.UserDetailDB(ctx, uid)
	if err != nil {
		return nil, err
	}
	likesMap := utils.ToMap(u.Likes)
	collectionMap := utils.ToMap(u.Collections)
	for i := 0; i < len(videos); i++ {
		v := &videos[i]
		if _, ok := likesMap[v.ID]; ok {
			v.Liked = true
		}
		if _, ok := collectionMap[v.ID]; ok {
			v.Collected = true
		}
	}
	return videos, nil
}

type RedisActionService struct {
	ActionService
	parent *Service
	rdb    *redis.Client
	mongo  *xmongo.Database
}

func NewRedisActionService(parent *Service) *RedisActionService {
	return &RedisActionService{
		parent: parent,
		rdb:    parent.Rdb,
		mongo:  parent.Mongo,
	}
}
