"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { ArrowLeft } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Post } from "@/components/post";
import { Repost } from "@/components/repost";
import { Reply } from "@/components/reply";
import { useAuth } from "@/contexts/auth-context";
import { use } from "react";

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

	const fetchPost = async () => {
		if (!idToken) return;

		setIsLoading(true);
		try {
			// params.idの取得
			const id = (await params).id;
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
		fetchPost(); // Refetch the post to update replies
		setReplyingTo(null);
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
						<h1 className="text-xl font-bold">ポスト</h1>
					</div>
				</div>

				{renderPost(post)}

				<div className="flex justify-around border-y border-[#2f3336] py-">
					<Reply
						postId={post.id}
						username={post.username}
						onReplySuccess={handleReplySuccess}
					/>
					<Repost postId={post.id} username={post.username} />
				</div>
			</div>
		</div>
	);
}
