"use client";

import { useEffect, useState } from "react";
import { useAuth } from "@/contexts/auth-context";
import { Post } from "./post";



interface TimelinePost {
	id: string;
	display_name: string;
	username: string;
	created_at: string;
	content: string;
	replies_count: number;
	reposts_count: number;
	likes_count: number;
}

export function Timeline() {
	const { idToken } = useAuth();
	const [posts, setPosts] = useState<TimelinePost[]>([]);
	const [isLoading, setIsLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		async function fetchPosts() {
			// if (!idToken) return;

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
				setPosts(data);
				setError(null);
			} catch (error) {
				console.error("Error fetching timeline:", error);
				setError("投稿の取得に失敗しました。");
			} finally {
				setIsLoading(false);
			}
		}

		fetchPosts();
	}, [idToken]);

	if (isLoading) {
		return <div className="p-4 text-center">読み込み中...</div>;
	}

	if (error) {
		return <div className="p-4 text-center text-red-500">{error}</div>;
	}

	if (posts.length === 0) {
		return <div className="p-4 text-center">投稿がありません。</div>;
	}

	return (
		<div>
			{posts.map((post) => (
				<Post key={post.id} {...post} />
			))}
		</div>
	);
}
