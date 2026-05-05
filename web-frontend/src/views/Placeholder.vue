<template>
  <div class="placeholder-page">
    <div class="placeholder-header">
      <h2>{{ title }}</h2>
    </div>
    <div class="placeholder-content">
      <p>{{ title }}功能即将上线</p>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  title: string
}>()
</script>

<style scoped>
.placeholder-page {
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.placeholder-header {
  position: sticky;
  top: 0;
  padding: 16px;
  border-bottom: 1px solid #eff3f4;
  background-color: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(12px);
  z-index: 100;
}

.placeholder-header h2 {
  font-size: 20px;
  font-weight: 700;
  color: #0f1419;
  margin: 0;
}

.placeholder-content {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #536471;
  font-size: 18px;
}
</style>
