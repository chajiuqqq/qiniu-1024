'use client'
import Profile from "@/app/ui/my/profile";
import Menu from "../ui/my/menu";
import { useState, useReducer, useEffect } from "react";
import VideoItemList from "../ui/video/list";
import { initalVideos } from "../lib/data";
import { ProfileTab } from "../lib/const";
import Popup from "../ui/Popup";
import VideoPlayerComponent from "../ui/VideoPlayerComponent";
import { VideoType } from "../lib/video";
import { flushSync } from "react-dom";
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
  const [videos, setVideos] = useState<VideoType[]>(initalVideos)
  const [startedVideoID, setStartedVideoID] = useState<number>(-1)
  const myUrl = "http://47.106.228.5:9133/v1/main/videos?category_id=1";
  const dev = true
  useEffect(() => {
    let ignore = false;
    if (!dev) {
      fetch(myUrl)
        .then((response) => response.json())
        .then((data) => {
          if (!ignore) {
            setVideos(data);
          }
        });
    }
    return () => {
      ignore = true;
    };
  }, []);
  const handleUpdateVideos = () => {
    if (dev) {
        setVideos(initalVideos)
    } else {
      fetch(myUrl)
        .then((response) => response.json())
        .then((data) => {
          setVideos(data)
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
      <div className="p-4">
        <VideoItemList videos={initalVideos} type={getProfileType(menuIndex)} onClick={handleVideoItemClick} />
      </div>
      {popStatue.isPopupVisible && (
        <Popup onClose={closePopup}>
          <VideoPlayerComponent videos={videos} updateVideos={handleUpdateVideos} startedVideoID={startedVideoID}></VideoPlayerComponent>
          {/* <h1>hello</h1> */}
        </Popup>
      )}
    </>
  );
};

export default My;
