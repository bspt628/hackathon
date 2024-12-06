'use client'

import { createContext, useContext, useEffect, useState } from 'react'
import { auth } from '@/lib/firebase'
import { User, signInWithEmailAndPassword } from 'firebase/auth'
import { useRouter } from 'next/navigation'


interface AuthContextType {
  user: User | null
  idToken: string | null
  getIdToken: () => Promise<string | null>
  login: (email: string, password: string) => Promise<void>
}

const AuthContext = createContext<AuthContextType>({
  user: null,
  idToken: null,
  getIdToken: async () => null,
  login: async () => {},
})

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [idToken, setIdToken] = useState<string | null>(null)
  const router = useRouter()

  useEffect(() => {
    const unsubscribe = auth.onAuthStateChanged(async (user) => {
      setUser(user)
      if (user) {
        const token = await user.getIdToken()
        setIdToken(token)
      } else {
        setIdToken(null)
      }
    })

    return () => unsubscribe()
  }, [])

  const getIdToken = async () => {
    if (user) {
      const token = await user.getIdToken(true)
      setIdToken(token)
      return token
    }
    return null
  }

  const login = async (email: string, password: string) => {
    console.log('Login function called with:', email, password);
    try {
      const userCredential = await signInWithEmailAndPassword(auth, email, password);
      console.log('User signed in:', userCredential.user);
      router.push('/home');
    } catch (error) {
      console.error('Login error:', error);
      throw error; // Re-throw the error to be caught in the component
    }
  }

  return (
    <AuthContext.Provider value={{ user, idToken, getIdToken, login}}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => useContext(AuthContext)
