import { fastapiClient } from './client'
import type { UploadResponse } from './types'

export async function uploadDocuments(userKbId: string, files: File[]) {
  const formData = new FormData()
  files.forEach((file) => formData.append('files', file))

  const res = await fastapiClient.post<UploadResponse>(
    `/knowledge-bases/${userKbId}/documents`,
    formData,
    { headers: { 'Content-Type': 'multipart/form-data' } },
  )
  return res.data
}
