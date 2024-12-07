interface YT {
  Player: {
    new (elementId: string, options: YT.PlayerOptions): YT.Player;
  };
}

declare namespace YT {
  interface PlayerOptions {
    height?: string | number;
    width?: string | number;
    videoId?: string;
    playerVars?: {
      autoplay?: 0 | 1;
      controls?: 0 | 1;
      loop?: 0 | 1;
      mute?: 0 | 1;
      playlist?: string;
      rel?: 0 | 1;
      start?: number;
      end?: number;
      modestbranding?: 0 | 1;
      playsinline?: 0 | 1;
      origin?: string;
      enablejsapi?: 0 | 1;
    } & Record<string, string | number | boolean>;
    events?: {
      onReady?: (event: OnReadyEvent) => void;
      onStateChange?: (event: OnStateChangeEvent) => void;
      onPlaybackQualityChange?: (event: OnPlaybackQualityChangeEvent) => void;
      onPlaybackRateChange?: (event: OnPlaybackRateChangeEvent) => void;
      onError?: (event: OnErrorEvent) => void;
      onApiChange?: (event: OnApiChangeEvent) => void;
    };
  }

  interface Player {
    destroy(): void;
    getVideoData(): VideoData;
    playVideo(): void;
    pauseVideo(): void;
    stopVideo(): void;
    seekTo(seconds: number, allowSeekAhead: boolean): void;
    getPlayerState(): number;
    getCurrentTime(): number;
    getDuration(): number;
    getVideoUrl(): string;
    getVideoEmbedCode(): string;
    getVolume(): number;
    setVolume(volume: number): void;
    mute(): void;
    unMute(): void;
    isMuted(): boolean;
    setSize(width: number, height: number): void;
    getPlaybackRate(): number;
    setPlaybackRate(suggestedRate: number): void;
    getAvailablePlaybackRates(): number[];
    getPlaybackQuality(): string;
    setPlaybackQuality(suggestedQuality: string): void;
    getAvailableQualityLevels(): string[];
    addEventListener<T extends keyof YT.Events>(event: T, listener: (event: YT.Events[T]) => void): void;
    removeEventListener<T extends keyof YT.Events>(event: T, listener: (event: YT.Events[T]) => void): void;
  }

  interface VideoData {
    title: string;
    author: string;
    video_id: string;
  }

  interface PlayerEvent {
    target: Player;
  }

  interface OnReadyEvent extends PlayerEvent {}
  interface OnStateChangeEvent extends PlayerEvent {
    data: number;
  }
  interface OnPlaybackQualityChangeEvent extends PlayerEvent {
    data: string;
  }
  interface OnPlaybackRateChangeEvent extends PlayerEvent {
    data: number;
  }
  interface OnErrorEvent extends PlayerEvent {
    data: number;
  }
  interface OnApiChangeEvent extends PlayerEvent {}

  interface Events {
    OnReady: OnReadyEvent;
    OnStateChange: OnStateChangeEvent;
    OnPlaybackQualityChange: OnPlaybackQualityChangeEvent;
    OnPlaybackRateChange: OnPlaybackRateChangeEvent;
    OnError: OnErrorEvent;
    OnApiChange: OnApiChangeEvent;
  }

  enum PlayerState {
    UNSTARTED = -1,
    ENDED = 0,
    PLAYING = 1,
    PAUSED = 2,
    BUFFERING = 3,
    CUED = 5
  }
}

interface Window {
  YT: typeof YT;
  onYouTubeIframeAPIReady: () => void;
}

