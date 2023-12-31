"use client";
import React, { ChangeEvent, useEffect, useRef, useState } from "react";
import axios from "axios";
import api from "../lib/api/api-client";
import { ArrowPathIcon, ArrowUpTrayIcon } from "@heroicons/react/24/outline";
import {
  Category,
  MainVideoItem,
  MainVideoSubmit,
  UploadResponse,
} from "../lib/api/types";
import { TIMEOUT } from "dns";

import { useRouter } from "next/navigation";
import AutoDismissAlert from "../ui/alert";
import { url } from "inspector";
import { text } from "stream/consumers";

const FileUpload = () => {
  const intervalRef = useRef<NodeJS.Timeout>();
  const [uploadVideo, setUploadVideo] = useState<MainVideoItem>();
  const [cates, setCates] = useState<Category[]>();
  const [loading, setLoading] = useState<boolean>(false);
  const [alertText, setAlertText] = useState("");
  const [submitLoading, setSubmitLoading] = useState(false);
  const videoIDRef = useRef(0)
  const [selectedID, setSelectedID] = useState(0)
  const [textDesc, setTextDesc] = useState('')

  const router = useRouter();
  useEffect(() => {
    if (!cates) {
      api.category.getCategories().then((res) => {
        if (!cates) {
          setCates(res.data);
        }
      });
    }
    return () => {
      if (intervalRef.current) {
        clearInterval(intervalRef.current);
      }
    };
  }, []);
  const handleSelect = (e: ChangeEvent<HTMLSelectElement>) => {
    setSelectedID(Number(e.target.value))
  };
  const handleDescChange = (e: ChangeEvent<HTMLTextAreaElement>) => {
    setTextDesc(e.target.value)
  };
  const onFileChange = (e: any) => {
    const file = e.target.files[0];
    if (!file) {
      return;
    }

    const formData = new FormData();
    formData.append("file", file);

    setLoading(true);
    api.video
      .uploadVideo(formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then((res) => {
        videoIDRef.current = res.data.vid
        intervalRef.current = setInterval(() => {
          console.log("query upload status...");
          if (!uploadVideo) {
            api.video
              .getVideo(res.data.vid)
              .then((res) => {
                if (
                  res.data.status == "New" &&
                  res.data.cover_status == "Success"
                ) {
                  setLoading(false);
                  setAlertText("上传成功!");
                  console.log(res.data);
                  setUploadVideo(res.data);
                  clearInterval(intervalRef.current); // 清除定时器
                }
              })
              .catch((err) => {
                clearInterval(intervalRef.current); // 清除定时器
              });
          } else {
            clearInterval(intervalRef.current); // 清除定时器
          }
        }, 2000);
      }).catch(e => {
        setLoading(false);
        setAlertText("上传失败！");
        console.error("Error uploading file:", e);
      })
  };
  const onUpload = async () => {
    if (selectedID == 0 || textDesc == '' || videoIDRef.current == 0) {
      setAlertText("视频信息不完整哦！");
      return
    }
    setSubmitLoading(true);
    api.video
      .postVideo({
        video_id: videoIDRef.current,
        category_id: selectedID,
        desc: textDesc,
      })
      .then((res) => {
        setSubmitLoading(false);
        setAlertText("发布成功！");
        setLoading(false);
        setTimeout(() => {
          router.push("/my");
        }, 1500);
      })
      .catch((err) => {
        setSubmitLoading(false);
        setLoading(false);
        setAlertText("发布失败！");
      });
  };

  return (
    <>
      {alertText != "" ? (
        <AutoDismissAlert message={alertText} key={alertText} />
      ) : (
        ""
      )}
        <div className="flex space-x-5 w-full mt-10">
          <div className="w-3/4 h-6/12 border rounded-md flex flex-col justify-center items-center">
            {uploadVideo && (
              <img
                className="w-full h-full object-contain rounded-md"
                src={uploadVideo.cover_url}
                alt=""
              />
            )}
            {/* <img className="w-full h-full object-contain rounded-md" src='http://cdn.chajiuqqq.cn/100000032_cover.jpg' alt="" /> */}
          </div>
          <div className="flex flex-col space-y-4">
              
            <input
              id="file-input"
              type="file"
              onChange={onFileChange}
              className="hidden"
            />

            {loading ? (
              <div className="z-100">
                <ArrowPathIcon className="w-4 text-black animate-spin ml-2 inline"></ArrowPathIcon>
                <p>正在上传...</p>
              </div>
            ) : (
              <label
                htmlFor="file-input"
                className="cursor-pointer bg-white/70 shadow w-16 h-16 flex flex-col items-center p-2 rounded-md  hover:text-sky-500"
              >
                <ArrowUpTrayIcon className="w-8"></ArrowUpTrayIcon>
                <p>上传</p>
              </label>
            )}
            <select
              name="category_id"
              onChange={handleSelect}
              className="rounded-md border p-2 w-6/12"
              value={selectedID}
            >
              <option value="0" key="0">
                选择分类
              </option>
              {cates?.map((c) => {
                return (
                  <>
                    <option value={c.id} key={c.id}>
                      {c.name}
                    </option>
                  </>
                );
              })}
            </select>
            <textarea
              className="h-64 border rounded-md px-5 py-2"
              placeholder="视频描述"
              onChange={handleDescChange}
              value={textDesc}
            />
            
            <button
              onClick={onUpload}
              className={`w-full bg-blue-600 text-white rounded-md px-5 py-2   ${loading ? "opacity-50 cursor-not-allowed" : "hover:bg-blue-700"
                }`}
            >
              发布
              {submitLoading && (
                <ArrowPathIcon className="w-4 text-black animate-spin ml-2 inline"></ArrowPathIcon>
              )}
            </button>

          </div>
        </div>
    </>
  );
};

export default FileUpload;
