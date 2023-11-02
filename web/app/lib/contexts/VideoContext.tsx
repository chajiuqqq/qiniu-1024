// contexts/UserContext.tsx
import { createContext, useContext, ReactNode, useState } from 'react';
import { VideoType } from '../video';
import { initalVideos } from '../data';

type VideosContextType = {
  videos: VideoType[];
  setVideos: React.Dispatch<React.SetStateAction<VideoType[]>>;
};

const VideosContext = createContext<VideosContextType | undefined>(undefined);

export const VideoProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [videos, setVideos] = useState<VideoType[]>(initalVideos);

  return (
    <VideosContext.Provider value={{ videos, setVideos }}>
      {children}
    </VideosContext.Provider>
  );
};

export const useVideos = () => {
  const context = useContext(VideosContext);
  if (context === undefined) {
    throw new Error('useUser must be used within a UserProvider');
  }
  return context;
};
