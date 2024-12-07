'use client'

import { useAuth } from "@/contexts/auth-context";
import { YouTubeProvider } from "@/contexts/youtube-context";
import { usePathname } from 'next/navigation';

export function ConditionalYouTubeProvider({ children }: { children: React.ReactNode }) {
  const { idToken } = useAuth();
  const pathname = usePathname();

  const isAuthRoute = ['/', '/signup', '/login'].includes(pathname);
  const isYouTubeEnabled = !!idToken && !isAuthRoute;

  return (
    <YouTubeProvider isEnabled={isYouTubeEnabled}>
      {children}
    </YouTubeProvider>
  );
}

