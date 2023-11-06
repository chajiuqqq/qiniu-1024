'use client';
// components/SearchBar.tsx
import React, { useState, FormEvent } from 'react';

import { MagnifyingGlassIcon } from '@heroicons/react/24/solid'; // 从heroicons中引入图标
interface SearchBarProps {
  onSearch: (searchTerm: string) => void;  // 定义传入的onSearch函数的类型
}

const SearchBar: React.FC<SearchBarProps> = ({ onSearch }) => {
  const [searchTerm, setSearchTerm] = useState('');

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    onSearch(searchTerm);  // 触发传入的搜索处理函数
  };

  return (
    <div className="flex border-2 rounded-lg pl-5">
      <input
        type="text"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        className="p-2 flex-grow outline-none border-none bg-transparent"
        placeholder="coming soon..."
      />
      <button onClick={handleSubmit} className="p-2">
            <MagnifyingGlassIcon className="h-6 w-6 text-gray-500" />
      </button>
    </div>
  );
};

export default SearchBar;
