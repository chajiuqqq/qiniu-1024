"use client";
import React, { useEffect, useState } from "react";
import VideoPlayerComponent from "@/app/ui/VideoPlayerComponent";
import Loading from "./ui/loading";
import { MainVideoItem, Video } from "./lib/api/types";
import api from "./lib/api/api-client";
import { useUser } from "./lib/contexts/UserContext";
import { redirect } from "next/navigation";
const Page = () => {
  const [videos, setVideos] = useState<MainVideoItem[] | undefined>()
  const [loading, setLoading] = useState(true)
  const {user} = useUser()
  const dev = false
  if (!user){
    redirect('/login')
  }
  useEffect(() => {
    let ignore = false;
    if (!dev) {
      api.video.getVideos()
        .then((res) => {
          if (!ignore) {
            setLoading(false)
            setVideos(res.data);
          }
        }).catch((err)=>{
          console.log(err)
          throw err
        });
    }
    return () => {
      ignore = true;
    };
  }, []);
  const handleUpdateVideos = () => {
    setLoading(true)
    if (dev) {
    } else {
      api.video.getVideos()
        .then((res) => {
          setLoading(false)
          setVideos(res.data);
        });
    }
  }
  return (
    <>
      {loading ? (<Loading></Loading>) : ''}
      {videos ? (
        <VideoPlayerComponent videos={videos} updateVideos={handleUpdateVideos}></VideoPlayerComponent>) : ''}
    </>
  );
};

export default Page;
