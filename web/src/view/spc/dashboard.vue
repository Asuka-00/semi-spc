<template>
  <div class="space-y-4">
    <SpcPageHeader :kicker="$t('spc.kicker')" :title="$t('spc.dashboard.title')" :subtitle="$t('spc.dashboard.subtitle')">
      <template #actions>
        <el-select v-model="days" style="width: 140px" @change="loadAll">
          <el-option :value="7" :label="$t('spc.dashboard.days7')" />
          <el-option :value="14" :label="$t('spc.dashboard.days14')" />
          <el-option :value="30" :label="$t('spc.dashboard.days30')" />
        </el-select>
        <el-button type="primary" icon="refresh" @click="loadAll">{{ $t('common.refresh') }}</el-button>
      </template>
    </SpcPageHeader>

    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 xl:grid-cols-4">
      <SpcStatCard :title="$t('spc.dashboard.todayAlarms')" :value="overview.todayAlarms" icon="Warning" color="#f43f5e" :hint="$t('spc.dashboard.todayHint')" />
      <SpcStatCard :title="$t('spc.dashboard.openAlarms')" :value="overview.openAlarms" icon="Bell" color="#f59e0b" :hint="$t('spc.dashboard.openHint')" />
      <SpcStatCard :title="$t('spc.dashboard.oocRate')" :value="overview.oocRate" :precision="2" suffix="%" icon="TrendCharts" color="#0ea5e9" :hint="$t('spc.dashboard.oocHint')" />
      <SpcStatCard :title="$t('spc.dashboard.activeCharts')" :value="overview.activeCharts" icon="DataAnalysis" color="#10b981" :hint="$t('spc.dashboard.chartHint')" />
    </div>

    <div class="grid grid-cols-1 gap-4 xl:grid-cols-12">
      <div class="spc-panel xl:col-span-8">
        <div class="mb-3">
          <div class="text-sm font-semibold text-slate-800 dark:text-slate-100">{{ $t('spc.dashboard.trend') }}</div>
          <div class="text-xs text-slate-400 mt-1">{{ $t('spc.dashboard.trendHint') }}</div>
        </div>
        <Chart v-if="trendOption" :option="trendOption" height="280px" />
      </div>
      <div class="spc-panel xl:col-span-4">
        <div class="mb-3 text-sm font-semibold text-slate-800 dark:text-slate-100">{{ $t('spc.dashboard.composition') }}</div>
        <Chart v-if="pieOption" :option="pieOption" height="280px" />
      </div>
    </div>

    <div class="grid grid-cols-1 gap-4 xl:grid-cols-12">
      <div class="spc-panel xl:col-span-7">
        <div class="mb-3 flex items-center justify-between">
          <span class="text-sm font-semibold text-slate-800 dark:text-slate-100">{{ $t('spc.dashboard.recent') }}</span>
          <el-button link type="primary" @click="$router.push({ name: 'spcAlarm' })">{{ $t('spc.dashboard.gotoAlarm') }}</el-button>
        </div>
        <el-table :data="overview.recentAlarms || []" size="large" class="spc-table-wrap">
          <el-table-column :label="$t('common.type')" width="100">
            <template #default="{ row }">
              <el-tag :type="alarmTypeTag(row.alarmType)" effect="light" round>{{ $t(`spc.option.alarmType.${row.alarmType}`, row.alarmType) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column :label="$t('spc.dashboard.chart')" min-width="140">
            <template #default="{ row }">{{ row.chart?.name || row.chart?.code || '-' }}</template>
          </el-table-column>
          <el-table-column :label="$t('spc.dashboard.rule')" prop="ruleCode" width="110" />
          <el-table-column :label="$t('common.status')" width="110">
            <template #default="{ row }">
              <el-tag :type="alarmStatusTag(row.status)" round>{{ alarmStatusText(row.status, t) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column :label="$t('common.time')" min-width="160">
            <template #default="{ row }">{{ formatDate(row.CreatedAt) }}</template>
          </el-table-column>
        </el-table>
        <el-empty v-if="!(overview.recentAlarms || []).length" :description="$t('spc.dashboard.noAlarm')" class="spc-empty" />
      </div>
      <div class="spc-panel xl:col-span-5">
        <div class="mb-3 text-sm font-semibold text-slate-800 dark:text-slate-100">{{ $t('spc.dashboard.topEqp') }}</div>
        <el-table :data="overview.topEquipment || []">
          <el-table-column :label="$t('spc.equipment.equipment')" min-width="140">
            <template #default="{ row }">
              <div class="font-medium">{{ row.equipmentName }}</div>
              <div class="text-xs text-slate-400">{{ row.equipmentCode }}</div>
            </template>
          </el-table-column>
          <el-table-column :label="$t('spc.dashboard.total')" prop="alarmCount" width="80" />
          <el-table-column label="OOC" prop="oocCount" width="70" />
          <el-table-column label="OOS" prop="oosCount" width="70" />
        </el-table>
        <el-empty v-if="!(overview.topEquipment || []).length" :description="$t('spc.dashboard.noEqp')" class="spc-empty" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { getDashboardOverview } from '@/api/spc'
import { formatDate } from '@/utils/format'
import { alarmStatusTag, alarmStatusText, alarmTypeTag } from './constants'
import SpcPageHeader from './components/SpcPageHeader.vue'
import SpcStatCard from './components/SpcStatCard.vue'
import Chart from '@/components/charts/index.vue'
import useChartOption from '@/hooks/charts'

const { t } = useI18n()
const days = ref(7)
const overview = ref({
  todayAlarms: 0,
  openAlarms: 0,
  oocRate: 0,
  activeCharts: 0,
  trend: [],
  byType: [],
  recentAlarms: [],
  topEquipment: []
})

const trendSource = ref([])
const pieSource = ref([])

const { chartOption: trendOption } = useChartOption((isDark) => ({
  backgroundColor: 'transparent',
  tooltip: { trigger: 'axis' },
  grid: { left: 36, right: 16, top: 24, bottom: 28 },
  xAxis: {
    type: 'category',
    data: trendSource.value.map((i) => i.day),
    axisLabel: { color: isDark ? '#94a3b8' : '#64748b' }
  },
  yAxis: {
    type: 'value',
    minInterval: 1,
    splitLine: { lineStyle: { color: isDark ? '#1e293b' : '#e2e8f0' } }
  },
  series: [
    {
      type: 'line',
      smooth: true,
      data: trendSource.value.map((i) => i.count),
      areaStyle: { color: isDark ? 'rgba(14,165,233,0.18)' : 'rgba(14,165,233,0.12)' },
      lineStyle: { color: '#0ea5e9', width: 2 },
      itemStyle: { color: '#0284c7' }
    }
  ]
}))

const { chartOption: pieOption } = useChartOption((isDark) => ({
  backgroundColor: 'transparent',
  tooltip: { trigger: 'item' },
  legend: { bottom: 0, textStyle: { color: isDark ? '#94a3b8' : '#64748b' } },
  series: [
    {
      type: 'pie',
      radius: ['48%', '72%'],
      center: ['50%', '46%'],
      data: pieSource.value.length
        ? pieSource.value
        : [{ name: t('common.noData'), value: 1, itemStyle: { color: '#cbd5e1' } }],
      label: { color: isDark ? '#e2e8f0' : '#334155' }
    }
  ]
}))

const loadAll = async () => {
  const res = await getDashboardOverview({ days: days.value })
  if (res.code === 0) {
    overview.value = res.data || {}
    trendSource.value = res.data?.trend || []
    pieSource.value = (res.data?.byType || []).map((item) => ({
      name: item.alarmType ? t(`spc.option.alarmType.${item.alarmType}`, item.alarmType) : t('common.noData'),
      value: item.count
    }))
  }
}

onMounted(loadAll)
</script>
