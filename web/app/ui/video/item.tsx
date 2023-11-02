// VideoItem.tsx
import React from "react";

import {
  LikeIcon,
  MyLikeIcon,
  MyCommentIcon,
  MyCollectionIcon,
  MyShareIcon,
  MyPlayIcon
} from "@/app/ui/icon";
import { Video } from "@/app/lib/video";
import { ProfileTab } from "@/app/lib/const";
interface VideoItemProps {
  imgUrl: string;
  type: ProfileTab;
  curVideo: Video;
}

const VideoItem: React.FC<VideoItemProps> = ({ imgUrl, type, curVideo }) => {
  let label;
  switch (type) {
    case ProfileTab.My:
      label = <div className="flex space-x-2" ><MyPlayIcon></MyPlayIcon><span> {curVideo.play_count}</span></div>
      break;
    case ProfileTab.Likes:
      label = <div className="flex space-x-2" ><MyLikeIcon></MyLikeIcon><span> {curVideo.likes_count}</span></div>
      break;
    case ProfileTab.Collection:
      label = <div className="flex space-x-2" ><MyCollectionIcon></MyCollectionIcon><span> {curVideo.collect_count}</span></div>
      break;
    default:
      break;
  }
  return (
    <div className="relative">
      <div className="w-40 h-64  bg-gray-200 rounded-md">
        <img
          src={imgUrl}
          alt="Video Thumbnail"
          className="w-full h-full object-contain"
        />
      </div>

      <div className="absolute bottom-0 left-0 text-white bg-black/50 text-shadow p-2 rounded-md">
        {label}
      </div>
    </div>
  );
};

export default VideoItem;
