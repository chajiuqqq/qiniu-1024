"use client";
import React, { useEffect, useState } from "react";
import VideoPlayerComponent from "@/app/ui/VideoPlayerComponent";
import { VideoType } from "./lib/video";
import { initalVideos } from "./lib/data";
import Loading from "./ui/loading";
const Page = () => {
  const [videos, setVideos] = useState<VideoType[]>(initalVideos)
  const hotUrl = "http://47.106.228.5:9133/v1/main/videos?category_id=1";
  const dev = true
  useEffect(() => {
    let ignore = false;
    if (!dev) {
      fetch(hotUrl)
        .then((response) => response.json())
        .then((data) => {
          if (!ignore) {
            setVideos(data);
          }
        });
    }
    return () => {
      ignore = true;
    };
  }, []);
  const handleUpdateVideos = () => {
    if (dev) {
      setVideos(initalVideos)
    } else {
      fetch(hotUrl)
        .then((response) => response.json())
        .then((data) => {
          setVideos(data)
        });
    }
  }
  return (
    <>
    {/* <Loading></Loading>  */}
      <VideoPlayerComponent videos={videos} updateVideos={handleUpdateVideos}></VideoPlayerComponent>
    </>
  );
};

export default Page;
