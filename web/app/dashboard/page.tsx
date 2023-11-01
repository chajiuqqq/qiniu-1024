"use client";
import PlyrComponent from "@/app/ui/video-player/player";
import { useState, useEffect } from "react";

type Video = {
  id: number;
  play_url: string;
  description: string;
};
const url = "http://47.106.228.5:9133/v1/main/videos?category_id=1";
let dev = true
const initalVideos = [
  {
    "id": 100000015,
    "number": 15,
    "user_id": 100001,
    "category_id": 1,
    "category": "旅游",
    "play_url": "http://cdn.chajiuqqq.cn/100000015.mp4",
    "cover_url": "http://cdn.chajiuqqq.cn/100000015_cover.jpg",
    "description": "",
    "play_count": 0,
    "likes_count": 0,
    "collect_count": 0,
    "comments": null,
    "status": "OnShow",
    "cover_status": "Success",
    "is_deleted": false,
    "created_at": "2023-10-28T13:23:47.64Z",
    "updated_at": "2023-10-28T13:23:47.64Z",
    "score": 0
  },
  {
    "id": 100000016,
    "number": 16,
    "user_id": 100001,
    "category_id": 1,
    "category": "旅游",
    "play_url": "http://cdn.chajiuqqq.cn/100000016.mp4",
    "cover_url": "http://cdn.chajiuqqq.cn/100000016_cover.jpg",
    "description": "",
    "play_count": 0,
    "likes_count": 0,
    "collect_count": 0,
    "comments": null,
    "status": "OnShow",
    "cover_status": "Success",
    "is_deleted": false,
    "uploaded_at": "2023-10-28T13:26:57.467Z",
    "cover_uploaded_at": "2023-10-28T13:27:11.098Z",
    "created_at": "2023-10-28T13:26:56.479Z",
    "updated_at": "2023-10-28T13:27:11.14Z",
    "score": 0
  },
  {
    "id": 100000017,
    "number": 17,
    "user_id": 100001,
    "category_id": 1,
    "category": "旅游",
    "play_url": "http://cdn.chajiuqqq.cn/100000017.mp4",
    "cover_url": "http://cdn.chajiuqqq.cn/100000017_cover.jpg",
    "description": "热门爆款推荐",
    "play_count": 0,
    "likes_count": 0,
    "collect_count": 0,
    "comments": null,
    "status": "OnShow",
    "cover_status": "Success",
    "is_deleted": false,
    "uploaded_at": "2023-10-28T13:52:36.606Z",
    "cover_uploaded_at": "2023-10-28T13:52:42.708Z",
    "created_at": "2023-10-28T13:52:04.148Z",
    "updated_at": "2023-10-28T14:04:54.512Z",
    "score": 0
  }
]
const Page = () => {
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
        case 'ArrowUp':
          console.log('上箭头键被按下了');
          lastVideo()
          break;
        case 'ArrowDown':
          console.log('下箭头键被按下了');
          nextVideo()
          break;
        default:
          // 其他按键可以在此处理
          break;
      }
    };

    // 在组件挂载时添加事件监听器
    window.addEventListener('keydown', handleKeyDown);
    // 组件卸载时移除滚动事件监听器
    return () => {
      window.removeEventListener('keydown', handleKeyDown);
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
      ignore = true
    };

  },[])
  return (
    <>
      {videos.length > 0 && index >= 0 && index < videos.length ? (
        <>
          <div className="w-full">
            <PlyrComponent url={videos[index].play_url} desc={videos[index].description} />
          </div>
        </>
      ) : (
        <p>Loading or invalid index...</p>
      )}
    </>
  );
};

export default Page;
