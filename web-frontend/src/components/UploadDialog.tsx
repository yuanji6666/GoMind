import { useState, useRef } from 'react'
import { X, Upload, FileText, Trash2, CheckCircle2, Loader2 } from 'lucide-react'
import toast from 'react-hot-toast'
import { uploadDocuments } from '@/api/fastapi'
import { useChatStore } from '@/store/useChatStore'

interface Props {
  open: boolean
  onClose: () => void
}

export default function UploadDialog({ open, onClose }: Props) {
  const { activeKB } = useChatStore()
  const [files, setFiles] = useState<File[]>([])
  const [uploading, setUploading] = useState(false)
  const [done, setDone] = useState(false)
  const inputRef = useRef<HTMLInputElement>(null)

  if (!open) return null

  function handleFilesSelected(e: React.ChangeEvent<HTMLInputElement>) {
    const selected = e.target.files
    if (!selected) return
    setFiles((prev) => [...prev, ...Array.from(selected)])
    setDone(false)
    if (inputRef.current) inputRef.current.value = ''
  }

  function removeFile(index: number) {
    setFiles((prev) => prev.filter((_, i) => i !== index))
  }

  async function handleUpload() {
    if (!activeKB) {
      toast.error('请先选择知识库')
      return
    }
    if (files.length === 0) {
      toast.error('请选择至少一个文件')
      return
    }
    setUploading(true)
    try {
      const res = await uploadDocuments(activeKB.user_kb_id, files)
      if (res.status === 'ok') {
        toast.success(`成功上传 ${res.files_ingested} 个文件`)
        setDone(true)
        setFiles([])
      } else {
        toast.error('上传失败')
      }
    } catch (err: unknown) {
      const msg =
        err instanceof Error ? err.message : '上传失败，请检查 FastAPI 服务是否运行'
      toast.error(msg)
    } finally {
      setUploading(false)
    }
  }

  function handleClose() {
    setFiles([])
    setDone(false)
    onClose()
  }

  function formatSize(bytes: number) {
    if (bytes < 1024) return bytes + ' B'
    if (bytes < 1048576) return (bytes / 1024).toFixed(1) + ' KB'
    return (bytes / 1048576).toFixed(1) + ' MB'
  }

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      <div className="absolute inset-0 bg-black/50 backdrop-blur-sm" onClick={handleClose} />
      <div className="relative bg-white rounded-2xl shadow-2xl w-full max-w-lg mx-4 p-6">
        <button
          onClick={handleClose}
          className="absolute top-4 right-4 p-1 rounded-lg hover:bg-gray-100 text-gray-400 hover:text-gray-600 transition"
        >
          <X className="w-5 h-5" />
        </button>

        <div className="flex items-center gap-3 mb-5">
          <div className="w-10 h-10 rounded-xl bg-emerald-100 flex items-center justify-center">
            <Upload className="w-5 h-5 text-emerald-600" />
          </div>
          <div>
            <h3 className="text-lg font-semibold text-gray-800">上传文档</h3>
            <p className="text-xs text-gray-500">
              上传至知识库：{activeKB?.name || '—'}
            </p>
          </div>
        </div>

        {done ? (
          <div className="text-center py-8">
            <CheckCircle2 className="w-14 h-14 text-emerald-500 mx-auto mb-3" />
            <p className="text-gray-700 font-medium">上传完成！</p>
            <p className="text-sm text-gray-500 mt-1">文档正在被处理，稍后即可开始问答</p>
            <button
              onClick={handleClose}
              className="mt-5 px-6 py-2 bg-indigo-600 text-white rounded-xl text-sm font-medium hover:bg-indigo-500 transition"
            >
              完成
            </button>
          </div>
        ) : (
          <>
            {/* Drop zone */}
            <label className="block border-2 border-dashed border-gray-200 hover:border-indigo-400 rounded-xl p-8 text-center cursor-pointer transition mb-4 group">
              <input
                ref={inputRef}
                type="file"
                multiple
                accept=".pdf,.txt,.md,.docx,.doc,.html,.htm,.csv,.json"
                onChange={handleFilesSelected}
                className="hidden"
              />
              <Upload className="w-8 h-8 text-gray-300 group-hover:text-indigo-400 mx-auto mb-2 transition" />
              <p className="text-sm text-gray-500">
                点击选择文件或拖拽到此处
              </p>
              <p className="text-xs text-gray-400 mt-1">
                支持 PDF、TXT、MD、DOCX 等格式
              </p>
            </label>

            {/* File list */}
            {files.length > 0 && (
              <div className="max-h-48 overflow-y-auto space-y-2 mb-4">
                {files.map((file, i) => (
                  <div
                    key={`${file.name}-${i}`}
                    className="flex items-center gap-3 px-3 py-2 bg-gray-50 rounded-lg"
                  >
                    <FileText className="w-4 h-4 text-gray-400 flex-shrink-0" />
                    <div className="flex-1 min-w-0">
                      <p className="text-sm text-gray-700 truncate">{file.name}</p>
                      <p className="text-xs text-gray-400">{formatSize(file.size)}</p>
                    </div>
                    <button
                      onClick={() => removeFile(i)}
                      className="p-1 rounded hover:bg-gray-200 text-gray-400 hover:text-red-500 transition"
                    >
                      <Trash2 className="w-3.5 h-3.5" />
                    </button>
                  </div>
                ))}
              </div>
            )}

            <div className="flex gap-3">
              <button
                type="button"
                onClick={handleClose}
                className="flex-1 py-2.5 border border-gray-200 text-gray-600 rounded-xl text-sm font-medium hover:bg-gray-50 transition"
              >
                取消
              </button>
              <button
                onClick={handleUpload}
                disabled={uploading || files.length === 0}
                className="flex-1 flex items-center justify-center gap-2 py-2.5 bg-indigo-600 hover:bg-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed text-white rounded-xl text-sm font-medium transition"
              >
                {uploading ? (
                  <>
                    <Loader2 className="w-4 h-4 animate-spin" />
                    上传中...
                  </>
                ) : (
                  <>
                    <Upload className="w-4 h-4" />
                    上传 ({files.length})
                  </>
                )}
              </button>
            </div>
          </>
        )}
      </div>
    </div>
  )
}
