package model

type User struct {
	ID            int64   `bson:"id" json:"id"` // 1e5+number
	Name          string  `bson:"name" json:"name"`
	Username      string  `bson:"username" json:"username"`
	Password      string  `bson:"password" json:"password"`
	Phone         string  `bson:"phone" json:"phone"`
	AvatarUrl     string  `bson:"avatar_url" json:"avatar_url"`
	Description   string  `bson:"description" json:"description"`
	GithubAccount string  `bson:"github_account" json:"github_account"`
	WechatAccount string  `bson:"wechat_account" json:"wechat_account"`
	UserLikes     []int64 `bson:"user_likes" json:"user_likes"` // 喜欢该用户的列表
	Follows       []int64 `bson:"follows" json:"follows"`
	Followers     []int64 `bson:"followers" json:"followers"`
	Likes         []int64 `bson:"likes" json:"likes"` // 喜欢的视频
	Collections   []int64 `bson:"collections" json:"collections"`
	CreatedAt     int64   `bson:"created_at" json:"created_at"`
	UpdatedAt     int64   `bson:"updated_at" json:"updated_at"`
}

func (u User) Collection() string {
	return "users"
}

type UserFollowLog struct {
	From      int64 `bson:"from" json:"from"`
	To        int64 `bson:"to" json:"to"`
	CreatedAt int64 `bson:"created_at" json:"created_at"`
}

func (u UserFollowLog) Collection() string {
	return "userFollowLogs"
}
