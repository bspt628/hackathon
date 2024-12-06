"use client";

import { useState } from "react";
import { Home, Search, Bell, Mail, User, List} from "lucide-react";
import { Timeline } from "@/components/timeline";
import { YouTubeSearch } from "@/components/youtube-search";
import { AudioPlayer } from "@/components/audio-player";
import Link from "next/link";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useAuth, AuthProvider } from "@/contexts/auth-context";

export default function HomePage() {
	const [currentVideoId, setCurrentVideoId] = useState<string | null>(null);
	const [showTimeline, setShowTimeline] = useState(false);
	const { idToken } = useAuth();


	return (
		<div className="min-h-screen bg-black text-white">
			<div className="flex mx-auto max-w-7xl">
				{/* Left Sidebar */}
				<div className="w-64 fixed h-screen border-r border-[#2f3336] p-4">
					<div className="space-y-4">
						<Link
							href="/home"
							className="block p-3 hover:bg-white/10 rounded-full"
						>
							<Home className="w-7 h-7" />
						</Link>
						<Link href="#" className="block p-3 hover:bg-white/10 rounded-full">
							<Search className="w-7 h-7" />
						</Link>
						<Link href="#" className="block p-3 hover:bg-white/10 rounded-full">
							<Bell className="w-7 h-7" />
						</Link>
						<Link href="#" className="block p-3 hover:bg-white/10 rounded-full">
							<Mail className="w-7 h-7" />
						</Link>
						<Link href="#" className="block p-3 hover:bg-white/10 rounded-full">
							<User className="w-7 h-7" />
						</Link>
					</div>
				</div>

				{/* Main Content */}
				<main className="flex-1 min-h-screen border-r border-[#2f3336] ml-64 mr-80">
					<div className="sticky top-0 z-10 bg-black/80 backdrop-blur-md border-b border-[#2f3336]">
						<div className="flex items-center justify-between px-4 py-3">
							<h1 className="text-xl font-bold">ホーム</h1>
							<div className="space-x-2">
								<Button
									variant="outline"
									size="sm"
									className="text-[#1d9bf0] border-[#1d9bf0] hover:bg-[#1d9bf0]/10"
									onClick={() => {
										setShowTimeline(!showTimeline)
										console.log("showTimeline: ", showTimeline)
									}}
								>
									<List className="w-5 h-5 mr-2" />
									{showTimeline ? "タイムラインを隠す" : "タイムラインを表示"}
								</Button>
							</div>
						</div>
						<div className="flex border-b border-[#2f3336]">
							<button className="flex-1 hover:bg-white/[0.03] px-4 py-4 relative">
								<span>おすすめ</span>
								<div className="absolute bottom-0 left-0 right-0 h-1 bg-[#1d9bf0] rounded-full" />
							</button>
							<button className="flex-1 hover:bg-white/[0.03] px-4 py-4">
								フォロー中
							</button>
						</div>
					</div>
					{idToken && (
						<div className="p-4 border-b border-[#2f3336] break-all">
							<h2 className="font-bold mb-2">IDトークン:</h2>
							<p className="text-sm">{idToken}</p>
						</div>
					)}
					<div className="p-4 border-b border-[#2f3336]">
						<YouTubeSearch onVideoSelect={setCurrentVideoId} />
					</div>
					{currentVideoId && <AudioPlayer videoId={currentVideoId} />}
					<AuthProvider> {showTimeline && <Timeline />} </AuthProvider>
					
				</main>

				{/* Right Sidebar */}
				<div className="w-80 fixed right-0 h-screen p-4">
					<div className="sticky top-0">
						<div className="relative">
							<Search className="absolute left-3 top-3 h-5 w-5 text-gray-500" />
							<Input
								placeholder="検索"
								className="pl-10 bg-[#202327] border-transparent focus:border-[#1d9bf0] text-white"
							/>
						</div>
						<div className="mt-4 bg-[#16181c] rounded-2xl p-4">
							<h2 className="text-xl font-bold mb-4">トレンド</h2>
							<div className="space-y-4">
								<div className="hover:bg-white/[0.03] cursor-pointer">
									<div className="text-sm text-[#71767b]">トレンド</div>
									<div className="font-bold">インナーカラー</div>
									<div className="text-sm text-[#71767b]">1,234 posts</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}
