<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.chart.title')" :subtitle="$t('spc.chart.subtitle')" />
    <el-tabs v-model="activeTab" class="spc-panel !px-4 !pt-2" @tab-change="onTabChange">
      <el-tab-pane :label="$t('spc.chart.tabChart')" name="chart">
        <div class="gva-search-box !bg-transparent !p-0 !my-3">
          <el-form :inline="true" :model="chart.searchInfo">
            <el-form-item :label="$t('common.keyword')"><el-input v-model="chart.searchInfo.keyword" clearable @keyup.enter="chart.onSubmit" /></el-form-item>
            <el-form-item>
              <el-button type="primary" icon="search" @click="chart.onSubmit">{{ $t('common.query') }}</el-button>
              <el-button icon="refresh" @click="chart.onReset()">{{ $t('common.reset') }}</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openChart('add')">{{ $t('spc.chart.addChart') }}</el-button></div>
        <el-table v-loading="chart.loading" :data="chart.tableData" row-key="ID">
          <el-table-column :label="$t('common.code')" prop="code" min-width="120" />
          <el-table-column :label="$t('common.name')" prop="name" min-width="150" />
          <el-table-column :label="$t('spc.chart.chartType')" min-width="130"><template #default="{ row }">{{ $t(`spc.option.chartType.${row.chartType}`, row.chartType) }}</template></el-table-column>
          <el-table-column :label="$t('spc.chart.parameter')" min-width="120"><template #default="{ row }">{{ row.parameter?.name || row.parameterId }}</template></el-table-column>
          <el-table-column :label="$t('spc.chart.subgroup')" prop="subgroupSize" width="100" />
          <el-table-column :label="$t('common.status')" width="90"><template #default="{ row }"><el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag></template></el-table-column>
          <el-table-column :label="$t('common.action')" fixed="right" width="240">
            <template #default="{ row }">
              <el-button type="primary" link @click="openChart('edit', row)">{{ $t('common.edit') }}</el-button>
              <el-button type="primary" link @click="gotoRules(row)">{{ $t('spc.chart.configRules') }}</el-button>
              <el-button type="danger" link @click="chart.deleteRow(row, $t('spc.chart.tabChart'))">{{ $t('common.delete') }}</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="chart.page" :page-size="chart.pageSize" :total="chart.total" layout="total, prev, pager, next" @current-change="chart.handleCurrentChange" @size-change="chart.handleSizeChange" /></div>
      </el-tab-pane>

      <el-tab-pane :label="$t('spc.chart.tabLimit')" name="limit">
        <div class="gva-search-box !bg-transparent !p-0 !my-3">
          <el-form :inline="true">
            <el-form-item :label="$t('spc.chart.tabChart')">
              <el-select v-model="limit.searchInfo.chartId" filterable clearable :placeholder="$t('spc.chart.selectChart')" @change="limit.onSubmit">
                <el-option v-for="item in charts" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" />
              </el-select>
            </el-form-item>
            <el-form-item><el-button type="primary" icon="plus" @click="openLimit('add')">{{ $t('spc.chart.addLimit') }}</el-button></el-form-item>
            <el-form-item><el-button :disabled="!limit.searchInfo.chartId" :loading="recalcing" @click="recalcLimits">{{ $t('spc.chart.recalc') }}</el-button></el-form-item>
          </el-form>
        </div>
        <el-table v-loading="limit.loading" :data="limit.tableData" row-key="ID">
          <el-table-column :label="$t('spc.chart.tabChart')" min-width="140"><template #default="{ row }">{{ row.chart?.name || row.chartId }}</template></el-table-column>
          <el-table-column label="UCL" width="100"><template #default="{ row }">{{ num(row.ucl) }}</template></el-table-column>
          <el-table-column label="CL" width="100"><template #default="{ row }">{{ num(row.cl) }}</template></el-table-column>
          <el-table-column label="LCL" width="100"><template #default="{ row }">{{ num(row.lcl) }}</template></el-table-column>
          <el-table-column label="UCL-S" width="100"><template #default="{ row }">{{ num(row.uclS) }}</template></el-table-column>
          <el-table-column :label="$t('spc.chart.source')" width="110"><template #default="{ row }">{{ $t(`spc.option.limitSource.${row.source}`, row.source) }}</template></el-table-column>
          <el-table-column :label="$t('spc.chart.calcN')" prop="calcN" width="110" />
          <el-table-column :label="$t('common.action')" width="160"><template #default="{ row }"><el-button type="primary" link @click="openLimit('edit', row)">{{ $t('common.edit') }}</el-button><el-button type="danger" link @click="limit.deleteRow(row, $t('spc.chart.tabLimit'))">{{ $t('common.delete') }}</el-button></template></el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="limit.page" :page-size="limit.pageSize" :total="limit.total" layout="total, prev, pager, next" @current-change="limit.handleCurrentChange" @size-change="limit.handleSizeChange" /></div>
      </el-tab-pane>

      <el-tab-pane :label="$t('spc.chart.tabRule')" name="rule">
        <div class="mb-4 flex flex-col gap-3 lg:flex-row lg:items-center lg:justify-between">
          <el-select v-model="ruleChartId" filterable :placeholder="$t('spc.chart.selectChart')" style="width: 320px" @change="loadRules">
            <el-option v-for="item in charts" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" />
          </el-select>
          <div class="flex gap-2">
            <el-button @click="applyDefaults">{{ $t('spc.chart.applyDefault') }}</el-button>
            <el-button type="primary" :loading="savingRules" :disabled="!ruleChartId" @click="saveRules">{{ $t('spc.chart.saveRules') }}</el-button>
          </div>
        </div>
        <p class="text-sm text-slate-500 mb-4">{{ $t('spc.chart.ruleHint') }}</p>
        <el-empty v-if="!ruleChartId" :description="$t('spc.chart.selectChart')" />
        <div v-else class="space-y-6">
          <div v-for="family in ruleFamilies" :key="family.key">
            <div class="mb-3 text-sm font-semibold text-slate-800 dark:text-slate-100">{{ family.label }}</div>
            <div class="grid grid-cols-1 gap-3">
              <div v-for="item in family.items" :key="item.ruleCode" class="rounded-xl border border-slate-200 dark:border-slate-700 p-4">
                <div class="flex flex-col gap-3 lg:flex-row lg:items-start lg:justify-between">
                  <div class="flex items-start gap-3">
                    <el-switch v-model="item.enabled" />
                    <div>
                      <div class="font-medium text-slate-800 dark:text-slate-100">{{ $t(item.nameKey) }}</div>
                      <div class="text-xs text-slate-400 mt-1 max-w-2xl">{{ $t(item.descKey) }}</div>
                    </div>
                  </div>
                  <div class="flex flex-wrap gap-3">
                    <el-form-item v-if="item.params.includes('n')" :label="$t('spc.chart.n')" class="!mb-0">
                      <el-input-number v-model="item.n" :min="1" :max="50" :disabled="!item.enabled" />
                    </el-form-item>
                    <el-form-item v-if="item.params.includes('hits')" :label="$t('spc.chart.hits')" class="!mb-0">
                      <el-input-number v-model="item.hits" :min="1" :max="item.n || 50" :disabled="!item.enabled" />
                    </el-form-item>
                    <el-form-item v-if="item.params.includes('k')" :label="$t('spc.chart.k')" class="!mb-0">
                      <el-input-number v-model="item.k" :min="0" :max="6" :step="0.1" :precision="2" :disabled="!item.enabled" />
                    </el-form-item>
                    <el-form-item v-if="item.params.includes('severity')" :label="$t('spc.chart.severity')" class="!mb-0">
                      <el-select v-model="item.severity" style="width: 120px" :disabled="!item.enabled">
                        <el-option v-for="sev in severities" :key="sev.value" :label="sev.label" :value="sev.value" />
                      </el-select>
                    </el-form-item>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>

    <el-drawer v-model="chart.drawerVisible" :title="chart.drawerTitle" size="560px" destroy-on-close>
      <el-form :ref="bindFormRef(chart)" :model="chart.formData" :rules="chartRules" label-width="130px">
        <el-form-item :label="$t('spc.chart.chartCode')" prop="code"><el-input v-model="chart.formData.code" /></el-form-item>
        <el-form-item :label="$t('spc.chart.chartName')" prop="name"><el-input v-model="chart.formData.name" /></el-form-item>
        <el-form-item :label="$t('spc.chart.parameter')" prop="parameterId"><el-select v-model="chart.formData.parameterId" class="w-full" filterable><el-option v-for="item in params" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.chart.spec')" prop="specId"><el-select v-model="chart.formData.specId" class="w-full" filterable><el-option v-for="item in specs" :key="item.ID" :label="specLabel(item)" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.chart.chartType')" prop="chartType"><el-select v-model="chart.formData.chartType" class="w-full"><el-option v-for="item in chartTypes" :key="item.value" :label="item.label" :value="item.value" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.chart.subgroup')"><el-input-number v-model="chart.formData.subgroupSize" :min="1" :max="25" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('spc.chart.limitMethod')"><el-select v-model="chart.formData.limitMethod" class="w-full"><el-option v-for="item in limitMethods" :key="item.value" :label="item.label" :value="item.value" /></el-select></el-form-item>
        <el-form-item v-if="chart.formData.chartType === 'EWMA'" :label="$t('spc.chart.ewmaLambda')"><el-input-number v-model="chart.formData.ewmaLambda" :min="0.01" :max="0.99" :step="0.05" :precision="2" class="!w-full" /></el-form-item>
        <el-form-item v-if="chart.formData.chartType === 'CUSUM'" :label="$t('spc.chart.cusumK')"><el-input-number v-model="chart.formData.cusumK" :precision="3" class="!w-full" /></el-form-item>
        <el-form-item v-if="chart.formData.chartType === 'CUSUM'" :label="$t('spc.chart.cusumH')"><el-input-number v-model="chart.formData.cusumH" :precision="3" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('common.status')"><el-radio-group v-model="chart.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="chart.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="chart.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="chart.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
    <el-drawer v-model="limit.drawerVisible" :title="limit.drawerTitle" size="480px" destroy-on-close>
      <el-form :ref="bindFormRef(limit)" :model="limit.formData" :rules="limitRules" label-width="110px">
        <el-form-item :label="$t('spc.chart.tabChart')" prop="chartId"><el-select v-model="limit.formData.chartId" class="w-full" filterable><el-option v-for="item in charts" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item label="UCL"><el-input-number v-model="limit.formData.ucl" :precision="4" class="!w-full" /></el-form-item>
        <el-form-item label="CL"><el-input-number v-model="limit.formData.cl" :precision="4" class="!w-full" /></el-form-item>
        <el-form-item label="LCL"><el-input-number v-model="limit.formData.lcl" :precision="4" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('spc.chart.source')"><el-select v-model="limit.formData.source" class="w-full"><el-option v-for="item in limitSources" :key="item.value" :label="item.label" :value="item.value" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.chart.calcN')"><el-input-number v-model="limit.formData.calcN" :min="0" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="limit.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="limit.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="limit.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ElMessage } from 'element-plus'
