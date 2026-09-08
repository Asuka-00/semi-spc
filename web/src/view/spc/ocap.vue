<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.ocap.title')" :subtitle="$t('spc.ocap.subtitle')" />
    <el-tabs v-model="activeTab" class="spc-panel !px-4 !pt-2">
      <el-tab-pane :label="$t('spc.ocap.plan')" name="plan">
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openPlan('add')">{{ $t('spc.ocap.addPlan') }}</el-button></div>
        <el-table v-loading="plan.loading" :data="plan.tableData" row-key="ID" :empty-text="$t('common.noData')">
          <el-table-column :label="$t('spc.ocap.planName')" prop="name" min-width="160" />
          <el-table-column :label="$t('spc.chart.tabChart')" min-width="140"><template #default="{ row }">{{ row.chart?.name || row.chartId }}</template></el-table-column>
          <el-table-column :label="$t('spc.ocap.trigger')" width="120"><template #default="{ row }">{{ $t(`spc.option.trigger.${row.triggerType}`) }}</template></el-table-column>
          <el-table-column :label="$t('common.status')" width="90"><template #default="{ row }"><el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag></template></el-table-column>
          <el-table-column :label="$t('common.action')" width="160"><template #default="{ row }"><el-button type="primary" link @click="openPlan('edit', row)">{{ $t('common.edit') }}</el-button><el-button type="danger" link @click="plan.deleteRow(row, $t('spc.ocap.plan'))">{{ $t('common.delete') }}</el-button></template></el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="plan.page" :page-size="plan.pageSize" :total="plan.total" layout="total, prev, pager, next" @current-change="plan.handleCurrentChange" @size-change="plan.handleSizeChange" /></div>
      </el-tab-pane>
      <el-tab-pane :label="$t('spc.ocap.exec')" name="exec">
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openExec('add')">{{ $t('spc.ocap.addExec') }}</el-button></div>
        <el-table v-loading="exec.loading" :data="exec.tableData" row-key="ID" :empty-text="$t('common.noData')">
          <el-table-column :label="$t('spc.ocap.plan')" min-width="140"><template #default="{ row }">{{ row.ocap?.name || row.ocapId }}</template></el-table-column>
          <el-table-column :label="$t('spc.ocap.alarm')" prop="alarmId" width="90" />
          <el-table-column :label="$t('spc.ocap.owner')" prop="owner" width="120" />
          <el-table-column :label="$t('common.status')" width="120"><template #default="{ row }">{{ $t(`spc.option.ocapStatus.${row.status}`) }}</template></el-table-column>
          <el-table-column :label="$t('spc.ocap.startedAt')" min-width="160"><template #default="{ row }">{{ formatDate(row.startedAt) }}</template></el-table-column>
          <el-table-column :label="$t('common.action')" width="200">
            <template #default="{ row }">
              <el-button v-if="row.status === 'PENDING'" type="primary" link @click="startExec(row)">{{ $t('spc.ocap.start') }}</el-button>
              <el-button v-if="row.status === 'IN_PROGRESS'" type="primary" link @click="completeExec(row)">{{ $t('spc.ocap.complete') }}</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="exec.page" :page-size="exec.pageSize" :total="exec.total" layout="total, prev, pager, next" @current-change="exec.handleCurrentChange" @size-change="exec.handleSizeChange" /></div>
      </el-tab-pane>
    </el-tabs>

    <el-drawer v-model="plan.drawerVisible" :title="plan.drawerTitle" size="560px" destroy-on-close>
      <el-form :ref="bindFormRef(plan)" :model="plan.formData" :rules="planRules" label-width="120px">
        <el-form-item :label="$t('spc.chart.tabChart')" prop="chartId"><el-select v-model="plan.formData.chartId" class="w-full" filterable><el-option v-for="item in charts" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.ocap.planName')" prop="name"><el-input v-model="plan.formData.name" /></el-form-item>
        <el-form-item :label="$t('spc.ocap.trigger')"><el-select v-model="plan.formData.triggerType" class="w-full"><el-option v-for="item in triggers" :key="item.value" :label="item.label" :value="item.value" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.ocap.steps')">
          <div class="w-full space-y-2">
            <div v-for="(step, idx) in planSteps" :key="idx" class="flex items-center gap-2">
              <span class="w-6 shrink-0 text-xs text-slate-400">{{ idx + 1 }}</span>
              <el-input v-model="step.action" :placeholder="$t('spc.ocap.stepAction')" />
              <el-button link type="danger" @click="removePlanStep(idx)">{{ $t('spc.ocap.removeAction') }}</el-button>
            </div>
            <el-button size="small" @click="addPlanStep">{{ $t('spc.ocap.addAction') }}</el-button>
          </div>
        </el-form-item>
        <el-form-item :label="$t('common.status')"><el-radio-group v-model="plan.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="plan.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="plan.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="savePlan">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
    <el-drawer v-model="exec.drawerVisible" :title="exec.drawerTitle" size="480px" destroy-on-close>
      <el-form :ref="bindFormRef(exec)" :model="exec.formData" :rules="execRules" label-width="110px">
        <el-form-item :label="$t('spc.ocap.plan')" prop="ocapId"><el-select v-model="exec.formData.ocapId" class="w-full" filterable><el-option v-for="item in plans" :key="item.ID" :label="item.name" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.ocap.alarm')" prop="alarmId"><el-input-number v-model="exec.formData.alarmId" :min="1" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('spc.ocap.owner')"><el-input v-model="exec.formData.owner" /></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="exec.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="exec.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="exec.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ElMessage, ElMessageBox } from 'element-plus'
