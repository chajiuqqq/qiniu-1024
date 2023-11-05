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

export type MainVideoSubmit = {
    categoryId: number;
    videoId: number;
    desc: string;
};

export type Comment = {
    id: number;
    content: string;
    userId: number;
    videoId: number;
    likesCount: number;
    isDeleted: boolean;
    updatedAt: Date;
    createdAt: Date;
};

export type Video = {
    id: number;
    number: number;
    userId: number;
    categoryId: number;
    category: string;
    playUrl: string;
    coverUrl: string;
    description: string;
    playCount: number;
    likesCount: number;
    collectCount: number;
    comments: Comment[];
    status: string; // ["Uploading", "New", "OnShow", "UnderShow"]
    coverStatus: string; // ["Uploading", "Success", "Failed"]
    isDeleted: boolean;
    uploadedAt?: Date | null;
    coverUploadedAt?: Date | null;
    submittedAt?: Date | null;
    deletedAt?: Date | null;
    createdAt: Date;
    updatedAt: Date;
};
export type UploadResponse = {
    vid:number,
}