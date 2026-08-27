<template>
  <!-- SPC告警中心 - 列表和处理页面 -->
  <div>
    <div class="gva-search-box">
      <el-form ref="searchFormRef" :inline="true" :model="searchInfo">
        <el-form-item label="告警类型">
          <el-select v-model="searchInfo.alarmType" placeholder="请选择" clearable>
            <el-option label="OOC失控" value="OOC" />
            <el-option label="OOS超规格" value="OOS" />
          </el-select>
        </el-form-item>
        <el-form-item label="告警状态">
          <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
            <el-option label="未处理" value="OPEN" />
            <el-option label="已确认" value="ACK" />
            <el-option label="已关闭" value="CLOSED" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    
    <div class="gva-table-box">
      <el-table :data="tableData" ref="multipleTable" row-key="ID">
        <el-table-column label="ID" prop="ID" min-width="80" />
        <el-table-column label="告警类型" prop="alarmType" min-width="100">
          <template #default="scope">
            <el-tag :type="scope.row.alarmType === 'OOC' ? 'warning' : 'danger'">
              {{ scope.row.alarmType }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="严重度" prop="severity" min-width="100">
          <template #default="scope">
            <el-tag :type="scope.row.severity === 'CRIT' ? 'danger' : (scope.row.severity === 'WARN' ? 'warning' : 'info')">
              {{ scope.row.severity }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="规则代码" prop="ruleCode" min-width="100" />
        <el-table-column label="状态" prop="status" min-width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 'OPEN' ? 'danger' : (scope.row.status === 'ACK' ? 'warning' : 'success')">
              {{ scope.row.status === 'OPEN' ? '未处理' : (scope.row.status === 'ACK' ? '已确认' : '已关闭') }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="备注" prop="remark" min-width="200" />
        <el-table-column label="创建时间" prop="CreatedAt" min-width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" min-width="250">
          <template #default="scope">
            <el-button v-if="scope.row.status === 'OPEN'" type="warning" link icon="check" @click="ackAlarm(scope.row)">确认</el-button>
            <el-button v-if="scope.row.status !== 'CLOSED'" type="success" link icon="close" @click="closeAlarm(scope.row)">关闭</el-button>
            <el-button v-if="scope.row.status !== 'CLOSED'" type="primary" link icon="document" @click="startOcap(scope.row)">启动OCAP</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-dialog v-model="remarkDialogVisible" :title="remarkDialogTitle" width="40%">
      <el-input v-model="remarkText" type="textarea" :rows="4" placeholder="请输入处理备注" />
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="remarkDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmRemark">确定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 启动OCAP对话框 -->
    <el-dialog v-model="ocapDialogVisible" title="启动OCAP" width="50%">
      <el-form :model="ocapForm" label-width="120px">
        <el-form-item label="OCAP模板">
          <el-select v-model="ocapForm.ocapId" placeholder="选择OCAP模板" style="width: 100%">
            <el-option
              v-for="ocap in ocapList"
              :key="ocap.ID"
              :label="`${ocap.code} - ${ocap.name}`"
              :value="ocap.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="执行人">
          <el-input v-model="ocapForm.assignee" placeholder="请输入执行人" />
        </el-form-item>
        <el-form-item label="说明">
          <el-input v-model="ocapForm.remark" type="textarea" :rows="3" placeholder="请输入说明" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="ocapDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmStartOcap">启动</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getAlarmList, acknowledgeAlarm, closeAlarm as closeAlarmApi } from '@/api/spc/collect'
import { getOcapList, startOcap as startOcapApi } from '@/api/spc/ocap'
import { formatDate } from '@/utils/format'

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = reactive({})

const remarkDialogVisible = ref(false)
const remarkDialogTitle = ref('')
const remarkText = ref('')
let currentAlarm = null
let currentAction = ''

const ocapDialogVisible = ref(false)
const ocapList = ref([])
const ocapForm = reactive({
  ocapId: null,
  assignee: '',
  remark: ''
})

const loadOcapList = async () => {
  try {
    const res = await getOcapList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      ocapList.value = res.data.list || []
    }
  } catch (error) {
    ElMessage.error('加载OCAP模板失败')
  }
}

const getTableData = async() => {
  const table = await getAlarmList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

onMounted(() => {
  getTableData()
  loadOcapList()
})

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.alarmType = ''
  searchInfo.status = ''
  onSubmit()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const ackAlarm = (row) => {
  currentAlarm = row
  currentAction = 'ack'
  remarkDialogTitle.value = '确认告警'
  remarkText.value = ''
  remarkDialogVisible.value = true
}

const closeAlarm = (row) => {
  currentAlarm = row
  currentAction = 'close'
  remarkDialogTitle.value = '关闭告警'
  remarkText.value = ''
  remarkDialogVisible.value = true
}

const startOcap = (row) => {
  currentAlarm = row
  ocapForm.ocapId = null
  ocapForm.assignee = ''
  ocapForm.remark = `告警ID: ${row.ID}, 类型: ${row.alarmType}`
  ocapDialogVisible.value = true
}

const confirmRemark = async() => {
  if (!currentAlarm) return
  
  const data = { ID: currentAlarm.ID, remark: remarkText.value }
  let res
  
  if (currentAction === 'ack') {
    res = await acknowledgeAlarm(data)
  } else {
    res = await closeAlarmApi(data)
  }
  
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: currentAction === 'ack' ? '确认成功' : '关闭成功'
    })
    remarkDialogVisible.value = false
    getTableData()
  }
}

const confirmStartOcap = async () => {
  if (!ocapForm.ocapId) {
    ElMessage.error('请选择OCAP模板')
    return
  }
  
  try {
    const data = {
      alarmId: currentAlarm.ID,
      ocapId: ocapForm.ocapId,
      assignee: ocapForm.assignee,
      remark: ocapForm.remark
    }
    
    const res = await startOcapApi(data)
    if (res.code === 0) {
      ElMessage.success('OCAP启动成功')
      ocapDialogVisible.value = false
      getTableData()
    }
  } catch (error) {
    ElMessage.error(error.message || 'OCAP启动失败')
  }
}
</script>

<style scoped>
</style>