import { chartApi, calculateControlLimit, controlLimitApi, getRuleCatalog, parameterApi, ruleApi, saveChartRules, specApi } from '@/api/spc'
import { CHART_TYPE_VALUES, LIMIT_METHOD_VALUES, LIMIT_SOURCE_VALUES, SEVERITY_VALUES, num, optionOf, requiredRule, selectRule, statusTag, statusText } from './constants'
import { bindFormRef, loadOptions, useSpcCrud } from './composables/useSpcCrud'
import SpcPageHeader from './components/SpcPageHeader.vue'

const { t } = useI18n()
const activeTab = ref('chart')
const charts = ref([])
const params = ref([])
const specs = ref([])
const catalog = ref([])
const ruleRows = ref([])
const ruleChartId = ref()
const savingRules = ref(false)
const recalcing = ref(false)
const chartTypes = computed(() => optionOf(t, 'spc.option.chartType', CHART_TYPE_VALUES))
const limitMethods = computed(() => optionOf(t, 'spc.option.limitMethod', LIMIT_METHOD_VALUES))
const limitSources = computed(() => optionOf(t, 'spc.option.limitSource', LIMIT_SOURCE_VALUES))
const severities = computed(() => optionOf(t, 'spc.option.severity', SEVERITY_VALUES))
const chart = useSpcCrud({ listApi: chartApi.getList, createApi: chartApi.create, updateApi: chartApi.update, deleteApi: chartApi.remove })
const limit = useSpcCrud({ listApi: controlLimitApi.getList, createApi: controlLimitApi.create, updateApi: controlLimitApi.update, deleteApi: controlLimitApi.remove })
const chartRules = computed(() => ({ code: requiredRule(t), name: requiredRule(t), parameterId: selectRule(t), specId: selectRule(t), chartType: selectRule(t) }))
const limitRules = computed(() => ({ chartId: selectRule(t) }))
const specLabel = (item) => `${item.parameter?.code || item.parameterId} / ${item.version || 'v'} USL ${item.usl ?? '-'}`
const ruleFamilies = computed(() => [
  { key: 'WE', label: t('spc.chart.familyWE'), items: ruleRows.value.filter((i) => i.family === 'WE') },
  { key: 'NELSON', label: t('spc.chart.familyNelson'), items: ruleRows.value.filter((i) => i.family === 'NELSON') }
])

