"use client";
// LoginComponent.tsx
import React, { useState, FormEvent } from "react";
import { useUser } from "../lib/contexts/UserContext";
import Link from 'next/link'
import api from "../lib/api/api-client";
import { User } from "../lib/api/types";
import Cookies from 'js-cookie';
import Loading from "../ui/loading";
import { useRouter } from 'next/navigation'
import { ArrowPathIcon } from "@heroicons/react/24/outline";
import AutoDismissAlert from "../ui/alert";
const LoginComponent: React.FC = () => {
  const { setUser } = useUser();
  const [alertText, setAlertText] = useState('')
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });
  const [loading, setLoading] = useState<boolean>(false)
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };
  const router = useRouter()
  const handleSubmit = (e: FormEvent) => {
    setLoading(true)
    let token = 'someToken';
    let u: User = {};
    e.preventDefault();
    console.log("Form data submitted:", formData);

    // 登录请求
    api.user.login(formData).then((res) => {
      setLoading(false)
      token = res.data.token
      u = res.data.user
      console.log(u)
      // 设置 cookie
      Cookies.set('token', token);
      localStorage.setItem('user', JSON.stringify(u));
      // 更新用户状态
      setUser(u);
      
      setAlertText('登录成功')
      setTimeout(() => {
        router.push("/my")  
      }, 1500);
    }).catch((err) => {
      if (err.response.status === 500) {
        setLoading(false)
        setAlertText('用户名或密码错误,'+err.response.data.message)
      }
      console.log(err)
    });

  };

  return (
    <>
    
    {alertText != '' ? (<AutoDismissAlert message={alertText} key={Date.now()} />) : ''}
      <div className="h-full flex items-center justify-center bg-gray-100 py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full space-y-8">
          <div>
            <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
              登录New视频
            </h2>
          </div>
          <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
            <div className="rounded-md shadow-sm space-y-4">
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
                  className="appearance-none  relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                  placeholder="Username"
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
                  autoComplete="current-password"
                  required
                  className="appearance-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                  placeholder="Password"
                  value={formData.password}
                  onChange={handleChange}
                />
              </div>
            </div>
            <div className="flex  space-x-2">
              <button
                type="submit"
                className="group relative w-9/12  py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              >
                进入{loading ? (<ArrowPathIcon className="w-4 text-white animate-spin ml-2 inline"></ArrowPathIcon>) : ''}
              
              </button>
              <div className="w-3/12">
                <Link href='/register'>
                  <div className="w-full flex justify-center text-center border text-sm font-medium border-indigo-400 rounded-md text-indigo-400 py-2 px-4 hover:text-indigo-600 hover:border-indigo-600">
                    创建
                  </div>
                </Link>
              </div>

            </div>
          </form>
        </div>
      </div>
    </>

  );
};

export default LoginComponent;
