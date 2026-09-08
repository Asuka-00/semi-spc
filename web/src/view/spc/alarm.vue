<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.alarm.title')" :subtitle="$t('spc.alarm.subtitle')" />
    <div class="spc-panel">
      <div class="gva-search-box !bg-transparent !p-0 !my-3">
        <el-form :inline="true" :model="search">
          <el-form-item :label="$t('spc.alarm.type')">
            <el-select v-model="search.alarmType" clearable>
              <el-option v-for="item in alarmTypes" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('spc.alarm.status')">
            <el-select v-model="search.status" clearable>
              <el-option :label="$t('spc.alarm.open')" value="OPEN" />
              <el-option :label="$t('spc.alarm.acked')" value="ACK" />
              <el-option :label="$t('spc.alarm.closed')" value="CLOSED" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="search" @click="loadList">{{ $t('common.query') }}</el-button>
            <el-button icon="refresh" @click="reset">{{ $t('common.reset') }}</el-button>
          </el-form-item>
        </el-form>
      </div>
      <el-table v-loading="loading" :data="tableData" row-key="ID" :empty-text="$t('spc.alarm.empty')">
        <el-table-column :label="$t('spc.alarm.type')" width="110">
          <template #default="{ row }"><el-tag :type="alarmTypeTag(row.alarmType)" round>{{ $t(`spc.option.alarmType.${row.alarmType}`) }}</el-tag></template>
        </el-table-column>
        <el-table-column :label="$t('spc.alarm.severity')" width="100">
          <template #default="{ row }"><el-tag :type="severityTag(row.severity)" round>{{ $t(`spc.option.severity.${row.severity}`) }}</el-tag></template>
        </el-table-column>
        <el-table-column :label="$t('spc.alarm.chart')" min-width="140"><template #default="{ row }">{{ row.chart?.name || row.chartId }}</template></el-table-column>
        <el-table-column :label="$t('spc.alarm.rule')" prop="ruleCode" width="110" />
        <el-table-column :label="$t('spc.alarm.status')" width="110">
          <template #default="{ row }"><el-tag :type="alarmStatusTag(row.status)" round>{{ alarmStatusText(row.status, t) }}</el-tag></template>
        </el-table-column>
        <el-table-column :label="$t('spc.alarm.time')" min-width="160"><template #default="{ row }">{{ formatDate(row.CreatedAt) }}</template></el-table-column>
        <el-table-column :label="$t('common.action')" width="180" fixed="right">
          <template #default="{ row }">
            <el-button v-if="row.status === 'OPEN'" type="primary" link @click="act(row, 'ack')">{{ $t('spc.alarm.ack') }}</el-button>
            <el-button v-if="row.status !== 'CLOSED'" type="primary" link @click="act(row, 'close')">{{ $t('spc.alarm.close') }}</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination mt-4">
        <el-pagination :current-page="page" :page-size="pageSize" :total="total" layout="total, prev, pager, next" @current-change="(v) => { page = v; loadList() }" @size-change="(v) => { pageSize = v; loadList() }" />
      </div>
    </div>
    <el-dialog v-model="dialog.visible" :title="dialog.title" width="420px">
      <el-input v-model="dialog.remark" type="textarea" :rows="4" :placeholder="$t('spc.alarm.remark')" />
      <template #footer>
        <el-button @click="dialog.visible = false">{{ $t('common.cancel') }}</el-button>
        <el-button type="primary" @click="confirmAct">{{ $t('common.confirm') }}</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ElMessage } from 'element-plus'
import { acknowledgeAlarm, closeAlarm, getAlarmList } from '@/api/spc'
import { formatDate } from '@/utils/format'
import { ALARM_TYPE_VALUES, alarmStatusTag, alarmStatusText, alarmTypeTag, optionOf, severityTag } from './constants'
import SpcPageHeader from './components/SpcPageHeader.vue'

const { t } = useI18n()
const loading = ref(false)
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const search = reactive({ alarmType: undefined, status: undefined })
const alarmTypes = computed(() => optionOf(t, 'spc.option.alarmType', ALARM_TYPE_VALUES))
const dialog = reactive({ visible: false, title: '', remark: '', row: null, mode: 'ack' })

const loadList = async () => {
  loading.value = true
  try {
    const res = await getAlarmList({ page: page.value, pageSize: pageSize.value, ...search })
    if (res.code === 0) {
      tableData.value = res.data?.list || []
      total.value = res.data?.total || 0
    }
  } finally {
    loading.value = false
  }
}
const reset = () => {
  search.alarmType = undefined
  search.status = undefined
  page.value = 1
  loadList()
}
const act = (row, mode) => {
  dialog.row = row
  dialog.mode = mode
  dialog.remark = ''
  dialog.title = mode === 'ack' ? t('spc.alarm.ackTitle') : t('spc.alarm.closeTitle')
  dialog.visible = true
}
const confirmAct = async () => {
  const api = dialog.mode === 'ack' ? acknowledgeAlarm : closeAlarm
  const res = await api({ ID: dialog.row.ID, id: dialog.row.ID, remark: dialog.remark })
  if (res.code === 0) {
    ElMessage.success(dialog.mode === 'ack' ? t('spc.alarm.ackOk') : t('spc.alarm.closeOk'))
    dialog.visible = false
    loadList()
  }
}
onMounted(loadList)
</script>
