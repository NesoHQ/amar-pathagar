'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Layout from '@/components/Layout'
import { useAuthStore } from '@/store/authStore'
import { bookmarksAPI } from '@/lib/api'

export default function MyLibraryPage() {
  const router = useRouter()
  const { isAuthenticated, user, _hasHydrated } = useAuthStore()
  const [bookmarks, setBookmarks] = useState<any[]>([])
  const [activeTab, setActiveTab] = useState('all')

  useEffect(() => {
    if (_hasHydrated && !isAuthenticated) {
      router.push('/login')
    } else if (isAuthenticated) {
      loadBookmarks()
    }
  }, [isAuthenticated, _hasHydrated, router])

  const loadBookmarks = async () => {
    try {
      const response = await bookmarksAPI.getAll()
      const bookmarksData = response.data.data || response.data || []
      setBookmarks(Array.isArray(bookmarksData) ? bookmarksData : [])
    } catch (error) {
      console.error('Failed to load bookmarks:', error)
      setBookmarks([])
    }
  }

  const filteredBookmarks = activeTab === 'all'
    ? bookmarks
    : bookmarks.filter(b => b.bookmark_type === activeTab)

  if (!_hasHydrated || !isAuthenticated || !user) {
    return null
  }

  return (
    <Layout>
      <div className="space-y-6">
        {/* Header */}
        <div className="classic-card">
          <h1 className="text-4xl font-bold uppercase tracking-wider mb-2">📖 My Library</h1>
          <p className="text-old-grey">Your personal collection and bookmarks</p>
        </div>

        {/* Stats */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div className="classic-card text-center">
            <p className="text-4xl font-bold">{user.books_received}</p>
            <p className="text-sm uppercase text-old-grey mt-2">Books Read</p>
          </div>
          <div className="classic-card text-center">
            <p className="text-4xl font-bold">{user.books_shared}</p>
            <p className="text-sm uppercase text-old-grey mt-2">Books Shared</p>
          </div>
          <div className="classic-card text-center">
            <p className="text-4xl font-bold">{bookmarks.length}</p>
            <p className="text-sm uppercase text-old-grey mt-2">Bookmarked</p>
          </div>
        </div>

        {/* Tabs */}
        <div className="flex flex-wrap gap-2">
          {[
            { key: 'all', label: 'All' },
            { key: 'like', label: 'Liked' },
            { key: 'bookmark', label: 'Bookmarked' },
            { key: 'priority', label: 'Priority' },
          ].map((tab) => (
            <button
              key={tab.key}
              onClick={() => setActiveTab(tab.key)}
              className={`px-4 py-2 font-bold uppercase text-sm tracking-wider transition-colors
                ${activeTab === tab.key
                  ? 'bg-old-ink text-old-paper'
                  : 'bg-white border-2 border-old-border hover:border-old-ink'
                }`}
            >
              {tab.label}
            </button>
          ))}
        </div>

        {/* Bookmarks */}
        <div className="classic-card">
          {filteredBookmarks.length === 0 ? (
            <p className="text-center text-old-grey py-8">No bookmarks yet</p>
          ) : (
            <div className="space-y-3">
              {filteredBookmarks.map((bookmark: any) => (
                <div key={bookmark.id} className="p-4 border-2 border-old-border hover:border-old-ink transition-colors">
                  <div className="flex justify-between items-start">
                    <div>
                      <h3 className="font-bold uppercase">{bookmark.book?.title || 'Unknown Book'}</h3>
                      <p className="text-sm text-old-grey">
                        {bookmark.book?.author || 'Unknown Author'}
                      </p>
                      <div className="mt-2 flex gap-2">
                        <span className="vintage-badge">{bookmark.bookmark_type}</span>
                        {bookmark.priority_level > 0 && (
                          <span className="vintage-badge">Priority: {bookmark.priority_level}</span>
                        )}
                      </div>
                    </div>
                    <button
                      onClick={() => router.push(`/books/${bookmark.book_id}`)}
                      className="classic-button-secondary text-sm"
                    >
                      View
                    </button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
    </Layout>
  )
}
