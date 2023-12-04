<template>
  <div class="h-full">
    <iframe class="wh-full" :src="src"></iframe>
  </div>
</template>

<script setup lang="ts">
const src = ref('https://cn.vitejs.dev/');
</script>

<style scoped></style>
