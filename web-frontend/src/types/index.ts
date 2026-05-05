// 用户类型
export interface User {
  id: number
  username: string
  email: string
  avatar?: string
  cover?: string
  bio?: string
  location?: string
  website?: string
  createdAt: string
  followersCount: number
  followingCount: number
  postsCount: number
}

// 帖子类型
export interface Post {
  id: string
  userId: number
  author: User
  content: string
  images?: string[]
  createdAt: string
  likesCount: number
  commentsCount: number
  retweetsCount: number
  isLiked?: boolean
  isRetweeted?: boolean
}

// 评论类型
export interface Comment {
  id: string
  postId: string
  userId: number
  author: User
  content: string
  createdAt: string
  likesCount: number
  repliesCount: number
  isLiked?: boolean
}

// 关注关系
export interface Follow {
  followerId: number
  followingId: number
  createdAt: string
}

// 点赞记录
export interface Like {
  userId: number
  postId: string
  createdAt: string
}

// 登录请求
export interface LoginRequest {
  username: string
  password: string
}

// 注册请求
export interface RegisterRequest {
  email: string
  username: string
  password: string
  captcha: string
}

// 创建帖子请求
export interface CreatePostRequest {
  content: string
  images?: File[]
  replyToId?: string
}

// API 响应包装
export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

// 分页响应
export interface PaginatedResponse<T> {
  items: T[]
  total: number
  page: number
  pageSize: number
  hasMore: boolean
}
