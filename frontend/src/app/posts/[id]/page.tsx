"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { ArrowLeft } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Post } from "@/components/post";
import { useAuth } from "@/contexts/auth-context";

interface PostDetail {
	id: string;
	display_name: string;
	username: string;
	created_at: string;
	content: string;
	replies_count: number;
	reposts_count: number;
	likes_count: number;
}

export default function PostDetailPage({ params }: { params: { id: string } }) {
	const router = useRouter();
	const { idToken } = useAuth();
	const [post, setPost] = useState<PostDetail | null>(null);
	const [error, setError] = useState<string | null>(null);
	const [isLoading, setIsLoading] = useState(true);

	useEffect(() => {
		async function fetchPost() {
			if (!idToken) return;

			setIsLoading(true);
			try {
				const response = await fetch(
					`https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/posts/timeline/one/${params.id}`,
					{
						headers: {
							Authorization: `Bearer ${idToken}`,
						},
					}
				);

				if (!response.ok) {
					throw new Error("Failed to fetch post");
				}

				const data = await response.json();
				setPost(data);
				setError(null);
			} catch (error) {
				console.error("Error fetching post:", error);
				setError("投稿の取得に失敗しました。");
			} finally {
				setIsLoading(false);
			}
		}

		fetchPost();
	}, [idToken, params.id]);

	if (isLoading) {
		return <div className="p-4 text-center">読み込み中...</div>;
	}

	if (error) {
		return <div className="p-4 text-center text-red-500">{error}</div>;
	}

	if (!post) {
		return <div className="p-4 text-center">投稿が見つかりません。</div>;
	}

	return (
		<div className="min-h-screen bg-black text-white">
			<div className="max-w-2xl mx-auto">
				<div className="sticky top-0 z-10 bg-black/80 backdrop-blur-md border-b border-[#2f3336]">
					<div className="flex items-center gap-4 px-4 py-3">
						<Button
							variant="ghost"
							size="icon"
							className="rounded-full hover:bg-white/10"
							onClick={() => router.back()}
						>
							<ArrowLeft className="h-5 w-5" />
						</Button>
						<h1 className="text-xl font-bold">ポストする</h1>
					</div>
				</div>

				<Post {...post} />

				<div className="p-4 border-b border-[#2f3336]">
					<div className="flex gap-4">
						<div className="w-10 h-10 rounded-full bg-[#2f3336]" />
						<div className="flex-1">
							<textarea
								placeholder="返信をポスト"
								className="w-full bg-transparent border-none resize-none focus:ring-0 placeholder:text-[#71767b]"
								rows={4}
							/>
							<div className="flex justify-end mt-2">
								<Button className="rounded-full bg-[#1d9bf0] hover:bg-[#1a8cd8] px-4">
									返信
								</Button>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}
