<template>
  <div class="post-detail-page">
    <!-- 返回按钮 -->
    <div class="detail-header">
      <button class="back-btn" @click="$router.back()">
        <BackIcon />
      </button>
      <span>帖子</span>
    </div>

    <!-- 帖子详情 -->
    <div v-if="post" class="post-detail">
      <!-- 用户信息 -->
      <div class="post-user">
        <RouterLink :to="`/profile/${post.userId}`">
          <img :src="post.author.avatar || defaultAvatar" :alt="post.author.username" />
        </RouterLink>
        <div class="user-info">
          <RouterLink :to="`/profile/${post.userId}`" class="username">
            {{ post.author.username }}
          </RouterLink>
          <span class="handle">@{{ post.author.username }}</span>
        </div>
      </div>

      <!-- 帖子内容 -->
      <div class="post-content-detail">
        <p class="content">{{ post.content }}</p>
        <div v-if="post.images && post.images.length > 0" class="images">
          <img v-for="(img, idx) in post.images" :key="idx" :src="img" />
        </div>
      </div>

      <!-- 时间和统计 -->
      <div class="post-stats">
        <span class="time">{{ formatTime(post.createdAt) }}</span>
        <div class="divider" />
        <span><strong>{{ post.likesCount }}</strong> 赞</span>
        <span><strong>{{ post.commentsCount }}</strong> 条回复</span>
      </div>

      <!-- 交互按钮 -->
      <div class="post-actions">
        <button class="action-btn" @click="scrollToComments">
          <CommentIcon />
        </button>
        <button class="action-btn">
          <RetweetIcon />
        </button>
        <button
          class="action-btn"
          :class="{ liked: post.isLiked }"
          @click="toggleLike"
        >
          <HeartIcon :filled="post.isLiked" />
        </button>
        <button class="action-btn">
          <ShareIcon />
        </button>
      </div>
    </div>

    <!-- 评论区 -->
    <div ref="commentsSection" class="comments-section">
      <div class="comment-composer">
        <img :src="userStore.user?.avatar || defaultAvatar" :alt="userStore.user?.username" />
        <div class="comment-input-area">
          <textarea
            v-model="commentContent"
            placeholder="发表你的想法"
            rows="3"
          />
          <button
            class="submit-btn"
            @click="submitComment"
            :disabled="!commentContent.trim() || isCommentLoading"
          >
            {{ isCommentLoading ? '发布中...' : '回复' }}
          </button>
        </div>
      </div>

      <!-- 评论列表 -->
      <div class="comments-list">
        <div v-if="isLoadingComments" class="loading">加载中...</div>
        <div
          v-for="comment in comments"
          :key="comment.id"
          class="comment-item"
        >
          <RouterLink :to="`/profile/${comment.userId}`">
            <img :src="comment.author.avatar || defaultAvatar" :alt="comment.author.username" />
          </RouterLink>
          <div class="comment-content">
            <div class="comment-header">
              <RouterLink :to="`/profile/${comment.userId}`" class="username">
                {{ comment.author.username }}
              </RouterLink>
              <span class="handle">@{{ comment.author.username }}</span>
              <span class="dot">·</span>
              <span class="time">{{ formatTime(comment.createdAt) }}</span>
            </div>
            <p class="comment-text">{{ comment.content }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'
import type { Post, Comment } from '@/types'
import { postAPI, commentAPI, likeAPI } from '@/api'
import BackIcon from '@/components/icons/BackIcon.vue'
import CommentIcon from '@/components/icons/CommentIcon.vue'
import RetweetIcon from '@/components/icons/RetweetIcon.vue'
import HeartIcon from '@/components/icons/HeartIcon.vue'
import ShareIcon from '@/components/icons/ShareIcon.vue'
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import 'dayjs/locale/zh-cn'

dayjs.extend(relativeTime)
dayjs.locale('zh-cn')

const route = useRoute()
const userStore = useUserStore()

const post = ref<Post | null>(null)
const comments = ref<Comment[]>([])
const commentContent = ref('')
const isLoading = ref(false)
const isLoadingComments = ref(false)
const isCommentLoading = ref(false)
const commentsSection = ref<HTMLElement>()

const defaultAvatar =
  'https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png'

const postId = route.params.postId as string

const formatTime = (date: string) => {
  return dayjs(date).fromNow()
}

const scrollToComments = () => {
  commentsSection.value?.scrollIntoView({ behavior: 'smooth' })
}

onMounted(async () => {
  try {
    isLoading.value = true
    const res = await postAPI.getPost(postId)
    post.value = res.data

    // 加载评论
    loadComments()
  } catch (error) {
    console.error('Failed to load post:', error)
  } finally {
    isLoading.value = false
  }
})

const loadComments = async () => {
  try {
    isLoadingComments.value = true
    const res = await commentAPI.getPostComments(postId)
    comments.value = res.data.items
  } catch (error) {
    console.error('Failed to load comments:', error)
  } finally {
    isLoadingComments.value = false
  }
}

const submitComment = async () => {
  if (!commentContent.value.trim()) return

  try {
    isCommentLoading.value = true
    const res = await commentAPI.createComment(postId, commentContent.value)
    comments.value.unshift(res.data)
    commentContent.value = ''
    if (post.value) {
      post.value.commentsCount++
    }
  } catch (error) {
    console.error('Failed to submit comment:', error)
  } finally {
    isCommentLoading.value = false
  }
}

const toggleLike = async () => {
  if (!post.value) return

  try {
    if (post.value.isLiked) {
      await likeAPI.unlikePost(postId)
      post.value.isLiked = false
      post.value.likesCount--
    } else {
      await likeAPI.likePost(postId)
      post.value.isLiked = true
      post.value.likesCount++
    }
  } catch (error) {
    console.error('Failed to toggle like:', error)
  }
}
</script>

<style scoped>
.post-detail-page {
  width: 100%;
}

.detail-header {
  position: sticky;
  top: 0;
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px 16px;
  border-bottom: 1px solid #eff3f4;
  background-color: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(12px);
  z-index: 100;
}

.back-btn {
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

.back-btn:hover {
  background-color: rgba(29, 155, 240, 0.1);
}

.back-btn svg {
  width: 20px;
  height: 20px;
}

.post-detail {
  padding: 16px;
  border-bottom: 1px solid #eff3f4;
}

.post-user {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
}

.post-user img {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
}

.user-info {
  display: flex;
  flex-direction: column;
}

.username {
  font-weight: 700;
  color: #0f1419;
}

.handle {
  color: #536471;
  font-size: 14px;
}

.post-content-detail {
  margin-bottom: 16px;
}

.content {
  font-size: 20px;
  line-height: 1.5;
  color: #0f1419;
  white-space: pre-wrap;
  word-wrap: break-word;
  margin-bottom: 12px;
}

.images {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 8px;
  border-radius: 12px;
  overflow: hidden;
}

.images img {
  width: 100%;
  height: auto;
  border-radius: 8px;
}

.post-stats {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 0;
  border-top: 1px solid #eff3f4;
  border-bottom: 1px solid #eff3f4;
  margin-bottom: 16px;
  font-size: 13px;
  color: #536471;
}

.divider {
  width: 1px;
  height: 16px;
  background-color: #eff3f4;
}

.post-actions {
  display: flex;
  justify-content: space-around;
  margin-top: 16px;
}

.action-btn {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background-color: transparent;
  color: #536471;
  transition: all 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.action-btn:hover {
  background-color: rgba(29, 155, 240, 0.1);
  color: #1d9bf0;
}

.action-btn.liked {
  color: #f91880;
}

.action-btn svg {
  width: 18px;
  height: 18px;
}

.comments-section {
  padding: 16px;
}

.comment-composer {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
  padding-bottom: 16px;
  border-bottom: 1px solid #eff3f4;
}

.comment-composer img {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  object-fit: cover;
  flex-shrink: 0;
}

.comment-input-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.comment-input-area textarea {
  width: 100%;
  padding: 12px;
  border: 1px solid #cfd9de;
  border-radius: 12px;
  font-size: 15px;
  font-family: inherit;
  resize: vertical;
}

.comment-input-area textarea:focus {
  border-color: #1d9bf0;
}

.submit-btn {
  align-self: flex-end;
  padding: 8px 24px;
  background-color: #1d9bf0;
  color: #ffffff;
  font-weight: 700;
  border-radius: 20px;
  font-size: 14px;
  transition: all 0.2s;
}

.submit-btn:hover:not(:disabled) {
  background-color: #1a91da;
}

.submit-btn:disabled {
  opacity: 0.5;
}

.comments-list {
  min-height: 200px;
}

.comment-item {
  display: flex;
  gap: 12px;
  padding: 16px 0;
  border-bottom: 1px solid #eff3f4;
}

.comment-item img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
  flex-shrink: 0;
}

.comment-content {
  flex: 1;
}

.comment-header {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 14px;
  margin-bottom: 4px;
}

.comment-text {
  font-size: 15px;
  color: #0f1419;
  white-space: pre-wrap;
  word-wrap: break-word;
}

.loading {
  padding: 32px;
  text-align: center;
  color: #536471;
}

.dot {
  color: #536471;
}

.time {
  color: #536471;
}
</style>
