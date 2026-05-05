<template>
  <div class="home-page">
    <!-- 主页头部 -->
    <div class="home-header">
      <h2>主页</h2>
    </div>

    <!-- 发布区域 -->
    <div class="compose-area">
      <img :src="userStore.user?.avatar || defaultAvatar" :alt="userStore.user?.username" />
      <div class="compose-input">
        <textarea
          placeholder="有什么新鲜事？"
          rows="3"
          @click="showComposeModal = true"
          readonly
        />
        <button class="compose-btn" @click="showComposeModal = true">发布</button>
      </div>
    </div>

    <!-- 帖子列表 -->
    <div class="posts-container">
      <div v-if="isLoading && postStore.posts.length === 0" class="loading">
        加载中...
      </div>

      <template v-else>
        <Post
          v-for="post in postStore.posts"
          :key="post.id"
          :post="post"
        />
      </template>

      <!-- 加载更多 -->
      <div v-if="postStore.hasMore && postStore.posts.length > 0" class="load-more">
        <button @click="loadMore" :disabled="isLoading">
          {{ isLoading ? '加载中...' : '加载更多' }}
        </button>
      </div>
    </div>

    <!-- 发布模态框 -->
    <ComposeModal v-if="showComposeModal" @close="closeComposeModal" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { usePostStore } from '@/stores/post'
import Post from '@/components/Post.vue'
import ComposeModal from '@/components/ComposeModal.vue'

const userStore = useUserStore()
const postStore = usePostStore()

const isLoading = ref(false)
const showComposeModal = ref(false)
const currentPage = ref(1)

const defaultAvatar =
  'https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png'

onMounted(async () => {
  try {
    isLoading.value = true
    await postStore.fetchHomeFeed(1, true)
  } catch (error) {
    console.error('Failed to fetch home feed:', error)
  } finally {
    isLoading.value = false
  }
})

const loadMore = async () => {
  try {
    isLoading.value = true
    currentPage.value++
    await postStore.fetchHomeFeed(currentPage.value, false)
  } catch (error) {
    console.error('Failed to load more posts:', error)
    currentPage.value--
  } finally {
    isLoading.value = false
  }
}

const closeComposeModal = async () => {
  showComposeModal.value = false
  // 刷新时间线
  currentPage.value = 1
  await postStore.fetchHomeFeed(1, true)
}
</script>

<style scoped>
.home-page {
  width: 100%;
  max-width: 600px;
  margin: 0 auto;
}

.home-header {
  position: sticky;
  top: 0;
  padding: 16px;
  border-bottom: 1px solid #eff3f4;
  background-color: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(12px);
  z-index: 100;
}

.home-header h2 {
  font-size: 20px;
  font-weight: 700;
  color: #0f1419;
  margin: 0;
}

.compose-area {
  display: flex;
  gap: 16px;
  padding: 16px;
  border-bottom: 1px solid #eff3f4;
}

.compose-area img {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  flex-shrink: 0;
}

.compose-input {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.compose-input textarea {
  width: 100%;
  padding: 12px 0;
  border: none;
  font-size: 20px;
  color: #0f1419;
  resize: none;
  outline: none;
  cursor: text;
}

.compose-input textarea::placeholder {
  color: #8a8d91;
}

.compose-btn {
  align-self: flex-end;
  padding: 12px 32px;
  background-color: #1d9bf0;
  color: #ffffff;
  font-weight: 700;
  border-radius: 24px;
  font-size: 15px;
  transition: all 0.2s;
}

.compose-btn:hover {
  background-color: #1a91da;
}

.posts-container {
  min-height: 400px;
}

.loading {
  padding: 32px;
  text-align: center;
  color: #536471;
}

.load-more {
  padding: 16px;
  text-align: center;
}

.load-more button {
  padding: 12px 32px;
  background-color: #1d9bf0;
  color: #ffffff;
  font-weight: 700;
  border-radius: 24px;
  font-size: 15px;
  transition: all 0.2s;
}

.load-more button:hover:not(:disabled) {
  background-color: #1a91da;
}

.load-more button:disabled {
  opacity: 0.6;
}
</style>
