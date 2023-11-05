import { User } from "@/app/lib/api/types";
import React from "react";
// Name 组件
const ProfileInfo: React.FC<User> = ({followers,follows,user_likes}) => {

  return (
    <div className="flex items-center space-x-4 w-full">
      <div className="flex space-x-2">
        <span>关注</span>
        <span>{follows?.length}</span>
      </div>
      <div className="flex space-x-2">
        <span>粉丝</span>
        <span>{followers?.length}</span>
      </div>
      <div className="flex space-x-2">
        <span>获赞</span>
        <span>{user_likes?.length}</span>
      </div>
    </div>
  );
};

export default ProfileInfo;
