// VideoItem.tsx
import React from "react";
import { Video } from "./list";
import {
  LikeIcon,
  MyLikeIcon,
  MyCommentIcon,
  MyCollectionIcon,
  MyShareIcon,
} from "@/app/ui/icon";
interface VideoItemProps {
  imgUrl: string;
  type: string;
  curVideo: Video;
}

const VideoItem: React.FC<VideoItemProps> = ({ imgUrl, type, curVideo }) => {
  let icon;
  switch (type) {
    case "my":
      break;

    default:
      break;
  }
  return (
    <div className="relative">
      <div className="w-40 h-64  bg-gray-200">
        <img
          src={imgUrl}
          alt="Video Thumbnail"
          className="w-full h-full object-contain"
        />
      </div>

      <div className="absolute bottom-0 left-0 bg-black bg-opacity-50 text-white p-2">
        100
      </div>
    </div>
  );
};

export default VideoItem;
