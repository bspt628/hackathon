"use client";

import { useState, useEffect } from "react";
import { Home, Search, Bell, Mail, User, LogOut } from "lucide-react";
import { Timeline } from "@/components/timeline";
import { CreatePost } from "@/components/create-post";
import { YouTubeSearch } from "@/components/youtube-search";
import { AudioPlayer } from "@/components/audio-player";
import Link from "next/link";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useAuth } from "@/contexts/auth-context";
import { useYouTube } from "@/contexts/youtube-context";
import { useRouter } from "next/navigation";

export default function HomePage() {
	const [refreshTrigger, setRefreshTrigger] = useState(0);
	const [isLoading, setIsLoading] = useState(true);
	const [postContent, setPostContent] = useState("");
	const [currentVideoId, setCurrentVideoId] = useState<string | null>(null); // Update 1

	const { user, idToken, logout } = useAuth();
	const { copiedText } = useYouTube();
	const router = useRouter();

	useEffect(() => {
		if (idToken === null && !localStorage.getItem("idToken")) {
			router.push("/login");
		} else {
			setIsLoading(false);
		}
	}, [idToken, router]);

	const handleSignOut = async () => {
		await logout();
		router.push("/login");
	};

	const handlePostSuccess = () => {
		setRefreshTrigger((prev) => prev + 1);
		setPostContent(""); // Clear the post content after successful post
	};

	useEffect(() => {
		if (copiedText) {
			setPostContent(copiedText);
		}
	}, [copiedText]);

	if (isLoading) {
		return (
			<div className="min-h-screen bg-black text-white flex items-center justify-center">
				読み込み中...
			</div>
		);
	}

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
									className="text-red-500 border-red-500 hover:bg-red-500/10"
									onClick={handleSignOut}
								>
									<LogOut className="w-5 h-5 mr-2" />
									ログアウト
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
					<CreatePost
						onPostSuccess={handlePostSuccess}
						initialContent={postContent}
					/>
					<Timeline refreshTrigger={refreshTrigger} />
				</main>

				{/* Right Sidebar */}
				<div className="w-80 fixed right-0 h-screen overflow-y-auto p-4">
					<div className="sticky top-0 bg-black pb-4">
						<div className="relative mb-4">
							<Search className="absolute left-3 top-3 h-5 w-5 text-gray-500" />
							<Input
								placeholder="検索"
								className="pl-10 bg-[#202327] border-transparent focus:border-[#1d9bf0] text-white"
							/>
						</div>
						<YouTubeSearch onVideoSelect={setCurrentVideoId} />
					</div>
				</div>
			</div>
			{currentVideoId && (
				<AudioPlayer
					key={currentVideoId} // Update 2
					videoId={currentVideoId}
					onClose={() => setCurrentVideoId(null)}
					onCopy={(text) => setPostContent(text)}
				/>
			)}
		</div>
	);
}
