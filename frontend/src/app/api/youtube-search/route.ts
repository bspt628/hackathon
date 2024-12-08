import { NextResponse } from 'next/server'

const NEXT_PUBLIC_YOUTUBE_API_KEY = "AIzaSyCB-aNbuHZiBjdFtuL3LmRu_hqJE3seoSw"

console.log('NEXT_PUBLIC_YOUTUBE_API_KEY:', NEXT_PUBLIC_YOUTUBE_API_KEY) 

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url)
  const query = searchParams.get('q')

  if (!query) {
    return NextResponse.json({ error: 'Query parameter is required' }, { status: 400 })
  }

  const response = await fetch(
    `https://www.googleapis.com/youtube/v3/search?part=snippet&q=${encodeURIComponent(
      query
    )}&key=${NEXT_PUBLIC_YOUTUBE_API_KEY}&type=video&maxResults=20`
  )
  

  const data = await response.json()

  return NextResponse.json(data)
}

