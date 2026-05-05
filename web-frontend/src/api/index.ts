import axios, { AxiosInstance } from 'axios'
import type { ApiResponse, PaginatedResponse, Post, Comment, User } from '@/types'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || '/api/v1'

const api: AxiosInstance = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
})

// 添加请求拦截器
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// 添加响应拦截器
api.interceptors.response.use(
  (response) => response.data,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

// ==================== 认证接口 ====================
export const authAPI = {
  // 获取验证码
  getCaptcha: (email: string) =>
    api.post<any, ApiResponse>('/user/captcha', { email }),

  // 注册
  register: (email: string, username: string, password: string, captcha: string) =>
    api.post<any, ApiResponse<{ id: number; token: string }>>('/user/register', {
      email,
      username,
      password,
      captcha,
    }),

  // 登录
  login: (username: string, password: string) =>
    api.post<any, ApiResponse<{ id: number; token: string; user: User }>>('/user/login', {
      username,
      password,
    }),
}

// ==================== 用户接口 ====================
export const userAPI = {
  // 获取用户信息
  getUserInfo: (userId: number) =>
    api.get<any, ApiResponse<User>>(`/user/${userId}`),

  // 获取当前用户信息
  getCurrentUser: () =>
    api.get<any, ApiResponse<User>>('/user/me'),

  // 更新用户信息
  updateUserInfo: (userId: number, data: Partial<User>) =>
    api.put<any, ApiResponse<User>>(`/user/${userId}`, data),

  // 获取用户的帖子列表
  getUserPosts: (userId: number, page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<Post>>>(`/user/${userId}/posts`, {
      params: { page, pageSize },
    }),

  // 获取用户的赞过的帖子
  getUserLikes: (userId: number, page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<Post>>>(`/user/${userId}/likes`, {
      params: { page, pageSize },
    }),

  // 上传头像
  uploadAvatar: (userId: number, file: File) => {
    const formData = new FormData()
    formData.append('avatar', file)
    return api.post<any, ApiResponse<{ avatar: string }>>(
      `/user/${userId}/avatar`,
      formData,
      {
        headers: { 'Content-Type': 'multipart/form-data' },
      }
    )
  },

  // 上传封面
  uploadCover: (userId: number, file: File) => {
    const formData = new FormData()
    formData.append('cover', file)
    return api.post<any, ApiResponse<{ cover: string }>>(
      `/user/${userId}/cover`,
      formData,
      {
        headers: { 'Content-Type': 'multipart/form-data' },
      }
    )
  },
}

// ==================== 帖子接口 ====================
export const postAPI = {
  // 获取主页时间线（推荐帖子）
  getHomeFeed: (page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<Post>>>('/posts/feed', {
      params: { page, pageSize },
    }),

  // 获取探索页面时间线
  getExploreFeed: (page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<Post>>>('/posts/explore', {
      params: { page, pageSize },
    }),

  // 获取单个帖子
  getPost: (postId: string) =>
    api.get<any, ApiResponse<Post>>(`/posts/${postId}`),

  // 创建帖子
  createPost: (content: string, images?: File[], replyToId?: string) => {
    const formData = new FormData()
    formData.append('content', content)
    if (images && images.length > 0) {
      images.forEach((img, idx) => {
        formData.append(`images[${idx}]`, img)
      })
    }
    if (replyToId) {
      formData.append('replyToId', replyToId)
    }
    return api.post<any, ApiResponse<Post>>('/posts', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
  },

  // 删除帖子
  deletePost: (postId: string) =>
    api.delete<any, ApiResponse>(`/posts/${postId}`),

  // 搜索帖子
  searchPosts: (query: string, page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<Post>>>('/posts/search', {
      params: { query, page, pageSize },
    }),
}

// ==================== 点赞接口 ====================
export const likeAPI = {
  // 点赞帖子
  likePost: (postId: string) =>
    api.post<any, ApiResponse>(`/posts/${postId}/like`, {}),

  // 取消点赞帖子
  unlikePost: (postId: string) =>
    api.delete<any, ApiResponse>(`/posts/${postId}/like`),

  // 点赞评论
  likeComment: (commentId: string) =>
    api.post<any, ApiResponse>(`/comments/${commentId}/like`, {}),

  // 取消点赞评论
  unlikeComment: (commentId: string) =>
    api.delete<any, ApiResponse>(`/comments/${commentId}/like`),

  // 获取赞过某个帖子的用户列表
  getPostLikes: (postId: string, page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<User>>>(`/posts/${postId}/likes`, {
      params: { page, pageSize },
    }),
}

// ==================== 评论接口 ====================
export const commentAPI = {
  // 获取帖子的评论列表
  getPostComments: (postId: string, page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<Comment>>>(`/posts/${postId}/comments`, {
      params: { page, pageSize },
    }),

  // 创建评论
  createComment: (postId: string, content: string, images?: File[]) => {
    const formData = new FormData()
    formData.append('content', content)
    if (images && images.length > 0) {
      images.forEach((img, idx) => {
        formData.append(`images[${idx}]`, img)
      })
    }
    return api.post<any, ApiResponse<Comment>>(`/posts/${postId}/comments`, formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
  },

  // 删除评论
  deleteComment: (commentId: string) =>
    api.delete<any, ApiResponse>(`/comments/${commentId}`),

  // 获取评论的回复
  getCommentReplies: (commentId: string, page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<Comment>>>(`/comments/${commentId}/replies`, {
      params: { page, pageSize },
    }),

  // 回复评论
  replyToComment: (commentId: string, content: string, images?: File[]) => {
    const formData = new FormData()
    formData.append('content', content)
    if (images && images.length > 0) {
      images.forEach((img, idx) => {
        formData.append(`images[${idx}]`, img)
      })
    }
    return api.post<any, ApiResponse<Comment>>(`/comments/${commentId}/replies`, formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
  },
}

// ==================== 关注接口 ====================
export const followAPI = {
  // 关注用户
  followUser: (userId: number) =>
    api.post<any, ApiResponse>(`/user/${userId}/follow`, {}),

  // 取消关注用户
  unfollowUser: (userId: number) =>
    api.delete<any, ApiResponse>(`/user/${userId}/follow`),

  // 获取用户的粉丝列表
  getFollowers: (userId: number, page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<User>>>(`/user/${userId}/followers`, {
      params: { page, pageSize },
    }),

  // 获取用户的关注列表
  getFollowing: (userId: number, page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<User>>>(`/user/${userId}/following`, {
      params: { page, pageSize },
    }),

  // 检查是否关注某个用户
  checkFollowing: (userId: number) =>
    api.get<any, ApiResponse<boolean>>(`/user/${userId}/following-status`),
}

// ==================== 通知接口 ====================
export const notificationAPI = {
  // 获取通知列表
  getNotifications: (page = 1, pageSize = 20) =>
    api.get<any, ApiResponse<PaginatedResponse<any>>>('/notifications', {
      params: { page, pageSize },
    }),

  // 标记通知为已读
  markNotificationAsRead: (notificationId: string) =>
    api.put<any, ApiResponse>(`/notifications/${notificationId}/read`, {}),

  // 标记所有通知为已读
  markAllNotificationsAsRead: () =>
    api.put<any, ApiResponse>('/notifications/read-all', {}),
}

export default api
