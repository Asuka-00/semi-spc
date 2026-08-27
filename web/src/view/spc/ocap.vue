<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-search-box">
        <el-form :inline="true">
          <el-form-item label="状态">
            <el-select v-model="searchInfo.status" placeholder="选择状态" clearable @change="getList">
              <el-option label="未执行" value="OPEN" />
              <el-option label="执行中" value="IN_PROGRESS" />
              <el-option label="已关闭" value="CLOSED" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="getList">查询</el-button>
          </el-form-item>
        </el-form>
      </div>
      <el-table :data="tableData" style="width: 100%" v-loading="loading">
        <el-table-column prop="alarmId" label="告警ID" width="100" />
        <el-table-column prop="ocapName" label="OCAP名称" width="200" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="statusType(row.status)">
              {{ statusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="assignee" label="责任人" width="120" />
        <el-table-column prop="startedAt" label="开始时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.startedAt) }}
          </template>
        </el-table-column>
        <el-table-column prop="completedAt" label="完成时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.completedAt) }}
          </template>
        </el-table-column>
        <el-table-column prop="actionsTaken" label="处理措施" show-overflow-tooltip />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button v-if="row.status === 'OPEN'" type="primary" link @click="handleStart(row)">开始执行</el-button>
            <el-button v-if="row.status === 'IN_PROGRESS'" type="success" link @click="handleComplete(row)">完成</el-button>
            <el-button type="info" link @click="handleView(row)">详情</el-button>
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

    <el-dialog v-model="dialogVisible" title="OCAP执行详情" width="60%">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="告警ID">{{ currentRow.alarmId }}</el-descriptions-item>
        <el-descriptions-item label="OCAP名称">{{ currentRow.ocapName }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ statusText(currentRow.status) }}</el-descriptions-item>
        <el-descriptions-item label="责任人">{{ currentRow.assignee }}</el-descriptions-item>
        <el-descriptions-item label="开始时间">{{ formatDate(currentRow.startedAt) }}</el-descriptions-item>
        <el-descriptions-item label="完成时间">{{ formatDate(currentRow.completedAt) }}</el-descriptions-item>
        <el-descriptions-item label="处理措施" :span="2">
          <el-input v-model="currentRow.actionsTaken" type="textarea" :rows="4" v-if="currentRow.status !== 'CLOSED'" />
          <span v-else>{{ currentRow.actionsTaken }}</span>
        </el-descriptions-item>
        <el-descriptions-item label="根本原因" :span="2">
          <el-input v-model="currentRow.rootCause" type="textarea" :rows="4" v-if="currentRow.status !== 'CLOSED'" />
          <span v-else>{{ currentRow.rootCause }}</span>
        </el-descriptions-item>
      </el-descriptions>
      <template #footer>
        <el-button @click="dialogVisible = false">关闭</el-button>
        <el-button v-if="currentRow.status !== 'CLOSED'" type="primary" @click="handleUpdate">更新</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getOcapExecutionList,
  updateOcapExecution
} from '@/api/spc/ocap'

const loading = ref(false)
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const searchInfo = ref({
  status: ''
})

const dialogVisible = ref(false)
const currentRow = ref({})

const statusType = (status) => {
  const types = {
    'OPEN': 'info',
    'IN_PROGRESS': 'warning',
    'CLOSED': 'success'
  }
  return types[status] || 'info'
}

const statusText = (status) => {
  const texts = {
    'OPEN': '未执行',
    'IN_PROGRESS': '执行中',
    'CLOSED': '已关闭'
  }
  return texts[status] || status
}

const formatDate = (date) => {
  if (!date) return '-'
  return new Date(date).toLocaleString('zh-CN')
}

const getList = async () => {
  loading.value = true
  try {
    const params = {
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    }
    const res = await getOcapExecutionList(params)
    if (res.code === 0) {
      tableData.value = res.data.list || []
      total.value = res.data.total || 0
    }
  } catch (error) {
    ElMessage.error('获取列表失败')
  } finally {
    loading.value = false
  }
}

const handleStart = async (row) => {
  try {
    await updateOcapExecution({ ID: row.ID, status: 'IN_PROGRESS', startedAt: new Date() })
    ElMessage.success('OCAP执行已开始')
    await getList()
  } catch (error) {
    ElMessage.error('操作失败')
  }
}

const handleComplete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要完成该OCAP执行吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await updateOcapExecution({ ID: row.ID, status: 'CLOSED', completedAt: new Date() })
    ElMessage.success('OCAP执行已完成')
    await getList()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('操作失败')
    }
  }
}

const handleView = (row) => {
  currentRow.value = { ...row }
  dialogVisible.value = true
}

const handleUpdate = async () => {
  try {
    await updateOcapExecution(currentRow.value)
    ElMessage.success('更新成功')
    dialogVisible.value = false
    await getList()
  } catch (error) {
    ElMessage.error('更新失败')
  }
}

const handleCurrentChange = (val) => {
  page.value = val
  getList()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getList()
}

onMounted(() => {
  getList()
})
</script>

<style scoped>
.gva-table-box {
  padding: 20px;
}
.gva-search-box {
  margin-bottom: 15px;
}
.gva-pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
