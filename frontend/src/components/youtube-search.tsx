"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Loader2 } from "lucide-react";

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

interface SearchResult {
	id: string;
	title: string;
	thumbnail: string;
}

export function YouTubeSearch({
	onVideoSelect,
}: {
	onVideoSelect: (videoId: string) => void;
}) {
	const [query, setQuery] = useState("");
	const [results, setResults] = useState<SearchResult[]>([]);
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
			setResults(
				data.items.map((item) => ({
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
					className="flex-grow"
				/>
				<Button onClick={handleSearch} disabled={isSearching}>
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
				{results.map((result) => (
					<div
						key={result.id}
						className="cursor-pointer hover:bg-gray-800 p-2 rounded"
						onClick={() => onVideoSelect(result.id)}
					>
						<img src={result.thumbnail} alt={result.title} className="w-full" />
						<p className="mt-2 text-sm">{result.title}</p>
					</div>
				))}
			</div>
		</div>
	);
}
