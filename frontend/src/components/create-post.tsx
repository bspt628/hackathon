"use client";

import { useState, useEffect } from "react";
import { useAuth } from "@/contexts/auth-context";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { Image, SmilePlus, MapPin, Calendar, ListFilter } from "lucide-react";
import { createPost } from "@/app/actions/create-post";

interface CreatePostProps {
	onPostSuccess?: () => void;
	initialContent?: string;
}

export function CreatePost({ onPostSuccess, initialContent="" }: CreatePostProps) {
	const [content, setContent] = useState(initialContent);
	const [isLoading, setIsLoading] = useState(false);
	const { idToken } = useAuth();

	useEffect(() => {
		setContent(initialContent);
	},	[initialContent]);

	const handleSubmit = async () => {
		if (!content || !idToken) return;

		setIsLoading(true);
		try {
			const result = await createPost(content, idToken);
			if (result.success) {
				setContent("");
				onPostSuccess?.();
			}
		} catch (error) {
			console.error("Error creating post:", error);
		} finally {
			setIsLoading(false);
		}
	};

	return (
		<div className="border-b border-[#2f3336] p-4">
			<div className="flex gap-4">
				<div className="w-10 h-10 rounded-full bg-[#2f3336] shrink-0" />
				<div className="flex-1 min-w-0">
					<Textarea
						value={content}
						onChange={(e) => setContent(e.target.value)}
						placeholder="いまどうしてる？"
						className="min-h-[80px] w-full resize-none bg-transparent border-none p-0 placeholder:text-[#71767b] focus-visible:ring-0"
						rows={6}
					/>
					<div className="flex items-center justify-between mt-4">
						<div className="flex -ml-2">
							<Button
								size="icon"
								variant="ghost"
								className="rounded-full text-[#1d9bf0] hover:bg-[#1d9bf0]/10"
							>
								<Image className="w-5 h-5" />
								<span className="sr-only">Add image</span>
							</Button>
							<Button
								size="icon"
								variant="ghost"
								className="rounded-full text-[#1d9bf0] hover:bg-[#1d9bf0]/10"
							>
								<ListFilter className="w-5 h-5" />
								<span className="sr-only">Add GIF</span>
							</Button>
							<Button
								size="icon"
								variant="ghost"
								className="rounded-full text-[#1d9bf0] hover:bg-[#1d9bf0]/10"
							>
								<Calendar className="w-5 h-5" />
								<span className="sr-only">Add poll</span>
							</Button>
							<Button
								size="icon"
								variant="ghost"
								className="rounded-full text-[#1d9bf0] hover:bg-[#1d9bf0]/10"
							>
								<SmilePlus className="w-5 h-5" />
								<span className="sr-only">Add emoji</span>
							</Button>
							<Button
								size="icon"
								variant="ghost"
								className="rounded-full text-[#1d9bf0] hover:bg-[#1d9bf0]/10"
							>
								<MapPin className="w-5 h-5" />
								<span className="sr-only">Add location</span>
							</Button>
						</div>
						<Button
							onClick={handleSubmit}
							disabled={!content || isLoading}
							className="rounded-full bg-[#1d9bf0] hover:bg-[#1a8cd8] text-white px-4"
						>
							{isLoading ? "投稿中..." : "ポストする"}
						</Button>
					</div>
				</div>
			</div>
		</div>
	);
}