const mergeRules = (saved = []) => {
  const byCode = Object.fromEntries(saved.map((item) => [item.ruleCode, item]))
  ruleRows.value = catalog.value.map((def) => {
    const hit = byCode[def.code]
    return {
      ruleCode: def.code,
      family: def.family,
      params: def.params || [],
      nameKey: def.nameKey,
      descKey: def.descKey,
      enabled: hit ? Boolean(hit.enabled) : false,
      n: hit?.n || def.n,
      hits: hit?.hits || def.hits,
      k: hit?.k ?? def.k,
      severity: hit?.severity || def.severity
    }
  })
}

const loadRules = async () => {
  if (!ruleChartId.value) {
    mergeRules([])
    return
  }
  const res = await ruleApi.getList({ page: 1, pageSize: 100, chartId: ruleChartId.value })
  mergeRules(res.code === 0 ? res.data?.list || [] : [])
}

const applyDefaults = () => mergeRules([])
const saveRules = async () => {
  savingRules.value = true
  try {
    const res = await saveChartRules({
      chartId: ruleChartId.value,
      rules: ruleRows.value.map((item) => ({
        chartId: ruleChartId.value,
        ruleCode: item.ruleCode,
        enabled: item.enabled,
        n: item.n,
        hits: item.hits,
        k: item.k,
        severity: item.severity
      }))
    })
    if (res.code === 0) ElMessage.success(t('spc.chart.rulesSaved'))
  } finally {
    savingRules.value = false
  }
}

