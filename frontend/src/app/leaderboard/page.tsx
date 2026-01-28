'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Layout from '@/components/Layout'
import { useAuthStore } from '@/store/authStore'
import { userAPI } from '@/lib/api'

export default function LeaderboardPage() {
  const router = useRouter()
  const { isAuthenticated, _hasHydrated } = useAuthStore()
  const [data, setData] = useState<any>(null)
  const [activeTab, setActiveTab] = useState('readers')

  useEffect(() => {
    if (_hasHydrated && !isAuthenticated) {
      router.push('/login')
    } else if (isAuthenticated) {
      loadLeaderboard()
    }
  }, [isAuthenticated, _hasHydrated, router])

  const loadLeaderboard = async () => {
    try {
      const response = await userAPI.getLeaderboard()
      setData(response.data.data || response.data)
    } catch (error) {
      console.error('Failed to load leaderboard:', error)
    }
  }

  if (!_hasHydrated || !isAuthenticated || !data) {
    return null
  }

  const tabs = [
    { key: 'readers', label: 'Top Readers', data: data.top_readers, metric: 'books_received' },
    { key: 'sharers', label: 'Top Sharers', data: data.top_sharers, metric: 'books_shared' },
    { key: 'donors', label: 'Top Donors', data: data.top_donors, metric: 'books_shared' },
    { key: 'scores', label: 'Highest Scores', data: data.highest_scores, metric: 'success_score' },
    { key: 'ideas', label: 'Top Idea Writers', data: data.top_idea_writers, metric: 'ideas_posted' },
  ]

  const activeData = tabs.find(t => t.key === activeTab)

  return (
    <Layout>
      <div className="space-y-6">
        {/* Header */}
        <div className="classic-card">
          <h1 className="text-4xl font-bold uppercase tracking-wider mb-2">🏆 Leaderboard</h1>
          <p className="text-old-grey">Celebrating our top contributors</p>
        </div>

        {/* Tabs */}
        <div className="flex flex-wrap gap-2">
          {tabs.map((tab) => (
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

        {/* Leaderboard */}
        <div className="classic-card">
          <div className="space-y-3">
            {activeData?.data?.map((user: any, index: number) => (
              <div
                key={user.id}
                className="flex items-center gap-4 p-4 border-2 border-old-border hover:border-old-ink transition-colors"
              >
                {/* Rank */}
                <div className="text-center w-16">
                  <div className={`text-3xl font-bold ${
                    index === 0 ? 'text-yellow-600' :
                    index === 1 ? 'text-gray-400' :
                    index === 2 ? 'text-orange-600' :
                    'text-old-grey'
                  }`}>
                    #{index + 1}
                  </div>
                </div>

                {/* User Info */}
                <div className="flex-1">
                  <h3 className="font-bold uppercase tracking-wider">
                    {user.full_name || user.username}
                  </h3>
                  <p className="text-sm text-old-grey">@{user.username}</p>
                  {user.is_donor && (
                    <span className="inline-block mt-1 stamp border-yellow-600 text-yellow-600">
                      Donor
                    </span>
                  )}
                </div>

                {/* Score */}
                <div className="text-right">
                  <div className="text-3xl font-bold">
                    {user[activeData.metric as keyof typeof user]}
                  </div>
                  <p className="text-xs uppercase text-old-grey">
                    {activeTab === 'readers' && 'Books Read'}
                    {activeTab === 'sharers' && 'Books Shared'}
                    {activeTab === 'donors' && 'Donations'}
                    {activeTab === 'scores' && 'Success Score'}
                    {activeTab === 'ideas' && 'Ideas Posted'}
                  </p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </Layout>
  )
}
