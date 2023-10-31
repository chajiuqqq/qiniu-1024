"use client";
import { useRef } from "react";
import Plyr, { APITypes } from "plyr-react";
import "plyr-react/plyr.css";

const videoOptions: Plyr.Options = {
  autoplay: true,
  volume: 0.5,
  loop: { active: true },
};
interface PlyrOption {
  url: string;
}
const PlyrComponent: React.FC<PlyrOption> = ({ url }) => {
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
    <div className="w-full">
      {plyrVideo}
      <button onClick={enterVideo}>fullscreen</button>
      <button onClick={make2x}>2x</button>
    </div>
  );
};

export default PlyrComponent;
