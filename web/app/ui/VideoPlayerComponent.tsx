"use client";
import PlyrComponent, { PlyrAttach } from "@/app/ui/video-player/player";
import React, { useState, useEffect } from "react";
import Loading from "./loading";
import { MainVideoItem } from "../lib/api/types";
import api from "../lib/api/api-client";

interface VideoPlayerProps {
  videos: MainVideoItem[];
  setVideos: React.Dispatch<React.SetStateAction<MainVideoItem[] | undefined>>;
  dev?: boolean;
  updateVideos: () => void;
  startedVideoID?: number;
}
const VideoPlayerComponent: React.FC<VideoPlayerProps> = ({
  videos,
  setVideos,
  updateVideos,
  dev = true,
  startedVideoID,
}) => {
  const startedIndex = videos.findIndex((video) => video.id === startedVideoID);
  const [index, setIndex] = useState<number>(
    startedIndex == -1 ? 0 : startedIndex
  );
  const nextVideo = () => {
    setIndex((n) => {
      if (index < videos.length - 1) {
        return n + 1;
      } else {
        updateVideos();
        return 0;
      }
    });
  };

  const lastVideo = () => {
    setIndex((n) => {
      if (n > 0) {
        return n - 1;
      }
      return n;
    });
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

  useEffect(() => {
    setVideos((vs) => {
      if (vs) {
        let newVs = [...vs];
        newVs[index].play_count += 1;
        setVideos(newVs);
        api.action.playVideo(newVs[index].id).then((res) => {
          console.log("play：" + newVs[index].id);
        });
      }
      return vs;
    });
  }, [index]);
  const handleLike = () => {
    let vs = [...videos];
    vs[index].likes_count += 1;
    vs[index].liked = true;
    setVideos(vs);
  };
  const handleCancelLike = () => {
    let vs = [...videos];
    vs[index].likes_count -= 1;
    vs[index].liked = false;
    setVideos(vs);
  };
  const handleCollect = () => {
    let vs = [...videos];
    vs[index].collect_count += 1;
    vs[index].collected = true;
    setVideos(vs);
  };
  const handleCancelCollect = () => {
    let vs = [...videos];
    vs[index].collect_count -= 1;
    vs[index].collected = false;
    setVideos(vs);
  };

  return (
    <>
      {videos.length > 0 && index >= 0 && index < videos.length ? (
        <>
          <PlyrComponent play_url={videos[index].play_url}>
            <PlyrAttach
              v={videos[index]}
              onLike={handleLike}
              onCancelLike={handleCancelLike}
              onCollect={handleCollect}
              onCancelCollect={handleCancelCollect}
            ></PlyrAttach>
          </PlyrComponent>
        </>
      ) : (
        <Loading></Loading>
      )}
    </>
  );
};

export default VideoPlayerComponent;
