"use client";
import { useRef } from "react";
import Plyr, { APITypes } from "plyr-react";
import "plyr-react/plyr.css";

import {
  HeartIcon,
  StarIcon,
  ChatBubbleBottomCenterTextIcon,
  ShareIcon
} from '@heroicons/react/24/outline';


import {
  HeartIcon as SolidHeartIcon,
  StarIcon as SolidStarIcon,
} from '@heroicons/react/24/solid';
import { MainVideoItem, Video } from "@/app/lib/api/types";
const videoOptions: Plyr.Options = {
  autoplay: true,
  volume: 0.5,
  loop: { active: true },
};
const PlyrComponent: React.FC<MainVideoItem> = (v) => {
  const ref = useRef<APITypes>(null);

  return (
    <div className="relative">
      <Plyr
        // className="w-6/12"
        ref={ref}
        source={{
          type: "video",
          sources: [
            {
              src: v.play_url,
            },
          ],
        }}
        options={videoOptions}
      />
      <div className="absolute bottom-10 left-5">
        <h2 className="text-2xl text-white">{v.nickname}</h2>
        <div>{v.description}</div>
      </div>
      <div className="absolute bottom-10 right-2 p-4 flex flex-col justify-center items-center space-y-2 bg-black/50 rounded-lg">
        {
          v.liked ? (
            <SolidHeartIcon className="w-6 text-pink-500"></SolidHeartIcon>
          ) : (
            <HeartIcon className="w-6 text-white"></HeartIcon>
          )
        }

        <p className="text-white">{v.play_count}</p>

        {/* <ChatBubbleBottomCenterTextIcon className="w-6 text-white"></ChatBubbleBottomCenterTextIcon> */}
        {/* <p className="text-white">288</p> */}
        {
          v.collected?(
            <SolidStarIcon className="w-6 text-yellow-500"></SolidStarIcon>
          ):(
            <StarIcon className="w-6 text-white"></StarIcon>
          )
        }
        <p className="text-white">{v.collect_count}</p>

      </div>
    </div>
  );
};

export default PlyrComponent;
