// VideoItemList.tsx
import React from "react";
import VideoItem from "./item";

export interface Video {
  imgUrl: string;
}

interface VideoItemListProps {
  videos: Video[];
  type: string;
}

const VideoItemList: React.FC<VideoItemListProps> = ({ videos, type }) => {
  return (
    <div className="grid grid-cols-9 gap-4 w-full">
      {videos.map((video, idx) => (
        <VideoItem
          key={idx}
          imgUrl={video.imgUrl}
          curVideo={video}
          type={type}
        />
      ))}
    </div>
  );
};

export default VideoItemList;
