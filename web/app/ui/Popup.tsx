// components/Popup.tsx
import React from "react";
import { MyCloseIcon } from "../lib/icon";
interface PopupProps {
  onClose: () => void;
  children: React.ReactNode;
}

const Popup: React.FC<PopupProps> = ({ onClose, children }) => {
  return (
    <div className="fixed top-0 left-0 w-full h-full bg-gray-500 bg-opacity-50 flex justify-center items-center">
      <div className="bg-white p-4 rounded-md relative w-11/12">
        {children}
        <button onClick={onClose} className=" absolute top-5 right-5 shadow-md">
          <MyCloseIcon></MyCloseIcon>
        </button>
      </div>
    </div>
  );
};

export default Popup;
