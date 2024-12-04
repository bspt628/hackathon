'use server'

interface LoginData {
  username: string
  password: string
}

export async function loginUser(formData: FormData) {
  const data: LoginData = {
    username: formData.get('username') as string,
    password: formData.get('password') as string,
  }

  try {
    const response = await fetch(
      'https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/users/signin',
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      }
    )

    if (!response.ok) {
      throw new Error('Login failed')
    }

    return { success: true }
  } catch (error) {
    console.error('Login error:', error);
    return { success: false, error: error instanceof Error ? error.message : 'ログインに失敗しました。' }
  }
}
