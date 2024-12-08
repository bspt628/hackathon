import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { replyToPost } from "@/app/actions/reply-post";
import { useAuth } from "@/contexts/auth-context";

interface ReplyProps {
	postId: string;
	username: string;
	onClose?: () => void;
	onReplySuccess?: () => void;
}

export function Reply({
	postId,
	username,
	onClose,
	onReplySuccess,
}: ReplyProps) {
	const [replyContent, setReplyContent] = useState("");
	const [isSubmitting, setIsSubmitting] = useState(false);
	const { idToken } = useAuth();

	const handleReply = async () => {
		if (!idToken || !replyContent.trim()) return;

		setIsSubmitting(true);
		try {
			const result = await replyToPost(postId, replyContent, idToken);
			if (result.success) {
				setReplyContent("");
				onClose?.();
				onReplySuccess?.();
			} else {
				// Handle error
				console.error("Failed to reply:", result.error);
			}
		} catch (error) {
			console.error("Error replying:", error);
		} finally {
			setIsSubmitting(false);
		}
	};

	return (
		<div className="p-4 border-b border-[#2f3336]">
			<div className="flex gap-4">
				<div className="w-10 h-10 rounded-full bg-[#2f3336]" />
				<div className="flex-1">
					<Textarea
						placeholder={`@${username}に返信`}
						value={replyContent}
						onChange={(e) => setReplyContent(e.target.value)}
						className="w-full bg-transparent border-none resize-none focus:ring-0 placeholder:text-[#71767b]"
						rows={6}
					/>
					<div className="flex justify-end mt-2 space-x-2">
						<Button variant="ghost" onClick={onClose}>
							キャンセル
						</Button>
						<Button
							onClick={handleReply}
							disabled={!replyContent.trim() || isSubmitting}
							className="rounded-full bg-primary hover:bg-secondary px-4"
						>
							{isSubmitting ? "送信中..." : "返信"}
						</Button>
					</div>
				</div>
			</div>
		</div>
	);
}
