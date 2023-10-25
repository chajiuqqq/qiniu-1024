package model

import "time"

type Comment struct {
	Content    string    `bson:"content" json:"content"`
	UserID     int64     `bson:"user_id" json:"user_id"`
	VideoID    int64     `bson:"video_id" json:"video_id"`
	LikesCount int64     `bson:"likes_count" json:"likes_count"`
	IsDeleted  bool      `bson:"is_deleted" json:"is_deleted"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
}
type CommentLog struct {
	CommentID int64     `bson:"comment_id" json:"comment_id"`
	Op        string    `bson:"op" json:"op"` // [Like,Delete]
	UserID    int64     `bson:"user_id" json:"user_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

func (c CommentLog) Collection() string {
	return "commentLogs"
}

type Video struct {
	ID           int64     `bson:"id" json:"id"` // CategoryID*1e8 + Number
	Number       int64     `bson:"number" json:"number"`
	UserID       int64     `bson:"user_id" json:"user_id"`
	CategoryID   int64     `bson:"category_id" json:"category_id"`
	Category     string    `bson:"category" json:"category"`
	PlayUrl      string    `bson:"play_url" json:"play_url"`
	CoverUrl     string    `bson:"cover_url" json:"cover_url"`
	Description  string    `bson:"description" json:"description"`
	PlayCount    int64     `bson:"play_count" json:"play_count"`
	LikesCount   int64     `bson:"likes_count" json:"likes_count"`
	CollectCount int64     `bson:"collect_count" json:"collect_count"`
	Comments     []Comment `bson:"comments" json:"comments"`
	Status       string    `bson:"status" json:"status"` // [New新上传, OnShow通过审核,UnderShow需要修改的]
	IsDeleted    bool      `bson:"is_deleted" json:"is_deleted"`
	DeletedAt    time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
}

func (v Video) Collection() string {
	return "videos"
}

type VideoLog struct {
	Op        string    `bson:"op" json:"op"` // [Like,Collect,Play]
	VideoID   int64     `bson:"video_id" json:"video_id"`
	OwnerID   int64     `bson:"owner_id" json:"owner_id"`
	UserID    int64     `bson:"user_id" json:"user_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

func (v VideoLog) Collection() string {
	return "videoLogs"
}

type Category struct {
	ID        int64     `bson:"id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (c Category) Collection() string {
	return "categories"
}
