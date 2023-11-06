export type UserRegisterPayload = {
    name: string;
    username: string;
    password: string;
    phone?: string;
    avatar_url?: string;
    description?: string;
}

type UserLikeItem = {
    user_id: number;
    created_at: string;
}

type FollowItem = {
    user_id: number;
    created_at: string;
}

type LikeItem = {
    video_id: number;
    created_at: string;
}

type CollectionItem = {
    video_id: number;
    created_at: string;
}

export type User = {
    id?: number;
    name?: string;
    username?: string;
    password?: string;
    phone?: string;
    avatar_url?: string;
    description?: string;
    github_account?: string;
    wechat_account?: string;
    user_likes?: UserLikeItem[];
    follows?: FollowItem[];
    followers?: FollowItem[];
    likes?: LikeItem[];
    collections?: CollectionItem[];
    created_at?: string;
    updated_at?: string;
}

export type MainVideoSubmit = {
    category_id: number;
    video_id: number;
    desc: string;
};

export type Comment = {
    id: number;
    content: string;
    user_id: number;
    video_id: number;
    likes_count: number;
    is_deleted: boolean;
    updated_at: string;
    created_at: string;
};

export type Video = {
    id: number;
    number: number;
    user_id: number;
    category_id: number;
    category: string;
    play_url: string;
    cover_url: string;
    description: string;
    play_count: number;
    likes_count: number;
    collect_count: number;
    comments?: Comment[];
    status: string; // ["Uploading", "New", "OnShow", "UnderShow"]
    cover_status: string; // ["Uploading", "Success", "Failed"]
    is_deleted: boolean;
    uploaded_at?: string | null;
    cover_uploaded_at?: string | null;
    submitted_at?: string | null;
    deleted_at?: string | null;
    created_at: string;
    updated_at: string;
};
export type UploadResponse = {
    vid:number,
}

export type MainVideoItem = Video &{
    user_id: number;
    nickname: string;
    avatar_url: string;
    follower_cnt: number;
    published_cnt: number;
    liked: boolean;
    collected: boolean;
    score: number;
  };
  export type VideoQuery = {
    category_id?:number,
    user_id?:number
  }

  export type Category = {
    id: number;
    name: string;
    order: number;
    onShow: boolean;
  };
  