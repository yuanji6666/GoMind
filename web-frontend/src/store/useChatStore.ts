import { create } from 'zustand'
import type { KnowledgeBaseInfo, SessionInfo, ChatMessage } from '@/api/types'

interface ChatState {
  knowledgeBases: KnowledgeBaseInfo[]
  sessions: SessionInfo[]
  activeKB: KnowledgeBaseInfo | null
  activeSession: SessionInfo | null
  messages: Record<string, ChatMessage[]>
  loading: boolean
  /** 该会话是否还有更早消息可拉（游标分页） */
  historyHasMore: Record<string, boolean>
  historyLoadingMore: Record<string, boolean>

  setKnowledgeBases: (kbs: KnowledgeBaseInfo[]) => void
  setSessions: (sessions: SessionInfo[]) => void
  setActiveKB: (kb: KnowledgeBaseInfo | null) => void
  setActiveSession: (session: SessionInfo | null) => void
  addSession: (session: SessionInfo) => void
  addKB: (kb: KnowledgeBaseInfo) => void
  pushMessage: (sessionId: string, msg: ChatMessage) => void
  /** 覆盖某会话消息列表（如首次从服务端拉历史） */
  setSessionMessages: (sessionId: string, msgs: ChatMessage[]) => void
  /** 在列表头部插入更早的消息 */
  prependMessages: (sessionId: string, msgs: ChatMessage[]) => void
  setHistoryHasMore: (sessionId: string, hasMore: boolean) => void
  setHistoryLoadingMore: (sessionId: string, loading: boolean) => void
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
  historyHasMore: {},
  historyLoadingMore: {},

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

  setSessionMessages: (sessionId, msgs) =>
    set((s) => ({
      messages: { ...s.messages, [sessionId]: msgs },
    })),

  prependMessages: (sessionId, older) =>
    set((s) => ({
      messages: {
        ...s.messages,
        [sessionId]: [...older, ...(s.messages[sessionId] || [])],
      },
    })),

  setHistoryHasMore: (sessionId, hasMore) =>
    set((s) => ({
      historyHasMore: { ...s.historyHasMore, [sessionId]: hasMore },
    })),

  setHistoryLoadingMore: (sessionId, loading) =>
    set((s) => ({
      historyLoadingMore: { ...s.historyLoadingMore, [sessionId]: loading },
    })),

  setLoading: (v) => set({ loading: v }),

  clearMessages: () =>
    set({ messages: {}, historyHasMore: {}, historyLoadingMore: {} }),
}))
