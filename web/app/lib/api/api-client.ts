// api-client.js
import axios, { AxiosRequestConfig } from 'axios';
import { Category, MainVideoItem, MainVideoSubmit, UploadResponse, User, UserRegisterPayload, Video, VideoQuery } from './types';
import Cookies from 'js-cookie';
const apiClient = axios.create({
  baseURL: 'http://47.106.228.5:9133/v1',
});

// 请求拦截器
apiClient.interceptors.request.use(
  config => {
    // 获取存储在客户端的 JWT
    const token = Cookies.get('token');

    // 如果有 token，则在每个请求的头部添加 Authorization
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }else{
      console.log('no token')
    }

    // 返回修改后的配置
    return config;
  },
  error => {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);


const login= async (data:{username:string,password:string}) => {
  let url = '/user/login'
  return apiClient.post<{token:string,user:User}>(url,data)
}
const register= async (data:UserRegisterPayload) => {
  let url = '/user/register'
  return apiClient.post<User>(url,data)
}

const curUser =async () => {
  const url = '/current-user'
  return apiClient.get<User>(url)
}
const userAction =async (id:number,action: 'Like'| 'Follow'| 'LikeCancel'| 'FollowCancel') => {
  const url = `/user/${id}/action`
  return apiClient.post(url,null,{
    params:{
      action:action
    }
  })
}

const getVideos = async (q?:VideoQuery) => {
  let url= '/videos'
  return apiClient.get<MainVideoItem[]>(url,{
    params:q,
  })
}

const getVideo = async (id:number) => {
  let url= `/video/${id}`
  return apiClient.get(url)
}
const postVideo = async (d:MainVideoSubmit) => {
  let url= `/video`
  return apiClient.post<Video>(url,d)
}
const uploadVideo = async (file:FormData,args?:AxiosRequestConfig) => {
  let url= `/upload`
  return apiClient.post<UploadResponse>(url,file,args)
}

const getCategories =async () => {
  const url  = '/categories'
  return apiClient.get<Category[]>(url)
}
const playVideo =async (id:number) => {
  const url = `/action/play/video/${id}`
  return apiClient.post(url)
}

const likeVideo =async (id:number) => {
  const url = `/action/like/video/${id}`
  return apiClient.post(url)
}

const collectVideo =async (id:number) => {
  const url = `/action/collect/video/${id}`
  return apiClient.post(url)
}

const cancelLikeVideo =async (id:number) => {
  const url = `/action/like/video/${id}`
  return apiClient.delete(url)
}

const cancelCollectVideo =async (id:number) => {
  const url = `/action/collect/video/${id}`
  return apiClient.delete(url)
}

const user = {
  login,
  register,
  curUser,
  userAction
};
const video={
  getVideos,
  getVideo,
  postVideo,
  uploadVideo
}
const category={
  getCategories
}
const action = {
  playVideo,
  likeVideo,
  collectVideo,
  cancelLikeVideo,
  cancelCollectVideo,
}
const api={
  user,
  video,
  category,
  action
}

export default api
