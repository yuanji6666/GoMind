<template>
  <aside class="trends">
    <!-- 搜索框 -->
    <div class="search-container">
      <input
        v-model="searchQuery"
        type="text"
        placeholder="搜索 GoMind"
        class="search-input"
        @keyup.enter="handleSearch"
      />
      <SearchIcon />
    </div>

    <!-- 趋势列表 -->
    <div class="trends-container">
      <div class="trends-header">你可能感兴趣的内容</div>
      <div class="trend-item" v-for="(trend, idx) in trends" :key="idx">
        <div class="trend-tag">{{ trend.category }}</div>
        <div class="trend-title">{{ trend.title }}</div>
        <div class="trend-count">{{ trend.count }} 条帖子</div>
      </div>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import SearchIcon from './icons/SearchIcon.vue'

const router = useRouter()
const searchQuery = ref('')

const trends = [
  { category: '科技', title: 'AI 应用', count: '1.2M' },
  { category: '社交', title: '知识分享', count: '856K' },
  { category: '热点', title: '今日关注', count: '432K' },
  { category: '科技', title: '开发工具', count: '321K' },
  { category: '生活', title: '学习分享', count: '218K' },
]

const handleSearch = () => {
  if (searchQuery.value.trim()) {
    router.push(`/search?q=${encodeURIComponent(searchQuery.value)}`)
  }
}
</script>

<style scoped>
.trends {
  display: flex;
  flex-direction: column;
  padding: 16px;
  background-color: #f7f9f9;
  height: 100vh;
  position: sticky;
  top: 0;
  overflow-y: auto;
}

.search-container {
  position: relative;
  margin-bottom: 16px;
}

.search-input {
  width: 100%;
  padding: 12px 16px 12px 40px;
  border-radius: 24px;
  border: 1px solid #cfd9de;
  background-color: #ffffff;
  font-size: 15px;
}

.search-input:focus {
  border-color: #1d9bf0;
  background-color: #ffffff;
}

.search-container svg {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  width: 20px;
  height: 20px;
  color: #536471;
  pointer-events: none;
}

.trends-container {
  background-color: #ffffff;
  border-radius: 16px;
  overflow: hidden;
}

.trends-header {
  padding: 16px;
  font-size: 20px;
  font-weight: 700;
  color: #0f1419;
  border-bottom: 1px solid #eff3f4;
}

.trend-item {
  padding: 16px;
  cursor: pointer;
  transition: all 0.2s;
  border-bottom: 1px solid #eff3f4;
}

.trend-item:hover {
  background-color: #f7f9f9;
}

.trend-tag {
  font-size: 13px;
  color: #536471;
  margin-bottom: 4px;
}

.trend-title {
  font-size: 15px;
  font-weight: 700;
  color: #0f1419;
  margin-bottom: 4px;
}

.trend-count {
  font-size: 13px;
  color: #536471;
}

@media (max-width: 1200px) {
  .trends {
    display: none;
  }
}
</style>
