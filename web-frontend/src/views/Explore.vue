<template>
  <div class="explore-page">
    <div class="explore-header">
      <h2>探索</h2>
    </div>

    <div class="explore-content">
      <div v-if="isLoading" class="loading">
        加载中...
      </div>

      <template v-else>
        <Post
          v-for="post in postStore.posts"
          :key="post.id"
          :post="post"
        />
      </template>

      <div v-if="postStore.hasMore && postStore.posts.length > 0" class="load-more">
        <button @click="loadMore" :disabled="isLoading">
          {{ isLoading ? '加载中...' : '加载更多' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { usePostStore } from '@/stores/post'
import Post from '@/components/Post.vue'

const postStore = usePostStore()
const isLoading = ref(false)
const currentPage = ref(1)

onMounted(async () => {
  try {
    isLoading.value = true
    await postStore.fetchExploreFeed(1, true)
  } catch (error) {
    console.error('Failed to fetch explore feed:', error)
  } finally {
    isLoading.value = false
  }
})

const loadMore = async () => {
  try {
    isLoading.value = true
    currentPage.value++
    await postStore.fetchExploreFeed(currentPage.value, false)
  } catch (error) {
    console.error('Failed to load more posts:', error)
    currentPage.value--
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.explore-page {
  width: 100%;
}

.explore-header {
  position: sticky;
  top: 0;
  padding: 16px;
  border-bottom: 1px solid #eff3f4;
  background-color: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(12px);
  z-index: 100;
}

.explore-header h2 {
  font-size: 20px;
  font-weight: 700;
  color: #0f1419;
  margin: 0;
}

.explore-content {
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
