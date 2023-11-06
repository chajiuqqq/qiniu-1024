'use client'
import React, { ChangeEvent, useEffect, useRef, useState } from 'react';
import axios from 'axios';
import api from '../lib/api/api-client';
import {
  ArrowPathIcon,
  ArrowUpTrayIcon
} from "@heroicons/react/24/outline";
import { Category, MainVideoItem, MainVideoSubmit, UploadResponse } from '../lib/api/types';
import { TIMEOUT } from 'dns';

import { useRouter } from 'next/navigation'
import AutoDismissAlert from '../ui/alert';
import { url } from 'inspector';

const FileUpload = () => {
  const [uploadProgress, setUploadProgress] = useState(0);
  const [uploadStatus, setUploadStatus] = useState('');
  const intervalRef = useRef<NodeJS.Timeout>()
  const [uploadVideo, setUploadVideo] = useState<MainVideoItem>()
  const [cates, setCates] = useState<Category[]>()
  const [submitReq, setSubmitReq] = useState<MainVideoSubmit>({ category_id: 0, video_id: 0, desc: '' })
  const [loading, setLoading] = useState<boolean>(false)
  const [alertText, setAlertText] = useState('')
  const [submitLoading, setSubmitLoading] = useState(false)

  const router = useRouter()
  useEffect(() => {
    if (!cates) {
      api.category.getCategories().then(res => {
        setCates(res.data)
      })
    }
    return () => {
      if (intervalRef.current) {
        clearInterval(intervalRef.current);
      }
    };
  }, []);
  const handleSelect = (e: ChangeEvent<HTMLSelectElement>) => {
    setSubmitReq({
      ...submitReq,
      category_id: Number(e.target.value),
    })
  }
  const handleDescChange = (e: ChangeEvent<HTMLTextAreaElement>) => {
    setSubmitReq({
      ...submitReq,
      desc: e.target.value,
    })
  }
  const onFileChange = (e: any) => {
    const file = e.target.files[0]
    if (!file) {
      return;
    }

    const formData = new FormData();
    formData.append('file', file);

    try {
      setLoading(true)
      api.video.uploadVideo(formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
        onUploadProgress: (progressEvent) => {
          let percentCompleted: number = 0
          if (progressEvent.total) {
            percentCompleted = Math.round((progressEvent.loaded * 100) / progressEvent.total);
          }
          setUploadProgress(percentCompleted);
        },
      }).then(res => {
        setSubmitReq({
          ...submitReq,
          video_id: res.data.vid,
        })
        intervalRef.current = setInterval(() => {
          console.log('query upload status...')
          if (!uploadVideo) {
            api.video.getVideo(res.data.vid).then(res => {
              if (res.data.status == 'New' && res.data.cover_status == 'Success') {
                setLoading(false)
                setAlertText('上传成功!')
                console.log(res.data);
                setUploadVideo(res.data)
                clearInterval(intervalRef.current); // 清除定时器
              }
            }).catch(err => {
              clearInterval(intervalRef.current); // 清除定时器
            })
          } else {
            clearInterval(intervalRef.current); // 清除定时器
          }
        }, 2000)
      })
    } catch (error) {
      setAlertText('上传失败！')
      setLoading(false)
      console.error('Error uploading file:', error);
    }
  };

  const onUpload = async () => {
    setSubmitLoading(true)
    api.video.postVideo(submitReq).then(res => {
      setSubmitLoading(false)
      setAlertText('发布成功！')
      setTimeout(() => {
        router.push("/my")
      }, 1500);
    }).catch(err => {
      setSubmitLoading(false)
      setAlertText('发布失败！')
    })
  };

  return (
    <>
      {alertText != '' ? (<AutoDismissAlert message={alertText} key={alertText} />) : ''}

      <div className='flex flex-col justify-center items-center space-y-5'>
        <div className='flex justify-center space-x-5'>
          <div className='w-64 h-6/12 border rounded-md flex flex-col justify-center items-center'>
            {
              uploadVideo?.avatar_url&&(
                <img className="w-full h-full object-contain rounded-md" src='http://cdn.chajiuqqq.cn/100000032_cover.jpg'  alt="" />
              )
            }
            {/* <img className="w-full h-full object-contain rounded-md" src='http://cdn.chajiuqqq.cn/100000032_cover.jpg' alt="" /> */}
          </div>
          <div className='flex flex-col space-y-4'>

            <input id="file-input" type="file" onChange={onFileChange} className='hidden' />

            {loading ? (
              <div className='z-100'>
                <ArrowPathIcon className="w-4 text-black animate-spin ml-2 inline"></ArrowPathIcon>
                <p>正在上传...</p>
              </div>
            ) : (
              <label htmlFor="file-input" className='cursor-pointer bg-white/70 shadow w-16 h-16 flex flex-col items-center p-2 rounded-md  hover:text-sky-500'>
                <ArrowUpTrayIcon className='w-8'>
                </ArrowUpTrayIcon>
                <p>上传</p>
              </label>
            )}
            <select name="category_id" onChange={handleSelect} className='rounded-md border p-2 w-6/12'>
              <option value="">选择分类</option>
              {cates?.map((c) => {
                return (
                  <>
                    <option value={c.id} key={c.id}>{c.name}</option>
                  </>
                )
              })}
            </select>
            <textarea className='h-64 border rounded-md px-5 py-2' placeholder='视频描述' onChange={handleDescChange} />
          </div>
        </div>
        <button onClick={onUpload} className={`w-5/12 bg-blue-600 text-white rounded-md px-5 py-2   ${loading ? 'opacity-50 cursor-not-allowed' : 'hover:bg-blue-700'}`}>
          发布
          {
            submitLoading && (
              <ArrowPathIcon className="w-4 text-black animate-spin ml-2 inline"></ArrowPathIcon>
            )
          }
        </button>

      </div>
    </>

  );
};

export default FileUpload;
