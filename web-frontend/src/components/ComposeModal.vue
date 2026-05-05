<template>
  <div class="modal-overlay" @click="close">
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <button class="close-btn" @click="close">✕</button>
      </div>

      <div class="modal-body">
        <div class="compose-header">
          <img :src="userStore.user?.avatar || defaultAvatar" :alt="userStore.user?.username" />
          <div class="compose-right">
            <div class="text-gray">发布到</div>
            <div class="audience-selector">
              公开 <ChevronDownIcon />
            </div>
          </div>
        </div>

        <textarea
          v-model="content"
          placeholder="有什么新鲜事？"
          class="compose-textarea"
          @input="autoResize"
        />

        <!-- 图片预览 -->
        <div v-if="images.length > 0" class="images-preview">
          <div
            v-for="(img, idx) in images"
            :key="idx"
            class="image-preview-item"
            :style="{ backgroundImage: `url(${img.preview})` }"
          >
            <button class="remove-image-btn" @click="removeImage(idx)">✕</button>
          </div>
        </div>

        <!-- 工具栏 -->
        <div class="compose-toolbar">
          <button class="toolbar-btn" @click="triggerImageUpload" title="添加图片">
            <ImageIcon />
          </button>
          <input
            ref="fileInput"
            type="file"
            multiple
            accept="image/*"
            style="display: none"
            @change="handleImageSelect"
          />
          <button class="toolbar-btn" title="表情">
            <EmojiIcon />
          </button>
          <button class="toolbar-btn" title="日期">
            <CalendarIcon />
          </button>
          <button class="toolbar-btn" title="位置">
            <LocationIcon />
          </button>
        </div>

        <div class="modal-footer">
          <button class="post-btn" @click="handlePost" :disabled="!content.trim() || isLoading">
            {{ isLoading ? '发布中...' : '发布' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'
import { usePostStore } from '@/stores/post'
import ChevronDownIcon from './icons/ChevronDownIcon.vue'
import ImageIcon from './icons/ImageIcon.vue'
import EmojiIcon from './icons/EmojiIcon.vue'
import CalendarIcon from './icons/CalendarIcon.vue'
import LocationIcon from './icons/LocationIcon.vue'

const emit = defineEmits(['close'])

const userStore = useUserStore()
const postStore = usePostStore()

const content = ref('')
const images = ref<any[]>([])
const isLoading = ref(false)
const fileInput = ref<HTMLInputElement>()
const textareaRef = ref<HTMLTextAreaElement>()

const defaultAvatar =
  'https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png'

const close = () => {
  content.value = ''
  images.value = []
  emit('close')
}

const autoResize = () => {
  if (textareaRef.value) {
    textareaRef.value.style.height = 'auto'
    textareaRef.value.style.height = textareaRef.value.scrollHeight + 'px'
  }
}

const triggerImageUpload = () => {
  fileInput.value?.click()
}

const handleImageSelect = (e: Event) => {
  const files = (e.target as HTMLInputElement).files
  if (files) {
    for (let file of files) {
      const reader = new FileReader()
      reader.onload = (event) => {
        images.value.push({
          file,
          preview: event.target?.result,
        })
      }
      reader.readAsDataURL(file)
    }
  }
}

const removeImage = (idx: number) => {
  images.value.splice(idx, 1)
}

const handlePost = async () => {
  if (!content.value.trim()) return

  try {
    isLoading.value = true
    const imageFiles = images.value.map((img) => img.file)
    await postStore.createPost(content.value, imageFiles)
    close()
  } catch (error) {
    console.error('Failed to create post:', error)
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background-color: #ffffff;
  border-radius: 16px;
  width: 90%;
  max-width: 600px;
  max-height: 90vh;
  overflow-y: auto;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
}

.modal-header {
  display: flex;
  justify-content: flex-end;
  padding: 16px;
  border-bottom: 1px solid #eff3f4;
}

.close-btn {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: transparent;
  color: #0f1419;
  font-size: 18px;
  transition: all 0.2s;
}

.close-btn:hover {
  background-color: #f7f9f9;
}

.modal-body {
  padding: 16px;
}

.compose-header {
  display: flex;
  gap: 16px;
  margin-bottom: 16px;
}

.compose-header img {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
}

.compose-right {
  flex: 1;
}

.text-gray {
  font-size: 13px;
  color: #536471;
  margin-bottom: 4px;
}

.audience-selector {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px 12px;
  border: 1px solid #cfd9de;
  border-radius: 20px;
  font-size: 15px;
  font-weight: 500;
  color: #1d9bf0;
  cursor: pointer;
  transition: all 0.2s;
}

.audience-selector:hover {
  background-color: #f7f9f9;
}

.compose-textarea {
  width: 100%;
  min-height: 120px;
  padding: 0;
  border: none;
  font-size: 20px;
  font-family: inherit;
  color: #0f1419;
  resize: none;
  outline: none;
  margin-bottom: 16px;
}

.compose-textarea::placeholder {
  color: #8a8d91;
}

.images-preview {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 8px;
  margin-bottom: 16px;
}

.image-preview-item {
  position: relative;
  width: 100%;
  padding-bottom: 100%;
  background-size: cover;
  background-position: center;
  border-radius: 12px;
  overflow: hidden;
}

.remove-image-btn {
  position: absolute;
  top: 8px;
  right: 8px;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: rgba(0, 0, 0, 0.7);
  color: #ffffff;
  font-size: 18px;
  transition: all 0.2s;
}

.remove-image-btn:hover {
  background-color: rgba(0, 0, 0, 0.9);
}

.compose-toolbar {
  display: flex;
  gap: 8px;
  margin-bottom: 16px;
  padding-bottom: 16px;
  border-bottom: 1px solid #eff3f4;
}

.toolbar-btn {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background-color: transparent;
  color: #1d9bf0;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.toolbar-btn:hover {
  background-color: rgba(29, 155, 240, 0.1);
}

.toolbar-btn svg {
  width: 20px;
  height: 20px;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
}

.post-btn {
  padding: 12px 32px;
  border-radius: 24px;
  background-color: #1d9bf0;
  color: #ffffff;
  font-weight: 700;
  font-size: 15px;
  transition: all 0.2s;
}

.post-btn:hover:not(:disabled) {
  background-color: #1a91da;
}

.post-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>
