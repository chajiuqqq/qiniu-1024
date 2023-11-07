'use client'
import Profile from "@/app/ui/my/profile";
import Menu from "../ui/my/menu";
import { useState, useReducer, useEffect } from "react";
import VideoItemList from "../ui/video/list";
import { ProfileTab } from "../lib/const";
import Popup from "../ui/Popup";
import VideoPlayerComponent from "../ui/VideoPlayerComponent";
import { MainVideoItem } from "../lib/api/types";
import api from "../lib/api/api-client";
import { useUser } from "../lib/contexts/UserContext";
import Loading from "../ui/loading";
const getProfileType = function (index: number): ProfileTab {
  switch (index) {
    case 0:
      return ProfileTab.My
    case 1:
      return ProfileTab.Likes
    case 2:
      return ProfileTab.Collection
    default:
      return ProfileTab.My
  }
}

type State = {
  isPopupVisible: boolean;
};

type Action =
  | { type: 'SHOW_POPUP' }
  | { type: 'CLOSE_POPUP' };

const initialState: State = {
  isPopupVisible: false,
};

const popReducer = (state: State, action: Action): State => {
  switch (action.type) {
    case 'SHOW_POPUP':
      return { ...state, isPopupVisible: true };
    case 'CLOSE_POPUP':
      return { ...state, isPopupVisible: false };
    default:
      return state;
  }
};
const My = () => {
  const [menuIndex, setMenuIndex] = useState(0);
  const [popStatue, popDispatch] = useReducer(popReducer, initialState);
  const [videos, setVideos] = useState<MainVideoItem[]>()
  const [startedVideoID, setStartedVideoID] = useState<number>(-1)
  const dev = false
  const { user } = useUser()
  const [loading, setLoading] = useState(true)
  useEffect(() => {
    let ignore = false;
    if (!dev) {
      const q = user ? { user_id: user.id } : undefined
      api.video.getVideos(q)
        .then((res) => {
          if (!ignore) {
            setLoading(false)
            setVideos(res.data);
          }
        });
    }
    return () => {
      ignore = true;
    };
  }, []);
  const handleUpdateVideos = () => {
    setLoading(true)
    if (dev) {
    } else {
      api.video.getVideos()
        .then((res) => {
          setLoading(false)
          setVideos(res.data);
        });
    }
  }
  const showPopup = () => {
    popDispatch({ type: 'SHOW_POPUP' });
  };

  const closePopup = () => {
    popDispatch({ type: 'CLOSE_POPUP' });
  };
  const handleVideoItemClick: (videoID: number) => void = (videoID) => {
    setStartedVideoID(videoID)
    showPopup()
  }
  return (
    <>
      <Profile></Profile>
      <Menu index={menuIndex} setIndex={setMenuIndex} />
      {videos ?(
        <div className="p-4">
          <VideoItemList videos={videos} type={getProfileType(menuIndex)} onClick={handleVideoItemClick} />
        </div>):(
          <>
          <Loading></Loading>
          </>
        )
      }
      {popStatue.isPopupVisible && videos && (
        <Popup onClose={closePopup}>
          <VideoPlayerComponent videos={videos} setVideos={setVideos} updateVideos={handleUpdateVideos} startedVideoID={startedVideoID}></VideoPlayerComponent>
        </Popup>
      )}
    </>
  );
};

export default My;
