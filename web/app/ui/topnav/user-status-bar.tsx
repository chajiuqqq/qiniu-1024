// components/UserStatusBar.tsx
import React from "react";
import {
  VideoCameraIcon,
  HeartIcon,
  StarIcon,
} from "@heroicons/react/24/solid"; // 从heroicons中引入图标
import { useUser } from "@/app/lib/contexts/UserContext";
import Link from "next/link";

const UserStatusBar: React.FC = () => {
  const { user, setUser } = useUser()

  return (
    <div className="flex items-center space-x-4">
      {/* 用户头像 */}
      {/* <img
        src={user.avatar_url}
        alt="用户头像"
        className="h-10 w-10 rounded-full object-cover"
      /> */}
      {
        user ? (
          <>
            <div>{user.name}</div>
            <div className="flex flex-grow items-center justify-around space-x-4">

              <div className="flex items-center space-x-2">
                <VideoCameraIcon className="h-6 w-6 text-gray-500" />
                <span>{`作品`}</span>
              </div>

              <div className="flex items-center space-x-2">
                <HeartIcon className="h-6 w-6 text-gray-500" />
                <span>{`喜欢 ${user.likes?.length}`}</span>
              </div>

              <div className="flex items-center space-x-2">
                <StarIcon className="h-6 w-6 text-gray-500" />
                <span>{`收藏 ${user.collections?.length}`}</span>
              </div>
            </div>
          </>

        ) : (
          <><Link href='/register'>
            <div className="cursor-pointer w-24 flex justify-center text-center border text-sm font-medium border-indigo-400 rounded-md text-indigo-400 py-2 px-4 hover:text-indigo-600 hover:border-indigo-600">
              创建
            </div>
          </Link>
            <Link href='/login'>
              <div className=" w-32 flex justify-center text-center rounded-lg px-4 py-2 text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                登录
              </div>
            </Link>
          </>
        )
      }
    </div>
  );
};

export default UserStatusBar;
