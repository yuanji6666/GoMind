import { create } from 'zustand'

function decodeUsername(token: string): string {
  try {
    const payload = JSON.parse(atob(token.split('.')[1]))
    return payload.username || ''
  } catch {
    return ''
  }
}

interface AuthState {
  token: string | null
  username: string
  setAuth: (token: string) => void
  logout: () => void
  isLoggedIn: () => boolean
}

export const useAuthStore = create<AuthState>((set, get) => ({
  token: localStorage.getItem('token'),
  username: localStorage.getItem('username') || '',

  setAuth: (token: string) => {
    const username = decodeUsername(token)
    localStorage.setItem('token', token)
    localStorage.setItem('username', username)
    set({ token, username })
  },

  logout: () => {
    localStorage.removeItem('token')
    localStorage.removeItem('username')
    set({ token: null, username: '' })
  },

  isLoggedIn: () => !!get().token,
}))
