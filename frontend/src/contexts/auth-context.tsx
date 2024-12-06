"use client";

import { createContext, useContext, useEffect, useState } from "react";
import { auth, signOut } from "@/lib/firebase";
import { User } from "firebase/auth";

interface AuthContextType {
  user: User | null;
  idToken: string | null;
  logout: () => Promise<void>;
  setCurrentUser: React.Dispatch<React.SetStateAction<User | null>>; // 追加
  setIdToken: React.Dispatch<React.SetStateAction<string | null>>; // 追加
}

const AuthContext = createContext<AuthContextType>({
  user: null,
  idToken: null,
  logout: async () => {
    throw new Error("logout function must be overridden in the provider.");
  },
  setCurrentUser: () => {}, // 初期値は空の関数
  setIdToken: () => {}, // 初期値は空の関数
});

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [currentUser, setCurrentUser] = useState<User | null>(null);
  const [idToken, setIdToken] = useState<string | null>(null);

  useEffect(() => {
    const unsubscribe = auth.onAuthStateChanged(async (user) => {
      console.log("User state changed:");
      console.log(user);
      if (user) {
        const token = await user.getIdToken();
        setCurrentUser(user); // ユーザー情報を更新
        setIdToken(token); // IDトークンを更新
      } else {
        setCurrentUser(null); // ユーザーがログアウトした場合、nullを設定
        setIdToken(null); // IDトークンもnullに設定
      }
    });

    return () => unsubscribe(); // クリーンアップ
  }, []);

  // ログアウト関数
  const logout = async () => {
    try {
      await signOut(auth); // FirebaseのsignOutを呼び出してログアウト
      setCurrentUser(null); // ユーザー情報をnullに設定
      setIdToken(null); // IDトークンをnullに設定
      console.log("User logged out successfully");
    } catch (error) {
      console.error("Error logging out: ", error);
    }
  };

  return (
    <AuthContext.Provider
      value={{ user: currentUser, idToken, logout, setCurrentUser, setIdToken }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => useContext(AuthContext);
