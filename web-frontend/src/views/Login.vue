<template>
  <div class="login-container">
    <div class="login-box">
      <!-- Logo -->
      <div class="login-header">
        <svg viewBox="0 0 24 24" class="logo">
          <path
            fill="currentColor"
            d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24h-6.627l-5.1-6.694L2.2 21.75H.897l7.671-8.757L.244 2.25h6.814l4.612 6.096L17.07 2.25h.174zm-1.161 17.52h1.833L5.83 4.125H3.863L17.083 19.77z"
          />
        </svg>
        <h1>GoMind</h1>
      </div>

      <!-- 登录表单 -->
      <form @submit.prevent="handleLogin" class="login-form">
        <h2>登录 GoMind</h2>

        <div class="form-group">
          <input
            v-model="username"
            type="text"
            placeholder="用户名或邮箱"
            required
            @keyup.enter="handleLogin"
          />
        </div>

        <div class="form-group">
          <input
            v-model="password"
            type="password"
            placeholder="密码"
            required
            @keyup.enter="handleLogin"
          />
        </div>

        <div v-if="error" class="error-message">{{ error }}</div>

        <button type="submit" class="login-btn" :disabled="isLoading">
          {{ isLoading ? '登录中...' : '登录' }}
        </button>
      </form>

      <!-- 注册链接 -->
      <p class="signup-link">
        没有账号？<RouterLink to="/register">注册</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const username = ref('')
const password = ref('')
const isLoading = ref(false)
const error = ref('')

const handleLogin = async () => {
  if (!username.value || !password.value) {
    error.value = '请输入用户名和密码'
    return
  }

  try {
    isLoading.value = true
    error.value = ''
    await userStore.login(username.value, password.value)
    router.push('/')
  } catch (err: any) {
    error.value = err.response?.data?.message || '登录失败，请重试'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #ffffff;
}

.login-box {
  background-color: #ffffff;
  border-radius: 16px;
  padding: 40px;
  width: 100%;
  max-width: 400px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.login-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 32px;
  text-align: center;
}

.logo {
  width: 40px;
  height: 40px;
  color: #1d9bf0;
}

.login-header h1 {
  font-size: 32px;
  font-weight: 700;
  color: #0f1419;
  margin: 0;
}

.login-form h2 {
  font-size: 24px;
  font-weight: 700;
  margin-bottom: 24px;
  color: #0f1419;
}

.form-group {
  margin-bottom: 16px;
}

.form-group input {
  width: 100%;
  padding: 16px;
  font-size: 16px;
  border: 1px solid #cfd9de;
  border-radius: 8px;
  transition: all 0.2s;
}

.form-group input:focus {
  border-color: #1d9bf0;
  box-shadow: 0 0 0 2px rgba(29, 155, 240, 0.1);
}

.error-message {
  padding: 12px;
  background-color: #ffe6e6;
  color: #c41e3a;
  border-radius: 8px;
  font-size: 14px;
  margin-bottom: 16px;
}

.login-btn {
  width: 100%;
  padding: 16px;
  background-color: #1d9bf0;
  color: #ffffff;
  font-weight: 700;
  font-size: 16px;
  border-radius: 8px;
  transition: all 0.2s;
  margin-bottom: 16px;
}

.login-btn:hover:not(:disabled) {
  background-color: #1a91da;
}

.login-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.signup-link {
  text-align: center;
  color: #536471;
  font-size: 14px;
}

.signup-link a {
  color: #1d9bf0;
  font-weight: 700;
  text-decoration: none;
}

.signup-link a:hover {
  text-decoration: underline;
}
</style>
