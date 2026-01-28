'use client'

import { useEffect, useState } from 'react'
import { useParams, useRouter } from 'next/navigation'
import Layout from '@/components/Layout'
import { useAuthStore } from '@/store/authStore'
import { booksAPI, ideasAPI, reviewsAPI } from '@/lib/api'

export default function BookDetailPage() {
  const params = useParams()
  const router = useRouter()
  const { isAuthenticated, user, _hasHydrated } = useAuthStore()
  const [book, setBook] = useState<any>(null)
  const [ideas, setIdeas] = useState<any[]>([])
  const [showIdeaForm, setShowIdeaForm] = useState(false)
  const [ideaForm, setIdeaForm] = useState({ title: '', content: '' })

  useEffect(() => {
    if (_hasHydrated && !isAuthenticated) {
      router.push('/login')
    } else if (isAuthenticated && params.id) {
      loadBook()
      loadIdeas()
    }
  }, [isAuthenticated, _hasHydrated, params.id, router])

  const loadBook = async () => {
    try {
      const response = await booksAPI.getById(params.id as string)
      setBook(response.data.data || response.data)
    } catch (error) {
      console.error('Failed to load book:', error)
    }
  }

  const loadIdeas = async () => {
    try {
      const response = await ideasAPI.getByBook(params.id as string)
      const ideasData = response.data.data || response.data || []
      setIdeas(Array.isArray(ideasData) ? ideasData : [])
    } catch (error) {
      console.error('Failed to load ideas:', error)
      setIdeas([])
    }
  }

  const handleRequest = async () => {
    if (!user || user.success_score < 20) {
      alert('Your success score must be at least 20 to request books.')
      return
    }
    try {
      await booksAPI.request(params.id as string)
      alert('Book requested successfully!')
      loadBook()
    } catch (error: any) {
      alert(error.response?.data?.error || 'Failed to request book')
    }
  }

  const handleSubmitIdea = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await ideasAPI.create({
        book_id: params.id as string,
        ...ideaForm,
      })
      setIdeaForm({ title: '', content: '' })
      setShowIdeaForm(false)
      loadIdeas()
      alert('Idea posted! +3 points')
    } catch (error) {
      alert('Failed to post idea')
    }
  }

  const handleVote = async (ideaId: string, voteType: 'upvote' | 'downvote') => {
    try {
      await ideasAPI.vote(ideaId, voteType)
      loadIdeas()
    } catch (error) {
      console.error('Failed to vote:', error)
    }
  }

  if (!_hasHydrated || !isAuthenticated || !book) {
    return null
  }

  return (
    <Layout>
      <div className="space-y-6">
        {/* Book Header */}
        <div className="classic-card">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {/* Cover */}
            <div className="aspect-[3/4] bg-old-border flex items-center justify-center border-4 border-old-ink">
              {book.cover_url ? (
                <img src={book.cover_url} alt={book.title} className="w-full h-full object-cover" />
              ) : (
                <span className="text-8xl">📖</span>
              )}
            </div>

            {/* Details */}
            <div className="md:col-span-2 space-y-4">
              <div>
                <h1 className="text-4xl font-bold uppercase tracking-wider mb-2">{book.title}</h1>
                <p className="text-xl text-old-grey uppercase tracking-wider">By {book.author}</p>
              </div>

              {book.description && (
                <p className="text-old-grey leading-relaxed">{book.description}</p>
              )}

              <div className="flex flex-wrap gap-2">
                {book.category && (
                  <span className="vintage-badge">{book.category}</span>
                )}
                {book.tags?.map((tag: string) => (
                  <span key={tag} className="vintage-badge">{tag}</span>
                ))}
              </div>

              <div className="grid grid-cols-3 gap-4 pt-4 border-t-2 border-old-border">
                <div>
                  <p className="text-2xl font-bold">{book.total_reads}</p>
                  <p className="text-xs uppercase text-old-grey">Total Reads</p>
                </div>
                <div>
                  <p className="text-2xl font-bold">
                    {book.average_rating > 0 ? `★ ${book.average_rating.toFixed(1)}` : 'N/A'}
                  </p>
                  <p className="text-xs uppercase text-old-grey">Rating</p>
                </div>
                <div>
                  <p className="text-2xl font-bold uppercase">{book.status}</p>
                  <p className="text-xs uppercase text-old-grey">Status</p>
                </div>
              </div>

              {/* Actions */}
              <div className="space-y-2">
                {book.status === 'available' && (
                  <button onClick={handleRequest} className="w-full classic-button">
                    Request This Book
                  </button>
                )}
                {book.current_holder && (
                  <div className="p-3 border-2 border-old-ink bg-white">
                    <p className="text-sm font-bold uppercase">Currently With:</p>
                    <p className="text-old-grey">{book.current_holder.full_name || book.current_holder.username}</p>
                  </div>
                )}
              </div>
            </div>
          </div>
        </div>

        {/* Reading Ideas Section */}
        <div className="classic-card">
          <div className="flex justify-between items-center mb-6">
            <h2 className="text-2xl font-bold uppercase tracking-wider">Reading Ideas</h2>
            <button
              onClick={() => setShowIdeaForm(!showIdeaForm)}
              className="classic-button-secondary text-sm"
            >
              {showIdeaForm ? 'Cancel' : 'Share Your Thoughts'}
            </button>
          </div>

          {showIdeaForm && (
            <form onSubmit={handleSubmitIdea} className="mb-6 p-4 border-2 border-old-border bg-white">
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Title</label>
                  <input
                    type="text"
                    value={ideaForm.title}
                    onChange={(e) => setIdeaForm({ ...ideaForm, title: e.target.value })}
                    className="classic-input"
                    required
                  />
                </div>
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Your Thoughts</label>
                  <textarea
                    value={ideaForm.content}
                    onChange={(e) => setIdeaForm({ ...ideaForm, content: e.target.value })}
                    className="classic-input"
                    rows={4}
                    required
                  />
                </div>
                <button type="submit" className="classic-button">
                  Post Idea (+3 Points)
                </button>
              </div>
            </form>
          )}

          {ideas.length === 0 ? (
            <p className="text-center text-old-grey py-8">No ideas yet. Be the first to share!</p>
          ) : (
            <div className="space-y-4">
              {ideas.map((idea) => (
                <div key={idea.id} className="p-4 border-2 border-old-border bg-white">
                  <h3 className="font-bold uppercase mb-2">{idea.title}</h3>
                  <p className="text-old-grey mb-3">{idea.content}</p>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-old-grey">
                      By {idea.user?.username || 'Anonymous'}
                    </span>
                    <div className="flex gap-2">
                      <button
                        onClick={() => handleVote(idea.id, 'upvote')}
                        className="px-3 py-1 border border-old-ink hover:bg-old-ink hover:text-old-paper"
                      >
                        👍 {idea.upvotes}
                      </button>
                      <button
                        onClick={() => handleVote(idea.id, 'downvote')}
                        className="px-3 py-1 border border-old-ink hover:bg-old-ink hover:text-old-paper"
                      >
                        👎 {idea.downvotes}
                      </button>
                    </div>
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
