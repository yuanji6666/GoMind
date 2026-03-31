import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import { User, Bot } from 'lucide-react'
import type { ChatMessage } from '@/api/types'

interface MessageBubbleProps {
  message: ChatMessage
}

export default function MessageBubble({ message }: MessageBubbleProps) {
  if (message.isUser) {
    return (
      <div className="flex justify-end gap-3 px-4">
        <div className="max-w-[70%]">
          <div className="bg-indigo-600 text-white px-4 py-3 rounded-2xl rounded-br-md text-sm leading-relaxed">
            {message.content}
          </div>
        </div>
        <div className="w-8 h-8 rounded-full bg-indigo-600 flex items-center justify-center flex-shrink-0">
          <User className="w-4 h-4 text-white" />
        </div>
      </div>
    )
  }

  return (
    <div className="flex gap-3 px-4">
      <div className="w-8 h-8 rounded-full bg-emerald-600 flex items-center justify-center flex-shrink-0">
        <Bot className="w-4 h-4 text-white" />
      </div>
      <div className="max-w-[70%]">
        <div className="bg-white border border-gray-200 px-4 py-3 rounded-2xl rounded-bl-md text-sm shadow-sm">
          <div className="markdown-body">
            <ReactMarkdown remarkPlugins={[remarkGfm]}>
              {message.content}
            </ReactMarkdown>
          </div>
        </div>
      </div>
    </div>
  )
}
