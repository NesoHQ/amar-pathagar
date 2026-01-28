'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Layout from '@/components/Layout'
import { useAuthStore } from '@/store/authStore'
import { booksAPI } from '@/lib/api'

export default function AdminPage() {
  const router = useRouter()
  const { isAuthenticated, user, _hasHydrated } = useAuthStore()
  const [showAddBook, setShowAddBook] = useState(false)
  const [bookForm, setBookForm] = useState({
    title: '',
    author: '',
    isbn: '',
    description: '',
    category: '',
    physical_code: '',
  })

  useEffect(() => {
    if (_hasHydrated && !isAuthenticated) {
      router.push('/login')
    } else if (_hasHydrated && isAuthenticated && user?.role !== 'admin') {
      router.push('/dashboard')
    }
  }, [isAuthenticated, user, _hasHydrated, router])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await booksAPI.create(bookForm)
      setBookForm({
        title: '',
        author: '',
        isbn: '',
        description: '',
        category: '',
        physical_code: '',
      })
      setShowAddBook(false)
      alert('Book added successfully!')
    } catch (error: any) {
      alert(error.response?.data?.error || 'Failed to add book')
    }
  }

  if (!_hasHydrated || !isAuthenticated || user?.role !== 'admin') {
    return null
  }

  return (
    <Layout>
      <div className="space-y-6">
        {/* Header */}
        <div className="classic-card">
          <h1 className="text-4xl font-bold uppercase tracking-wider mb-2">⚙️ Admin Panel</h1>
          <p className="text-old-grey">Manage the library system</p>
        </div>

        {/* Quick Stats */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
          <div className="classic-card text-center">
            <p className="text-3xl font-bold">--</p>
            <p className="text-xs uppercase text-old-grey mt-2">Total Users</p>
          </div>
          <div className="classic-card text-center">
            <p className="text-3xl font-bold">--</p>
            <p className="text-xs uppercase text-old-grey mt-2">Total Books</p>
          </div>
          <div className="classic-card text-center">
            <p className="text-3xl font-bold">--</p>
            <p className="text-xs uppercase text-old-grey mt-2">Active Requests</p>
          </div>
          <div className="classic-card text-center">
            <p className="text-3xl font-bold">--</p>
            <p className="text-xs uppercase text-old-grey mt-2">In Circulation</p>
          </div>
        </div>

        {/* Add Book Section */}
        <div className="classic-card">
          <div className="flex justify-between items-center mb-4">
            <h2 className="text-2xl font-bold uppercase tracking-wider">Book Management</h2>
            <button
              onClick={() => setShowAddBook(!showAddBook)}
              className="classic-button"
            >
              {showAddBook ? 'Cancel' : 'Add New Book'}
            </button>
          </div>

          {showAddBook && (
            <form onSubmit={handleSubmit} className="space-y-4 p-4 border-2 border-old-border">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Title *</label>
                  <input
                    type="text"
                    value={bookForm.title}
                    onChange={(e) => setBookForm({ ...bookForm, title: e.target.value })}
                    className="classic-input"
                    required
                  />
                </div>
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Author *</label>
                  <input
                    type="text"
                    value={bookForm.author}
                    onChange={(e) => setBookForm({ ...bookForm, author: e.target.value })}
                    className="classic-input"
                    required
                  />
                </div>
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">ISBN</label>
                  <input
                    type="text"
                    value={bookForm.isbn}
                    onChange={(e) => setBookForm({ ...bookForm, isbn: e.target.value })}
                    className="classic-input"
                  />
                </div>
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Physical Code *</label>
                  <input
                    type="text"
                    value={bookForm.physical_code}
                    onChange={(e) => setBookForm({ ...bookForm, physical_code: e.target.value })}
                    className="classic-input"
                    required
                    placeholder="Unique identifier"
                  />
                </div>
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Category</label>
                  <input
                    type="text"
                    value={bookForm.category}
                    onChange={(e) => setBookForm({ ...bookForm, category: e.target.value })}
                    className="classic-input"
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm font-bold uppercase mb-2">Description</label>
                <textarea
                  value={bookForm.description}
                  onChange={(e) => setBookForm({ ...bookForm, description: e.target.value })}
                  className="classic-input"
                  rows={3}
                />
              </div>
              <button type="submit" className="classic-button">
                Add Book to Library
              </button>
            </form>
          )}
        </div>

        {/* Admin Actions */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="classic-card">
            <h3 className="font-bold uppercase tracking-wider mb-4">User Management</h3>
            <div className="space-y-2">
              <button className="w-full classic-button-secondary text-sm py-2">
                View All Users
              </button>
              <button className="w-full classic-button-secondary text-sm py-2">
                Adjust Success Scores
              </button>
              <button className="w-full classic-button-secondary text-sm py-2">
                View Audit Logs
              </button>
            </div>
          </div>

          <div className="classic-card">
            <h3 className="font-bold uppercase tracking-wider mb-4">Book Requests</h3>
            <div className="space-y-2">
              <button className="w-full classic-button-secondary text-sm py-2">
                Pending Requests
              </button>
              <button className="w-full classic-button-secondary text-sm py-2">
                Approve/Reject
              </button>
              <button className="w-full classic-button-secondary text-sm py-2">
                Dispute Resolution
              </button>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  )
}
