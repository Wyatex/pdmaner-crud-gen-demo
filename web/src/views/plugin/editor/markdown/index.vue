<template>
  <div class="h-full">
    <n-card title="markdown插件" :bordered="false" class="rounded-8px shadow-sm">
      <div ref="domRef"></div>
      <template #footer>
        <github-link link="https://github.com/Vanessa219/vditor" />
      </template>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import Vditor from 'vditor';
import 'vditor/dist/index.css';

const theme = useThemeStore();

const vditor = ref<Vditor>();
const domRef = ref<HTMLElement>();

function renderVditor() {
  if (!domRef.value) return;
  vditor.value = new Vditor(domRef.value, {
    minHeight: 400,
    theme: theme.darkMode ? 'dark' : 'classic',
    icon: 'material',
    cache: { enable: false }
  });
}

const stopHandle = watch(
  () => theme.darkMode,
  newValue => {
    const themeMode = newValue ? 'dark' : 'classic';
    vditor.value?.setTheme(themeMode);
  }
);

onMounted(() => {
  renderVditor();
});

onUnmounted(() => {
  stopHandle();
});
</script>

<style scoped></style>
