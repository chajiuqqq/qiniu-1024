"use client";
import React, { useEffect, useState } from "react";
import VideoPlayerComponent from "@/app/ui/VideoPlayerComponent";
import Main from '@/app/ui/main/page'
import { useParams } from "next/navigation";
import { VideoQuery } from "@/app/lib/api/types";
const Page = () => {
    const params = useParams()
    const query:VideoQuery = {
        category_id:Number(params.id),
    }
  return (
    <>
      <Main {...query}></Main>
    </>
  );
};

export default Page;
