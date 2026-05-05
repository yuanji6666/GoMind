<template>
  <aside class="sidebar">
    <!-- Logo -->
    <div class="sidebar-logo">
      <svg viewBox="0 0 24 24" class="logo-icon">
        <path
          fill="currentColor"
          d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24h-6.627l-5.1-6.694L2.2 21.75H.897l7.671-8.757L.244 2.25h6.814l4.612 6.096L17.07 2.25h.174zm-1.161 17.52h1.833L5.83 4.125H3.863L17.083 19.77z"
        />
      </svg>
    </div>

    <!-- 导航菜单 -->
    <nav class="sidebar-nav">
      <RouterLink to="/" class="nav-item" :class="{ active: isActive('/') }">
        <HomeIcon />
        <span>主页</span>
      </RouterLink>

      <RouterLink to="/explore" class="nav-item" :class="{ active: isActive('/explore') }">
        <ExploreIcon />
        <span>探索</span>
      </RouterLink>

      <RouterLink
        to="/notifications"
        class="nav-item"
        :class="{ active: isActive('/notifications') }"
      >
        <NotificationIcon />
        <span>通知</span>
      </RouterLink>

      <RouterLink to="/messages" class="nav-item" :class="{ active: isActive('/messages') }">
        <MessageIcon />
        <span>消息</span>
      </RouterLink>

      <RouterLink to="/bookmarks" class="nav-item" :class="{ active: isActive('/bookmarks') }">
        <BookmarkIcon />
        <span>书签</span>
      </RouterLink>

      <RouterLink
        :to="`/profile/${userStore.user?.id}`"
        class="nav-item"
        :class="{ active: isActive('/profile') }"
      >
        <ProfileIcon />
        <span>个人资料</span>
      </RouterLink>
    </nav>

    <!-- 发布按钮 -->
    <button class="compose-btn" @click="showComposeModal = true">发布</button>

    <!-- 用户信息 -->
    <div v-if="userStore.user" class="sidebar-user">
      <img :src="userStore.user.avatar || defaultAvatar" :alt="userStore.user.username" />
      <div class="user-info">
        <div class="user-name">{{ userStore.user.username }}</div>
        <div class="user-handle">@{{ userStore.user.username }}</div>
      </div>
      <button class="user-menu-btn" @click="showUserMenu = !showUserMenu">...</button>
      <div v-if="showUserMenu" class="user-menu">
        <button @click="logout">退出登录</button>
      </div>
    </div>

    <!-- 发布模态框 -->
    <ComposeModal v-if="showComposeModal" @close="showComposeModal = false" />
  </aside>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import ComposeModal from './ComposeModal.vue'
import HomeIcon from './icons/HomeIcon.vue'
import ExploreIcon from './icons/ExploreIcon.vue'
import NotificationIcon from './icons/NotificationIcon.vue'
import MessageIcon from './icons/MessageIcon.vue'
import BookmarkIcon from './icons/BookmarkIcon.vue'
import ProfileIcon from './icons/ProfileIcon.vue'

const router = useRouter()
const userStore = useUserStore()
const showComposeModal = ref(false)
const showUserMenu = ref(false)

const defaultAvatar =
  'https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png'

const isActive = (path: string) => {
  return router.currentRoute.value.path.startsWith(path)
}

const logout = () => {
  userStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.sidebar {
  display: flex;
  flex-direction: column;
  padding: 16px;
  border-right: 1px solid #eff3f4;
  height: 100vh;
  position: sticky;
  top: 0;
  overflow-y: auto;
  background-color: #ffffff;
}

.sidebar-logo {
  width: 56px;
  height: 56px;
  margin-bottom: 32px;
  color: #1d9bf0;
}

.logo-icon {
  width: 100%;
  height: 100%;
}

.sidebar-nav {
  display: flex;
  flex-direction: column;
  gap: 16px;
  flex: 1;
  margin-bottom: 16px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px 16px;
  border-radius: 28px;
  font-size: 20px;
  font-weight: 500;
  color: #0f1419;
  transition: all 0.2s;
}

.nav-item:hover {
  background-color: #f7f9f9;
}

.nav-item.active {
  color: #1d9bf0;
}

.nav-item svg {
  width: 24px;
  height: 24px;
}

.compose-btn {
  width: 100%;
  padding: 12px 24px;
  border-radius: 24px;
  background-color: #1d9bf0;
  color: #ffffff;
  font-weight: 700;
  font-size: 15px;
  margin-bottom: 16px;
  transition: all 0.2s;
}

.compose-btn:hover {
  background-color: #1a91da;
}

.sidebar-user {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  border-radius: 12px;
  background-color: #f7f9f9;
  position: relative;
}

.sidebar-user img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.user-info {
  flex: 1;
  min-width: 0;
}

.user-name {
  font-weight: 700;
  font-size: 14px;
  color: #0f1419;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.user-handle {
  font-size: 13px;
  color: #536471;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.user-menu-btn {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: transparent;
  color: #536471;
  font-size: 18px;
  transition: all 0.2s;
}

.user-menu-btn:hover {
  background-color: rgba(29, 155, 240, 0.1);
  color: #1d9bf0;
}

.user-menu {
  position: absolute;
  bottom: -50px;
  right: 0;
  background-color: #ffffff;
  border: 1px solid #eff3f4;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 100;
}

.user-menu button {
  display: block;
  width: 100%;
  padding: 12px 16px;
  text-align: left;
  color: #0f1419;
  font-size: 15px;
  transition: all 0.2s;
}

.user-menu button:hover {
  background-color: #f7f9f9;
}
</style>
