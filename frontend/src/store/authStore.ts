import { create } from 'zustand'

interface User {
  id: string
  username: string
  email: string
  full_name: string
  role: string
  avatar_url?: string
  success_score: number
  books_shared: number
  books_received: number
}

interface AuthState {
  user: User | null
  accessToken: string | null
  isAuthenticated: boolean
  setAuth: (user: User, accessToken: string) => void
  logout: () => void
  loadFromStorage: () => void
}

export const useAuthStore = create<AuthState>((set) => ({
  user: null,
  accessToken: null,
  isAuthenticated: false,
  
  setAuth: (user, accessToken) => {
    if (typeof window !== 'undefined') {
      localStorage.setItem('user', JSON.stringify(user))
      localStorage.setItem('access_token', accessToken)
    }
    set({ user, accessToken, isAuthenticated: true })
  },
  
  logout: () => {
    if (typeof window !== 'undefined') {
      localStorage.removeItem('user')
      localStorage.removeItem('access_token')
    }
    set({ user: null, accessToken: null, isAuthenticated: false })
  },
  
  loadFromStorage: () => {
    if (typeof window !== 'undefined') {
      const userStr = localStorage.getItem('user')
      const token = localStorage.getItem('access_token')
      if (userStr && token) {
        try {
          const user = JSON.parse(userStr)
          set({ user, accessToken: token, isAuthenticated: true })
        } catch (e) {
          console.error('Failed to parse user from storage')
        }
      }
    }
  },
}))
