<template>
  <n-space :vertical="true" :size="16">
    <n-card title="Tab Detail" :bordered="false" size="small" class="rounded-8px shadow-sm">
      <n-space :vertical="true" :size="12">
        <div>当前路由的描述数据(meta)：</div>
        <div>{{ route.meta }}</div>
        <n-button @click="handleToTab">返回Tab</n-button>
      </n-space>
    </n-card>
  </n-space>
</template>

<script setup lang="ts">
import { routeName } from '@/router';

const { routerPush } = useRouterPush();

const route = useRoute();

function handleToTab() {
  routerPush({ name: routeName('function_tab') });
}
</script>

<style scoped></style>
