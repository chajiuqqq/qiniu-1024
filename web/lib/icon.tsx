import {
  HeartIcon,
  StarIcon,
  ChatBubbleBottomCenterTextIcon,
  ShareIcon,
  PlayIcon,
  XMarkIcon,
  ArrowPathIcon,
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
const MyPlayIcon = () => {
  return <PlayIcon className="w-6 text-white"></PlayIcon>;
};

const MyCloseIcon = () => {
  return <XMarkIcon className="w-6 text-white hover:text-blue-700"></XMarkIcon>;
};
const MyArrowPathIcon = () => {
  return <ArrowPathIcon className="w-6 text-white"></ArrowPathIcon>;
};

export default () => {
  return <ArrowPathIcon className="w-6 text-white"></ArrowPathIcon>;
};
export {
  LikeIcon,
  MyLikeIcon,
  MyCommentIcon,
  MyCollectionIcon,
  MyShareIcon,
  MyPlayIcon,
  MyCloseIcon,
};