import { chartApi, ocapApi, ocapExecutionApi } from '@/api/spc'
import { formatDate } from '@/utils/format'
import { TRIGGER_VALUES, optionOf, requiredRule, selectRule, statusTag, statusText } from './constants'
import { bindFormRef, loadOptions, useSpcCrud } from './composables/useSpcCrud'
import SpcPageHeader from './components/SpcPageHeader.vue'

const { t } = useI18n()
const activeTab = ref('plan')
const charts = ref([])
const plans = ref([])
const planSteps = ref([{ action: '' }])
const triggers = computed(() => optionOf(t, 'spc.option.trigger', TRIGGER_VALUES))
const plan = useSpcCrud({ listApi: ocapApi.getList, createApi: ocapApi.create, updateApi: ocapApi.update, deleteApi: ocapApi.remove })
const exec = useSpcCrud({ listApi: ocapExecutionApi.getList, createApi: ocapExecutionApi.create, updateApi: ocapExecutionApi.update, deleteApi: ocapExecutionApi.remove })
const planRules = computed(() => ({ chartId: selectRule(t), name: requiredRule(t) }))
const execRules = computed(() => ({ ocapId: selectRule(t), alarmId: requiredRule(t, 'change') }))

const parseSteps = (raw) => {
  try {
    const parsed = typeof raw === 'string' ? JSON.parse(raw || '[]') : raw
    if (Array.isArray(parsed) && parsed.length) {
      return parsed.map((item) => ({ action: item.action || item || '' }))
    }
  } catch (_) { /* ignore invalid json */ }
  return [{ action: '' }]
}

const openPlan = (type, row) => {
  plan.openDrawer(type === 'add' ? t('spc.ocap.addPlan') : t('spc.ocap.editPlan'), type === 'add' ? { chartId: charts.value[0]?.ID, name: '', triggerType: 'OOC', stepsJson: '[]', status: 1, remark: '' } : row)
  planSteps.value = parseSteps(plan.formData.stepsJson)
}
const addPlanStep = () => planSteps.value.push({ action: '' })
const removePlanStep = (idx) => {
  planSteps.value.splice(idx, 1)
  if (!planSteps.value.length) planSteps.value = [{ action: '' }]
}
const savePlan = () => {
  plan.formData.stepsJson = JSON.stringify(planSteps.value.map((item, idx) => ({ step: idx + 1, action: item.action })))
  plan.enterDrawer()
}
const openExec = (type, row) => exec.openDrawer(t('spc.ocap.addExec'), type === 'add' ? { ocapId: plans.value[0]?.ID, alarmId: undefined, owner: '', status: 'PENDING', remark: '' } : row)

const startExec = async (row) => {
  const res = await ocapExecutionApi.start({ ID: row.ID, id: row.ID, owner: row.owner })
  if (res.code === 0) {
    ElMessage.success(t('common.saved'))
    exec.getTableData()
  }
}
const completeExec = async (row) => {
  const { value } = await ElMessageBox.prompt(t('spc.ocap.comment'), t('spc.ocap.complete'), { confirmButtonText: t('common.confirm'), cancelButtonText: t('common.cancel') })
  const res = await ocapExecutionApi.complete({ ID: row.ID, id: row.ID, comment: value })
  if (res.code === 0) {
    ElMessage.success(t('common.saved'))
    exec.getTableData()
  }
}

onMounted(async () => {
  plan.getTableData(); exec.getTableData()
  charts.value = await loadOptions(chartApi.getList)
  plans.value = await loadOptions(ocapApi.getList)
})
</script>
