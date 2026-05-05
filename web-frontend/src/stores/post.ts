import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Post } from '@/types'
import { postAPI } from '@/api'

export const usePostStore = defineStore('post', () => {
  const posts = ref<Post[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)
  const currentPage = ref(1)
  const hasMore = ref(true)

  const postsCount = computed(() => posts.value.length)

  const fetchHomeFeed = async (page = 1, reset = false) => {
    try {
      isLoading.value = true
      error.value = null
      const response = await postAPI.getHomeFeed(page, 20)
      if (reset) {
        posts.value = response.data.items
      } else {
        posts.value.push(...response.data.items)
      }
      currentPage.value = page
      hasMore.value = response.data.hasMore
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to fetch feed'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const fetchExploreFeed = async (page = 1, reset = false) => {
    try {
      isLoading.value = true
      error.value = null
      const response = await postAPI.getExploreFeed(page, 20)
      if (reset) {
        posts.value = response.data.items
      } else {
        posts.value.push(...response.data.items)
      }
      currentPage.value = page
      hasMore.value = response.data.hasMore
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to fetch explore'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const createPost = async (content: string, images?: File[]) => {
    try {
      isLoading.value = true
      error.value = null
      const response = await postAPI.createPost(content, images)
      posts.value.unshift(response.data)
      return response.data
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to create post'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const deletePost = async (postId: string) => {
    try {
      isLoading.value = true
      posts.value = posts.value.filter((p) => p.id !== postId)
      await postAPI.deletePost(postId)
    } catch (err: any) {
      error.value = err.response?.data?.message || 'Failed to delete post'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const clearPosts = () => {
    posts.value = []
    currentPage.value = 1
    hasMore.value = true
  }

  return {
    posts,
    isLoading,
    error,
    currentPage,
    hasMore,
    postsCount,
    fetchHomeFeed,
    fetchExploreFeed,
    createPost,
    deletePost,
    clearPosts,
  }
})
