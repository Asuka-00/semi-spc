<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.collect.title')" :subtitle="$t('spc.collect.subtitle')" />
    <div class="grid grid-cols-1 gap-4 xl:grid-cols-12">
      <div class="spc-panel xl:col-span-7">
        <el-form ref="formRef" :model="form" :rules="rules" label-width="120px">
          <el-form-item :label="$t('spc.collect.chartCode')" prop="chartCode">
            <el-select v-model="form.chartCode" class="w-full" filterable @change="onChartChange">
              <el-option v-for="item in charts" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.code" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('spc.collect.lotId')"><el-select v-model="form.lotId" class="w-full" clearable filterable><el-option v-for="item in lots" :key="item.ID" :label="item.lotId" :value="item.lotId" /></el-select></el-form-item>
          <el-form-item :label="$t('spc.collect.waferId')"><el-select v-model="form.waferId" class="w-full" clearable filterable><el-option v-for="item in wafers" :key="item.ID" :label="item.waferId" :value="item.waferId" /></el-select></el-form-item>
          <el-form-item :label="$t('spc.collect.equipment')"><el-select v-model="form.equipmentId" class="w-full" clearable filterable><el-option v-for="item in eqs" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
          <el-form-item :label="$t('spc.collect.chamber')"><el-select v-model="form.chamberId" class="w-full" clearable filterable><el-option v-for="item in chambers" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
          <el-form-item :label="$t('spc.collect.recipe')"><el-select v-model="form.recipeId" class="w-full" clearable filterable><el-option v-for="item in recipes" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
          <el-form-item :label="$t('spc.collect.sampleTime')"><el-date-picker v-model="form.sampleTime" type="datetime" class="!w-full" /></el-form-item>
          <el-form-item :label="$t('spc.collect.subgroupNo')"><el-input-number v-model="form.subgroupNo" :min="1" class="!w-full" /></el-form-item>
          <el-form-item :label="$t('spc.collect.values')" prop="valuesText">
            <el-input v-model="form.valuesText" type="textarea" :rows="3" />
            <div class="mt-1 flex items-center justify-between gap-2">
              <div class="text-xs text-slate-400">{{ $t('spc.collect.valuesHint') }}</div>
              <el-button link type="primary" @click="fillExample">{{ $t('spc.collect.fillExample') }}</el-button>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" :loading="submitting" @click="submit">{{ $t('spc.collect.submit') }}</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="spc-panel xl:col-span-5">
        <div class="mb-3 text-sm font-semibold">{{ $t('spc.collect.result') }}</div>
        <el-empty v-if="!result" :description="$t('spc.collect.noResult')" />
        <div v-else class="space-y-3">
          <el-descriptions :column="1" border>
            <el-descriptions-item :label="$t('spc.collect.sampleId')">{{ result.sampleId }}</el-descriptions-item>
            <el-descriptions-item :label="$t('spc.collect.ooc')"><el-tag :type="result.oocFlag ? 'warning' : 'success'">{{ result.oocFlag ? $t('spc.collect.yes') : $t('spc.collect.no') }}</el-tag></el-descriptions-item>
            <el-descriptions-item :label="$t('spc.collect.oos')"><el-tag :type="result.oosFlag ? 'danger' : 'success'">{{ result.oosFlag ? $t('spc.collect.yes') : $t('spc.collect.no') }}</el-tag></el-descriptions-item>
            <el-descriptions-item :label="$t('spc.collect.alarms')">{{ (result.alarms || []).join(', ') || '-' }}</el-descriptions-item>
          </el-descriptions>
          <div class="text-sm font-medium">{{ $t('spc.collect.violations') }}</div>
          <el-empty v-if="!(result.violations || []).length" :description="$t('spc.collect.ok')" :image-size="60" />
          <el-alert v-for="(item, idx) in result.violations || []" :key="idx" :title="`${item.ruleCode} · ${translateViolation(item)}`" :type="item.severity === 'CRIT' ? 'error' : 'warning'" show-icon class="mb-2" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ElMessage } from 'element-plus'
import { chamberApi, chartApi, collectData, equipmentApi, lotApi, recipeApi, waferApi } from '@/api/spc'
import { translateViolation } from '@/i18n'
import { requiredRule, selectRule } from './constants'
import { loadOptions } from './composables/useSpcCrud'
import SpcPageHeader from './components/SpcPageHeader.vue'

const { t } = useI18n()
const formRef = ref()
const submitting = ref(false)
const result = ref(null)
const charts = ref([])
const lots = ref([])
const wafers = ref([])
const eqs = ref([])
const chambers = ref([])
const recipes = ref([])
const form = reactive({
  chartCode: '',
  lotId: undefined,
  waferId: undefined,
  equipmentId: undefined,
  chamberId: undefined,
  recipeId: undefined,
  sampleTime: new Date(),
  subgroupNo: 1,
  valuesText: ''
})
const rules = computed(() => ({ chartCode: selectRule(t), valuesText: requiredRule(t) }))

const parseValues = () => {
  const parts = String(form.valuesText || '').split(/[,，\s]+/).filter(Boolean)
  const values = parts.map(Number)
  if (!values.length || values.some((n) => Number.isNaN(n))) return null
  return values
}

const currentChart = () => charts.value.find((item) => item.code === form.chartCode)

const fillExample = () => {
  const chart = currentChart()
  const n = Math.max(1, Number(chart?.subgroupSize) || 5)
  const center = chart?.chartType === 'I_MR' ? 1500 : 45
  const step = chart?.chartType === 'I_MR' ? 2 : 0.15
  form.valuesText = Array.from({ length: n }, (_, i) => (center + ((i % 3) - 1) * step).toFixed(2)).join(', ')
}

const onChartChange = () => fillExample()

const submit = async () => {
  await formRef.value?.validate()
  const values = parseValues()
  if (!values) {
    ElMessage.warning(t('spc.collect.badValues'))
    return
  }
  submitting.value = true
  try {
    const res = await collectData({
      chartCode: form.chartCode,
      lotId: form.lotId || undefined,
      waferId: form.waferId || undefined,
      equipmentId: form.equipmentId || undefined,
      chamberId: form.chamberId || undefined,
      recipeId: form.recipeId || undefined,
      sampleTime: form.sampleTime,
      subgroupNo: form.subgroupNo,
      values
    })
    if (res.code === 0) {
      result.value = res.data
      ElMessage.success(res.data?.oocFlag || res.data?.oosFlag ? t('spc.collect.abnormal') : t('spc.collect.ok'))
    }
  } finally {
    submitting.value = false
  }
}

onMounted(async () => {
  charts.value = await loadOptions(chartApi.getList)
  lots.value = await loadOptions(lotApi.getList)
  wafers.value = await loadOptions(waferApi.getList)
  eqs.value = await loadOptions(equipmentApi.getList)
  chambers.value = await loadOptions(chamberApi.getList)
  recipes.value = await loadOptions(recipeApi.getList)
  if (charts.value[0]) {
    form.chartCode = charts.value[0].code
    fillExample()
  }
})
</script>
