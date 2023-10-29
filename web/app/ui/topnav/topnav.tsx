'use client';
// components/TopNav.tsx
import React from 'react';
import SearchBar from './searchbar'; // 请根据实际情况调整路径
import UserStatusBar, { UserStatusBarProps } from './user-status-bar'; // 调整路径

interface TopNavProps {
  onSearch: (searchTerm: string) => void;
  userStatus: UserStatusBarProps;
}

const TopNav: React.FC<TopNavProps> = ({ onSearch, userStatus }) => {
  return (
    <div className="flex justify-between items-center shadow-md w-full px-5 rounded bg-gray-50">
      {/* 搜索栏组件，放在左边 */}
      <SearchBar onSearch={onSearch} />

      {/* 用户状态栏，放在右边 */}
      <UserStatusBar {...userStatus} />
    </div>
  );
};

export default TopNav;
