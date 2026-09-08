<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.runtime.title')" :subtitle="$t('spc.runtime.subtitle')">
      <template #actions>
        <el-select v-model="chartId" filterable :placeholder="$t('spc.runtime.selectChart')" style="width: 280px" @change="loadRuntime">
          <el-option v-for="item in charts" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" />
        </el-select>
        <el-button type="primary" icon="refresh" @click="loadRuntime">{{ $t('common.refresh') }}</el-button>
      </template>
    </SpcPageHeader>
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-3">
      <SpcStatCard :title="$t('spc.runtime.points')" :value="samples.length" icon="DataLine" color="#0ea5e9" />
      <SpcStatCard :title="$t('spc.runtime.latest')" :value="latest" :precision="3" icon="Odometer" color="#10b981" />
      <SpcStatCard title="OOC / OOS" :value="`${oocCount} / ${oosCount}`" icon="Warning" color="#f59e0b" />
    </div>
    <div class="spc-panel">
      <SpcControlChart :samples="samples" :control-limit="runtime.controlLimit" :spec="runtime.spec" :title="$t('spc.runtime.chartTitle')" />
    </div>
    <div class="spc-panel">
      <div class="mb-3 text-sm font-semibold">{{ $t('spc.runtime.samples') }}</div>
      <el-table :data="samples" row-key="ID">
        <el-table-column :label="$t('spc.runtime.subgroup')" prop="subgroupNo" width="100" />
        <el-table-column :label="$t('spc.runtime.sampleTime')" min-width="170"><template #default="{ row }">{{ formatDate(row.sampleTime) }}</template></el-table-column>
        <el-table-column :label="$t('spc.runtime.mean')" width="110"><template #default="{ row }">{{ num(row.meanVal) }}</template></el-table-column>
        <el-table-column :label="$t('spc.runtime.range')" width="110"><template #default="{ row }">{{ num(row.rangeVal) }}</template></el-table-column>
        <el-table-column :label="$t('spc.runtime.std')" width="110"><template #default="{ row }">{{ num(row.stdVal) }}</template></el-table-column>
        <el-table-column label="OOC" width="80"><template #default="{ row }"><el-tag :type="row.oocFlag ? 'warning' : 'success'" round>{{ row.oocFlag ? $t('common.yes') : $t('common.no') }}</el-tag></template></el-table-column>
        <el-table-column label="OOS" width="80"><template #default="{ row }"><el-tag :type="row.oosFlag ? 'danger' : 'success'" round>{{ row.oosFlag ? $t('common.yes') : $t('common.no') }}</el-tag></template></el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { chartApi, getChartRuntime } from '@/api/spc'
import { formatDate } from '@/utils/format'
import { num } from './constants'
import { loadOptions } from './composables/useSpcCrud'
import SpcPageHeader from './components/SpcPageHeader.vue'
import SpcStatCard from './components/SpcStatCard.vue'
import SpcControlChart from './components/SpcControlChart.vue'

const charts = ref([])
const chartId = ref()
const runtime = ref({ samples: [], controlLimit: null, spec: null })
const samples = computed(() => runtime.value.samples || [])
const latest = computed(() => samples.value[samples.value.length - 1]?.meanVal ?? 0)
const oocCount = computed(() => samples.value.filter((i) => i.oocFlag).length)
const oosCount = computed(() => samples.value.filter((i) => i.oosFlag).length)

const loadRuntime = async () => {
  if (!chartId.value) return
  const res = await getChartRuntime({ chartId: chartId.value, limit: 50 })
  if (res.code === 0) runtime.value = res.data || { samples: [] }
}

onMounted(async () => {
  charts.value = await loadOptions(chartApi.getList)
  if (charts.value[0]) {
    chartId.value = charts.value[0].ID
    await loadRuntime()
  }
})
</script>
