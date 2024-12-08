import { useState } from "react";
import {
	Dialog,
	DialogContent,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { repostPost } from "@/app/actions/repost-post";
import { useAuth } from "@/contexts/auth-context";

interface RepostProps {
	postId: string;
	username: string;
}

export function Repost({ postId, username }: RepostProps) {
	const [isOpen, setIsOpen] = useState(false);
	const [isQuoteRepost, setIsQuoteRepost] = useState(false);
	const [additionalComment, setAdditionalComment] = useState("");
	const { idToken } = useAuth();

	const handleRepost = async () => {
		if (!idToken) return;

		try {
			const result = await repostPost(
				postId,
				isQuoteRepost,
				additionalComment,
				idToken
			);
			if (result.success) {
				setIsOpen(false);
				setAdditionalComment("");
				// You might want to update the UI or show a success message here
			}
		} catch (error) {
			console.error("Error reposting:", error);
		}
	};

	return (
		<Dialog open={isOpen} onOpenChange={setIsOpen}>
			<DialogTrigger asChild>
				<Button
					variant="ghost"
					className="text-[#00ba7c] hover:bg-[#00ba7c]/10"
				>
					リポスト
				</Button>
			</DialogTrigger>
			<DialogContent className="sm:max-w-[425px] bg-black text-white">
				<DialogHeader>
					<DialogTitle>リポスト</DialogTitle>
				</DialogHeader>
				<div className="grid gap-4 py-4">
					<Button
						onClick={() => {
							setIsQuoteRepost(false);
							handleRepost();
						}}
						className="w-full"
					>
						リポスト
					</Button>
					<Button onClick={() => setIsQuoteRepost(true)} className="w-full">
						引用リポスト
					</Button>
					{isQuoteRepost && (
						<Textarea
							placeholder={`repost from @${username}`}
							value={additionalComment}
							onChange={(e) => setAdditionalComment(e.target.value)}
							className="bg-black border-[#536471] focus:border-primary text-white"
						/>
					)}
					{isQuoteRepost && (
						<Button onClick={handleRepost} className="w-full">
							引用リポストを投稿
						</Button>
					)}
				</div>
			</DialogContent>
		</Dialog>
	);
}
