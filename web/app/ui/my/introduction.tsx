import React from "react";
import EditIcon from "./edit";
// 定义传入的 props 类型
interface IntroductionProps {
  introduction: string;
  onEdit: () => void;
}

const Introduction: React.FC<IntroductionProps> = ({
  introduction,
  onEdit,
}) => {
  return (
    <div className="flex items-center">
      <p className="m-4 text-grey-100">{introduction}</p>
      <button onClick={onEdit} className="text-blue-500 hover:text-blue-700">
        <EditIcon></EditIcon>
      </button>
    </div>
  );
};

export default Introduction;
