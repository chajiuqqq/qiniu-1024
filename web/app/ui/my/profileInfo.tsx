import React from "react";
import {UserType} from '@/app/lib/contexts/UserContext'
// Name 组件
const ProfileInfo: React.FC<UserType> = ({followCnt,followerCnt,userLikesCnt}) => {

  return (
    <div className="flex items-center space-x-4 w-full">
      <div className="flex space-x-2">
        <span>关注</span>
        <span>{followCnt}</span>
      </div>
      <div className="flex space-x-2">
        <span>粉丝</span>
        <span>{followerCnt}</span>
      </div>
      <div className="flex space-x-2">
        <span>获赞</span>
        <span>{userLikesCnt}</span>
      </div>
    </div>
  );
};

export default ProfileInfo;
