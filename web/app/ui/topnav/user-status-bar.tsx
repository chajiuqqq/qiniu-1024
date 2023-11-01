// components/UserStatusBar.tsx
import React from "react";
import {
  VideoCameraIcon,
  HeartIcon,
  StarIcon,
} from "@heroicons/react/24/solid"; // 从heroicons中引入图标

export interface UserStatusBarProps {
  // 定义传入的props类型
  avatarUrl: string; // 用户头像URL
  worksCount: number; // “我的作品”数量
  likesCount: number; // “我的喜欢”数量
  followersCount: number; // “我的粉丝”数量
}

const UserStatusBar: React.FC<UserStatusBarProps> = ({
  avatarUrl,
  worksCount,
  likesCount,
  followersCount,
}) => {
  return (
    <div className="flex items-center space-x-4 p-4">
      {/* 用户头像 */}
      <img
        src={avatarUrl}
        alt="用户头像"
        className="h-10 w-10 rounded-full object-cover"
      />

      {/* 用户信息栏 */}
      <div className="flex flex-grow items-center justify-around space-x-4">
        {/* 我的作品 */}
        <div className="flex items-center space-x-2">
          <VideoCameraIcon className="h-6 w-6 text-gray-500" />
          <span>{`作品 ${worksCount}`}</span>
        </div>

        {/* 我的喜欢 */}
        <div className="flex items-center space-x-2">
          <HeartIcon className="h-6 w-6 text-gray-500" />
          <span>{`喜欢 ${likesCount}`}</span>
        </div>

        {/* 我的收藏 */}
        <div className="flex items-center space-x-2">
          <StarIcon className="h-6 w-6 text-gray-500" />
          <span>{`收藏 ${followersCount}`}</span>
        </div>
      </div>
    </div>
  );
};

export default UserStatusBar;
