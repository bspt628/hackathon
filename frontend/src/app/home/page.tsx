"use client";

import { useState, useEffect } from "react";
import { Home, LogOut } from "lucide-react";
import { Timeline } from "@/components/timeline";
import { CreatePost } from "@/components/create-post";
import { YouTubeSearch } from "@/components/youtube-search";
import { AudioPlayer } from "@/components/audio-player";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { useAuth } from "@/contexts/auth-context";
import { useYouTube } from "@/contexts/youtube-context";
import { useRouter } from "next/navigation";

export default function HomePage() {
	const [refreshTrigger, setRefreshTrigger] = useState(0);
	const [isLoading, setIsLoading] = useState(true);
	const [postContent, setPostContent] = useState("");
	const { currentVideoId, setCurrentVideoId, copiedText } = useYouTube();


	const { idToken, logout } = useAuth();
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
			<div className="min-h-screen bg-[#F8FAFF] text-foreground flex items-center justify-center">
				読み込み中...
			</div>
		);
	}

	return (
		<div className="min-h-screen bg-gradient-to-br from-[#E6EFFF] to-[#F8FAFF] text-foreground">
			<div className="flex mx-auto max-w-7xl">
				{/* Left Sidebar */}
				<div className="w-64 fixed h-screen border-r border-[#E1E8FF] p-4">
					<div className="space-y-4">
						<Link
							href="/home"
							className="block p-3 hover:bg-primary/[0.03] rounded-full"
						>
							<Home className="w-7 h-7" />
						</Link>
					</div>
				</div>
				{/* Main Content */}
				<main className="flex-1 min-h-screen border-r border-[#E1E8FF] ml-64 mr-80 pb-24">
					<div className="sticky top-0 z-20 bg-gradient-to-br from-[#E6EFFF] to-[#F8FAFF] text-foreground">
						<div className="flex items-center justify-between px-4 py-3">
							<h1 className="text-xl font-bold">ホーム</h1>
							<div className="space-x-2">
								<Button
									variant="outline"
									size="sm"
									className="text-secondary border-secondary hover:bg-secondary/10"
									onClick={handleSignOut}
								>
									<LogOut className="w-5 h-5 mr-2" />
									ログアウト
								</Button>
							</div>
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
					<div className="sticky top-0 bg-gradient-to-br from-[#E6EFFF] to-[#F8FAFF] text-foreground">
						
						<YouTubeSearch />
					</div>
				</div>
			</div>
			<div className="fixed bottom-0 left-0 right-0 z-10">
				{currentVideoId && (
					<AudioPlayer
						key={currentVideoId}
						videoId={currentVideoId}
						onClose={() => setCurrentVideoId(null)}
						onCopy={(text) => setPostContent(text)}
					/>
				)}
			</div>
		</div>
	);
}
