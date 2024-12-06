"use client";

import { useEffect, useRef } from "react";

export function AudioPlayer({ videoId }: { videoId: string }) {
	const playerRef = useRef<YT.Player | null>(null);

	useEffect(() => {
		if (typeof window !== "undefined" && !window.YT) {
			const tag = document.createElement("script");
			tag.src = "https://www.youtube.com/iframe_api";
			const firstScriptTag = document.getElementsByTagName("script")[0];
			firstScriptTag.parentNode?.insertBefore(tag, firstScriptTag);

			window.onYouTubeIframeAPIReady = initPlayer;
		} else {
			initPlayer();
		}

		return () => {
			if (playerRef.current) {
				playerRef.current.destroy();
			}
		};
	}, [videoId]);

	function initPlayer() {
		if (typeof window.YT !== "undefined" && window.YT.Player) {
			playerRef.current = new window.YT.Player("youtube-audio-player", {
				height: "0",
				width: "0",
				videoId: videoId,
				playerVars: {
					autoplay: 1,
					controls: 0,
				},
				events: {
					onReady: (event: YT.OnReadyEvent) => {
						event.target.playVideo();
					},
				},
			});
		}
	}

	return <div id="youtube-audio-player"></div>;
}
