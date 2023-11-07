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
import { ProfileTab } from "@/app/lib/const";
import { MainVideoItem } from "@/app/lib/api/types";
interface VideoItemProps {
  imgUrl: string;
  type: ProfileTab;
  curVideo: MainVideoItem;
  onClick:(videoID:number)=>void
}

const VideoItem: React.FC<VideoItemProps> = ({ imgUrl, type, curVideo,onClick }) => {
  let label;
  switch (type) {
    case ProfileTab.My:
      label = <><MyPlayIcon></MyPlayIcon><span> {curVideo.play_count}</span> </>
      break;
    case ProfileTab.Likes:
      label = <><MyLikeIcon></MyLikeIcon><span> {curVideo.likes_count}</span></>
      break;
    case ProfileTab.Collection:
      label = <><MyCollectionIcon></MyCollectionIcon><span> {curVideo.collect_count}</span></>
      break;
    default:
      break;
  }
  return (
    <div className="relative cursor-pointer" onClick={()=>{ console.log(curVideo.id,'clicked');onClick(curVideo.id)}}>
      <div className=" w-64 bg-gray-200 rounded-md">
        <img
          src={imgUrl}
          alt="Video Thumbnail"
          className="w-full h-full object-contain"
        />
      </div>

      <div className="absolute bottom-0 left-0 text-white bg-black/50 text-shadow p-2 rounded-md">
      <div className="flex space-x-2" >
        {label}
      </div>
        
      </div>
    </div>
  );
};

export default VideoItem;
