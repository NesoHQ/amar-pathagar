'use client'

import { useEffect } from 'react'
import Link from 'next/link'
import { useRouter, usePathname } from 'next/navigation'
import { useAuthStore } from '@/store/authStore'

export default function Layout({ children }: { children: React.ReactNode }) {
  const { user, isAuthenticated, logout } = useAuthStore()
  const router = useRouter()
  const pathname = usePathname()

  const handleLogout = () => {
    logout()
    router.push('/login')
  }

  const isActive = (path: string) => pathname === path

  return (
    <div className="min-h-screen old-paper-texture">
      {/* Classic Header */}
      <header className="bg-white border-b-4 border-old-ink shadow-md">
        <div className="max-w-7xl mx-auto px-4 py-4">
          <div className="flex justify-between items-center">
            {/* Logo */}
            <Link href="/dashboard" className="flex items-center space-x-3">
              <div className="text-4xl">📚</div>
              <div>
                <h1 className="text-2xl font-bold uppercase tracking-wider text-old-ink">
                  Amar Pathagar
                </h1>
                <p className="text-xs text-old-grey uppercase tracking-widest">Community Library</p>
              </div>
            </Link>

            {/* Navigation */}
            {isAuthenticated && (
              <nav className="hidden md:flex space-x-1">
                <NavLink href="/dashboard" active={isActive('/dashboard')}>Dashboard</NavLink>
                <NavLink href="/books" active={isActive('/books')}>Books</NavLink>
                <NavLink href="/my-library" active={isActive('/my-library')}>My Library</NavLink>
                <NavLink href="/leaderboard" active={isActive('/leaderboard')}>Leaderboard</NavLink>
                <NavLink href="/donations" active={isActive('/donations')}>Donations</NavLink>
                {user?.role === 'admin' && (
                  <NavLink href="/admin" active={isActive('/admin')}>Admin</NavLink>
                )}
              </nav>
            )}

            {/* User Info */}
            {isAuthenticated && user && (
              <div className="flex items-center space-x-4">
                <div className="text-right hidden md:block">
                  <div className="font-bold text-old-ink">{user.full_name || user.username}</div>
                  <div className="text-xs text-old-grey uppercase">
                    Score: {user.success_score} • {user.role}
                  </div>
                </div>
                <button
                  onClick={handleLogout}
                  className="px-4 py-2 border-2 border-old-ink text-old-ink font-bold uppercase text-sm
                           hover:bg-old-ink hover:text-old-paper transition-colors"
                >
                  Logout
                </button>
              </div>
            )}
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 py-8">
        {children}
      </main>

      {/* Classic Footer */}
      <footer className="bg-white border-t-4 border-old-ink mt-16">
        <div className="max-w-7xl mx-auto px-4 py-8">
          <div className="text-center">
            <p className="text-old-grey text-sm uppercase tracking-wider">
              Est. 2026 • A Trust-Based Reading Network
            </p>
            <p className="text-old-grey text-xs mt-2">
              Knowledge &gt; Hoarding • Reputation Through Contribution
            </p>
          </div>
        </div>
      </footer>
    </div>
  )
}

function NavLink({ href, active, children }: { href: string; active: boolean; children: React.ReactNode }) {
  return (
    <Link
      href={href}
      className={`px-4 py-2 font-bold uppercase text-sm tracking-wider transition-colors
        ${active 
          ? 'bg-old-ink text-old-paper' 
          : 'text-old-ink hover:bg-old-grey hover:text-old-paper'
        }`}
    >
      {children}
    </Link>
  )
}