const gotoRules = (row) => {
  ruleChartId.value = row.ID
  activeTab.value = 'rule'
  loadRules()
}

const recalcLimits = async () => {
  if (!limit.searchInfo.chartId) {
    ElMessage.warning(t('spc.chart.needChart'))
    return
  }
  recalcing.value = true
  try {
    const res = await calculateControlLimit({ chartId: limit.searchInfo.chartId, sampleN: 50 })
    if (res.code === 0) {
      ElMessage.success(t('spc.chart.recalcOk'))
      limit.getTableData()
    }
  } finally {
    recalcing.value = false
  }
}

const onTabChange = (name) => {
  if (name === 'limit') limit.getTableData()
  if (name === 'rule' && ruleChartId.value) loadRules()
}

const openChart = (type, row) => chart.openDrawer(type === 'add' ? t('spc.chart.addChart') : t('spc.chart.editChart'), type === 'add'
  ? { code: '', name: '', parameterId: params.value[0]?.ID, specId: specs.value[0]?.ID, chartType: 'I_MR', subgroupSize: 5, limitMethod: 'CALC', status: 1, remark: '' }
  : row)
const openLimit = (type, row) => limit.openDrawer(type === 'add' ? t('spc.chart.addLimit') : t('spc.chart.editLimit'), type === 'add'
  ? { chartId: limit.searchInfo.chartId || charts.value[0]?.ID, ucl: undefined, cl: undefined, lcl: undefined, source: 'MANUAL', calcN: 0, remark: '' }
  : row)

onMounted(async () => {
  chart.getTableData()
  charts.value = await loadOptions(chartApi.getList)
  params.value = await loadOptions(parameterApi.getList)
  specs.value = await loadOptions(specApi.getList)
  const cat = await getRuleCatalog()
  if (cat.code === 0) catalog.value = cat.data || []
})
</script>
