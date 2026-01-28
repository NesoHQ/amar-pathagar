'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Layout from '@/components/Layout'
import { useAuthStore } from '@/store/authStore'
import { donationsAPI } from '@/lib/api'

export default function DonationsPage() {
  const router = useRouter()
  const { isAuthenticated, _hasHydrated } = useAuthStore()
  const [donations, setDonations] = useState<any[]>([])
  const [showForm, setShowForm] = useState(false)
  const [formData, setFormData] = useState({
    donation_type: 'money',
    amount: '',
    currency: 'USD',
    message: '',
    is_public: true,
  })

  useEffect(() => {
    if (_hasHydrated && !isAuthenticated) {
      router.push('/login')
    } else if (isAuthenticated) {
      loadDonations()
    }
  }, [isAuthenticated, _hasHydrated, router])

  const loadDonations = async () => {
    try {
      const response = await donationsAPI.getAll()
      const donationsData = response.data.data || response.data || []
      setDonations(Array.isArray(donationsData) ? donationsData : [])
    } catch (error) {
      console.error('Failed to load donations:', error)
      setDonations([])
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await donationsAPI.create({
        ...formData,
        amount: formData.amount ? parseFloat(formData.amount) : undefined,
      })
      setShowForm(false)
      setFormData({
        donation_type: 'money',
        amount: '',
        currency: 'USD',
        message: '',
        is_public: true,
      })
      loadDonations()
      alert('Thank you for your donation! +10 points')
    } catch (error) {
      alert('Failed to process donation')
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
          <div className="flex justify-between items-center">
            <div>
              <h1 className="text-4xl font-bold uppercase tracking-wider mb-2">🎁 Donations</h1>
              <p className="text-old-grey">Support our community library</p>
            </div>
            <button
              onClick={() => setShowForm(!showForm)}
              className="classic-button"
            >
              {showForm ? 'Cancel' : 'Make a Donation'}
            </button>
          </div>
        </div>

        {/* Donation Form */}
        {showForm && (
          <div className="classic-card">
            <h2 className="text-2xl font-bold uppercase tracking-wider mb-4">Make a Donation</h2>
            <form onSubmit={handleSubmit} className="space-y-4">
              <div>
                <label className="block text-sm font-bold uppercase mb-2">Type</label>
                <select
                  value={formData.donation_type}
                  onChange={(e) => setFormData({ ...formData, donation_type: e.target.value })}
                  className="classic-input"
                >
                  <option value="money">Money</option>
                  <option value="book">Book</option>
                </select>
              </div>

              {formData.donation_type === 'money' && (
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Amount</label>
                  <input
                    type="number"
                    value={formData.amount}
                    onChange={(e) => setFormData({ ...formData, amount: e.target.value })}
                    className="classic-input"
                    placeholder="0.00"
                    step="0.01"
                  />
                </div>
              )}

              <div>
                <label className="block text-sm font-bold uppercase mb-2">Message (Optional)</label>
                <textarea
                  value={formData.message}
                  onChange={(e) => setFormData({ ...formData, message: e.target.value })}
                  className="classic-input"
                  rows={3}
                  placeholder="Share why you're donating..."
                />
              </div>

              <div className="flex items-center gap-2">
                <input
                  type="checkbox"
                  checked={formData.is_public}
                  onChange={(e) => setFormData({ ...formData, is_public: e.target.checked })}
                  className="w-4 h-4"
                />
                <label className="text-sm">Make this donation public</label>
              </div>

              <button type="submit" className="w-full classic-button">
                Submit Donation
              </button>
            </form>
          </div>
        )}

        {/* Donations List */}
        <div className="classic-card">
          <h2 className="text-2xl font-bold uppercase tracking-wider mb-4">Recent Donations</h2>
          {donations.length === 0 ? (
            <p className="text-center text-old-grey py-8">No donations yet</p>
          ) : (
            <div className="space-y-3">
              {donations.map((donation: any) => (
                <div key={donation.id} className="p-4 border-2 border-old-border">
                  <div className="flex justify-between items-start mb-2">
                    <div>
                      <h3 className="font-bold uppercase">{donation.donor_name}</h3>
                      <p className="text-sm text-old-grey">
                        {donation.donation_type === 'money'
                          ? `Donated ${donation.currency} ${donation.amount}`
                          : `Donated a book: ${donation.book_title}`}
                      </p>
                    </div>
                    <span className="text-sm text-old-grey">
                      {new Date(donation.created_at).toLocaleDateString()}
                    </span>
                  </div>
                  {donation.message && (
                    <p className="text-old-grey italic mt-2">"{donation.message}"</p>
                  )}
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
    </Layout>
  )
}
