"use client";
import { FunctionComponent, ReactNode, memo, useRef, useState } from "react";
import Plyr, { APITypes } from "plyr-react";
import "plyr-react/plyr.css";

import { HeartIcon, StarIcon } from "@heroicons/react/24/outline";

import {
  HeartIcon as SolidHeartIcon,
  StarIcon as SolidStarIcon,
} from "@heroicons/react/24/solid";
import { MainVideoItem, Video } from "@/app/lib/api/types";
import api from "@/app/lib/api/api-client";
const videoOptions: Plyr.Options = {
  autoplay: true,
  volume: 0.5,
  loop: { active: true },
};
interface opts {
  v: MainVideoItem;
  onLike: () => void;
  onCollect: () => void;
  onCancelLike: () => void;
  onCancelCollect: () => void;
}
export const PlyrAttach: React.FC<opts> = ({
  v,
  onLike,
  onCollect,
  onCancelLike,
  onCancelCollect,
}) => {
  return (
    <>
      <div className="absolute left-5 top-2 text-white z-30">{v.id}</div>
      <div className="absolute bottom-10 left-5 bg-black/50 rounded-lg p-5 text-white min-w-min">
        <h2 className="text-2xl inline">@{v.nickname}</h2>
        <button className="border rounded-lg border-white px-4 py-1 ml-5 hover:border-sky-400 hover:text-sky-400">
          {" "}
          + 关注
        </button>
        <div className="pt-2">
          {v.play_count} 播放 | {v.description}
        </div>
      </div>
      <div className="absolute bottom-10 right-2 p-4 flex flex-col justify-center items-center space-y-2 bg-black/50 rounded-lg">
        {v.liked ? (
          <SolidHeartIcon
            className="w-6 text-pink-500"
            onClick={onCancelLike}
          ></SolidHeartIcon>
        ) : (
          <HeartIcon className="w-6 text-white" onClick={onLike}></HeartIcon>
        )}

        <p className="text-white">{v.likes_count}</p>

        {/* <ChatBubbleBottomCenterTextIcon className="w-6 text-white"></ChatBubbleBottomCenterTextIcon> */}
        {/* <p className="text-white">288</p> */}
        {v.collected ? (
          <SolidStarIcon
            className="w-6 text-yellow-500"
            onClick={onCancelCollect}
          ></SolidStarIcon>
        ) : (
          <StarIcon className="w-6 text-white" onClick={onCollect}></StarIcon>
        )}
        <p className="text-white">{v.collect_count}</p>
      </div>
    </>
  );
};
type ParentComponentProps = {
  children: React.ReactNode;
  play_url: string;
};
const PlyrComponent: React.FC<ParentComponentProps> = ({
  children,
  play_url,
}) => {
  const ref = useRef<APITypes>(null);

  return (
    <div className="relative w-full">
      <Plyr
        ref={ref}
        source={{
          type: "video",
          sources: [
            {
              src: play_url,
            },
          ],
        }}
        options={videoOptions}
      />
      {children}
    </div>
  );
};

export default PlyrComponent;
