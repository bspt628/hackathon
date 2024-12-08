import React, { useEffect, useState } from "react";
import { MessageSquare, Repeat2, Heart, Trash2 } from "lucide-react";
import { formatDistanceToNow } from "date-fns";
import { ja } from "date-fns/locale";
import { useRouter } from "next/navigation";
import { useLikes } from "@/contexts/like-context";
import { useAuth } from "@/contexts/auth-context";
import { cn } from "@/lib/utils";
import { deletePost } from "@/app/actions/delete-post";
import {
	Dialog,
	DialogContent,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
	DialogClose,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";

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
	isReplyTo: boolean;
	hasReplies: boolean;
	className?: string;
	onDelete?: () => void;
	user_id: string;
	is_deleted: boolean;
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
	isReplyTo,
	hasReplies,
	className,
	onDelete,
	user_id,
	is_deleted,
}: PostProps) {
	const router = useRouter();
	const { isLiked, toggleLike, fetchLikeStatus, likedPosts } = useLikes();
	const [liked, setLiked] = useState(isLiked(id));
	const [likesCount, setLikesCount] = useState(initialLikesCount);
	const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);
	const { idToken, backendUserId } = useAuth();
	const [isLikeLoading, setIsLikeLoading] = useState(false);


	useEffect(() => {
		(async () => {
			const likedStatus = await fetchLikeStatus(id); // fetchLikeStatus の結果を利用
			setLiked(likedStatus);
			console.log(likedStatus)
		})();
	}, [id, fetchLikeStatus]);
	

	const timeAgo = formatDistanceToNow(new Date(created_at), {
		addSuffix: true,
		locale: ja,
	});

	const handleClick = (e: React.MouseEvent) => {
		if ((e.target as HTMLElement).closest("button")) {
			return;
		}
		router.push(`/posts/${id}`);
	};

	const handleLikeToggle = async (e: React.MouseEvent) => {
		e.stopPropagation();
		setIsLikeLoading(true); // ローディング状態を開始
		try {
			await toggleLike(id);
			const newLikedState = !likedPosts.has(id);
			setLiked(newLikedState);
			setLikesCount((prev) => (newLikedState ? prev + 1 : prev - 1));
		} catch (error) {
			console.error("Error toggling like:", error);
		} finally {
			setIsLikeLoading(false); // ローディング状態を終了
		}
		
	};

	const handleDelete = async () => {
		if (idToken) {
			const result = await deletePost(id, idToken);
			if (result.success) {
				setIsDeleteDialogOpen(false);
				onDelete?.();
			} else {
				console.error("Failed to delete post:", result.error);
				// You might want to show an error message to the user here
			}
		}
	};

	const isOwnPost = backendUserId === user_id;

	return (
		<div
			className={cn(
				"border-b border-[#2f3336] hover:bg-white/[0.03] cursor-pointer relative",
				className
			)}
		>
			{/* Thread lines */}
			{isReplyTo && (
				<div className="absolute top-0 left-5 w-0.5 h-12 bg-[#2f3336]" />
			)}
			{hasReplies && (
				<div className="absolute top-12 left-5 bottom-0 w-0.5 bg-[#2f3336]" />
			)}
			<div className="p-4" onClick={handleClick}>
				<div className="flex gap-4">
					<div className="w-10 h-10 rounded-full bg-[#2f3336] shrink-0 relative z-10" />
					<div className="flex-1">
						<div className="flex items-center gap-2">
							<span className="font-bold">{display_name}</span>
							<span className="text-[#71767b]">@{username}</span>
							<span className="text-[#71767b]">·</span>
							<span className="text-[#71767b]">{timeAgo}</span>
						</div>
						{is_deleted ? (
							<p className="mt-2 text-[#71767b] italic">
								この投稿は@{username}によって削除されました
							</p>
						) : (
							<p className="mt-2 break-words whitespace-pre-wrap">
								{content.split("\n").map((line, index) => (
									<React.Fragment key={index}>
										{line}
										{index < content.split("\n").length - 1 && <br />}
									</React.Fragment>
								))}
							</p>
						)}
						{!is_deleted && (
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
									className={`flex items-center gap-2 hover:text-[#f91880] ${liked ? "text-[#f91880]" : ""
										}`}
									onClick={handleLikeToggle}
									disabled={isLikeLoading}
								>
									<Heart className={`w-5 h-5 ${liked ? "fill-current" : ""}`} />
									<span>{likesCount}</span>
								</button>
								{isOwnPost && (
									<Dialog
										open={isDeleteDialogOpen}
										onOpenChange={setIsDeleteDialogOpen}
									>
										<DialogTrigger asChild>
											<button
												className="flex items-center gap-2 hover:text-red-500"
												onClick={(e) => {
													e.stopPropagation();
												}}
											>
												<Trash2 className="w-5 h-5" />
											</button>
										</DialogTrigger>
										<DialogContent className="sm:max-w-[425px] bg-black text-white">
											<DialogHeader>
												<DialogTitle>投稿を削除しますか？</DialogTitle>
											</DialogHeader>
											<div className="grid gap-4 py-4">
												<p>この操作は取り消せません。本当に削除しますか？</p>
											</div>
											<div className="flex justify-end gap-4">
												<DialogClose asChild>
													<Button variant="outline">キャンセル</Button>
												</DialogClose>
												<Button onClick={handleDelete} variant="destructive">
													削除
												</Button>
											</div>
										</DialogContent>
									</Dialog>
								)}
							</div>
						)}
					</div>
				</div>
			</div>
		</div>
	);
}
