"use client";
// components/TopNav.tsx
import React from "react";
import SearchBar from "./searchbar"; // 请根据实际情况调整路径
import UserStatusBar from "./user-status-bar"; // 调整路径

interface TopNavProps {
  onSearch: (searchTerm: string) => void;
}

const TopNav: React.FC<TopNavProps> = ({ onSearch }) => {
  return (
    <div className="flex justify-between items-center w-full">
      {/* 搜索栏组件，放在左边 */}
      <SearchBar onSearch={onSearch} />

      {/* 用户状态栏，放在右边 */}
      <UserStatusBar />
    </div>
  );
};

export default TopNav;
