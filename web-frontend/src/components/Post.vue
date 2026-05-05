<template>
  <div class="post" @click="goToPost">
    <!-- 头部：用户信息 -->
    <div class="post-header">
      <RouterLink :to="`/profile/${post.userId}`" @click.stop>
        <img :src="post.author.avatar || defaultAvatar" :alt="post.author.username" />
      </RouterLink>

      <div class="post-header-content">
        <RouterLink :to="`/profile/${post.userId}`" @click.stop class="author-name">
          {{ post.author.username }}
        </RouterLink>
        <span class="author-handle">@{{ post.author.username }}</span>
        <span class="dot">·</span>
        <span class="post-time">{{ formatTime(post.createdAt) }}</span>
      </div>

      <button class="post-menu-btn" @click.stop>
        <MoreIcon />
      </button>
    </div>

    <!-- 内容 -->
    <div class="post-body">
      <p class="post-content">{{ post.content }}</p>

      <!-- 图片 -->
      <div v-if="post.images && post.images.length > 0" class="post-images">
        <div
          v-for="(img, idx) in post.images"
          :key="idx"
          class="post-image"
          :style="{ backgroundImage: `url(${img})` }"
        />
      </div>
    </div>

    <!-- 交互按钮 -->
    <div class="post-actions">
      <div class="action-group">
        <button class="action-btn" @click.stop>
          <CommentIcon />
          <span>{{ post.commentsCount }}</span>
        </button>
      </div>

      <div class="action-group">
        <button class="action-btn" @click.stop>
          <RetweetIcon />
          <span>{{ post.retweetsCount }}</span>
        </button>
      </div>

      <div class="action-group">
        <button
          class="action-btn"
          :class="{ liked: post.isLiked }"
          @click.stop="toggleLike"
        >
          <HeartIcon :filled="post.isLiked" />
          <span>{{ post.likesCount }}</span>
        </button>
      </div>

      <div class="action-group">
        <button class="action-btn" @click.stop>
          <ShareIcon />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import type { Post } from '@/types'
import { likeAPI } from '@/api'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
import relativeTime from 'dayjs/plugin/relativeTime'
import CommentIcon from './icons/CommentIcon.vue'
import RetweetIcon from './icons/RetweetIcon.vue'
import HeartIcon from './icons/HeartIcon.vue'
import ShareIcon from './icons/ShareIcon.vue'
import MoreIcon from './icons/MoreIcon.vue'

dayjs.extend(relativeTime)
dayjs.locale('zh-cn')

const props = defineProps<{
  post: Post
}>()

const emit = defineEmits(['like', 'unlike'])

const router = useRouter()
const isLiking = ref(false)

const defaultAvatar =
  'https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png'

const formatTime = (date: string) => {
  return dayjs(date).fromNow()
}

const goToPost = () => {
  router.push(`/post/${props.post.id}`)
}

const toggleLike = async () => {
  if (isLiking.value) return

  try {
    isLiking.value = true
    if (props.post.isLiked) {
      await likeAPI.unlikePost(props.post.id)
      props.post.isLiked = false
      props.post.likesCount--
      emit('unlike')
    } else {
      await likeAPI.likePost(props.post.id)
      props.post.isLiked = true
      props.post.likesCount++
      emit('like')
    }
  } catch (error) {
    console.error('Failed to toggle like:', error)
  } finally {
    isLiking.value = false
  }
}
</script>

<style scoped>
.post {
  padding: 16px;
  border-bottom: 1px solid #eff3f4;
  cursor: pointer;
  transition: all 0.2s;
}

.post:hover {
  background-color: #f7f9f9;
}

.post-header {
  display: flex;
  gap: 12px;
  margin-bottom: 12px;
}

.post-header img {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  flex-shrink: 0;
}

.post-header-content {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 15px;
}

.author-name {
  font-weight: 700;
  color: #0f1419;
}

.author-handle,
.dot,
.post-time {
  color: #536471;
}

.post-menu-btn {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: transparent;
  color: #536471;
  transition: all 0.2s;
  flex-shrink: 0;
}

.post-menu-btn:hover {
  background-color: rgba(29, 155, 240, 0.1);
  color: #1d9bf0;
}

.post-body {
  margin-left: 60px;
  margin-bottom: 12px;
}

.post-content {
  font-size: 15px;
  line-height: 1.5;
  color: #0f1419;
  word-wrap: break-word;
  white-space: pre-wrap;
  margin-bottom: 12px;
}

.post-images {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 8px;
  border-radius: 12px;
  overflow: hidden;
}

.post-image {
  width: 100%;
  padding-bottom: 100%;
  background-size: cover;
  background-position: center;
  border-radius: 8px;
  position: relative;
}

.post-actions {
  display: flex;
  justify-content: space-between;
  margin-left: 60px;
  margin-top: 12px;
  color: #536471;
  font-size: 13px;
}

.action-group {
  display: flex;
  align-items: center;
  flex: 1;
  max-width: 80px;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: transparent;
  color: #536471;
  font-size: 13px;
  transition: all 0.2s;
  padding: 8px 4px;
}

.action-btn:hover {
  color: #1d9bf0;
}

.action-btn.liked {
  color: #f91880;
}

.action-btn svg {
  width: 16px;
  height: 16px;
}
</style>
