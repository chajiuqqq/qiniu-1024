package model

import "time"

type UserLikeItem struct {
	UserID    int64     `bson:"user_id" json:"user_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

func (u UserLikeItem) Key() any {
	return u.UserID
}

type FollowItem struct {
	UserID    int64     `bson:"user_id" json:"user_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

func (u FollowItem) Key() any {
	return u.UserID
}

type LikeItem struct {
	VideoID   int64     `bson:"video_id" json:"video_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

func (u LikeItem) Key() any {
	return u.VideoID
}

type CollectionItem struct {
	VideoID   int64     `bson:"video_id" json:"video_id"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

func (u CollectionItem) Key() any {
	return u.VideoID
}

type User struct {
	ID            int64            `bson:"id" json:"id"` // 1e5+number
	Name          string           `bson:"name" json:"name"`
	Username      string           `bson:"username" json:"username"`
	Password      string           `bson:"password" json:"password"`
	Phone         string           `bson:"phone" json:"phone"`
	AvatarUrl     string           `bson:"avatar_url" json:"avatar_url"`
	Description   string           `bson:"description" json:"description"`
	GithubAccount string           `bson:"github_account" json:"github_account"`
	WechatAccount string           `bson:"wechat_account" json:"wechat_account"`
	UserLikes     []UserLikeItem   `bson:"user_likes" json:"user_likes"` // 喜欢该用户的列表
	Follows       []FollowItem     `bson:"follows" json:"follows"`
	Followers     []FollowItem     `bson:"followers" json:"followers"`
	Likes         []LikeItem       `bson:"likes" json:"likes"` // 喜欢的视频
	Collections   []CollectionItem `bson:"collections" json:"collections"`
	CreatedAt     time.Time        `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time        `bson:"updated_at" json:"updated_at"`
}

func (u User) Collection() string {
	return "users"
}
