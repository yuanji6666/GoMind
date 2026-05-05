<template>
  <div id="app" class="app-container" :class="{ 'auth-page': isAuthPage }">
    <Sidebar v-if="!isAuthPage" />
    <main class="main-content">
      <RouterView />
    </main>
    <Trends v-if="!isAuthPage" />
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import Sidebar from '@/components/Sidebar.vue'
import Trends from '@/components/Trends.vue'
import { onMounted, computed } from 'vue'

const router = useRouter()
const userStore = useUserStore()

const isAuthPage = computed(() => {
  const path = router.currentRoute.value.path
  return path === '/login' || path === '/register'
})

onMounted(async () => {
  // 检查是否已登录
  const token = localStorage.getItem('token')
  if (!token && router.currentRoute.value.path !== '/login' && router.currentRoute.value.path !== '/register') {
    router.push('/login')
  }
})
</script>

<style scoped>
.app-container {
  display: grid;
  grid-template-columns: 275px 1fr 350px;
  min-height: 100vh;
  background-color: #ffffff;
}

.app-container.auth-page {
  grid-template-columns: 1fr;
}

.main-content {
  border-left: 1px solid #eff3f4;
  border-right: 1px solid #eff3f4;
  max-width: 600px;
  min-height: 100vh;
  background-color: #ffffff;
}

.app-container.auth-page .main-content {
  border: none;
  max-width: 100%;
}

@media (max-width: 1200px) {
  .app-container {
    grid-template-columns: 275px 1fr;
  }
}

@media (max-width: 768px) {
  .app-container {
    grid-template-columns: 1fr;
  }
}
</style>
