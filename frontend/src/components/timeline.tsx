"use client";

import { useEffect, useState } from "react";
import { useAuth } from "@/app/contexts/auth-context";
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
	const [posts, setPosts] = useState<TimelinePost[]>([]);
	const [isLoading, setIsLoading] = useState(true);
	const { idToken } = useAuth();

	useEffect(() => {
		async function fetchPosts() {
			if (!idToken) return;

			try {
				const response = await fetch(
					"https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/posts/timeline/all",
					{
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
			} catch (error) {
				console.error("Error fetching timeline:", error);
			} finally {
				setIsLoading(false);
			}
		}

		fetchPosts();
	}, [idToken]);

	if (isLoading) {
		return <div className="p-4 text-center">読み込み中...</div>;
	}
	console.log(posts);

	return (
		<div>
			{posts.map((post) => (
				<Post key={post.id} {...post} />
			))}
		</div>
	);
}
