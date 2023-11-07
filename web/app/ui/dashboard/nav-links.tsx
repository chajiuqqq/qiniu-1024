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
import React, { useEffect, useState } from "react";
import api from "@/app/lib/api/api-client";
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
  // { name: "喜欢", href: "/likes", icon: HeartIcon },
  // { name: "收藏", href: "/collection", icon: StarIcon },
  // { name: "关注", href: "/follow", icon: UserPlusIcon },
];
type navLink = {
name:string,
href:string,
order:number,
icon?:any
}

export default function NavLinks() {
  const pathname = usePathname();
  const [cateLinks,setCateLinks] = useState<navLink[]>()
  useEffect(()=>{
    if(!cateLinks){
      api.category.getCategories().then(res=>{
        if (!cateLinks && res.data){
          let links:navLink[] = []
          for (let index = 0; index < res.data.length; index++) {
            const e = res.data[index]
            links.push({
              name:e.name,
              href:`/category/${e.id}`,
              order:e.order,
              icon: EllipsisHorizontalIcon
            })
          }
          links.sort((a,b)=>{
            return a.order-b.order
          })
          setCateLinks(links)
        }
      })
    }
    
  },[])
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

      {cateLinks && cateLinks.map((link) => {
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
