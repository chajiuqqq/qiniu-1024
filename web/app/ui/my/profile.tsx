import Name from "@/app/ui/my/name";
import ProfileInfo from "@/app/ui/my/profileInfo";
import Introduction from "@/app/ui/my/introduction";
import { useUser } from "@/app/lib/contexts/UserContext";
import { redirect } from "next/navigation";
import Link from "next/link";

const Profile = () => {
  const handleEditName = () => {};
  const handleEditIntroduction = () => {};
  const { user } = useUser();
  if (!user) {
    redirect("/login");
  }
  return (
    <>
      <div className="w-6/12 border border-gray-200 rounded-md shadow shadow-gray-50 p-5">
        <div className="flex flex-col justify-center space-y-2">
          <Name name={user.name || ""} onEdit={handleEditName}></Name>
          <ProfileInfo {...user}></ProfileInfo>
        </div>
        <div>
          <Introduction
            introduction={user.description || ""}
            onEdit={handleEditIntroduction}
          ></Introduction>
        </div>
      </div>
    </>
  );
};

export default Profile;
