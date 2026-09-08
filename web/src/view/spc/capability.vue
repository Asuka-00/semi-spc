<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.capability.title')" :subtitle="$t('spc.capability.subtitle')" />
    <div class="spc-panel">
      <el-form :inline="true" :model="form">
        <el-form-item :label="$t('spc.capability.chart')">
          <el-select v-model="form.chartId" filterable style="width: 260px">
            <el-option v-for="item in charts" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('spc.capability.from')"><el-date-picker v-model="form.windowFrom" type="datetime" /></el-form-item>
        <el-form-item :label="$t('spc.capability.to')"><el-date-picker v-model="form.windowTo" type="datetime" /></el-form-item>
        <el-form-item>
          <el-button type="primary" :loading="calculating" @click="calc">{{ $t('spc.capability.calculate') }}</el-button>
          <el-button @click="loadHistory">{{ $t('common.refresh') }}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div v-if="latest" class="grid grid-cols-2 gap-4 xl:grid-cols-6">
      <SpcStatCard title="Cp" :value="latest.cp" :precision="3" :hint="capabilityLevel(latest.cp, t).text" icon="DataAnalysis" color="#0ea5e9" />
      <SpcStatCard title="Cpk" :value="latest.cpk" :precision="3" :hint="capabilityLevel(latest.cpk, t).text" icon="TrendCharts" color="#10b981" />
      <SpcStatCard title="Pp" :value="latest.pp" :precision="3" icon="PieChart" color="#8b5cf6" />
      <SpcStatCard title="Ppk" :value="latest.ppk" :precision="3" icon="Histogram" color="#f59e0b" />
      <SpcStatCard :title="$t('spc.capability.n')" :value="latest.n" icon="Tickets" color="#64748b" />
      <SpcStatCard :title="$t('spc.capability.mean')" :value="latest.meanVal" :precision="3" icon="Odometer" color="#0284c7" />
    </div>
    <div class="spc-panel">
      <div class="mb-3 text-sm font-semibold">{{ $t('spc.capability.history') }}</div>
      <el-table :data="history" row-key="ID" :empty-text="$t('spc.capability.empty')">
        <el-table-column :label="$t('spc.capability.chart')" min-width="140"><template #default="{ row }">{{ row.chart?.name || row.chartId }}</template></el-table-column>
        <el-table-column :label="$t('spc.capability.window')" min-width="220"><template #default="{ row }">{{ formatDate(row.windowFrom) }} ~ {{ formatDate(row.windowTo) }}</template></el-table-column>
        <el-table-column label="N" prop="n" width="70" />
        <el-table-column label="Cp" width="90"><template #default="{ row }"><el-tag :type="capabilityLevel(row.cp, t).type">{{ num(row.cp, 3) }}</el-tag></template></el-table-column>
        <el-table-column label="Cpk" width="90"><template #default="{ row }"><el-tag :type="capabilityLevel(row.cpk, t).type">{{ num(row.cpk, 3) }}</el-tag></template></el-table-column>
        <el-table-column label="Pp" width="90"><template #default="{ row }">{{ num(row.pp, 3) }}</template></el-table-column>
        <el-table-column label="Ppk" width="90"><template #default="{ row }">{{ num(row.ppk, 3) }}</template></el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ElMessage } from 'element-plus'
import { calculateCapability, chartApi, getCapabilityList } from '@/api/spc'
import { formatDate } from '@/utils/format'
import { capabilityLevel, num } from './constants'
import { loadOptions } from './composables/useSpcCrud'
import SpcPageHeader from './components/SpcPageHeader.vue'
import SpcStatCard from './components/SpcStatCard.vue'

const { t } = useI18n()
const charts = ref([])
const history = ref([])
const latest = ref(null)
const calculating = ref(false)
const form = reactive({
  chartId: undefined,
  windowFrom: new Date(Date.now() - 7 * 24 * 3600 * 1000),
  windowTo: new Date()
})

const loadHistory = async () => {
  const res = await getCapabilityList({ page: 1, pageSize: 20, chartId: form.chartId })
  if (res.code === 0) history.value = res.data?.list || []
}

const calc = async () => {
  if (!form.chartId) {
    ElMessage.warning(t('common.pleaseSelect'))
    return
  }
  calculating.value = true
  try {
    const res = await calculateCapability({ chartId: form.chartId, windowFrom: form.windowFrom, windowTo: form.windowTo })
    if (res.code === 0) {
      latest.value = res.data
      ElMessage.success(t('common.success'))
      await loadHistory()
    }
  } finally {
    calculating.value = false
  }
}

onMounted(async () => {
  charts.value = await loadOptions(chartApi.getList)
  if (charts.value[0]) form.chartId = charts.value[0].ID
  await loadHistory()
})
</script>
