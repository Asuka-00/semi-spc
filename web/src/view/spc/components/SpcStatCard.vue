<template>
  <div class="spc-stat flex items-start justify-between gap-3">
    <div>
      <div class="spc-stat-label">{{ title }}</div>
      <div class="spc-stat-value">
        {{ display }}
        <span v-if="suffix" class="ml-1 text-sm font-normal text-slate-400">{{ suffix }}</span>
      </div>
      <div v-if="hint" class="mt-2 text-xs text-slate-400">{{ hint }}</div>
    </div>
    <div class="spc-icon-wrap" :style="{ background: color }">
      <el-icon :size="20"><component :is="icon" /></el-icon>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  title: { type: String, required: true },
  value: { type: [Number, String], default: 0 },
  suffix: { type: String, default: '' },
  hint: { type: String, default: '' },
  icon: { type: [String, Object], default: 'DataAnalysis' },
  color: { type: String, default: 'var(--el-color-primary)' },
  precision: { type: Number, default: null }
})

const display = computed(() => {
  if (props.value === null || props.value === undefined || props.value === '') return '-'
  if (props.precision != null && typeof props.value === 'number') {
    return props.value.toFixed(props.precision)
  }
  return props.value
})
</script>
