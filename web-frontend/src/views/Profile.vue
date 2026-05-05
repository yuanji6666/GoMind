<template>
  <div class="profile-page">
    <!-- 封面 -->
    <div class="profile-cover" :style="{ backgroundImage: `url(${userProfile?.cover})` }" />

    <!-- 个人信息 -->
    <div class="profile-info">
      <div class="profile-header">
        <img :src="userProfile?.avatar || defaultAvatar" :alt="userProfile?.username" />
        <button class="follow-btn" v-if="!isOwnProfile" @click="toggleFollow">
          {{ isFollowing ? '取消关注' : '关注' }}
        </button>
        <button class="edit-btn" v-else @click="showEditModal = true">编辑个人资料</button>
      </div>

      <div class="profile-details">
        <h2 class="profile-name">{{ userProfile?.username }}</h2>
        <p class="profile-handle">@{{ userProfile?.username }}</p>
        <p v-if="userProfile?.bio" class="profile-bio">{{ userProfile?.bio }}</p>

        <div class="profile-meta">
          <span v-if="userProfile?.location" class="meta-item">
            <LocationIcon /> {{ userProfile?.location }}
          </span>
          <span v-if="userProfile?.website" class="meta-item">
            <LinkIcon /> {{ userProfile?.website }}
          </span>
          <span class="meta-item">
            加入于 {{ formatDate(userProfile?.createdAt) }}
          </span>
        </div>

        <div class="profile-stats">
          <div class="stat">
            <span class="stat-count">{{ userProfile?.followingCount }}</span>
            <span class="stat-label">正在关注</span>
          </div>
          <div class="stat">
            <span class="stat-count">{{ userProfile?.followersCount }}</span>
            <span class="stat-label">粉丝</span>
          </div>
        </div>
      </div>
    </div>

    <!-- 选项卡 -->
    <div class="profile-tabs">
      <button
        class="tab"
        :class="{ active: activeTab === 'posts' }"
        @click="activeTab = 'posts'"
      >
        帖子
      </button>
      <button
        class="tab"
        :class="{ active: activeTab === 'likes' }"
        @click="activeTab = 'likes'"
      >
        赞过的帖子
      </button>
    </div>

    <!-- 帖子列表 -->
    <div class="profile-posts">
      <div v-if="isLoading" class="loading">加载中...</div>
      <template v-else>
        <Post
          v-for="post in displayPosts"
          :key="post.id"
          :post="post"
        />
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'
import type { User } from '@/types'
import { userAPI, followAPI } from '@/api'
import Post from '@/components/Post.vue'
import LocationIcon from '@/components/icons/LocationIcon.vue'
import LinkIcon from '@/components/icons/LinkIcon.vue'
import dayjs from 'dayjs'

const route = useRoute()
const userStore = useUserStore()

const userProfile = ref<User | null>(null)
const isLoading = ref(false)
const activeTab = ref('posts')
const isFollowing = ref(false)
const postsForProfile = ref<any[]>([])
const likesForProfile = ref<any[]>([])
const showEditModal = ref(false)

const defaultAvatar =
  'https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png'

const userId = computed(() => parseInt(route.params.userId as string))
const isOwnProfile = computed(() => userStore.user?.id === userId.value)

const displayPosts = computed(() => {
  return activeTab.value === 'posts' ? postsForProfile.value : likesForProfile.value
})

onMounted(async () => {
  try {
    isLoading.value = true

    // 获取用户信息
    const userRes = await userAPI.getUserInfo(userId.value)
    userProfile.value = userRes.data

    // 获取关注状态
    if (!isOwnProfile.value) {
      const followRes = await followAPI.checkFollowing(userId.value)
      isFollowing.value = followRes.data
    }

    // 获取用户帖子
    const postsRes = await userAPI.getUserPosts(userId.value)
    postsForProfile.value = postsRes.data.items

    // 获取赞过的帖子
    const likesRes = await userAPI.getUserLikes(userId.value)
    likesForProfile.value = likesRes.data.items
  } catch (error) {
    console.error('Failed to load profile:', error)
  } finally {
    isLoading.value = false
  }
})

const toggleFollow = async () => {
  try {
    if (isFollowing.value) {
      await followAPI.unfollowUser(userId.value)
      isFollowing.value = false
      if (userProfile.value) {
        userProfile.value.followersCount--
      }
    } else {
      await followAPI.followUser(userId.value)
      isFollowing.value = true
      if (userProfile.value) {
        userProfile.value.followersCount++
      }
    }
  } catch (error) {
    console.error('Failed to toggle follow:', error)
  }
}

const formatDate = (date: string) => {
  return dayjs(date).format('YYYY年M月')
}
</script>

<style scoped>
.profile-page {
  width: 100%;
}

.profile-cover {
  width: 100%;
  height: 200px;
  background-size: cover;
  background-position: center;
  background-color: #cfd9de;
  border-bottom: 1px solid #eff3f4;
}

.profile-info {
  padding: 16px;
  border-bottom: 1px solid #eff3f4;
}

.profile-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  margin-bottom: 16px;
  margin-top: -56px;
}

.profile-header img {
  width: 112px;
  height: 112px;
  border-radius: 50%;
  object-fit: cover;
  border: 4px solid #ffffff;
  background-color: #cfd9de;
}

.follow-btn,
.edit-btn {
  padding: 12px 32px;
  background-color: #1d9bf0;
  color: #ffffff;
  font-weight: 700;
  border-radius: 24px;
  font-size: 15px;
  transition: all 0.2s;
}

.follow-btn:hover,
.edit-btn:hover {
  background-color: #1a91da;
}

.profile-details {
  margin-bottom: 16px;
}

.profile-name {
  font-size: 20px;
  font-weight: 700;
  color: #0f1419;
  margin-bottom: 0;
}

.profile-handle {
  color: #536471;
  font-size: 15px;
  margin-bottom: 8px;
}

.profile-bio {
  font-size: 15px;
  color: #0f1419;
  margin-bottom: 12px;
}

.profile-meta {
  display: flex;
  gap: 16px;
  margin-bottom: 16px;
  font-size: 13px;
  color: #536471;
  flex-wrap: wrap;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.meta-item svg {
  width: 16px;
  height: 16px;
}

.profile-stats {
  display: flex;
  gap: 24px;
}

.stat {
  display: flex;
  align-items: baseline;
  gap: 4px;
}

.stat-count {
  font-weight: 700;
  color: #0f1419;
  font-size: 15px;
}

.stat-label {
  color: #536471;
  font-size: 13px;
}

.profile-tabs {
  display: flex;
  border-bottom: 1px solid #eff3f4;
  position: sticky;
  top: 0;
  background-color: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(12px);
  z-index: 100;
}

.tab {
  flex: 1;
  padding: 16px;
  background-color: transparent;
  color: #536471;
  font-weight: 500;
  font-size: 15px;
  transition: all 0.2s;
  border-bottom: 2px solid transparent;
}

.tab:hover {
  background-color: #f7f9f9;
}

.tab.active {
  color: #1d9bf0;
  border-bottom-color: #1d9bf0;
}

.profile-posts {
  min-height: 200px;
}

.loading {
  padding: 32px;
  text-align: center;
  color: #536471;
}
</style>
