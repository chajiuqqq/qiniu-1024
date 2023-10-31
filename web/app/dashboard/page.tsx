"use client";
import PlyrComponent from "@/app/ui/video-player/player";
import { useState, useEffect } from "react";
import useSWR from "swr";
type Video = {
  id: number;
  play_url: string;
  description: string;
};
const url = "http://47.106.228.5:9133/v1/main/videos?category_id=1";
const Page = () => {
  const [videos, setVideos] = useState<Video[]>([]);
  const [index, setIndex] = useState<number>(0);
  console.log("component load");
  useEffect(() => {
    let ignore = false;
    const clean = () => {
      ignore = true;
      window.removeEventListener("scroll", handleScroll);
    };
    // 组件挂载后添加滚动事件监听器
    window.addEventListener("scroll", handleScroll);

    console.log("useEffect called");

    fetch(url)
      .then((response) => response.json())
      .then((data) => {
        if (!ignore) {
          setVideos(data);
          setIndex(0);
        }
      });

    // 组件卸载时移除滚动事件监听器
    return clean;
  }, []); // 注意这里的空数组，这确保了 useEffect 只在组件挂载时运行

  const fetchVideos = async () => {
    console.log("fetchVideos called");
    fetch(url)
      .then((response) => response.json())
      .then((data) => {
        setVideos(data);
        setIndex(0);
      });
  };
  const nextVideos = () => {
    if (index < videos.length - 1) {
      setIndex((index) => index + 1);
    } else {
      fetchVideos();
    }
  };
  const handleScroll = () => {
    console.log("Window Scrolled", window.scrollY);
    nextVideos();
    // 你可以在这里添加更多逻辑 (例如，检测滚动位置等)
  };

  return (
    <>
      {videos.length > 0 && index >= 0 && index < videos.length ? (
        <>
          <h2 className="text-2xl text-purple-600">{videos[index].id}</h2>
          <div>{videos[index].description}</div>
          <div className="w-6/12">
            <PlyrComponent url={videos[index].play_url} />
          </div>
        </>
      ) : (
        <p>Loading or invalid index...</p>
      )}
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
        onClick={fetchVideos}
      >
        Reload Videos
      </button>
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
        onClick={nextVideos}
      >
        Next Video
      </button>
    </>
  );
};

export default Page;
