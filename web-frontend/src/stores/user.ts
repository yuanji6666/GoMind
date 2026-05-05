import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User } from '@/types'
import { authAPI, userAPI } from '@/api'

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(localStorage.getItem('token'))
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value && !!user.value)

  const extractToken = (response: any) => response?.token || response?.data?.token || null

  const login = async (username: string, password: string) => {
    try {
      isLoading.value = true
      error.value = null
      const response = await authAPI.login(username, password)

      const nextToken = extractToken(response)
      if (!nextToken) {
        throw new Error('Login response missing token')
      }

      token.value = nextToken
      localStorage.setItem('token', nextToken)

      await fetchCurrentUser()
      return user.value
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Login failed'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const register = async (
    email: string,
    username: string,
    password: string,
    captcha: string
  ) => {
    try {
      isLoading.value = true
      error.value = null
      const response = await authAPI.register(email, username, password, captcha)

      const nextToken = extractToken(response)
      if (nextToken) {
        token.value = nextToken
        localStorage.setItem('token', nextToken)
      }

      return response
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Registration failed'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const logout = () => {
    user.value = null
    token.value = null
    localStorage.removeItem('token')
  }

  const fetchCurrentUser = async () => {
    try {
      isLoading.value = true
      const response = await userAPI.getCurrentUser()
      user.value = response.data
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to fetch user'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const updateUser = async (data: Partial<User>) => {
    if (!user.value) return
    try {
      isLoading.value = true
      const response = await userAPI.updateUserInfo(user.value.id, data)
      user.value = response.data
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to update user'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    user,
    token,
    isLoading,
    error,
    isAuthenticated,
    login,
    register,
    logout,
    fetchCurrentUser,
    updateUser,
  }
})
