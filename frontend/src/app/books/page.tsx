'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Layout from '@/components/Layout'
import { useAuthStore } from '@/store/authStore'
import { booksAPI, bookmarksAPI } from '@/lib/api'

interface Book {
  id: string
  title: string
  author: string
  cover_url: string
  category: string
  status: string
  current_holder?: any
  total_reads: number
  average_rating: number
}

export default function BooksPage() {
  const router = useRouter()
  const { isAuthenticated, _hasHydrated } = useAuthStore()
  const [books, setBooks] = useState<Book[]>([])
  const [loading, setLoading] = useState(true)
  const [search, setSearch] = useState('')
  const [statusFilter, setStatusFilter] = useState('')

  useEffect(() => {
    if (_hasHydrated && !isAuthenticated) {
      router.push('/login')
    } else if (isAuthenticated) {
      loadBooks()
    }
  }, [isAuthenticated, _hasHydrated, router])

  const loadBooks = async () => {
    try {
      const response = await booksAPI.getAll({ search, status: statusFilter })
      const booksData = response.data.data || response.data || []
      setBooks(Array.isArray(booksData) ? booksData : [])
    } catch (error) {
      console.error('Failed to load books:', error)
      setBooks([])
    } finally {
      setLoading(false)
    }
  }

  const handleSearch = () => {
    loadBooks()
  }

  const handleBookmark = async (bookId: string, type: string) => {
    try {
      await bookmarksAPI.create({ book_id: bookId, bookmark_type: type })
      alert(`Book ${type}ed successfully!`)
    } catch (error) {
      console.error('Failed to bookmark:', error)
    }
  }

  if (!_hasHydrated || !isAuthenticated) {
    return null
  }

  return (
    <Layout>
      <div className="space-y-6">
        {/* Header */}
        <div className="classic-card">
          <h1 className="text-4xl font-bold uppercase tracking-wider mb-2">Book Collection</h1>
          <p className="text-old-grey">Browse and discover books in our community library</p>
        </div>

        {/* Search & Filters */}
        <div className="classic-card">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div className="md:col-span-2">
              <label className="block text-sm font-bold uppercase tracking-wider mb-2">
                Search Books
              </label>
              <div className="flex gap-2">
                <input
                  type="text"
                  value={search}
                  onChange={(e) => setSearch(e.target.value)}
                  onKeyPress={(e) => e.key === 'Enter' && handleSearch()}
                  placeholder="Title, author, or topic..."
                  className="classic-input"
                />
                <button onClick={handleSearch} className="classic-button px-8">
                  Search
                </button>
              </div>
            </div>
            <div>
              <label className="block text-sm font-bold uppercase tracking-wider mb-2">
                Status Filter
              </label>
              <select
                value={statusFilter}
                onChange={(e) => setStatusFilter(e.target.value)}
                className="classic-input"
              >
                <option value="">All Books</option>
                <option value="available">Available</option>
                <option value="reading">In Circulation</option>
                <option value="requested">Requested</option>
              </select>
            </div>
          </div>
        </div>

        {/* Books Grid */}
        {loading ? (
          <div className="text-center py-12">
            <p className="text-old-grey uppercase tracking-wider">Loading books...</p>
          </div>
        ) : books.length === 0 ? (
          <div className="classic-card text-center py-12">
            <p className="text-old-grey uppercase tracking-wider">No books found</p>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {books.map((book) => (
              <BookCard key={book.id} book={book} onBookmark={handleBookmark} />
            ))}
          </div>
        )}
      </div>
    </Layout>
  )
}

function BookCard({ book, onBookmark }: { book: Book; onBookmark: (id: string, type: string) => void }) {
  const router = useRouter()
  
  const getStatusBadge = (status: string) => {
    const badges = {
      available: { text: 'Available', class: 'border-green-600 text-green-600' },
      reading: { text: 'In Circulation', class: 'border-blue-600 text-blue-600' },
      requested: { text: 'Requested', class: 'border-orange-600 text-orange-600' },
    }
    const badge = badges[status as keyof typeof badges] || badges.available
    return <span className={`stamp ${badge.class}`}>{badge.text}</span>
  }

  return (
    <div className="classic-card hover:shadow-[6px_6px_0px_0px_rgba(0,0,0,0.15)] transition-shadow">
      {/* Book Cover */}
      <div className="aspect-[3/4] bg-old-border mb-4 flex items-center justify-center border-2 border-old-ink">
        {book.cover_url ? (
          <img src={book.cover_url} alt={book.title} className="w-full h-full object-cover" />
        ) : (
          <span className="text-6xl">📖</span>
        )}
      </div>

      {/* Book Info */}
      <div className="space-y-2">
        <div className="flex items-start justify-between">
          <h3 className="font-bold uppercase tracking-wider text-lg leading-tight flex-1">
            {book.title}
          </h3>
          {getStatusBadge(book.status)}
        </div>
        
        <p className="text-old-grey text-sm uppercase tracking-wider">
          By {book.author}
        </p>

        {book.category && (
          <p className="text-xs text-old-grey uppercase">
            Category: {book.category}
          </p>
        )}

        <div className="flex items-center justify-between text-xs text-old-grey pt-2 border-t border-old-border">
          <span>{book.total_reads} reads</span>
          {book.average_rating > 0 && (
            <span>★ {book.average_rating.toFixed(1)}</span>
          )}
        </div>
      </div>

      {/* Actions */}
      <div className="mt-4 space-y-2">
        <button
          onClick={() => router.push(`/books/${book.id}`)}
          className="w-full classic-button text-sm py-2"
        >
          View Details
        </button>
        <div className="flex gap-2">
          <button
            onClick={() => onBookmark(book.id, 'like')}
            className="flex-1 classic-button-secondary text-xs py-1"
            title="Like"
          >
            ❤️
          </button>
          <button
            onClick={() => onBookmark(book.id, 'bookmark')}
            className="flex-1 classic-button-secondary text-xs py-1"
            title="Bookmark"
          >
            🔖
          </button>
          <button
            onClick={() => onBookmark(book.id, 'priority')}
            className="flex-1 classic-button-secondary text-xs py-1"
            title="Priority"
          >
            ⭐
          </button>
        </div>
      </div>
    </div>
  )
}
