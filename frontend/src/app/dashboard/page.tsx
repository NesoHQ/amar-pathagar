'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Layout from '@/components/Layout'
import { useAuthStore } from '@/store/authStore'
import { booksAPI } from '@/lib/api'

export default function DashboardPage() {
  const router = useRouter()
  const { user, isAuthenticated, _hasHydrated } = useAuthStore()
  const [stats, setStats] = useState({
    totalBooks: 0,
    availableBooks: 0,
    booksReading: 0,
  })

  useEffect(() => {
    if (_hasHydrated && !isAuthenticated) {
      router.push('/login')
    }
  }, [isAuthenticated, _hasHydrated, router])

  useEffect(() => {
    if (isAuthenticated) {
      loadStats()
    }
  }, [isAuthenticated])

  const loadStats = async () => {
    try {
      const response = await booksAPI.getAll()
      const books = response.data.data || []
      setStats({
        totalBooks: books.length,
        availableBooks: books.filter((b: any) => b.status === 'available').length,
        booksReading: books.filter((b: any) => b.status === 'reading').length,
      })
    } catch (error) {
      console.error('Failed to load stats:', error)
    }
  }

  if (!_hasHydrated || !isAuthenticated || !user) {
    return null
  }

  return (
    <Layout>
      <div className="space-y-8">
        {/* Welcome Section */}
        <div className="classic-card">
          <h1 className="text-4xl font-bold uppercase tracking-wider mb-4">
            Welcome, {user.full_name || user.username}
          </h1>
          <p className="text-old-grey text-lg">
            Your personal library dashboard
          </p>
        </div>

        {/* Success Score Card */}
        <div className="classic-card bg-old-ink text-old-paper">
          <div className="flex items-center justify-between">
            <div>
              <p className="text-sm uppercase tracking-widest opacity-75 mb-2">Success Score</p>
              <p className="text-6xl font-bold">{user.success_score}</p>
              <p className="text-sm uppercase tracking-wider mt-2 opacity-75">
                {user.success_score >= 100 ? 'Excellent Standing' : 
                 user.success_score >= 50 ? 'Good Standing' : 
                 user.success_score >= 20 ? 'Fair Standing' : 'Low Priority'}
              </p>
            </div>
            <div className="text-right">
              <div className="mb-4">
                <p className="text-3xl font-bold">{user.books_shared}</p>
                <p className="text-xs uppercase tracking-wider opacity-75">Books Shared</p>
              </div>
              <div>
                <p className="text-3xl font-bold">{user.books_received}</p>
                <p className="text-xs uppercase tracking-wider opacity-75">Books Read</p>
              </div>
            </div>
          </div>
        </div>

        {/* Stats Grid */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <StatCard
            title="Total Books"
            value={stats.totalBooks}
            subtitle="In Library"
          />
          <StatCard
            title="Available"
            value={stats.availableBooks}
            subtitle="Ready to Borrow"
          />
          <StatCard
            title="In Circulation"
            value={stats.booksReading}
            subtitle="Currently Reading"
          />
        </div>

        {/* Quick Actions */}
        <div className="classic-card">
          <h2 className="classic-heading text-2xl">Quick Actions</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <ActionButton
              title="Browse Books"
              description="Explore available books"
              href="/books"
              icon="📚"
            />
            <ActionButton
              title="My Library"
              description="View your books"
              href="/my-library"
              icon="📖"
            />
            <ActionButton
              title="Leaderboard"
              description="Top contributors"
              href="/leaderboard"
              icon="🏆"
            />
            <ActionButton
              title="Donate"
              description="Support the community"
              href="/donations"
              icon="🎁"
            />
          </div>
        </div>

        {/* Philosophy Box */}
        <div className="border-4 border-old-ink p-6 bg-white">
          <p className="text-center font-bold uppercase tracking-wider text-lg mb-4">
            Core Philosophy
          </p>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-center text-sm">
            <div>
              <p className="font-bold uppercase">Trust-Based</p>
              <p className="text-old-grey">Reading Network</p>
            </div>
            <div>
              <p className="font-bold uppercase">Knowledge &gt; Hoarding</p>
              <p className="text-old-grey">Share Wisdom</p>
            </div>
            <div>
              <p className="font-bold uppercase">Reputation</p>
              <p className="text-old-grey">Through Contribution</p>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  )
}

function StatCard({ title, value, subtitle }: { title: string; value: number; subtitle: string }) {
  return (
    <div className="classic-card text-center">
      <p className="text-sm uppercase tracking-widest text-old-grey mb-2">{title}</p>
      <p className="text-5xl font-bold mb-2">{value}</p>
      <p className="text-xs uppercase tracking-wider text-old-grey">{subtitle}</p>
    </div>
  )
}

function ActionButton({ title, description, href, icon }: { 
  title: string; 
  description: string; 
  href: string; 
  icon: string 
}) {
  return (
    <a
      href={href}
      className="block p-4 border-2 border-old-border hover:border-old-ink hover:shadow-[4px_4px_0px_0px_rgba(0,0,0,0.2)] transition-all"
    >
      <div className="flex items-center space-x-3">
        <span className="text-3xl">{icon}</span>
        <div>
          <p className="font-bold uppercase tracking-wider">{title}</p>
          <p className="text-sm text-old-grey">{description}</p>
        </div>
      </div>
    </a>
  )
}
