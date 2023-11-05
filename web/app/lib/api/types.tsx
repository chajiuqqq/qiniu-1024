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
    created_at: Date;
}

type FollowItem = {
    user_id: number;
    created_at: Date;
}

type LikeItem = {
    video_id: number;
    created_at: Date;
}

type CollectionItem = {
    video_id: number;
    created_at: Date;
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
    created_at?: Date;
    updated_at?: Date;
}