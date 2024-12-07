"use client";

import { useEffect, useRef, useState } from "react";
import { X, Play, Pause, Loader2, Copy } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Slider } from "@/components/ui/slider";
import { useAuth } from "@/contexts/auth-context";

interface AudioPlayerProps {
	videoId: string | null;
	onClose: () => void;
	onCopy: (text: string) => void;
}

export function AudioPlayer({ videoId, onClose, onCopy }: AudioPlayerProps) {
	const playerRef = useRef<YT.Player | null>(null);
	const [isPlaying, setIsPlaying] = useState(false);
	const [title, setTitle] = useState("");
	const [duration, setDuration] = useState(0);
	const [currentTime, setCurrentTime] = useState(0);
	const [isLoading, setIsLoading] = useState(true);
	const intervalRef = useRef<NodeJS.Timeout | null>(null);
	const { user } = useAuth();

	useEffect(() => {
		if (typeof window !== "undefined" && !window.YT) {
			const tag = document.createElement("script");
			tag.src = "https://www.youtube.com/iframe_api";
			const firstScriptTag = document.getElementsByTagName("script")[0];
			firstScriptTag.parentNode?.insertBefore(tag, firstScriptTag);

			window.onYouTubeIframeAPIReady = initPlayer;
		} else if (videoId) {
			setIsPlaying(false);
			setIsLoading(true);
			initPlayer();
		}

		return () => {
			if (playerRef.current) {
				playerRef.current.destroy();
			}
			if (intervalRef.current) {
				clearInterval(intervalRef.current);
			}
		};
	}, [videoId]);

	function initPlayer() {
		if (typeof window.YT !== "undefined" && window.YT.Player && videoId) {
			setIsLoading(true);
			playerRef.current = new window.YT.Player("youtube-audio-player", {
				height: "0",
				width: "0",
				videoId: videoId,
				playerVars: {
					autoplay: 0,
					controls: 0,
				},
				events: {
					onReady: (event: YT.OnReadyEvent) => {
						const player = event.target;
						setTitle(player.getVideoData().title);
						setDuration(player.getDuration());
						startTimeUpdate();
						setIsLoading(false);
						player.playVideo();
						setIsPlaying(true);
					},
					onStateChange: (event: YT.OnStateChangeEvent) => {
						const newState = event.data;
						setIsPlaying(
							newState === YT.PlayerState.PLAYING ||
								newState === YT.PlayerState.BUFFERING
						);
						if (newState === YT.PlayerState.PLAYING) {
							startTimeUpdate();
						} else if (
							newState === YT.PlayerState.PAUSED ||
							newState === YT.PlayerState.ENDED
						) {
							stopTimeUpdate();
						}
					},
				},
			});
		}
	}

	const startTimeUpdate = () => {
		if (intervalRef.current) clearInterval(intervalRef.current);
		intervalRef.current = setInterval(() => {
			if (playerRef.current) {
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
		if (playerRef.current) {
			if (isPlaying) {
				playerRef.current.pauseVideo();
			} else {
				playerRef.current.playVideo();
			}
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
		if (playerRef.current && user) {
			const currentTime = playerRef.current.getCurrentTime();
			const formattedTime = formatTime(currentTime);
			const videoUrl = `https://www.youtube.com/watch?v=${videoId}`;
			const copyText = `${user.displayName} is playing ${title} ${formattedTime} \n see from here \n (${videoUrl})`;

			navigator.clipboard
				.writeText(copyText)
				.then(() => {
					console.log("Text copied to clipboard");
				})
				.catch((err) => {
					console.error("Failed to copy text: ", err);
				});

			onCopy(copyText);
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
							<Button onClick={togglePlayPause} variant="ghost" size="icon">
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
		</div>
	);
}
