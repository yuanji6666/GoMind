import { useState } from 'react'
import { X, FolderPlus } from 'lucide-react'
import toast from 'react-hot-toast'
import { createKB } from '@/api/gin'
import { CODE_SUCCESS } from '@/api/types'
import { useAuthStore } from '@/store/useAuthStore'
import { useChatStore } from '@/store/useChatStore'

interface Props {
  open: boolean
  onClose: () => void
}

export default function CreateKBDialog({ open, onClose }: Props) {
  const { username } = useAuthStore()
  const { addKB } = useChatStore()
  const [name, setName] = useState('')
  const [submitting, setSubmitting] = useState(false)

  if (!open) return null

  async function handleCreate(e: React.FormEvent) {
    e.preventDefault()
    if (!name.trim()) {
      toast.error('请输入知识库名称')
      return
    }
    setSubmitting(true)
    try {
      const res = await createKB({ username, kb_name: name.trim() })
      if (res.status_code === CODE_SUCCESS && res.knowledge_base_info) {
        addKB(res.knowledge_base_info)
        toast.success('知识库创建成功')
        setName('')
        onClose()
      } else {
        toast.error(res.status_msg || '创建失败')
      }
    } catch {
      toast.error('网络错误')
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      <div className="absolute inset-0 bg-black/50 backdrop-blur-sm" onClick={onClose} />
      <div className="relative bg-white rounded-2xl shadow-2xl w-full max-w-md mx-4 p-6">
        <button
          onClick={onClose}
          className="absolute top-4 right-4 p-1 rounded-lg hover:bg-gray-100 text-gray-400 hover:text-gray-600 transition"
        >
          <X className="w-5 h-5" />
        </button>

        <div className="flex items-center gap-3 mb-6">
          <div className="w-10 h-10 rounded-xl bg-indigo-100 flex items-center justify-center">
            <FolderPlus className="w-5 h-5 text-indigo-600" />
          </div>
          <h3 className="text-lg font-semibold text-gray-800">创建知识库</h3>
        </div>

        <form onSubmit={handleCreate} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1.5">知识库名称</label>
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="例如：产品文档、技术手册"
              className="w-full px-4 py-2.5 border border-gray-200 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition"
              autoFocus
            />
          </div>

          <div className="flex gap-3 pt-2">
            <button
              type="button"
              onClick={onClose}
              className="flex-1 py-2.5 border border-gray-200 text-gray-600 rounded-xl text-sm font-medium hover:bg-gray-50 transition"
            >
              取消
            </button>
            <button
              type="submit"
              disabled={submitting}
              className="flex-1 py-2.5 bg-indigo-600 hover:bg-indigo-500 disabled:opacity-50 text-white rounded-xl text-sm font-medium transition"
            >
              {submitting ? '创建中...' : '创建'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}
