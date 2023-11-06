// VideoItemList.tsx
import React from "react";
import VideoItem from "./item";
import { ProfileTab } from "@/app/lib/const";
import { MainVideoItem } from "@/app/lib/api/types";

interface VideoItemListProps {
  videos: MainVideoItem[];
  type: ProfileTab;
  onClick:(videoID:number)=>void
}

const VideoItemList: React.FC<VideoItemListProps> = ({ videos, type,onClick }) => {
  return (
    <div className="grid  grid-cols-5 2xl:grid-cols-8 gap-4 w-full">
      {videos.map((video, idx) => (
        <VideoItem
          key={idx}
          imgUrl={video.cover_url}
          curVideo={video}
          type={type}
          onClick={onClick}
        />
      ))}
    </div>
  );
};

export default VideoItemList;
