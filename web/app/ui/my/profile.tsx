
import Name from "@/app/ui/my/name";
import ProfileInfo from "@/app/ui/my/profileInfo";
import Introduction from "@/app/ui/my/introduction";
import { useUser } from "@/app/lib/contexts/UserContext";
import { redirect } from "next/navigation";

const Profile = () => {
  const handleEditName = () => { };
  const handleEditIntroduction = () => { };
  const { user } = useUser()
  if (!user){
    redirect("/login")
  }
  return (
    <>
          <div className="w-6/12 border border-gray-200 rounded-md shadow shadow-gray-50">
            <div className="flex">
              <img
                src={user.avatar_url}
                alt="Circular Image"
                className="h-24 w-24 rounded-full object-cover"
              />
              <div className="flex flex-col justify-center space-y-2">
                <Name name={user.name||''} onEdit={handleEditName}></Name>
                <ProfileInfo {...user}></ProfileInfo>
              </div>
            </div>
            <div>
              <Introduction
                introduction={user.description||''}
                onEdit={handleEditIntroduction}
              ></Introduction>
            </div>
          </div>
        </>
  )
    
    
};

export default Profile;
