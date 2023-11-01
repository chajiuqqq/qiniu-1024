import React from "react";

// Name 组件
const ProfileInfo: React.FC = () => {
  return (
    <div className="flex items-center space-x-4 w-full">
      <div className="flex space-x-2">
        <span>关注</span>
        <span>100</span>
      </div>
      <div className="flex space-x-2">
        <span>粉丝</span>
        <span>100</span>
      </div>
      <div className="flex space-x-2">
        <span>获赞</span>
        <span>100</span>
      </div>
    </div>
  );
};

export default ProfileInfo;
