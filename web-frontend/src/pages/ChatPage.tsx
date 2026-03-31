import { useState } from 'react'
import Sidebar from '@/components/Sidebar'
import ChatArea from '@/components/ChatArea'
import CreateKBDialog from '@/components/CreateKBDialog'
import UploadDialog from '@/components/UploadDialog'

export default function ChatPage() {
  const [showCreateKB, setShowCreateKB] = useState(false)
  const [showUpload, setShowUpload] = useState(false)

  return (
    <div className="h-screen flex overflow-hidden">
      <Sidebar
        onCreateKB={() => setShowCreateKB(true)}
        onUpload={() => setShowUpload(true)}
      />
      <ChatArea />
      <CreateKBDialog open={showCreateKB} onClose={() => setShowCreateKB(false)} />
      <UploadDialog open={showUpload} onClose={() => setShowUpload(false)} />
    </div>
  )
}
