'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'
import { authAPI } from '@/lib/api'
import { useAuthStore } from '@/store/authStore'

export default function LoginPage() {
  const router = useRouter()
  const { setAuth } = useAuthStore()
  const [formData, setFormData] = useState({
    username: '',
    password: '',
  })
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError('')
    setLoading(true)

    try {
      const response = await authAPI.login(formData)
      const { data } = response.data
      const { user, access_token } = data
      setAuth(user, access_token)
      router.push('/dashboard')
    } catch (err: any) {
      setError(err.response?.data?.error || 'Login failed')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <div className="w-full max-w-md">
        {/* Header */}
        <div className="text-center mb-8">
          <div className="text-6xl mb-4">📚</div>
          <h1 className="text-4xl font-bold uppercase tracking-wider mb-2">Amar Pathagar</h1>
          <p className="text-old-grey uppercase text-sm tracking-widest">Community Library</p>
        </div>

        {/* Login Form */}
        <div className="classic-card">
          <h2 className="classic-heading text-2xl">Login</h2>

          {error && (
            <div className="mb-4 p-3 border-2 border-red-600 bg-red-50 text-red-600 font-bold">
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block text-sm font-bold uppercase tracking-wider mb-2">
                Username
              </label>
              <input
                type="text"
                value={formData.username}
                onChange={(e) => setFormData({ ...formData, username: e.target.value })}
                className="classic-input"
                required
              />
            </div>

            <div>
              <label className="block text-sm font-bold uppercase tracking-wider mb-2">
                Password
              </label>
              <input
                type="password"
                value={formData.password}
                onChange={(e) => setFormData({ ...formData, password: e.target.value })}
                className="classic-input"
                required
              />
            </div>

            <button
              type="submit"
              disabled={loading}
              className="w-full classic-button disabled:opacity-50"
            >
              {loading ? 'Logging in...' : 'Login'}
            </button>
          </form>

          <div className="mt-6 text-center">
            <p className="text-old-grey text-sm">
              Don't have an account?{' '}
              <Link href="/register" className="text-old-ink font-bold underline">
                Register here
              </Link>
            </p>
          </div>
        </div>

        {/* Info Box */}
        <div className="mt-6 p-4 border-2 border-old-border bg-white">
          <p className="text-xs text-old-grey uppercase tracking-wider text-center">
            A Trust-Based Reading Network
          </p>
        </div>
      </div>
    </div>
  )
}
