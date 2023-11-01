"use client";
import {
  HeartIcon,
  StarIcon,
  ChatBubbleBottomCenterTextIcon,
  ShareIcon,
} from "@heroicons/react/24/outline";
const LikeIcon = () => {
  return <HeartIcon className="w-6 text-white"></HeartIcon>;
};
const MyLikeIcon = () => {
  return <HeartIcon className="w-6 text-white"></HeartIcon>;
};

const MyCommentIcon = () => {
  return (
    <ChatBubbleBottomCenterTextIcon className="w-6 text-white"></ChatBubbleBottomCenterTextIcon>
  );
};

const MyCollectionIcon = () => {
  return <StarIcon className="w-6 text-white"></StarIcon>;
};

const MyShareIcon = () => {
  return <ShareIcon className="w-6 text-white"></ShareIcon>;
};

export { LikeIcon, MyLikeIcon, MyCommentIcon, MyCollectionIcon, MyShareIcon };
