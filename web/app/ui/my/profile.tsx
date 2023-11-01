"use client";
// LoginComponent.tsx
import React, { useState, FormEvent } from "react";
import Name from "@/app/ui/my/name";
import ProfileInfo from "@/app/ui/my/profileInfo";
import Introduction from "@/app/ui/my/introduction";
import Menu from "@/app/ui/my/menu";
import VideoItemList from "@/app/ui/video/list";
const videos = [
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  { imgUrl: "http://cdn.chajiuqqq.cn/100000016_cover.jpg" },
  // ... 其他视频
];

const LoginComponent: React.FC = () => {
  const [menuIndex, setMenuIndex] = useState(0);
  const handleEditName = () => {};
  const handleEditIntroduction = () => {};
  return (
    <div className="w-full">
      <div className="flex">
        <img
          src="/avatar.jpg"
          alt="Circular Image"
          className="h-24 w-24 rounded-full object-cover"
        />
        <div className="flex flex-col justify-center space-y-2">
          <Name name="李四" onEdit={handleEditName}></Name>
          <ProfileInfo></ProfileInfo>
        </div>
      </div>
      <div>
        <Introduction
          introduction="我是个活泼开朗的小孩"
          onEdit={handleEditIntroduction}
        ></Introduction>
      </div>
      <Menu index={menuIndex} setIndex={setMenuIndex} />
      <div className="p-4">
        <VideoItemList videos={videos} />
      </div>
    </div>
  );
};

export default LoginComponent;
