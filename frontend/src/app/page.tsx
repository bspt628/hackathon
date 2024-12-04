'use client'
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardFooter, CardHeader } from "@/components/ui/card"
import Image from "next/image"
import Link from "next/link"
import { useRouter } from "next/navigation"

export default function LoginPage() {
  const router = useRouter()
  return (
    <div className="min-h-screen bg-black text-white flex items-center justify-center p-4">
      <Card className="w-full max-w-md bg-black border-none text-white">
        <CardHeader className="space-y-12">
          <div className="w-12 h-12 relative mx-auto">
            <Image
              src="/placeholder.svg"
              alt="Logo"
              fill
              className="object-contain"
            />
          </div>
          <div className="space-y-4 text-center">
            <h1 className="text-3xl font-bold">すべての話題が、ここに。</h1>
            <p className="text-xl">今すぐ参加しましょう。</p>
          </div>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="relative">
            <div className="absolute inset-0 flex items-center">
              <span className="w-full border-t border-[#536471]" />
            </div>
            <div className="relative flex justify-center text-xs uppercase">
              <span className="bg-black px-2 text-[#536471]">または</span>
            </div>
          </div>
          <Button 
            className="w-full bg-[#1d9bf0] hover:bg-[#1a8cd8]"
            onClick={() => router.push('/signup')}
          >
            アカウントを作成
          </Button>
          <p className="text-xs text-[#536471] text-center">
            アカウントを登録することにより、
            <Link href="#" className="text-[#1d9bf0] hover:underline">利用規約</Link>
            と
            <Link href="#" className="text-[#1d9bf0] hover:underline">プライバシーポリシー</Link>
            （Cookieの使用を含む）に同意したとみなされます。
          </p>
        </CardContent>
        <CardFooter className="flex flex-col space-y-4">
          <h2 className="font-bold">アカウントをお持ちの場合</h2>
          <Button 
            variant="outline" 
            className="w-full bg-transparent text-[#1d9bf0] border-[#536471] hover:bg-[#1d9bf0]/10"
          >
            ログイン
          </Button>
        </CardFooter>
      </Card>
    </div>
  )
}

