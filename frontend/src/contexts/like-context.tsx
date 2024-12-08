"use client";

import React, {
	createContext,
	useContext,
	useState,
	useCallback
} from "react";
import { likePost, unlikePost } from "@/app/actions/like-post";
import { useAuth } from "./auth-context";

interface LikesContextType {
	likedPosts: Set<string>;
	toggleLike: (postId: string) => Promise<void>;
	isLiked: (postId: string) => boolean;
	fetchLikeStatus: (postId: string) => Promise<boolean>;
}

const LikesContext = createContext<LikesContextType | undefined>(undefined);

export function LikesProvider({ children }: { children: React.ReactNode }) {
	const [likedPosts, setLikedPosts] = useState<Set<string>>(new Set());
	const { idToken } = useAuth();

	const fetchLikeStatus = useCallback(
		async (postId: string): Promise<boolean> => {
			if (!idToken) return false;
	
			try {
				const response = await fetch(
					`https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/likes/${postId}/status`,
					{
						method: "GET",
						headers: {
							Authorization: `Bearer ${idToken}`,
						},
					}
				);
				
	
				if (!response.ok) {
					throw new Error("Failed to fetch like status");
				}
	
				const data = await response.json();
				if (data.like_status) {
					setLikedPosts((prev) => new Set(prev).add(postId));
				} else {
					setLikedPosts((prev) => {
						const newSet = new Set(prev);
						newSet.delete(postId);
						return newSet;
					});
				}
	
				return data.like_status; // 取得した結果を返す
			} catch (error) {
				console.error("Error fetching like status:", error);
				return false; // エラー時はデフォルト値として false を返す
			}
		},
		[idToken]
	);
	

	const toggleLike = useCallback(
		async (postId: string) => {
			if (!idToken) return;

			try {
				if (likedPosts.has(postId)) {
					await unlikePost(postId, idToken);
					setLikedPosts((prev) => {
						const newSet = new Set(prev);
						newSet.delete(postId);
						return newSet;
					});
				} else {
					await likePost(postId, idToken);
					setLikedPosts((prev) => new Set(prev).add(postId));
				}
			} catch (error) {
				console.error("Error toggling like:", error);
			}
		},
		[likedPosts, idToken]
	);

	const isLiked = useCallback(
		(postId: string) => {
			return likedPosts.has(postId);
		},
		[likedPosts]
	);

	return (
		<LikesContext.Provider
			value={{ likedPosts, toggleLike, isLiked, fetchLikeStatus }}
		>
			{children}
		</LikesContext.Provider>
	);
}

export function useLikes() {
	const context = useContext(LikesContext);
	if (context === undefined) {
		throw new Error("useLikes must be used within a LikesProvider");
	}
	return context;
}
