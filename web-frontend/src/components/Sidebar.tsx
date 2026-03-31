import { useEffect, useState } from 'react'
import {
  Database,
  MessageSquarePlus,
  ChevronRight,
  LogOut,
  Plus,
  Upload,
  BookOpen,
  MessagesSquare,
} from 'lucide-react'
import toast from 'react-hot-toast'
import { getKBList, getSessionList } from '@/api/gin'
import { CODE_SUCCESS } from '@/api/types'
import type { KnowledgeBaseInfo, SessionInfo } from '@/api/types'
import { useAuthStore } from '@/store/useAuthStore'
import { useChatStore } from '@/store/useChatStore'

interface SidebarProps {
  onCreateKB: () => void
  onUpload: () => void
}

export default function Sidebar({ onCreateKB, onUpload }: SidebarProps) {
  const { username, logout } = useAuthStore()
  const {
    knowledgeBases,
    sessions,
    activeKB,
    activeSession,
    setKnowledgeBases,
    setSessions,
    setActiveKB,
    setActiveSession,
  } = useChatStore()
  const [loadingKB, setLoadingKB] = useState(false)

  useEffect(() => {
    fetchKBList()
  }, [username])

  useEffect(() => {
    if (activeKB) fetchSessions()
  }, [activeKB?.user_kb_id])

  async function fetchKBList() {
    if (!username) return
    setLoadingKB(true)
    try {
      const res = await getKBList()
      if (res.status_code === CODE_SUCCESS) {
        setKnowledgeBases(res.knowledge_base_list || [])
      }
    } catch {
      toast.error('获取知识库列表失败')
    } finally {
      setLoadingKB(false)
    }
  }

  async function fetchSessions() {
    if (!username) return
    try {
      const res = await getSessionList()
      if (res.status_code === CODE_SUCCESS) {
        const all = res.sessions || []
        setSessions(activeKB ? all.filter((s) => s.user_kb_id === activeKB.user_kb_id) : all)
      }
    } catch {
      toast.error('获取会话列表失败')
    }
  }

  function handleSelectKB(kb: KnowledgeBaseInfo) {
    setActiveKB(kb)
  }

  function handleSelectSession(session: SessionInfo) {
    setActiveSession(session)
  }

  function handleNewChat() {
    setActiveSession(null)
  }

  return (
    <aside className="w-72 h-full flex flex-col bg-slate-900 text-white border-r border-slate-700/50">
      {/* Header */}
      <div className="p-4 border-b border-slate-700/50">
        <div className="flex items-center gap-2.5">
          <div className="w-9 h-9 rounded-xl bg-indigo-600 flex items-center justify-center">
            <Database className="w-5 h-5" />
          </div>
          <div className="flex-1 min-w-0">
            <h1 className="text-base font-semibold truncate">GopherAI</h1>
            <p className="text-xs text-slate-400 truncate">{username}</p>
          </div>
        </div>
      </div>

      {/* Knowledge Bases */}
      <div className="flex-1 overflow-y-auto">
        <div className="p-3">
          <div className="flex items-center justify-between mb-2 px-1">
            <span className="text-xs font-semibold text-slate-400 uppercase tracking-wider">知识库</span>
            <button
              onClick={onCreateKB}
              className="p-1 rounded-md hover:bg-slate-700/60 text-slate-400 hover:text-white transition"
              title="创建知识库"
            >
              <Plus className="w-4 h-4" />
            </button>
          </div>

          {loadingKB ? (
            <div className="space-y-2 px-1">
              {[1, 2, 3].map((i) => (
                <div key={i} className="h-9 bg-slate-800 rounded-lg animate-pulse" />
              ))}
            </div>
          ) : knowledgeBases.length === 0 ? (
            <p className="text-sm text-slate-500 px-1 py-4 text-center">暂无知识库</p>
          ) : (
            <div className="space-y-0.5">
              {knowledgeBases.map((kb) => (
                <button
                  key={kb.user_kb_id}
                  onClick={() => handleSelectKB(kb)}
                  className={`w-full flex items-center gap-2.5 px-3 py-2 rounded-lg text-sm transition group ${
                    activeKB?.user_kb_id === kb.user_kb_id
                      ? 'bg-indigo-600/20 text-indigo-300'
                      : 'hover:bg-slate-800 text-slate-300'
                  }`}
                >
                  <BookOpen className="w-4 h-4 flex-shrink-0" />
                  <span className="truncate flex-1 text-left">{kb.name}</span>
                  <ChevronRight
                    className={`w-3.5 h-3.5 flex-shrink-0 transition ${
                      activeKB?.user_kb_id === kb.user_kb_id ? 'opacity-100' : 'opacity-0 group-hover:opacity-50'
                    }`}
                  />
                </button>
              ))}
            </div>
          )}
        </div>

        {/* Sessions for active KB */}
        {activeKB && (
          <div className="p-3 border-t border-slate-700/50">
            <div className="flex items-center justify-between mb-2 px-1">
              <span className="text-xs font-semibold text-slate-400 uppercase tracking-wider">会话</span>
              <div className="flex gap-1">
                <button
                  onClick={onUpload}
                  className="p-1 rounded-md hover:bg-slate-700/60 text-slate-400 hover:text-white transition"
                  title="上传文档"
                >
                  <Upload className="w-4 h-4" />
                </button>
                <button
                  onClick={handleNewChat}
                  className="p-1 rounded-md hover:bg-slate-700/60 text-slate-400 hover:text-white transition"
                  title="新建会话"
                >
                  <MessageSquarePlus className="w-4 h-4" />
                </button>
              </div>
            </div>

            {sessions.length === 0 ? (
              <p className="text-sm text-slate-500 px-1 py-4 text-center">发送第一条消息开始对话</p>
            ) : (
              <div className="space-y-0.5">
                {sessions.map((s) => (
                  <button
                    key={s.session_id}
                    onClick={() => handleSelectSession(s)}
                    className={`w-full flex items-center gap-2.5 px-3 py-2 rounded-lg text-sm transition ${
                      activeSession?.session_id === s.session_id
                        ? 'bg-indigo-600/20 text-indigo-300'
                        : 'hover:bg-slate-800 text-slate-300'
                    }`}
                  >
                    <MessagesSquare className="w-4 h-4 flex-shrink-0" />
                    <span className="truncate flex-1 text-left">{s.title || '新会话'}</span>
                  </button>
                ))}
              </div>
            )}
          </div>
        )}
      </div>

      {/* Logout */}
      <div className="p-3 border-t border-slate-700/50">
        <button
          onClick={logout}
          className="w-full flex items-center gap-2.5 px-3 py-2 rounded-lg text-sm text-slate-400 hover:bg-slate-800 hover:text-white transition"
        >
          <LogOut className="w-4 h-4" />
          退出登录
        </button>
      </div>
    </aside>
  )
}
