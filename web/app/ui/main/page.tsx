"use client";
import React, { useEffect, useState } from "react";
import VideoPlayerComponent from "@/app/ui/VideoPlayerComponent";
import Loading from "@/app/ui/loading";
import { MainVideoItem, Video, VideoQuery } from "@/app/lib/api/types";
import api from "@/app/lib/api/api-client";
import { useUser } from "@/app/lib/contexts/UserContext";
import { redirect } from "next/navigation";
const Main = (query?:VideoQuery) => {
  const [videos, setVideos] = useState<MainVideoItem[]>()
  const [loading, setLoading] = useState(true)
  const {user} = useUser()
  const dev = false
  if (!user){
    redirect('/login')
  }
  useEffect(() => {
    let ignore = false;
    if (!dev) {
      api.video.getVideos(query)
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
      api.video.getVideos(query)
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
        <VideoPlayerComponent videos={videos} setVideos={setVideos} updateVideos={handleUpdateVideos}></VideoPlayerComponent>) : ''}
    </>
  );
};

export default Main;
