interface SignupData {
  email: string
  password: string
  username: string
  display_name: string
}

export async function signupUser(formData: FormData) {
  const data: SignupData = {
    email: formData.get('email') as string,
    password: formData.get('password') as string,
    username: formData.get('username') as string,
    display_name: formData.get('display_name') as string,
    
  }
    // コンソールに表示されるデータを確認する
    console.log(data)
    

  try {
    const response = await fetch(
      'https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/users/signup',
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      }
    )

    if (!response.ok) {
      throw new Error('Signup failed')
    }

    return { success: true }
  } catch (error) {
    console.error('Signup error:', error);
    return { success: false, error: error instanceof Error ? error.message : 'アカウントの作成に失敗しました。' }
  }
}

