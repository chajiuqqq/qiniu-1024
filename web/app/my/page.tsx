'use client'
import Profile from "@/app/ui/my/profile";
import Menu from "../ui/my/menu";
import { useState } from "react";
import VideoItemList from "../ui/video/list";
import { initalVideos } from "../lib/data";
import { ProfileTab } from "../lib/const";

const getProfileType = function (index: number): ProfileTab {
  switch (index) {
    case 0:
      return ProfileTab.My
    case 1:
      return ProfileTab.Likes
    case 2:
      return ProfileTab.Collection
    default:
      return ProfileTab.My
  }
}
const My = () => {
  const [menuIndex, setMenuIndex] = useState(0);
  return (
    <>
      <Profile></Profile>
      <Menu index={menuIndex} setIndex={setMenuIndex} />
      <div className="p-4">
        <VideoItemList videos={initalVideos} type={getProfileType(menuIndex)} />
      </div>
    </>
  );
};

export default My;
