"use client";

import { useEffect, useState } from "react";
import { useAuth } from "@/contexts/auth-context";
import { Post } from "./post";
import { Reply } from "./reply";

interface TimelinePost {
	id: string;
	display_name: string;
	username: string;
	created_at: string;
	content: string;
	replies_count: number;
	reposts_count: number;
	likes_count: number;
	reply_to_id: string;
	original_post_id: string;
	replies?: TimelinePost[];
	user_id: string;
	is_deleted: boolean;
}

interface TimelineProps {
	refreshTrigger: number;
}

export function Timeline({ refreshTrigger }: TimelineProps) {
	const { idToken } = useAuth();
	const [posts, setPosts] = useState<TimelinePost[]>([]);
	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);
	const [replyingTo, setReplyingTo] = useState<string>("");

	const fetchPosts = async () => {
		if (!idToken) return;

		setIsLoading(true);
		try {
			const response = await fetch(
				"https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/posts/timeline/all",
				{
					method: "GET",
					headers: {
						Authorization: `Bearer ${idToken}`,
					},
				}
			);

			if (!response.ok) {
				throw new Error("Failed to fetch posts");
			}

			const data = await response.json();

			const threadedPosts = organizeThreads(data);
			setPosts(threadedPosts);
			setError(null);
		} catch (error) {
			console.error("Error fetching timeline:", error);
			setError("投稿の取得に失敗しました。");
		} finally {
			setIsLoading(false);
		}
	};

	const organizeThreads = (posts: TimelinePost[]): TimelinePost[] => {
		if (!Array.isArray(posts)) {
			console.error("Expected posts to be an array, but got:", posts);
			return [];
		}

		const threads: { [key: string]: TimelinePost[] } = {};
		const rootPosts: TimelinePost[] = [];

		posts.forEach((post) => {
			if (post.reply_to_id) {
				if (!threads[post.reply_to_id]) {
					threads[post.reply_to_id] = [];
				}
				threads[post.reply_to_id].push(post);
			} else if (post.original_post_id === "") {
				rootPosts.push(post);
			}
		});

		const buildReplyTree = (post: TimelinePost): TimelinePost => {
			const replies = threads[post.id] || [];
			return {
				...post,
				replies: replies
					.map(buildReplyTree)
					.sort(
						(a, b) =>
							new Date(a.created_at).getTime() -
							new Date(b.created_at).getTime()
					),
			};
		};

		const threaded = rootPosts.map(buildReplyTree);

		return threaded.sort(
			(a, b) =>
				new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
		);
	};

	useEffect(() => {
		fetchPosts();
	}, [idToken, refreshTrigger]);

	const handleReply = (postId: string) => {
		setReplyingTo(postId);
	};

	const handleRepostClick = () => {
		// Implement repost functionality if needed
	};

	const handleDelete = (postId: string) => {
		setPosts((prevPosts) =>
			prevPosts.map((post) =>
				post.id === postId ? { ...post, is_deleted: true } : post
			)
		);
	};

	if (isLoading) {
		return <div className="p-4 text-center text-muted-foreground">読み込み中...</div>;
	}

	if (error) {
		return <div className="p-4 text-center text-destructive">{error}</div>;
	}
	console.log(posts.length, "posts fetched");

	if (posts.length === 0) {
		return (
			<div className="p-4 text-center">
				<p>投稿がありません。</p>
				<details>
					<summary>デバッグ情報</summary>
					<pre>
						{JSON.stringify(
							{ postsLength: posts.length, isLoading, error },
							null,
							2
						)}
					</pre>
				</details>
			</div>
		);
	}

	const renderPost = (post: TimelinePost, isReplyTo: boolean = false) => {
		return (
			<div key={post.id}>
				<Post
					{...post}
					is_deleted={post.is_deleted}
					isReplyTo={isReplyTo}
					hasReplies={!!post.replies && post.replies.length > 0}
					onReplyClick={() => handleReply(post.id)}
					onRepostClick={() => handleRepostClick()}
					onDelete={() => handleDelete(post.id)}
					className="relative z-0"
				/>
				{replyingTo === post.id && (
					<Reply
						postId={post.id}
						username={post.username}
						onClose={() => setReplyingTo("")}
						onReplySuccess={fetchPosts}
					/>
				)}
				{post.replies?.map((reply) => renderPost(reply, true))}
			</div>
		);
	};

	return <div>{posts.map((post) => renderPost(post))}</div>;
}
