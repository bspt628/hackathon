"use client";

import { use, useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { ArrowLeft, Search } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Post } from "@/components/post";
import { Repost } from "@/components/repost";
import { Reply } from "@/components/reply";
import { YouTubeSearch } from "@/components/youtube-search";
import { AudioPlayer } from "@/components/audio-player";
import { useAuth } from "@/contexts/auth-context";
import { deletePost } from "@/app/actions/delete-post";

interface PostDetail {
	id: string;
	display_name: string;
	username: string;
	created_at: string;
	content: string;
	replies_count: number;
	reposts_count: number;
	likes_count: number;
	is_liked: boolean;
	reply_to_id: string | null;
	replies?: PostDetail[];
	user_id: string;
	is_deleted: boolean;
}

export default function PostDetailPage({
	params,
}: {
	params: Promise<{ id: string }>;
}) {
	const { id } = use(params);
	const router = useRouter();
	const { idToken } = useAuth();
	const [post, setPost] = useState<PostDetail | null>(null);
	const [error, setError] = useState<string | null>(null);
	const [isLoading, setIsLoading] = useState(true);
	const [replyingTo, setReplyingTo] = useState<string | null>(null);
	const [currentVideoId, setCurrentVideoId] = useState<string | null>(null);

	const fetchPost = async () => {
		if (!idToken) return;

		setIsLoading(true);
		try {
			const response = await fetch(
				`https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/posts/timeline/one/${id}`,
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
	};

	useEffect(() => {
		fetchPost();
	}, [idToken, id]);

	const handleReplySuccess = () => {
		fetchPost();
		setReplyingTo(null);
	};

	const handleDelete = async () => {
		if (!idToken || !post) return;

		try {
			const result = await deletePost(post.id, idToken);
			if (result.success) {
				router.push("/home");
			} else {
				console.error("Failed to delete post:", result.error);
				setError("投稿の削除に失敗しました。");
			}
		} catch (error) {
			console.error("Error deleting post:", error);
			setError("投稿の削除中にエラーが発生しました。");
		}
	};

	if (isLoading) {
		return <div className="p-4 text-center">読み込み中...</div>;
	}

	if (error) {
		return <div className="p-4 text-center text-red-500">{error}</div>;
	}

	if (!post) {
		return <div className="p-4 text-center">投稿が見つかりません。</div>;
	}

	const renderPost = (post: PostDetail, isReplyTo: boolean = false) => {
		return (
			<div key={post.id}>
				<Post
					{...post}
					isReplyTo={isReplyTo}
					hasReplies={!!post.replies && post.replies.length > 0}
					onReplyClick={() => setReplyingTo(post.id)}
					onRepostClick={() => {}} // Implement repost functionality if needed
					onDelete={handleDelete}
					className="relative z-0"
				/>
				{replyingTo === post.id && (
					<Reply
						postId={post.id}
						username={post.username}
						onClose={() => setReplyingTo(null)}
						onReplySuccess={handleReplySuccess}
					/>
				)}
				{post.replies?.map((reply) => renderPost(reply, true))}
			</div>
		);
	};

	return (
		<div className="min-h-screen bg-black text-white">
			<div className="flex mx-auto max-w-7xl">
				<div className="flex-1 min-h-screen border-r border-[#2f3336] mr-80">
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
							<h1 className="text-xl font-bold">ポスト</h1>
						</div>
					</div>

					{renderPost(post)}

					<div className="flex justify-around border-y border-[#2f3336] py-2">
						<Reply
							postId={post.id}
							username={post.username}
							onReplySuccess={handleReplySuccess}
						/>
						<Repost postId={post.id} username={post.username} />
					</div>
				</div>

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
					key={currentVideoId}
					videoId={currentVideoId}
					onClose={() => setCurrentVideoId(null)}
				/>
			)}
		</div>
	);
}
