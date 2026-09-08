<template>
  <el-menu-item
    :index="routerInfo.name"
    :style="{
          height: sideHeight
        }"
  >
    <el-icon v-if="routerInfo.meta.icon">
      <component :is="routerInfo.meta.icon" />
    </el-icon>
    <template v-else>
      {{ isCollapse ? displayTitle[0] : "" }}
    </template>
    <template #title>
      {{ displayTitle }}
    </template>
  </el-menu-item>
</template>

<script setup>
import {computed, inject} from 'vue'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  import { translateMenuTitle } from '@/i18n'
  const appStore = useAppStore()
  const { config } = storeToRefs(appStore)

  defineOptions({
    name: 'MenuItem'
  })

  const props = defineProps({
    routerInfo: {
      default: function () {
        return null
      },
      type: Object
    }
  })

  const displayTitle = computed(() =>
    translateMenuTitle(props.routerInfo?.meta?.title, props.routerInfo?.name)
  )

const isCollapse = inject('isCollapse', {
  default: false
})

  const sideHeight = computed(() => {
    return config.value.layout_side_item_height + 'px'
  })
</script>

<style lang="scss"></style>
