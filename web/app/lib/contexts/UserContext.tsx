// contexts/UserContext.tsx
import { createContext, useContext, ReactNode, useState } from 'react';
import { User } from '../api/types';
import api from '../api/api-client';
import Cookies from 'js-cookie';
type UserContextType = {
  user: User|undefined;
  setUser: React.Dispatch<React.SetStateAction<User|undefined>>;
};

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const u = getLocalUser()
  // 获取存储在客户端的 JWT
  const token = Cookies.get('token')
  if (u == undefined && token!=undefined) {
    api.user.curUser().then((res) => {
      setUser(res.data)
    }).catch((err) => {
      console.log('get current user error', err)
    })
  }
  const [user, setUser] = useState<User|undefined>(u);
  return (
    <UserContext.Provider value={{ user, setUser }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => {
  const context = useContext(UserContext);
  if (context === undefined) {
    throw new Error('useUser must be used within a UserProvider');
  }
  return context;
};

export const getLocalUser: () => User | undefined = () => {
  if (typeof window !== 'undefined') {
    const u = localStorage.getItem('user');
    // 从 localStorage 获取数据时，再将 JSON 字符串转回 JavaScript 对象
    if(u){
      const user = JSON.parse(u);
      return user
    }
  }
  return undefined
}