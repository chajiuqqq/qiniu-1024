'use client'
import React, { useState, ChangeEvent, FormEvent } from 'react';
import api from '../lib/api/api-client';

import {
  ArrowPathIcon
} from "@heroicons/react/24/outline";
import { useRouter } from 'next/navigation'
import Loading from "../ui/loading";
import AutoDismissAlert from '../ui/alert';
interface FormData {
  nickname: string;
  username: string;
  password: string;
  phone: string;
  avatar: string;
  desc: string;
}

const Registration: React.FC = () => {
  const [loading, setLoading] = useState<boolean>(false)
  const [alertText, setAlertText] = useState('')
  const router = useRouter()
  const [formData, setFormData] = useState<FormData>({
    nickname: '',
    username: '',
    password: '',
    phone: '',
    avatar: '',
    desc: '',
  });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    // 提交注册逻辑
    console.log(formData);
    setLoading(true)
    api.user.register({
      name: formData.nickname,
      username: formData.username,
      password: formData.password,
      phone: formData.phone,
      avatar_url: formData.avatar,
      description: formData.desc
    }).then((res) => {
      setLoading(false)
      setAlertText('注册成功')
      setTimeout(() => {
        router.push('/login')
      }, 2000)
    }).catch((err) => {
      if (err.response.status != 200) {
        setLoading(false)
        setAlertText('注册失败,' + err.response.data.message)
      }
      console.log(err)
    })
  };

  return (
    <>
      {alertText != '' ? (<AutoDismissAlert message={alertText} key={alertText} />) : ''}
      <div className="min-h-full flex items-center justify-center bg-gray-100 py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full space-y-8">
          <div>
            <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
              注册New视频
            </h2>
          </div>
          <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
            <div className="rounded-md shadow-sm space-y-4">
              <div>
                <label htmlFor="nickname" className="sr-only">
                  Nickname
                </label>
                <input
                  id="nickname"
                  name="nickname"
                  type="text"
                  required
                  className="appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                  placeholder="昵称"
                  value={formData.nickname}
                  onChange={handleChange}
                />
              </div>
              <div>
                <label htmlFor="username" className="sr-only">
                  Username
                </label>
                <input
                  id="username"
                  name="username"
                  type="text"
                  autoComplete="username"
                  required
                  className="appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                  placeholder="用户名"
                  value={formData.username}
                  onChange={handleChange}
                />
              </div>
              <div>
                <label htmlFor="password" className="sr-only">
                  Password
                </label>
                <input
                  id="password"
                  name="password"
                  type="password"
                  autoComplete="new-password"
                  required
                  className="appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                  placeholder="密码"
                  value={formData.password}
                  onChange={handleChange}
                />
              </div>
              <div>
                <label htmlFor="phone" className="sr-only">
                  Phone
                </label>
                <input
                  id="phone"
                  name="phone"
                  type="tel"
                  required
                  className="appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                  placeholder="手机号"
                  value={formData.phone}
                  onChange={handleChange}
                />
              </div>
              <div>
                <label htmlFor="phone" className="sr-only">
                  描述
                </label>
                <input
                  id="desc"
                  name="desc"
                  type="text"
                  required
                  className="appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                  placeholder="描述"
                  value={formData.desc}
                  onChange={handleChange}
                />
              </div>
            </div>

            <div>
              <button
                type="submit"
                className="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              >
                注册
              {loading ? (<ArrowPathIcon className="w-4 text-white animate-spin ml-2"></ArrowPathIcon>) : ''}
              
              </button>
            </div>
          </form>
        </div>
      </div>
    </>

  );
};

export default Registration;
