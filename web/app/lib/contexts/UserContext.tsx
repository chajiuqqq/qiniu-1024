// contexts/UserContext.tsx
import { createContext, useContext, ReactNode, useState } from 'react';

export type UserType = {
  name: string;
  token: string;
  followCnt:number;
  followerCnt:number;
  userLikesCnt:number;
  desc:string;
  avatar:string;
};

type UserContextType = {
  user: UserType;
  setUser: React.Dispatch<React.SetStateAction<UserType>>;
};

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<UserType>({name:"张三",token:"1234",followCnt:100,followerCnt:20,userLikesCnt:89,desc:'我是乐观开朗的孩子',avatar:'/avatar.jpg'});

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
