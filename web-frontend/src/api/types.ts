// Go backend response envelope
export interface GinResponse {
  status_code: number
  status_msg?: string
}

export const CODE_SUCCESS = 1000
export const CODE_INVALID_TOKEN = 2006
export const CODE_NOT_LOGIN = 2007

// Auth
export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse extends GinResponse {
  token?: string
}

export interface RegisterRequest {
  email: string
  captcha: string
  password: string
}

export interface RegisterResponse extends GinResponse {
  token?: string
}

export interface CaptchaRequest {
  email: string
}

// Knowledge Base
export interface KnowledgeBaseInfo {
  user_kb_id: string
  name: string
}

export interface KBListResponse extends GinResponse {
  knowledge_base_list?: KnowledgeBaseInfo[]
}

export interface CreateKBRequest {
  username: string
  kb_name: string
}

export interface CreateKBResponse extends GinResponse {
  knowledge_base_info?: KnowledgeBaseInfo
}

// Session
export interface SessionInfo {
  title: string
  session_id: string
  user_kb_id: string
}

export interface SessionListResponse extends GinResponse {
  sessions?: SessionInfo[]
}

export interface CreateSessionRequest {
  username: string
  user_question: string
  user_kb_id: string
}

export interface CreateSessionResponse extends GinResponse {
  answer?: string
  session_id?: string
}

export interface SendMessageRequest {
  session_id: string
  user_question: string
}

export interface SendMessageResponse extends GinResponse {
  answer?: string
}

// FastAPI
export interface UploadResponse {
  status: string
  files_ingested: number
}

// Client-side message model
export interface ChatMessage {
  id: string
  content: string
  isUser: boolean
  timestamp: Date
}
