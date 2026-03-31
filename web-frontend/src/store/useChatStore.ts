import { create } from 'zustand'
import type { KnowledgeBaseInfo, SessionInfo, ChatMessage } from '@/api/types'

interface ChatState {
  knowledgeBases: KnowledgeBaseInfo[]
  sessions: SessionInfo[]
  activeKB: KnowledgeBaseInfo | null
  activeSession: SessionInfo | null
  messages: Record<string, ChatMessage[]>
  loading: boolean

  setKnowledgeBases: (kbs: KnowledgeBaseInfo[]) => void
  setSessions: (sessions: SessionInfo[]) => void
  setActiveKB: (kb: KnowledgeBaseInfo | null) => void
  setActiveSession: (session: SessionInfo | null) => void
  addSession: (session: SessionInfo) => void
  addKB: (kb: KnowledgeBaseInfo) => void
  pushMessage: (sessionId: string, msg: ChatMessage) => void
  setLoading: (v: boolean) => void
  clearMessages: () => void
}

export const useChatStore = create<ChatState>((set) => ({
  knowledgeBases: [],
  sessions: [],
  activeKB: null,
  activeSession: null,
  messages: {},
  loading: false,

  setKnowledgeBases: (kbs) => set({ knowledgeBases: kbs }),

  setSessions: (sessions) => set({ sessions }),

  setActiveKB: (kb) => set({ activeKB: kb, activeSession: null }),

  setActiveSession: (session) => set({ activeSession: session }),

  addSession: (session) =>
    set((s) => ({ sessions: [session, ...s.sessions] })),

  addKB: (kb) =>
    set((s) => ({ knowledgeBases: [...s.knowledgeBases, kb] })),

  pushMessage: (sessionId, msg) =>
    set((s) => ({
      messages: {
        ...s.messages,
        [sessionId]: [...(s.messages[sessionId] || []), msg],
      },
    })),

  setLoading: (v) => set({ loading: v }),

  clearMessages: () => set({ messages: {} }),
}))
