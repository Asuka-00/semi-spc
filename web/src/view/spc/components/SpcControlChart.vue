<template>
  <div class="w-full">
    <Chart v-if="option && samples.length" :option="option" height="360px" />
    <div v-else class="spc-empty flex items-center justify-center h-[360px] rounded-xl border border-dashed border-slate-200 dark:border-slate-700">
      <el-empty :description="$t('spc.runtime.empty')" />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import Chart from '@/components/charts/index.vue'
import useChartOption from '@/hooks/charts'

const { t } = useI18n()

const props = defineProps({
  samples: { type: Array, default: () => [] },
  controlLimit: { type: Object, default: null },
  spec: { type: Object, default: null },
  title: { type: String, default: '' }
})

const { chartOption: option } = useChartOption((isDark) => {
  const axis = props.samples.map((item, index) => item.subgroupNo || index + 1)
  const means = props.samples.map((item) => item.meanVal)
  const pointColors = props.samples.map((item) => {
    if (item.oosFlag) return '#f43f5e'
    if (item.oocFlag) return '#f59e0b'
    return isDark ? '#38bdf8' : '#0ea5e9'
  })
  const markLineData = []
  const addLine = (value, name, color) => {
    if (value === null || value === undefined) return
    markLineData.push({
      yAxis: value,
      name,
      label: { formatter: `${name} {c}`, color },
      lineStyle: { color, type: ['USL', 'LSL'].includes(name) ? 'dotted' : 'solid', width: 1.5 }
    })
  }
  addLine(props.controlLimit?.ucl, 'UCL', '#f43f5e')
  addLine(props.controlLimit?.cl, 'CL', '#64748b')
  addLine(props.controlLimit?.lcl, 'LCL', '#f43f5e')
  addLine(props.spec?.usl, 'USL', '#a855f7')
  addLine(props.spec?.lsl, 'LSL', '#a855f7')

  return {
    backgroundColor: 'transparent',
    tooltip: {
      trigger: 'axis',
      axisPointer: { type: 'cross' }
    },
    grid: { left: 48, right: 24, top: 36, bottom: 40 },
    title: {
      text: props.title || t('spc.runtime.chartTitle'),
      left: 0,
      textStyle: { fontSize: 13, fontWeight: 600, color: isDark ? '#e2e8f0' : '#0f172a' }
    },
    xAxis: {
      type: 'category',
      data: axis,
      name: t('spc.runtime.subgroup'),
      axisLabel: { color: isDark ? '#94a3b8' : '#64748b' }
    },
    yAxis: {
      type: 'value',
      scale: true,
      splitLine: { lineStyle: { color: isDark ? '#1e293b' : '#e2e8f0' } },
      axisLabel: { color: isDark ? '#94a3b8' : '#64748b' }
    },
    series: [
      {
        type: 'line',
        name: t('spc.runtime.mean'),
        data: means,
        smooth: true,
        showSymbol: true,
        symbolSize: 9,
        lineStyle: { width: 2, color: isDark ? '#38bdf8' : '#0284c7' },
        itemStyle: {
          color: (params) => pointColors[params.dataIndex]
        },
        markLine: {
          silent: true,
          symbol: 'none',
          data: markLineData
        }
      }
    ]
  }
})
</script>
