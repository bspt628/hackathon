'use server'

import { auth } from '@/lib/firebase'
import { signInWithEmailAndPassword } from 'firebase/auth'

export async function loginUser(formData: FormData) {
  const username = formData.get('username') as string
  const password = formData.get('password') as string

  try {
    // First, fetch the email for the username
    const emailResponse = await fetch(
      `https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/users/email/${username}`
    )
    
    if (!emailResponse.ok) {
      throw new Error('Failed to fetch email')
    }

    const emailData = await emailResponse.json()
    
    if (!emailData.email) {
      throw new Error('Email not found for username')
    }

    // Then sign in with Firebase
    const userCredential = await signInWithEmailAndPassword(auth, emailData.email, password)
    const idToken = await userCredential.user.getIdToken()
    // ログイン成功後に、idTokenを返す
      .then((idToken) => {
        console.log('User signed in:', userCredential.user, idToken)
      })

    return { success: true, idToken }
  } catch (error) {
    console.error('Login error:', error)
    return { 
      success: false, 
      error: error instanceof Error ? error.message : 'ログインに失敗しました。'
    }
  }
}

