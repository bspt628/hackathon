"use client";

import { useEffect, useRef, useState } from "react";
import { X, Play, Pause, Loader2, Copy } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Slider } from "@/components/ui/slider";
import { useAuth } from "@/contexts/auth-context";

interface AudioPlayerProps {
	videoId: string | null;
	onClose: () => void;
	onCopy?: (text: string) => void;
}

export function AudioPlayer({ videoId, onClose, onCopy }: AudioPlayerProps) {
	const playerRef = useRef<YT.Player | null>(null);
	const [isPlaying, setIsPlaying] = useState(false);
	const [title, setTitle] = useState("");
	const [duration, setDuration] = useState(0);
	const [currentTime, setCurrentTime] = useState(0);
	const [isLoading, setIsLoading] = useState(true);
	const [isPlayerReady, setIsPlayerReady] = useState(false);
	const intervalRef = useRef<NodeJS.Timeout | null>(null);
	const { user } = useAuth();

	useEffect(() => {
		const loadYouTubeAPI = () => {
			if (typeof window !== "undefined" && !window.YT) {
				const tag = document.createElement("script");
				tag.src = "https://www.youtube.com/iframe_api";
				const firstScriptTag = document.getElementsByTagName("script")[0];
				firstScriptTag.parentNode?.insertBefore(tag, firstScriptTag);

				return new Promise<void>((resolve) => {
					window.onYouTubeIframeAPIReady = () => {
						resolve();
					};
				});
			}
			return Promise.resolve();
		};

		const initializePlayer = async () => {
			setIsLoading(true);
			setIsPlaying(false);
			setCurrentTime(0);
			setDuration(0);
			setTitle("");
			setIsPlayerReady(false);

			try {
				await loadYouTubeAPI();
				if (playerRef.current) {
					playerRef.current.destroy();
				}
				initPlayer();
			} catch (error) {
				console.error("Failed to initialize player:", error);
				setIsLoading(false);
			}
		};

		if (videoId) {
			initializePlayer();
		}

		return () => {
			if (
				playerRef.current &&
				typeof playerRef.current.destroy === "function"
			) {
				playerRef.current.destroy();
			}
			playerRef.current = null;
			setIsPlayerReady(false);
			if (intervalRef.current) {
				clearInterval(intervalRef.current);
			}
		};
	}, [videoId]);

	function initPlayer() {
		if (typeof window.YT !== "undefined" && window.YT.Player && videoId) {
			console.log("Initializing YouTube player with video ID:", videoId);
			const videoIdOrUrl = videoId.includes("youtube.com")
				? videoId
				: `https://www.youtube.com/watch?v=${videoId}`;
			playerRef.current = new window.YT.Player("youtube-audio-player", {
				height: "0",
				width: "0",
				videoId: videoId.includes("youtube.com") ? undefined : videoId,
				playerVars: {
					autoplay: 1,
					controls: 0,
					...(videoId.includes("youtube.com")
						? { origin: window.location.origin }
						: {}),
				},
				events: {
					onReady: (event: YT.OnReadyEvent) => {
						console.log("YouTube player is ready");
						const player = event.target;
						if (videoId.includes("youtube.com")) {
							player.loadVideoByUrl(videoIdOrUrl);
						}
						setTitle(player.getVideoData().title);
						setDuration(player.getDuration());
						setIsLoading(false);
						setIsPlayerReady(true);
						setIsPlaying(false);
					},
					onStateChange: (event: YT.OnStateChangeEvent) => {
						console.log("YouTube player state changed:", event.data);
						const newState = event.data;
						setIsPlaying(
							newState === YT.PlayerState.PLAYING ||
								newState === YT.PlayerState.BUFFERING
						);
						if (newState === YT.PlayerState.PLAYING) {
							startTimeUpdate();
							setIsLoading(false);
						} else if (
							newState === YT.PlayerState.PAUSED ||
							newState === YT.PlayerState.ENDED
						) {
							stopTimeUpdate();
						}
					},
					onError: (event: YT.OnErrorEvent) => {
						console.error("YouTube player error:", event.data);
						setIsLoading(false);
						setIsPlaying(false);
					},
				},
			});
		} else {
			console.error("YouTube API is not loaded or videoId is missing");
			setIsLoading(false);
		}
	}

	const startTimeUpdate = () => {
		if (intervalRef.current) clearInterval(intervalRef.current);
		intervalRef.current = setInterval(() => {
			if (
				playerRef.current &&
				typeof playerRef.current.getCurrentTime === "function"
			) {
				setCurrentTime(playerRef.current.getCurrentTime());
			}
		}, 1000);
	};

	const stopTimeUpdate = () => {
		if (intervalRef.current) {
			clearInterval(intervalRef.current);
		}
	};

	const togglePlayPause = () => {
		if (
			isPlayerReady &&
			playerRef.current &&
			typeof playerRef.current.getPlayerState === "function"
		) {
			try {
				if (isPlaying) {
					playerRef.current.pauseVideo();
				} else {
					playerRef.current.playVideo();
				}
				setIsPlaying(!isPlaying);
			} catch (error) {
				console.error("Error toggling play/pause:", error);
				setIsPlaying(false);
			}
		} else {
			console.error("Player is not ready");
		}
	};

	const handleSeek = (value: number[]) => {
		if (playerRef.current) {
			playerRef.current.seekTo(value[0], true);
		}
	};

	const formatTime = (time: number) => {
		const minutes = Math.floor(time / 60);
		const seconds = Math.floor(time % 60);
		return `${minutes}:${seconds.toString().padStart(2, "0")}`;
	};

	const handleCopy = () => {
		if (
			playerRef.current &&
			user &&
			typeof playerRef.current.getCurrentTime === "function"
		) {
			const currentTime = playerRef.current.getCurrentTime();
			const formattedTime = formatTime(currentTime);
			if (!videoId) return;
			const videoUrl = videoId.includes("youtube.com")
				? videoId
				: `https://www.youtube.com/watch?v=${videoId}`;
				const copyText = `🎵 ${user.displayName} is enjoying *${title}* 🎶 \n ⏱️ ${formattedTime}\n 🔗 [Listen now](${videoUrl})`;
				

			navigator.clipboard
				.writeText(copyText)
				.then(() => {
					console.log("Text copied to clipboard");
				})
				.catch((err) => {
					console.error("Failed to copy text: ", err);
				});

			onCopy?.(copyText);
		}
	};

	if (!videoId) return null;

	return (
		<div className="fixed bottom-0 left-0 right-0 bg-gray-900 text-white p-4 flex flex-col">
			<div id="youtube-audio-player"></div>
			<div className="flex items-center justify-between mb-2">
				<div className="flex items-center space-x-4">
					{isLoading ? (
						<div className="flex items-center">
							<Loader2 className="animate-spin h-4 w-4 mr-2" />
							<span>ロード中...</span>
						</div>
					) : (
						<>
							<Button
								onClick={togglePlayPause}
								variant="ghost"
								size="icon"
								disabled={!isPlayerReady}
							>
								{isPlaying ? (
									<Pause className="h-4 w-4" />
								) : (
									<Play className="h-4 w-4" />
								)}
							</Button>
							<div className="flex flex-col">
								<span className="font-semibold">{title}</span>
								<span className="text-sm text-gray-400">
									{formatTime(currentTime)} / {formatTime(duration)}
								</span>
							</div>
						</>
					)}
				</div>
				<div className="flex items-center space-x-2">
					<Button onClick={handleCopy} variant="ghost" size="icon">
						<Copy className="h-4 w-4" />
					</Button>
					<Button onClick={onClose} variant="ghost" size="icon">
						<X className="h-4 w-4" />
					</Button>
				</div>
			</div>
			{!isLoading && (
				<Slider
					value={[currentTime]}
					max={duration}
					step={1}
					onValueChange={handleSeek}
					className="w-full"
				/>
			)}
			{(isLoading || !isPlayerReady) && (
				<div className="text-xs text-gray-400 mt-2">
					Debug: {isLoading ? "Loading..." : "Waiting for player to be ready"}
					Player state: {isPlayerReady ? "Ready" : "Not ready"}, Title: {title},
					Duration: {duration}, Current Time: {currentTime}, Is Playing:{" "}
					{isPlaying ? "Yes" : "No"}
				</div>
			)}
		</div>
	);
}
