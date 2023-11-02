// VideoItemList.tsx
import React from "react";
import VideoItem from "./item";
import { Video } from "@/app/lib/video";
import { ProfileTab } from "@/app/lib/const";

interface VideoItemListProps {
  videos: Video[];
  type: ProfileTab;
}

const VideoItemList: React.FC<VideoItemListProps> = ({ videos, type }) => {
  return (
    <div className="grid  grid-cols-5 2xl:grid-cols-8 gap-4 w-full">
      {videos.map((video, idx) => (
        <VideoItem
          key={idx}
          imgUrl={video.cover_url}
          curVideo={video}
          type={type}
        />
      ))}
    </div>
  );
};

export default VideoItemList;
