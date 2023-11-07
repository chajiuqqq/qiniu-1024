"use client";
import PlyrComponent, { PlyrAttach } from "@/app/ui/video-player/player";
import React, { useState, useEffect } from "react";
import Loading from "./loading";
import { MainVideoItem, User } from "../lib/api/types";
import api from "../lib/api/api-client";
import { useUser } from "../lib/contexts/UserContext";

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
  const {user,setUser} = useUser()
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
    const vid = vs[index].id
    vs[index].likes_count += 1;
    vs[index].liked = true;
    setVideos(vs);
    setUser(u=>{
      if(u){
        const tmp:User = {
          ...u,
          collections:[...u.collections],
          likes:[...u.likes,{
            video_id:vid,
            created_at:''
          }]
        }
        return tmp
      }
      return u
    })
  };
  const handleCancelLike = () => {
    let vs = [...videos];
    const vid = vs[index].id
    vs[index].likes_count -= 1;
    vs[index].liked = false;
    setVideos(vs);
    setUser(u=>{
      if(u){
        const tmp:User = {
          ...u,
          collections:[...u.collections],
          likes: u.likes.filter(v=>{
            v.video_id!=vid
          })
        }
        return tmp
      }
      return u
    })
  };
  const handleCollect = () => {
    let vs = [...videos];
    vs[index].collect_count += 1;
    vs[index].collected = true;
    const vid = vs[index].id
    setVideos(vs);
    setUser(u=>{
      if(u){
        const tmp:User = {
          ...u,
          collections:[...u.collections,{
            video_id:vid,
            created_at:''
          }],
          likes:[...u.likes]
        }
        return tmp
      }
      return u
    })
  };
  const handleCancelCollect = () => {
    let vs = [...videos];
    vs[index].collect_count -= 1;
    vs[index].collected = false;
    const vid = vs[index].id
    setVideos(vs);
    setUser(u=>{
      if(u){
        const tmp:User = {
          ...u,
          collections:u.collections.filter(v=>{
            v.video_id!=vid
          }),
          likes: [...u.likes]
        }
        return tmp
      }
      return u
    })
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
