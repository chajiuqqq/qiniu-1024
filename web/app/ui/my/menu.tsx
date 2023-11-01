// Menu.tsx
import React from "react";

interface MenuProps {
  index: number;
  setIndex: (index: number) => void;
}

const Menu: React.FC<MenuProps> = ({ index, setIndex }) => {
  return (
    <div className="flex">
      {["作品", "点赞", "收藏"].map((menuItem, idx) => (
        <div
          key={idx}
          onClick={() => setIndex(idx)}
          className={`cursor-pointer p-4 ${
            index === idx ? "text-blue-500 border-b-2 border-blue-500" : ""
          }`}
        >
          {menuItem}
        </div>
      ))}
    </div>
  );
};

export default Menu;
