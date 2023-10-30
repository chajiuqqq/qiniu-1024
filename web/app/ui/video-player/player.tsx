import { useRef } from "react";
import Plyr, { APITypes } from "plyr-react";
import "plyr-react/plyr.css";

const videoId = "yWtFb9LJs3o";
const provider = "youtube";
const videoOptions = undefined;

const PlyrComponent = () => {
  const ref = useRef<APITypes>(null);

  const enterVideo = () => {
    (ref.current?.plyr as Plyr)?.fullscreen.enter();
  };

  const make2x = () => {
    const plyrInstance = ref.current?.plyr as Plyr;
    if (plyrInstance) plyrInstance.speed = 2;
  };

  const plyrVideo =
    videoId && provider ? (
      <Plyr
        ref={ref}
        source={{
          type: "video",
          sources: [
            {
              src: videoId,
              provider: provider
            }
          ]
        }}
        options={videoOptions}
      />
    ) : null;

  return (
    <div>
      {plyrVideo}
      <button onClick={enterVideo}>fullscreen</button>
      <button onClick={make2x}>2x</button>
    </div>
  );
};

export default PlyrComponent;
