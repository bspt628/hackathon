import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { replyToPost } from "@/app/actions/reply-post";
import { useAuth } from "@/contexts/auth-context";

interface ReplyProps {
	postId: string;
	username: string;
	onClose?: () => void;
}

export function Reply({ postId, username, onClose }: ReplyProps) {
	const [replyContent, setReplyContent] = useState("");
	const { idToken } = useAuth();

	const handleReply = async () => {
		console.log("Replying to post:", postId, replyContent);
		if (!idToken || !replyContent.trim()) {
			console.error("Invalid reply content or idToken");
			return;
		}

		try {
			console.log("Replying create start");
			const result = await replyToPost(postId, replyContent, idToken);
			console.log("result", result);
			if (result.success) {
				setReplyContent("");
				onClose?.()
				console.log("Replied successfully");
				// You might want to update the UI or show a success message here
			}
		} catch (error) {
			console.error("Error replying:", error);
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
						rows={4}
					/>
					<div className="flex justify-end mt-2 space-x-2">
						<Button variant="ghost" onClick={() => onClose?.()}>
							キャンセル
						</Button>
						<Button
							onClick={handleReply}
							disabled={!replyContent.trim()}
							className="rounded-full bg-[#1d9bf0] hover:bg-[#1a8cd8] px-4"
						>
							返信を送信
						</Button>
					</div>
				</div>
			</div>
		</div>
	);
}
