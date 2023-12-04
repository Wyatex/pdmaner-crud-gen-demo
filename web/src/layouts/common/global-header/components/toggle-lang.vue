<template>
  <n-dropdown :options="options" trigger="hover" :value="language" @select="handleSelect">
    <hover-container class="w-40px h-full" :inverted="theme.header.inverted">
      <icon-cil:language class="text-18px outline-transparent" />
    </hover-container>
  </n-dropdown>
</template>

<script lang="ts" setup>
import { localStg } from '@/utils';

const theme = useThemeStore();
const { locale } = useI18n();

const language = ref<I18nType.LangType>(localStg.get('lang') || 'zh-CN');
const options = [
  {
    label: '中文',
    key: 'zh-CN'
  },
  {
    label: 'English',
    key: 'en'
  },
  {
    label: 'ភាសាខ្មែរ',
    key: 'km-KH'
  }
];
const handleSelect = (key: string) => {
  language.value = key as I18nType.LangType;
  locale.value = key;
  localStg.set('lang', key as I18nType.LangType);
};
</script>
<style scoped></style>
