"use client";
import PlyrComponent from "@/app/ui/video-player/player";
import { useState, useEffect } from "react";
import { Video } from "../lib/video";
import { initalVideos } from "../lib/data";
const url = "http://47.106.228.5:9133/v1/main/videos?category_id=1";
let dev = true;
const Hot = () => {
  const [videos, setVideos] = useState<Video[]>([]);
  const [index, setIndex] = useState<number>(0);
  const fetchVideos = async () => {
    if (dev) {
      setVideos(initalVideos);
      setIndex(0);
    } else {
      fetch(url)
        .then((response) => response.json())
        .then((data) => {
          setVideos(data);
          setIndex(0);
        });
    }
  };
  const nextVideo = () => {
    if (index < videos.length - 1) {
      setIndex((index) => index + 1);
    } else {
      fetchVideos();
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
  }, [videos, index]); // 注意这里的空数组，这确保了 useEffect 只在组件挂载时运行

  useEffect(() => {
    let ignore = false;
    if (dev) {
      setVideos(initalVideos);
      setIndex(0);
    } else {
      fetch(url)
        .then((response) => response.json())
        .then((data) => {
          if (!ignore) {
            setVideos(data);
            setIndex(0);
          }
        });
    }
    return () => {
      ignore = true;
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

export default Hot;
