export type VideoComment = {
    content: string,
}
export type VideoType = {
    id: number;
    play_url: string;
    cover_url: string,
    description: string;
    play_count: number,
    likes_count: number,
    collect_count: number,
    comments: VideoComment[] | null,
};