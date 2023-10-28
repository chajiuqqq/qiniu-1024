package types

import "qiniu-1024-server/model"

type MainVideoQuery struct {
	CategoryID int64 `json:"category_id"`
}

type MainVideoSubmit struct {
	CategoryID int64  `json:"category_id"`
	VideoID    int64  `json:"video_id"`
	Desc       string `json:"desc"`
}

type MainVideoUploadResponse struct {
	VideoID int64  `json:"video_id"`
	Url     string `json:"url"`
}

type MainVideoItem struct {
	model.Video
	Score int64 `json:"score"`
}
