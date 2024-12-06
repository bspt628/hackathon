'use client'

import { createContext, useContext, useEffect, useState } from 'react'
import { auth } from '@/lib/firebase'
import { User } from 'firebase/auth'

interface AuthContextType {
  user: User | null
  idToken: string | null
}

const AuthContext = createContext<AuthContextType>({
  user: null,
  idToken: null
})

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [idToken, setIdToken] = useState<string | null>(null)

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

  return (
    <AuthContext.Provider value={{ user, idToken }}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => useContext(AuthContext)

