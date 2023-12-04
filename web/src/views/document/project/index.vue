<template>
  <div class="h-full">
    <iframe class="wh-full" :src="src"></iframe>
  </div>
</template>

<script setup lang="ts">
const src = ref('https://admin-docs.soybeanjs.cn/');
</script>

<style scoped></style>
