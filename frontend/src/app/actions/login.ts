'use server'

import { auth } from '@/lib/firebase'
import { signInWithEmailAndPassword} from 'firebase/auth'

interface LoginResultSuccess {
  success: true;  
  user: string; // ユーザー情報
  id_token: string; // IDトークン
}

interface LoginResultError {
  success: false;
  error: string; // エラーメッセージ
}

type LoginResult = LoginResultSuccess | LoginResultError;

export async function loginUser(formData: FormData): Promise<LoginResult> {
  console.log('Logging in...')
  const username = formData.get('username') as string
  const password = formData.get('password') as string

  try {
    // First, fetch the email for the username
    console.log('Logging in with username:', username)
    const emailResponse = await fetch(
      `https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/users/email/${username}`
    )
    console.log('fetch request completed')

    if (!emailResponse.ok) {
      throw new Error('Failed to fetch email')
    }

    const emailData = await emailResponse.json()

    if (!emailData.email) {
      throw new Error('Email not found for username')
    }
    console.log('Email found:', emailData.email)

    // Then sign in with Firebase Auth
    const userCredential = await signInWithEmailAndPassword(auth, emailData.email, password)
    console.log('Login successful usercredential:')
    const user = JSON.stringify(userCredential.user);
    const idToken = await userCredential.user.getIdToken()
    console.log('Login successful myuser:', idToken)

    return {
      success: true,
      user: user,
      id_token: idToken,
    }
  } catch (error) {
    console.error('Login error:', error)
    return { 
      success: false,
      error: error instanceof Error ? error.message : "ログインに失敗しました。",
    }
  }
}
