// api-client.js
import axios from 'axios';
import { MainVideoSubmit, UploadResponse, User, UserRegisterPayload, Video } from './types';
import Cookies from 'js-cookie';
const apiClient = axios.create({
  baseURL: 'http://47.106.228.5:9133/v1',
});

// 请求拦截器
apiClient.interceptors.request.use(
  config => {
    // 获取存储在客户端的 JWT
    const token = Cookies.get('token')?.value;

    // 如果有 token，则在每个请求的头部添加 Authorization
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
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

const videos = async () => {
  let url= '/videos'
  return apiClient.get(url)
}

const video = async (id:number) => {
  let url= `/video/${id}`
  return apiClient.get(url)
}
const postVideo = async (d:MainVideoSubmit) => {
  let url= `/video`
  return apiClient.post<Video>(url,d)
}
const uploadVideo = async (file:File) => {
  let url= `/upload`
  return apiClient.post<UploadResponse>(url,file)
}

const user = {
  login,
  register,
  curUser,
};
const api={
  user
}

export default api
