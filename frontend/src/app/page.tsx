"use client";
import { Button } from "@/components/ui/button";
import {
	Card,
	CardContent,
	CardFooter,
	CardHeader,
} from "@/components/ui/card";
import Link from "next/link";
import { useRouter } from "next/navigation";

export default function LoginPage() {
	const router = useRouter();
	return (
		<div className="min-h-screen bg-gradient-to-r from-[#1A202C] to-[#2D3748] text-white flex items-center justify-center p-6">
			<Card className="w-full max-w-md bg-gray-800 border-none shadow-lg rounded-lg">
				<CardHeader className="space-y-6 text-center">
					<h1 className="text-4xl font-semibold text-white">Hackathon APP</h1>
					<p className="text-lg text-muted">今すぐ参加しましょう。</p>
				</CardHeader>
				<CardContent className="space-y-6">
					<div className="relative">
						<div className="absolute inset-0 flex items-center">
							<span className="w-full border-t border-[#4A5568]" />
						</div>
					</div>
					<Button
						className="w-full bg-primary text-white hover:bg-primary/80 focus:ring-2 focus:ring-primary"
						onClick={() => router.push("/signup")}
					>
						アカウントを作成
					</Button>
					<p className="text-xs text-[#A0AEC0] text-center">
						アカウントを登録することにより、
						<Link href="#" className="text-primary hover:underline">
							利用規約
						</Link>
						と
						<Link href="#" className="text-primary hover:underline">
							プライバシーポリシー
						</Link>
						（Cookieの使用を含む）に同意したとみなされます。
					</p>
				</CardContent>
				<CardFooter className="flex flex-col space-y-4">
					<h2 className="font-semibold text-lg text-white">アカウントをお持ちの場合</h2>
					<Button
						variant="outline"
						className="w-full bg-transparent text-primary border-[#4A5568] hover:bg-primary/10 focus:ring-2 focus:ring-primary"
						onClick={() => router.push("/login")}
					>
						ログイン
					</Button>
				</CardFooter>
			</Card>
		</div>
	);
}
