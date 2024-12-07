"use client";

import { useYouTube } from "@/contexts/youtube-context";
import { AudioPlayer } from "@/components/audio-player";

export function AudioPlayerWrapper() {
	const { currentVideoId, setCurrentVideoId, setCopiedText, isEnabled } =
		useYouTube();

	if (!isEnabled || !currentVideoId) {
		return null;
	}

	return (
		<AudioPlayer
			videoId={currentVideoId}
			onClose={() => setCurrentVideoId(null)}
			onCopy={(text) => setCopiedText(text)}
		/>
	);
}
