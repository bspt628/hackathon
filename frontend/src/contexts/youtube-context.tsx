'use client'

import React, { createContext, useContext, useState } from 'react'

interface YouTubeContextType {
  currentVideoId: string | null
  setCurrentVideoId: (id: string | null) => void
  searchResults: any[]
  setSearchResults: (results: any[]) => void
  copiedText: string
  setCopiedText: (text: string) => void
  isEnabled: boolean
}

const YouTubeContext = createContext<YouTubeContextType | undefined>(undefined)

export function YouTubeProvider({ children, isEnabled = true }: { children: React.ReactNode, isEnabled?: boolean }) {
  const [currentVideoId, setCurrentVideoId] = useState<string | null>(null)
  const [searchResults, setSearchResults] = useState<any[]>([])
  const [copiedText, setCopiedText] = useState<string>('')

  const contextValue = {
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

