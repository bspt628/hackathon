"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Loader2 } from "lucide-react";
import { useYouTube } from "@/contexts/youtube-context";

interface YouTubeSearchResultItem {
	id: {
		videoId: string;
	};
	snippet: {
		title: string;
		thumbnails: {
			default: {
				url: string;
			};
		};
	};
}

interface YouTubeSearchResponse {
	items: YouTubeSearchResultItem[];
}

export function YouTubeSearch({
	onVideoSelect,
}: {
	onVideoSelect: (videoId: string) => void;
}) {
	const [query, setQuery] = useState("");
	const { searchResults, setSearchResults } = useYouTube();
	const [isSearching, setIsSearching] = useState(false);

	const handleSearch = async () => {
		console.log("Searching YouTube for:", query);
		setIsSearching(true);
		try {
			const response = await fetch(
				`/api/youtube-search?q=${encodeURIComponent(query)}`
			);
			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}

			const data: YouTubeSearchResponse = await response.json();
			setSearchResults(
				data.items.map((item: YouTubeSearchResultItem) => ({
					id: item.id.videoId,
					title: item.snippet.title,
					thumbnail: item.snippet.thumbnails.default.url,
				}))
			);
		} catch (error) {
			console.error("Error searching YouTube:", error);
		} finally {
			setIsSearching(false);
		}
	};

	return (
		<div className="space-y-4">
			<div className="flex space-x-2">
				<Input
					type="text"
					placeholder="Search YouTube"
					value={query}
					onChange={(e) => setQuery(e.target.value)}
					className="flex-grow text-white placeholder-gray-400"
				/>
				<Button
					onClick={handleSearch}
					disabled={isSearching}
					className="bg-[#1d9bf0] hover:bg-[#1a8cd8] text-white"
				>
					{isSearching ? (
						<>
							<Loader2 className="mr-2 h-4 w-4 animate-spin" />
							検索中...
						</>
					) : (
						"検索"
					)}
				</Button>
			</div>
			<div className="grid grid-cols-2 gap-4">
				{searchResults.map((result) => (
					<div
						key={result.id}
						className="cursor-pointer hover:bg-[#1d9bf0]/10 p-2 rounded"
						onClick={() => onVideoSelect(result.id)}
					>
						<img src={result.thumbnail} alt={result.title} className="w-full" />
						<p className="mt-2 text-sm text-white">{result.title}</p>
					</div>
				))}
			</div>
		</div>
	);
}
