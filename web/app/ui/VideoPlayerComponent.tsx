"use client";
import PlyrComponent from "@/app/ui/video-player/player";
import React,{ useState, useEffect } from "react";
import { VideoType } from "../lib/video";

interface VideoPlayerProps {
  videos:VideoType[];
  dev?:boolean;
  updateVideos:()=>void
  startedVideoID?:number
}
const VideoPlayerComponent:React.FC<VideoPlayerProps> = ({videos,updateVideos,dev=true,startedVideoID}) => {
  const startedIndex = videos.findIndex(video => video.id === startedVideoID);
  const [index, setIndex] = useState<number>(startedIndex==-1?0:startedIndex);
  const nextVideo = () => {
    if (index < videos.length - 1) {
      setIndex((index) => index + 1);
    } else {
      updateVideos()
      setIndex(0)
    }
  };

  const lastVideo = () => {
    if (index > 0) {
      setIndex((index) => index - 1);
    }
  };

  useEffect(() => {
    // 定义一个处理键盘事件的函数
    const handleKeyDown = (event: KeyboardEvent) => {
      switch (event.key) {
        case "ArrowUp":
          console.log("上箭头键被按下了");
          lastVideo();
          break;
        case "ArrowDown":
          console.log("下箭头键被按下了");
          nextVideo();
          break;
        default:
          // 其他按键可以在此处理
          break;
      }
    };

    // 在组件挂载时添加事件监听器
    window.addEventListener("keydown", handleKeyDown);
    // 组件卸载时移除滚动事件监听器
    return () => {
      window.removeEventListener("keydown", handleKeyDown);
    };
  }, []); 

  return (
    <>
      {videos.length > 0 && index >= 0 && index < videos.length ? (
        <>
          <div className="w-full">
            <PlyrComponent
              url={videos[index].play_url}
              desc={videos[index].description}
            />
          </div>
        </>
      ) : (
        <p>Loading or invalid index...</p>
      )}
    </>
  );
};

export default VideoPlayerComponent;
