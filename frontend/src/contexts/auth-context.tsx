'use client'

import { createContext, useContext, useEffect, useState } from 'react'
import { auth } from '@/lib/firebase'
import { User, signInWithEmailAndPassword, signOut } from 'firebase/auth'
import { useRouter } from 'next/navigation'

interface AuthContextType {
  user: User | null
  idToken: string | null
  backendUserId: string | null
  getIdToken: () => Promise<string | null>
  login: (email: string, password: string) => Promise<void>
  logout: () => Promise<void>
}

const AuthContext = createContext<AuthContextType>({
  user: null,
  idToken: null,
  backendUserId: null,
  getIdToken: async () => null,
  login: async () => {},
  logout: async () => {}
})

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [idToken, setIdToken] = useState<string | null>(null)
  const [backendUserId, setBackendUserId] = useState<string | null>(null)
  const router = useRouter()

  useEffect(() => {
    const unsubscribe = auth.onAuthStateChanged(async (user) => {
      setUser(user)
      if (user) {
        const token = await user.getIdToken()
        setIdToken(token)
        localStorage.setItem('idToken', token)
      } else {
        setIdToken(null)
        localStorage.removeItem('idToken')
      }
    })

    // ページロード時にローカルストレージからトークンを復元
    const storedToken = localStorage.getItem('idToken')
    if (storedToken) {
      setIdToken(storedToken)
    }

    const storedBackendUserId = localStorage.getItem('backendUserId')
    if (storedBackendUserId) {
      console.log('Restoring backend user ID:', storedBackendUserId)
      setBackendUserId(storedBackendUserId)
    }

    return () => unsubscribe()
  }, [])

  const getIdToken = async () => {
    if (user) {
      const token = await user.getIdToken(true)
      setIdToken(token)
      localStorage.setItem('idToken', token)
      return token
    }
    return null
  }

  const login = async (email: string, password: string) => {
    try {
      const userCredential = await signInWithEmailAndPassword(auth, email, password)
      console.log('User signed in:', userCredential.user)
      const token = await userCredential.user.getIdToken()
      setIdToken(token)
      localStorage.setItem('idToken', token)

      // バックエンドからユーザーIDを取得
      const response = await fetch(
        'https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/users/firebase',
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      )
      if (response.ok) {
        const data = await response.json()
        console.log('Backend user ID:', data.user_id)
        setBackendUserId(data.user_id)
        localStorage.setItem('backendUserId', data.user_id)
      }

      router.push('/home')
    } catch (error) {
      console.error('Login error:', error)
      throw error
    }
  }

  const logout = async () => {
    await signOut(auth)
    setIdToken(null)
    localStorage.removeItem('idToken')
    setBackendUserId(null)
    localStorage.removeItem('backendUserId')
    router.push('/login')
  }

  return (
    <AuthContext.Provider value={{ user, idToken, backendUserId, getIdToken, login, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => useContext(AuthContext)

