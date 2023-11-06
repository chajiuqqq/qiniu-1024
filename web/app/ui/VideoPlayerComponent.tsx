"use client";
import PlyrComponent from "@/app/ui/video-player/player";
import React,{ useState, useEffect } from "react";
import Loading from "./loading";
import { MainVideoItem } from "../lib/api/types";

interface VideoPlayerProps {
  videos:MainVideoItem[];
  setVideos: React.Dispatch<React.SetStateAction<MainVideoItem[] | undefined>>
  dev?:boolean;
  updateVideos:()=>void
  startedVideoID?:number
}
const VideoPlayerComponent:React.FC<VideoPlayerProps> = ({videos,setVideos,updateVideos,dev=true,startedVideoID}) => {
  const startedIndex = videos.findIndex(video => video.id === startedVideoID);
  const [index, setIndex] = useState<number>(startedIndex==-1?0:startedIndex);
  const nextVideo = () => {
    if (index < videos.length - 1) {
      setIndex(n=>n+1);
    } else {
      updateVideos()
      setIndex(0)
    }
  };

  const lastVideo = () => {
    if (index > 0) {
      setIndex(n=>n-1);
    }
  };

  useEffect(() => {
    // 定义一个处理键盘事件的函数
    const handleKeyDown = (event: KeyboardEvent) => {
      switch (event.key) {
        case "ArrowUp":
          lastVideo();
          break;
        case "ArrowDown":
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
  const handleLike = ()=>{
    let vs =[...videos]
    vs[index].likes_count+=1
    vs[index].liked=true
    setVideos(vs)
  }
  const handleCancelLike = ()=>{
    let vs =[...videos]
    vs[index].likes_count-=1
    vs[index].liked=false
    setVideos(vs)
  }
  const handleCollect = ()=>{
    let vs =[...videos]
    vs[index].collect_count+=1
    vs[index].collected=true
    setVideos(vs)
  }
  const handleCancelCollect = ()=>{
    let vs =[...videos]
    vs[index].collect_count-=1
    vs[index].collected=false
    setVideos(vs)
  }
  return (
    <>
      {videos.length > 0 && index >= 0 && index < videos.length ? (
        <>
            <PlyrComponent
             v={videos[index]}
             onLike={handleLike}
             onCancelLike={handleCancelLike}
             onCollect={handleCollect}
             onCancelCollect={handleCancelCollect}
            />
        </>
      ) : (
        <Loading></Loading> 
      )}
    </>
  );
};

export default VideoPlayerComponent;
