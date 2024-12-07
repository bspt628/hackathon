import React, { useEffect, useState } from 'react'
import { MessageSquare, Repeat2, Heart } from "lucide-react";
import { formatDistanceToNow } from "date-fns";
import { ja } from "date-fns/locale";
import { useRouter } from "next/navigation";
import { useLikes } from "@/contexts/like-context";

interface PostProps {
	id: string;
	display_name: string;
	username: string;
	created_at: string;
	content: string;
	replies_count: number;
	reposts_count: number;
	likes_count: number;
	onReplyClick?: () => void;
	onRepostClick?: () => void;
}

export function Post({
	id,
	display_name,
	username,
	created_at,
	content,
	replies_count,
	reposts_count,
	likes_count: initialLikesCount,
	onReplyClick,
	onRepostClick,
}: PostProps) {
	const router = useRouter();
	const { toggleLike, isLiked, fetchLikeStatus } = useLikes();
	const [liked, setLiked] = useState(isLiked(id));
	const [likesCount, setLikesCount] = useState(initialLikesCount);

	useEffect(() => {
		fetchLikeStatus(id).then(() => {
			setLiked(isLiked(id));
		});
	}, [id, fetchLikeStatus, isLiked]);
	const timeAgo = formatDistanceToNow(new Date(created_at), {
		addSuffix: true,
		locale: ja,
	});

	const handleClick = (e: React.MouseEvent) => {
		// Don't navigate if clicking on action buttons
		if ((e.target as HTMLElement).closest("button")) {
			return;
		}
		router.push(`/posts/${id}`);
	};
	const handleLikeToggle = async (e: React.MouseEvent) => {
		e.stopPropagation();
		await toggleLike(id);
		const newLikedState = !liked;
		setLiked(newLikedState);
		setLikesCount((prev) => (newLikedState ? prev + 1 : prev - 1));
	};

	return (
		<div
			className="border-b border-[#2f3336] p-4 hover:bg-white/[0.03] cursor-pointer"
			onClick={handleClick}
		>
			<div className="flex gap-4">
				<div className="w-10 h-10 rounded-full bg-[#2f3336]" />
				<div className="flex-1">
					<div className="flex items-center gap-2">
						<span className="font-bold">{display_name}</span>
						<span className="text-[#71767b]">@{username}</span>
						<span className="text-[#71767b]">·</span>
						<span className="text-[#71767b]">{timeAgo}</span>
					</div>
					<p className="mt-2 break-words">{content}</p>
					<div className="flex justify-between mt-4 max-w-md text-[#71767b]">
						<button
							className="flex items-center gap-2 hover:text-[#1d9bf0]"
							onClick={(e) => {
								e.stopPropagation();
								onReplyClick?.();
							}}
						>
							<MessageSquare className="w-5 h-5" />
							<span>{replies_count}</span>
						</button>
						<button
							className="flex items-center gap-2 hover:text-[#00ba7c]"
							onClick={(e) => {
								e.stopPropagation();
								onRepostClick?.();
							}}
						>
							<Repeat2 className="w-5 h-5" />
							<span>{reposts_count}</span>
						</button>
						<button
							className={`flex items-center gap-2 hover:text-[#f91880] ${
								liked ? "text-[#f91880]" : ""
							}`}
							onClick={handleLikeToggle}
						>
							<Heart className={`w-5 h-5 ${liked ? "fill-current" : ""}`} />
							<span>{likesCount}</span>
						</button>
					</div>
				</div>
			</div>
		</div>
	);
}
