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
const videoOptions: Plyr.Options = {
  autoplay: true,
  volume: 0.5,
  loop: { active: true },
};
interface PlyrOption {
  url: string
  desc: string
}
const PlyrComponent: React.FC<PlyrOption> = ({ url, desc }) => {
  const ref = useRef<APITypes>(null);

  const enterVideo = () => {
    (ref.current?.plyr as Plyr)?.fullscreen.enter();
  };

  const make2x = () => {
    const plyrInstance = ref.current?.plyr as Plyr;
    if (plyrInstance) plyrInstance.speed = 2;
  };

  const plyrVideo = (
    <Plyr
      ref={ref}
      source={{
        type: "video",
        sources: [
          {
            src: url,
          },
        ],
      }}
      options={videoOptions}
    />
  );

  return (
    <div className="w-full relative">
      {plyrVideo}
      <div className="absolute bottom-10 left-5">
        <h2 className="text-2xl text-white">{url}</h2>
        <div>{desc}</div>
      </div>
      <div className="absolute bottom-10 right-0 p-4 flex flex-col justify-center items-center space-y-2">
        <HeartIcon className="w-6 text-white"></HeartIcon>
        <p className="text-white">1000</p>

        <ChatBubbleBottomCenterTextIcon className="w-6 text-white"></ChatBubbleBottomCenterTextIcon>
        <p className="text-white">288</p>

        <StarIcon className="w-6 text-white"></StarIcon>
        <p className="text-white">99</p>

        <ShareIcon className="w-6 text-white"></ShareIcon>
        <p className="text-white">200</p>
      </div>
    </div>
  );
};

export default PlyrComponent;
