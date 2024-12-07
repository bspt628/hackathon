'use client'

import React, { createContext, useContext, useState } from 'react'

interface SearchResult {
  id: string;
  title: string;
  thumbnail: string;
}

interface YouTubeContextType {
  currentVideoId: string | null;
  setCurrentVideoId: (id: string | null) => void;
  searchResults: SearchResult[];
  setSearchResults: (results: SearchResult[]) => void;
  copiedText: string;
  setCopiedText: (text: string) => void;
  isEnabled: boolean;
}

const YouTubeContext = createContext<YouTubeContextType | undefined>(undefined)

export function YouTubeProvider({ children, isEnabled = true }: { children: React.ReactNode, isEnabled?: boolean }) {
  const [currentVideoId, setCurrentVideoId] = useState<string | null>(null)
  const [searchResults, setSearchResults] = useState<SearchResult[]>([])
  const [copiedText, setCopiedText] = useState<string>('')

  const contextValue: YouTubeContextType = {
    currentVideoId,
    setCurrentVideoId: isEnabled ? setCurrentVideoId : () => {},
    searchResults,
    setSearchResults: isEnabled ? setSearchResults : () => {},
    copiedText,
    setCopiedText: isEnabled ? setCopiedText : () => {},
    isEnabled
  }

  return (
    <YouTubeContext.Provider value={contextValue}>
      {children}
    </YouTubeContext.Provider>
  )
}

export function useYouTube() {
  const context = useContext(YouTubeContext)
  if (context === undefined) {
    throw new Error('useYouTube must be used within a YouTubeProvider')
  }
  return context
}

