import { ginClient } from './client'
import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  RegisterResponse,
  CaptchaRequest,
  GinResponse,
  KBListResponse,
  CreateKBRequest,
  CreateKBResponse,
  SessionListResponse,
  CreateSessionRequest,
  CreateSessionResponse,
  SendMessageRequest,
  SendMessageResponse,
} from './types'

export async function login(data: LoginRequest) {
  const res = await ginClient.post<LoginResponse>('/api/v1/user/login', data)
  return res.data
}

export async function register(data: RegisterRequest) {
  const res = await ginClient.post<RegisterResponse>('/api/v1/user/register', data)
  return res.data
}

export async function sendCaptcha(data: CaptchaRequest) {
  const res = await ginClient.post<GinResponse>('/api/v1/user/captcha', data)
  return res.data
}

export async function getKBList() {
  const res = await ginClient.get<KBListResponse>('/api/v1/AI/kb/list')
  return res.data
}

export async function createKB(data: CreateKBRequest) {
  const res = await ginClient.post<CreateKBResponse>('/api/v1/AI/kb/create', data)
  return res.data
}

export async function getSessionList() {
  const res = await ginClient.get<SessionListResponse>('/api/v1/AI/session/list')
  return res.data
}

export async function createSession(data: CreateSessionRequest) {
  const res = await ginClient.post<CreateSessionResponse>('/api/v1/AI/session/create', data)
  return res.data
}

export async function sendMessage(data: SendMessageRequest) {
  const res = await ginClient.post<SendMessageResponse>('/api/v1/AI/session/send', data)
  return res.data
}
