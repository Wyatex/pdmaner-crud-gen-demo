<template>
  <n-drawer :show="app.settingDrawerVisible" display-directive="show" :width="330" @mask-click="app.closeSettingDrawer">
    <n-drawer-content :title="$t('layout.settingDrawer.title')" :native-scrollbar="false">
      <dark-mode />
      <layout-mode />
      <theme-color-select />
      <page-func />
      <page-view />
      <theme-config />
    </n-drawer-content>
  </n-drawer>
  <drawer-button v-if="showButton" />
</template>

<script setup lang="ts">
import { DarkMode, DrawerButton, LayoutMode, PageFunc, PageView, ThemeColorSelect, ThemeConfig } from './components';

defineOptions({ name: 'SettingDrawer' });

const app = useAppStore();

const showButton = import.meta.env.VITE_FLOATING_THEME_SETTING_BUTTON === 'Y';
</script>

<style scoped></style>
