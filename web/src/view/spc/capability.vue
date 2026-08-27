<template>
  <div>
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>过程能力分析</span>
        </div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="控制图" prop="chartId">
              <el-select
                v-model="formData.chartId"
                placeholder="选择控制图"
                filterable
                style="width: 100%"
              >
                <el-option
                  v-for="chart in chartList"
                  :key="chart.ID"
                  :label="`${chart.code} - ${chart.name}`"
                  :value="chart.ID"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="开始时间" prop="from">
              <el-date-picker
                v-model="formData.from"
                type="datetime"
                placeholder="选择开始时间"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="结束时间" prop="to">
              <el-date-picker
                v-model="formData.to"
                type="datetime"
                placeholder="选择结束时间"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item>
          <el-button type="primary" @click="calculate" :loading="calculating">计算能力指数</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 能力指数结果 -->
    <el-card shadow="hover" style="margin-top: 20px" v-if="capabilityResult">
      <template #header>
        <div class="card-header">
          <span>能力指数结果</span>
          <el-tag>样本数: {{ capabilityResult.n }}</el-tag>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="6">
          <el-statistic title="Cp (短期能力)" :value="capabilityResult.cp || 0" :precision="3">
            <template #suffix>
              <el-tag :type="getCapabilityType(capabilityResult.cp)" size="small">
                {{ getCapabilityLevel(capabilityResult.cp) }}
              </el-tag>
            </template>
          </el-statistic>
        </el-col>
        <el-col :span="6">
          <el-statistic title="Cpk (短期能力指数)" :value="capabilityResult.cpk || 0" :precision="3">
            <template #suffix>
              <el-tag :type="getCapabilityType(capabilityResult.cpk)" size="small">
                {{ getCapabilityLevel(capabilityResult.cpk) }}
              </el-tag>
            </template>
          </el-statistic>
        </el-col>
        <el-col :span="6">
          <el-statistic title="Pp (长期性能)" :value="capabilityResult.pp || 0" :precision="3">
            <template #suffix>
              <el-tag :type="getCapabilityType(capabilityResult.pp)" size="small">
                {{ getCapabilityLevel(capabilityResult.pp) }}
              </el-tag>
            </template>
          </el-statistic>
        </el-col>
        <el-col :span="6">
          <el-statistic title="Ppk (长期性能指数)" :value="capabilityResult.ppk || 0" :precision="3">
            <template #suffix>
              <el-tag :type="getCapabilityType(capabilityResult.ppk)" size="small">
                {{ getCapabilityLevel(capabilityResult.ppk) }}
              </el-tag>
            </template>
          </el-statistic>
        </el-col>
      </el-row>

      <el-divider />

      <el-descriptions :column="3" border>
        <el-descriptions-item label="均值">{{ capabilityResult.mean?.toFixed(4) }}</el-descriptions-item>
        <el-descriptions-item label="标准差 (Within)">{{ capabilityResult.sigma_within?.toFixed(4) }}</el-descriptions-item>
        <el-descriptions-item label="标准差 (Overall)">{{ capabilityResult.sigma_overall?.toFixed(4) }}</el-descriptions-item>
        <el-descriptions-item label="规格上限 USL">{{ capabilityResult.usl?.toFixed(4) }}</el-descriptions-item>
        <el-descriptions-item label="规格下限 LSL">{{ capabilityResult.lsl?.toFixed(4) }}</el-descriptions-item>
        <el-descriptions-item label="目标值">{{ capabilityResult.target?.toFixed(4) || '-' }}</el-descriptions-item>
      </el-descriptions>
    </el-card>

    <!-- 历史记录 -->
    <el-card shadow="hover" style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>历史能力分析记录</span>
          <el-button size="small" @click="loadHistory">刷新</el-button>
        </div>
      </template>
      <el-table :data="historyData" style="width: 100%" v-loading="loadingHistory">
        <el-table-column prop="ID" label="ID" width="80" />
        <el-table-column prop="chartId" label="控制图ID" width="120" />
        <el-table-column prop="cp" label="Cp" width="100">
          <template #default="{ row }">{{ row.cp?.toFixed(3) }}</template>
        </el-table-column>
        <el-table-column prop="cpk" label="Cpk" width="100">
          <template #default="{ row }">{{ row.cpk?.toFixed(3) }}</template>
        </el-table-column>
        <el-table-column prop="pp" label="Pp" width="100">
          <template #default="{ row }">{{ row.pp?.toFixed(3) }}</template>
        </el-table-column>
        <el-table-column prop="ppk" label="Ppk" width="100">
          <template #default="{ row }">{{ row.ppk?.toFixed(3) }}</template>
        </el-table-column>
        <el-table-column prop="n" label="样本数" width="100" />
        <el-table-column prop="windowFrom" label="开始时间" width="160">
          <template #default="{ row }">{{ formatDateTime(row.windowFrom) }}</template>
        </el-table-column>
        <el-table-column prop="windowTo" label="结束时间" width="160">
          <template #default="{ row }">{{ formatDateTime(row.windowTo) }}</template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getChartList } from '@/api/spc/chart'
import { getCapability, getCapabilityHistory } from '@/api/spc/runtime'

const formRef = ref(null)
const formData = ref({
  chartId: null,
  from: null,
  to: null
})

const chartList = ref([])
const capabilityResult = ref(null)
const historyData = ref([])
const calculating = ref(false)
const loadingHistory = ref(false)

const rules = reactive({
  chartId: [{ required: true, message: '请选择控制图', trigger: 'change' }],
  from: [{ required: true, message: '请选择开始时间', trigger: 'change' }],
  to: [{ required: true, message: '请选择结束时间', trigger: 'change' }]
})

const formatDateTime = (dateStr) => {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleString('zh-CN')
}

const loadChartList = async () => {
  try {
    const res = await getChartList({ page: 1, pageSize: 100 })
    if (res.code === 0 && res.data.list) {
      chartList.value = res.data.list
    }
  } catch (error) {
    ElMessage.error('加载图表列表失败')
  }
}

const calculate = async () => {
  formRef.value?.validate(async (valid) => {
    if (valid) {
      calculating.value = true
      try {
        const params = {
          chartId: formData.value.chartId,
          from: formData.value.from.toISOString(),
          to: formData.value.to.toISOString(),
          persist: true
        }

        const res = await getCapability(params)
        if (res.code === 0) {
          capabilityResult.value = res.data
          ElMessage.success('能力分析计算完成')
          loadHistory()
        }
      } catch (error) {
        ElMessage.error(error.message || '能力分析计算失败')
      } finally {
        calculating.value = false
      }
    }
  })
}

const loadHistory = async () => {
  loadingHistory.value = true
  try {
    const res = await getCapabilityHistory({ page: 1, pageSize: 20 })
    if (res.code === 0) {
      historyData.value = res.data.list || []
    }
  } catch (error) {
    ElMessage.error('加载历史记录失败')
  } finally {
    loadingHistory.value = false
  }
}

const resetForm = () => {
  formRef.value?.resetFields()
  capabilityResult.value = null
}

const getCapabilityType = (value) => {
  if (!value) return 'info'
  if (value >= 1.67) return 'success'
  if (value >= 1.33) return 'warning'
  return 'danger'
}

const getCapabilityLevel = (value) => {
  if (!value) return '未知'
  if (value >= 1.67) return '优秀'
  if (value >= 1.33) return '良好'
  if (value >= 1.00) return '一般'
  return '不足'
}

onMounted(() => {
  loadChartList()
  loadHistory()
})
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
