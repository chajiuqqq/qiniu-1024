package types

import "qiniu-1024-server/model"

type MainVideoSubmit struct {
	CategoryID int64  `json:"category_id"`
	VideoID    int64  `json:"video_id"`
	Desc       string `json:"desc"`
}

type MainVideoUploadResponse struct {
	VideoID int64  `json:"video_id"`
	Url     string `json:"url"`
}

type VideoQuery struct {
	CategoryID int64 `query:"category_id"` // 只显示该分类的视频
	UserID     int64 `query:"user_id"`     // 只显示该用户的视频
}

type MainVideoItem struct {
	model.Video
	UserID       int64  `json:"user_id"`
	Nickname     string `json:"nickname"`
	AvatarUrl    string `json:"avatar_url"`
	FollowerCnt  int    `json:"follower_cnt"`
	PublishedCnt int    `json:"published_cnt"`
	Liked        bool   `json:"liked"`
	Collected    bool   `json:"collected"`
	Score        int64  `json:"score"`
}
