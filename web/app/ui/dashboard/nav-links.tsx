"use client";
import {
  UserGroupIcon,
  HomeIcon,
  UserIcon,
  HeartIcon,
  StarIcon,
  UserPlusIcon,
  EllipsisHorizontalIcon,
  ArrowUpTrayIcon
} from "@heroicons/react/24/outline";
import Link from "next/link";
import { usePathname } from "next/navigation";
import clsx from "clsx";
// Map of links to display in the side navigation.
// Depending on the size of the application, this would be stored in a database.
const links = [
  { name: "发布视频", href: "/publish", icon: ArrowUpTrayIcon ,highlight:true},
  { name: "热门", href: "/", icon: HomeIcon },
  {
    name: "我的",
    href: "/my",
    icon: UserIcon,
  },
  { name: "喜欢", href: "/likes", icon: HeartIcon },
  { name: "收藏", href: "/collection", icon: StarIcon },
  { name: "关注", href: "/follow", icon: UserPlusIcon },
];
const links2 = [
  { name: "旅游", href: "/dashboard/customers", icon: EllipsisHorizontalIcon },
  { name: "美食", href: "/dashboard/customers", icon: EllipsisHorizontalIcon },
  { name: "户外", href: "/dashboard/customers", icon: EllipsisHorizontalIcon },
];

export default function NavLinks() {
  const pathname = usePathname();
  return (
    <>
      {links.map((link) => {
        const LinkIcon = link.icon;
        return (
          <Link
            key={link.name}
            href={link.href}
            className={clsx(
              "flex h-[48px] grow items-center justify-center gap-2 rounded-md bg-gray-50 p-3 text-sm font-medium hover:bg-sky-100 hover:text-blue-600 md:flex-none md:justify-start md:p-2 md:px-3",
              {
                "bg-sky-100 text-blue-600": pathname === link.href,
              }
            )}
          >
            <LinkIcon className="w-6" />
            <p className="hidden md:block">{link.name}</p>
          </Link>
        );
      })}
      <br />

      {links2.map((link) => {
        const LinkIcon = link.icon;
        return (
          <Link
            key={link.name}
            href={link.href}
            className={clsx(
              "flex h-[48px] grow items-center justify-center gap-2 rounded-md bg-gray-50 p-3 text-sm font-medium hover:bg-sky-100 hover:text-blue-600 md:flex-none md:justify-start md:p-2 md:px-3",
              {
                "bg-sky-100 text-blue-600": pathname === link.href,
              }
            )}
          >
            <LinkIcon className="w-6" />
            <p className="hidden md:block">{link.name}</p>
          </Link>
        );
      })}
    </>
  );
}
