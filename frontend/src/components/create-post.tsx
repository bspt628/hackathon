"use client";

import { useState, useEffect, useCallback } from "react";
import { useAuth } from "@/contexts/auth-context";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import {
	Sparkles,
	Loader2,
} from "lucide-react";
import { createPost } from "@/app/actions/create-post";
import { useYouTube } from "@/contexts/youtube-context";

interface CreatePostProps {
	onPostSuccess?: () => void;
	initialContent?: string;
}

export function CreatePost({
	onPostSuccess,
	initialContent = "",
}: CreatePostProps) {
	const [content, setContent] = useState(initialContent);
	const [isLoading, setIsLoading] = useState(false);
	const [isGenerating, setIsGenerating] = useState(false);
	const { idToken } = useAuth();
	const { currentVideoId, searchResults } = useYouTube();

	useEffect(() => {
		setContent(initialContent);
	}, [initialContent]);

	const generateContent = useCallback(async () => {
		console.log("Generating content...");
		if (!idToken || !currentVideoId) return;
		setIsGenerating(true);
		setIsLoading(true);  // ここで isLoading を true にセット
	
		const videoTitle =
			searchResults.find((result) => result.id === currentVideoId)?.title || "";
	
		try {
			const response = await fetch(
				"https://hackathon-uchida-hiroto-241499864821.us-central1.run.app/api/generate-content",
				{
					method: "POST",
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${idToken}`,
					},
					body: JSON.stringify({
						prompt: `Generate an appealing description of the video titled "${videoTitle}" in Japanese. Highlight its key features and why it's worth watching. translation to english is not required.`,
					}),
				}
			);
	
			if (!response.ok) {
				throw new Error("Failed to generate content");
			}
	
			const data = await response.json();
			const generatedContent = data.Candidates[0].Content.Parts[0];
			setContent((prevContent) => prevContent + generatedContent);
		} catch (error) {
			console.error("Error generating content:", error);
		} finally {
			setIsGenerating(false);  // isGenerating を false にセット
			setIsLoading(false);     // isLoading を false にセット
		}
	}, [idToken, currentVideoId, searchResults]);
	

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
						<div className="flex space-x-2">
							<Button
								onClick={generateContent}
								variant="outline"
								size="sm"
								className="rounded-full text-primary hover:bg-primary/10"
								disabled={!currentVideoId || isGenerating}
							>
								{isGenerating ? (
									<>
										<Loader2 className="w-4 h-4 mr-2 animate-spin" />
										生成中...
									</>
								) : (
									<>
										<Sparkles className="w-5 h-5 mr-2" />
										魅力を生成
									</>
								)}
							</Button>
							<Button
								onClick={handleSubmit}
								disabled={!content || isLoading}
								className="rounded-full bg-secondary hover:bg-secondary/90 text-white px-4"
							>
								{isLoading ? "投稿中..." : "ポストする"}
							</Button>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}
