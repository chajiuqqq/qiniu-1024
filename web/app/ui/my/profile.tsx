
import Name from "@/app/ui/my/name";
import ProfileInfo from "@/app/ui/my/profileInfo";
import Introduction from "@/app/ui/my/introduction";

const Profile = () => {
  const handleEditName = () => { };
  const handleEditIntroduction = () => { };
  return (
    <div className="w-full">
      <div className="flex">
        <img
          src="/avatar.jpg"
          alt="Circular Image"
          className="h-24 w-24 rounded-full object-cover"
        />
        <div className="flex flex-col justify-center space-y-2">
          <Name name="李四" onEdit={handleEditName}></Name>
          <ProfileInfo></ProfileInfo>
        </div>
      </div>
      <div>
        <Introduction
          introduction="我是个活泼开朗的小孩"
          onEdit={handleEditIntroduction}
        ></Introduction>
      </div>
    </div>
  );
};

export default Profile;
