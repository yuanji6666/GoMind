import axios from 'axios'
import { CODE_INVALID_TOKEN, CODE_NOT_LOGIN } from './types'

// 开发：直连本机 8000；生产构建未设置 VITE_FASTAPI_URL 时走同源（由 Nginx 反代到 service-ai）
const viteFastapi = import.meta.env.VITE_FASTAPI_URL
const FASTAPI_URL =
  viteFastapi !== undefined && viteFastapi !== ''
    ? viteFastapi
    : import.meta.env.DEV
      ? 'http://localhost:8000'
      : ''

export const ginClient = axios.create({
  baseURL: '',
  headers: { 'Content-Type': 'application/json' },
  timeout: 30000,
})

export const fastapiClient = axios.create({
  baseURL: FASTAPI_URL,
  timeout: 120000,
})

ginClient.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

ginClient.interceptors.response.use(
  (response) => {
    const code = response.data?.status_code
    if (code === CODE_INVALID_TOKEN || code === CODE_NOT_LOGIN) {
      localStorage.removeItem('token')
      localStorage.removeItem('username')
      window.location.href = '/login'
    }
    return response
  },
  (error) => Promise.reject(error),
)
