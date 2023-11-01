import React from "react";
import EditIcon from "./edit";
// 定义传入的 props 类型
interface NameProps {
  name: string;
  onEdit: () => void;
}

// Name 组件
const Name: React.FC<NameProps> = ({ name, onEdit }) => {
  return (
    <div className="flex items-center">
      <p className="mr-4 text-xl font-semibold">{name}</p>
      <button onClick={onEdit} className="text-blue-500 hover:text-blue-700">
        <EditIcon></EditIcon>
      </button>
    </div>
  );
};

export default Name;
