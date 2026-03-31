import { useState, useRef, useEffect } from 'react'
import { Send, Sparkles, MessageSquarePlus } from 'lucide-react'
import toast from 'react-hot-toast'
import { createSession, sendMessage } from '@/api/gin'
import { CODE_SUCCESS } from '@/api/types'
import type { ChatMessage } from '@/api/types'
import { useAuthStore } from '@/store/useAuthStore'
import { useChatStore } from '@/store/useChatStore'
import MessageBubble from './MessageBubble'

function ThinkingIndicator() {
  return (
    <div className="flex gap-3 px-4">
      <div className="w-8 h-8 rounded-full bg-emerald-600 flex items-center justify-center flex-shrink-0">
        <Sparkles className="w-4 h-4 text-white" />
      </div>
      <div className="bg-white border border-gray-200 px-4 py-3 rounded-2xl rounded-bl-md shadow-sm">
        <div className="flex items-center gap-1.5">
          <div className="w-2 h-2 bg-gray-400 rounded-full thinking-dot" />
          <div className="w-2 h-2 bg-gray-400 rounded-full thinking-dot" />
          <div className="w-2 h-2 bg-gray-400 rounded-full thinking-dot" />
        </div>
      </div>
    </div>
  )
}

export default function ChatArea() {
  const { username } = useAuthStore()
  const {
    activeKB,
    activeSession,
    messages,
    loading,
    setActiveSession,
    addSession,
    pushMessage,
    setLoading,
  } = useChatStore()

  const [input, setInput] = useState('')
  const bottomRef = useRef<HTMLDivElement>(null)
  const inputRef = useRef<HTMLTextAreaElement>(null)

  const currentMessages = activeSession ? messages[activeSession.session_id] || [] : []

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: 'smooth' })
  }, [currentMessages.length, loading])

  useEffect(() => {
    inputRef.current?.focus()
  }, [activeSession?.session_id])

  async function handleSend() {
    const question = input.trim()
    if (!question || loading) return
    if (!activeKB) {
      toast.error('请先选择一个知识库')
      return
    }

    setInput('')
    setLoading(true)

    const userMsg: ChatMessage = {
      id: `user-${Date.now()}`,
      content: question,
      isUser: true,
      timestamp: new Date(),
    }

    try {
      if (!activeSession) {
        const tempSessionId = `temp-${Date.now()}`
        pushMessage(tempSessionId, userMsg)
        setActiveSession({ session_id: tempSessionId, title: question.slice(0, 30), user_kb_id: activeKB.user_kb_id })

        const res = await createSession({
          username,
          user_question: question,
          user_kb_id: activeKB.user_kb_id,
        })

        if (res.status_code === CODE_SUCCESS && res.session_id) {
          const newSession = {
            session_id: res.session_id,
            title: question.slice(0, 30),
            user_kb_id: activeKB.user_kb_id,
          }
          // Migrate messages from temp to real session ID
          const tempMsgs = useChatStore.getState().messages[tempSessionId] || []
          const aiMsg: ChatMessage = {
            id: `ai-${Date.now()}`,
            content: res.answer || '暂无回答',
            isUser: false,
            timestamp: new Date(),
          }
          tempMsgs.forEach((m) => pushMessage(res.session_id!, m))
          pushMessage(res.session_id!, aiMsg)
          addSession(newSession)
          setActiveSession(newSession)
        } else {
          toast.error(res.status_msg || '创建会话失败')
        }
      } else {
        pushMessage(activeSession.session_id, userMsg)

        const res = await sendMessage({
          session_id: activeSession.session_id,
          user_question: question,
        })

        if (res.status_code === CODE_SUCCESS) {
          const aiMsg: ChatMessage = {
            id: `ai-${Date.now()}`,
            content: res.answer || '暂无回答',
            isUser: false,
            timestamp: new Date(),
          }
          pushMessage(activeSession.session_id, aiMsg)
        } else {
          toast.error(res.status_msg || '发送失败')
        }
      }
    } catch {
      toast.error('请求失败，请检查网络')
    } finally {
      setLoading(false)
    }
  }

  function handleKeyDown(e: React.KeyboardEvent) {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      handleSend()
    }
  }

  if (!activeKB) {
    return (
      <div className="flex-1 flex items-center justify-center bg-gray-50">
        <div className="text-center max-w-sm">
          <div className="w-20 h-20 mx-auto mb-6 rounded-2xl bg-indigo-50 flex items-center justify-center">
            <Sparkles className="w-10 h-10 text-indigo-400" />
          </div>
          <h2 className="text-xl font-semibold text-gray-800 mb-2">选择一个知识库开始</h2>
          <p className="text-gray-500 text-sm leading-relaxed">
            从左侧选择或创建一个知识库，上传文档后即可开始智能问答
          </p>
        </div>
      </div>
    )
  }

  return (
    <div className="flex-1 flex flex-col bg-gray-50 h-full">
      {/* Header */}
      <header className="h-14 flex items-center justify-between px-6 bg-white border-b border-gray-200 flex-shrink-0">
        <div className="flex items-center gap-2">
          <h2 className="font-medium text-gray-800">
            {activeKB.name}
          </h2>
          {activeSession && (
            <span className="text-gray-400 text-sm">/ {activeSession.title || '新会话'}</span>
          )}
        </div>
        <button
          onClick={() => setActiveSession(null)}
          className="flex items-center gap-1.5 px-3 py-1.5 text-sm text-indigo-600 hover:bg-indigo-50 rounded-lg transition"
        >
          <MessageSquarePlus className="w-4 h-4" />
          新对话
        </button>
      </header>

      {/* Messages */}
      <div className="flex-1 overflow-y-auto py-6 space-y-5">
        {currentMessages.length === 0 && !loading && (
          <div className="flex items-center justify-center h-full">
            <div className="text-center">
              <Sparkles className="w-12 h-12 text-gray-300 mx-auto mb-3" />
              <p className="text-gray-400 text-sm">输入您的问题，开始与知识库对话</p>
            </div>
          </div>
        )}
        {currentMessages.map((msg) => (
          <MessageBubble key={msg.id} message={msg} />
        ))}
        {loading && <ThinkingIndicator />}
        <div ref={bottomRef} />
      </div>

      {/* Input */}
      <div className="p-4 bg-white border-t border-gray-200 flex-shrink-0">
        <div className="max-w-3xl mx-auto flex items-end gap-3">
          <div className="flex-1 relative">
            <textarea
              ref={inputRef}
              value={input}
              onChange={(e) => setInput(e.target.value)}
              onKeyDown={handleKeyDown}
              rows={1}
              placeholder="输入您的问题... (Enter 发送, Shift+Enter 换行)"
              className="w-full resize-none px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition max-h-32"
              style={{ minHeight: '44px' }}
              onInput={(e) => {
                const target = e.target as HTMLTextAreaElement
                target.style.height = 'auto'
                target.style.height = Math.min(target.scrollHeight, 128) + 'px'
              }}
            />
          </div>
          <button
            onClick={handleSend}
            disabled={!input.trim() || loading}
            className="p-3 bg-indigo-600 hover:bg-indigo-500 disabled:opacity-40 disabled:cursor-not-allowed text-white rounded-xl transition flex-shrink-0"
          >
            <Send className="w-4 h-4" />
          </button>
        </div>
      </div>
    </div>
  )
}
