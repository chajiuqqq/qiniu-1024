// AutoDismissAlert.tsx
import React, { useEffect, useState } from 'react';

interface AutoDismissAlertProps {
  message: string;
}

const AutoDismissAlert: React.FC<AutoDismissAlertProps> = ({ message }) => {
  const [visible, setVisible] = useState(true);

  useEffect(() => {
    const timer = setTimeout(() => setVisible(false), 3000);
    return () => clearTimeout(timer);
  }, []);

  if (!visible) return null;

  return (
    <div className="fixed  top-10 left-1/2 flex items-center justify-center z-50">
      <div className="bg-white p-5 rounded shadow-lg">
        <div className="flex justify-between items-start space-x-2">
          <div className="text-gray-700">{message}</div>
          <button onClick={() => setVisible(false)} className="text-gray-700 hover:text-gray-900">
            ×
          </button>
        </div>
      </div>
    </div>
  );
};

export default AutoDismissAlert;
