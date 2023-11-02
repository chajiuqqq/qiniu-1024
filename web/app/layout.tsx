"use client";
import "@/app/ui/global.css";
import SideNav from "@/app/ui/dashboard/sidenav";
import TopNav from "@/app/ui/topnav/topnav";
import { UserProvider } from "./lib/contexts/UserContext";
// 搜索处理函数
const handleSearch = (searchTerm: string) => {
  console.log("执行搜索:", searchTerm);
  // 在这里添加具体的搜索逻辑
};

// 用户状态数据（示例）
const userStatusData = {
  avatarUrl: "/avatar.jpg", // 实际的头像路径
  worksCount: 10, // 示例数据
  likesCount: 20, // 示例数据
  followersCount: 15, // 示例数据
};
export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body>
        <UserProvider>
        <div className="flex h-screen flex-col md:flex-row md:overflow-hidden">
          <div className="w-full flex-none md:w-64">
            <SideNav />
          </div>
          <div className="flex h-full flex-col w-full space-y-2 p-5">
            <TopNav onSearch={handleSearch} userStatus={userStatusData} />
            <div className="flex-grow rounded-md">{children}</div>
          </div>
        </div>
        </UserProvider>
      </body>
    </html>
  );
}
