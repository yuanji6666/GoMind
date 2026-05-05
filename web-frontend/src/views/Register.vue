<template>
  <div class="register-container">
    <div class="register-box">
      <!-- Logo -->
      <div class="register-header">
        <svg viewBox="0 0 24 24" class="logo">
          <path
            fill="currentColor"
            d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24h-6.627l-5.1-6.694L2.2 21.75H.897l7.671-8.757L.244 2.25h6.814l4.612 6.096L17.07 2.25h.174zm-1.161 17.52h1.833L5.83 4.125H3.863L17.083 19.77z"
          />
        </svg>
        <h1>GoMind</h1>
      </div>

      <!-- 注册表单 -->
      <form @submit.prevent="handleRegister" class="register-form">
        <h2>创建账号</h2>

        <div class="form-group">
          <input
            v-model="email"
            type="email"
            placeholder="邮箱地址"
            required
          />
        </div>

        <div class="form-group">
          <div class="captcha-group">
            <input
              v-model="captcha"
              type="text"
              placeholder="验证码"
              required
            />
            <button
              type="button"
              class="captcha-btn"
              @click="sendCaptcha"
              :disabled="captchaLoading || captchaCount > 0"
            >
              {{ captchaCount > 0 ? `${captchaCount}s` : '获取验证码' }}
            </button>
          </div>
        </div>

        <div class="form-group">
          <input
            v-model="username"
            type="text"
            placeholder="用户名"
            required
          />
        </div>

        <div class="form-group">
          <input
            v-model="password"
            type="password"
            placeholder="密码"
            required
          />
        </div>

        <div class="form-group">
          <input
            v-model="confirmPassword"
            type="password"
            placeholder="确认密码"
            required
          />
        </div>

        <div v-if="error" class="error-message">{{ error }}</div>

        <button type="submit" class="register-btn" :disabled="isLoading">
          {{ isLoading ? '注册中...' : '注册' }}
        </button>
      </form>

      <!-- 登录链接 -->
      <p class="login-link">
        已有账号？<RouterLink to="/login">登录</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { authAPI } from '@/api'

const router = useRouter()
const userStore = useUserStore()

const email = ref('')
const username = ref('')
const password = ref('')
const confirmPassword = ref('')
const captcha = ref('')
const isLoading = ref(false)
const captchaLoading = ref(false)
const captchaCount = ref(0)
const error = ref('')

const sendCaptcha = async () => {
  if (!email.value) {
    error.value = '请先输入邮箱地址'
    return
  }

  try {
    captchaLoading.value = true
    await authAPI.getCaptcha(email.value)
    error.value = ''
    
    // 倒计时 60 秒
    captchaCount.value = 60
    const timer = setInterval(() => {
      captchaCount.value--
      if (captchaCount.value <= 0) {
        clearInterval(timer)
      }
    }, 1000)
  } catch (err: any) {
    error.value = err.response?.data?.message || '发送验证码失败'
  } finally {
    captchaLoading.value = false
  }
}

const handleRegister = async () => {
  if (!email.value || !username.value || !password.value || !captcha.value) {
    error.value = '请填写所有字段'
    return
  }

  if (password.value !== confirmPassword.value) {
    error.value = '两次输入的密码不一致'
    return
  }

  if (password.value.length < 6) {
    error.value = '密码长度至少为 6 位'
    return
  }

  try {
    isLoading.value = true
    error.value = ''
    await userStore.register(email.value, username.value, password.value, captcha.value)
    router.push('/')
  } catch (err: any) {
    error.value = err.response?.data?.message || '注册失败，请重试'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.register-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #ffffff;
}

.register-box {
  background-color: #ffffff;
  border-radius: 16px;
  padding: 40px;
  width: 100%;
  max-width: 400px;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.register-header {
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

.register-header h1 {
  font-size: 32px;
  font-weight: 700;
  color: #0f1419;
  margin: 0;
}

.register-form h2 {
  font-size: 24px;
  font-weight: 700;
  margin-bottom: 24px;
  color: #0f1419;
}

.form-group {
  margin-bottom: 16px;
}

.captcha-group {
  display: flex;
  gap: 8px;
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

.captcha-group input {
  flex: 1;
}

.captcha-btn {
  padding: 16px 12px;
  background-color: #f0f0f0;
  color: #0f1419;
  border-radius: 8px;
  font-size: 14px;
  white-space: nowrap;
  font-weight: 500;
  transition: all 0.2s;
}

.captcha-btn:hover:not(:disabled) {
  background-color: #e0e0e0;
}

.captcha-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.error-message {
  padding: 12px;
  background-color: #ffe6e6;
  color: #c41e3a;
  border-radius: 8px;
  font-size: 14px;
  margin-bottom: 16px;
}

.register-btn {
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

.register-btn:hover:not(:disabled) {
  background-color: #1a91da;
}

.register-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.login-link {
  text-align: center;
  color: #536471;
  font-size: 14px;
}

.login-link a {
  color: #1d9bf0;
  font-weight: 700;
  text-decoration: none;
}

.login-link a:hover {
  text-decoration: underline;
}
</style>
