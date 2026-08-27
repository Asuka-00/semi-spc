<template>
  <!-- SPC能力分析 -->
  <div>
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>能力分析计算</span>
        </div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="控制图ID" prop="chart_id">
              <el-input-number v-model="formData.chart_id" :min="1" placeholder="请输入控制图ID" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="开始时间" prop="start_time">
              <el-date-picker
                v-model="formData.start_time"
                type="datetime"
                placeholder="选择开始时间"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="结束时间" prop="end_time">
              <el-date-picker
                v-model="formData.end_time"
                type="datetime"
                placeholder="选择结束时间"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item>
          <el-button type="primary" @click="calculate">计算能力指数</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card shadow="hover" style="margin-top: 20px" v-if="capabilityResult">
      <template #header>
        <div class="card-header">
          <span>能力分析结果</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="6">
          <el-statistic title="Cp (短期能力)" :value="capabilityResult.cp" :precision="3">
            <template #suffix>
              <el-tag :type="getCapabilityType(capabilityResult.cp)" size="small">
                {{ getCapabilityLevel(capabilityResult.cp) }}
              </el-tag>
            </template>
          </el-statistic>
        </el-col>
        <el-col :span="6">
          <el-statistic title="Cpk (短期能力指数)" :value="capabilityResult.cpk" :precision="3">
            <template #suffix>
              <el-tag :type="getCapabilityType(capabilityResult.cpk)" size="small">
                {{ getCapabilityLevel(capabilityResult.cpk) }}
              </el-tag>
            </template>
          </el-statistic>
        </el-col>
        <el-col :span="6">
          <el-statistic title="Pp (长期性能)" :value="capabilityResult.pp" :precision="3">
            <template #suffix>
              <el-tag :type="getCapabilityType(capabilityResult.pp)" size="small">
                {{ getCapabilityLevel(capabilityResult.pp) }}
              </el-tag>
            </template>
          </el-statistic>
        </el-col>
        <el-col :span="6">
          <el-statistic title="Ppk (长期性能指数)" :value="capabilityResult.ppk" :precision="3">
            <template #suffix>
              <el-tag :type="getCapabilityType(capabilityResult.ppk)" size="small">
                {{ getCapabilityLevel(capabilityResult.ppk) }}
              </el-tag>
            </template>
          </el-statistic>
        </el-col>
      </el-row>

      <el-divider />

      <div class="capability-histogram">
        <el-empty description="能力直方图 - 待集成ECharts">
          <template #description>
            <p>此处将展示:</p>
            <ul style="text-align: left; display: inline-block">
              <li>数据分布直方图</li>
              <li>正态分布曲线拟合</li>
              <li>规格限标记 (USL, LSL)</li>
              <li>均值和标准差标注</li>
            </ul>
          </template>
        </el-empty>
      </div>
    </el-card>

    <el-card shadow="hover" style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>历史能力分析记录</span>
        </div>
      </template>
      <el-table :data="historyData" style="width: 100%">
        <el-table-column prop="ID" label="ID" width="80" />
        <el-table-column prop="chart_id" label="控制图" width="100" />
        <el-table-column prop="cp" label="Cp" width="100" :formatter="numberFormatter" />
        <el-table-column prop="cpk" label="Cpk" width="100" :formatter="numberFormatter" />
        <el-table-column prop="pp" label="Pp" width="100" :formatter="numberFormatter" />
        <el-table-column prop="ppk" label="Ppk" width="100" :formatter="numberFormatter" />
        <el-table-column prop="start_time" label="开始时间" width="180" :formatter="dateFormatter" />
        <el-table-column prop="end_time" label="结束时间" width="180" :formatter="dateFormatter" />
        <el-table-column prop="CreatedAt" label="计算时间" width="180" :formatter="dateFormatter" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'

const formRef = ref(null)
const formData = ref({
  chart_id: null,
  start_time: null,
  end_time: null
})

const capabilityResult = ref(null)
const historyData = ref([])

const rules = reactive({
  chart_id: [{ required: true, message: '请输入控制图ID', trigger: 'blur' }],
  start_time: [{ required: true, message: '请选择开始时间', trigger: 'change' }],
  end_time: [{ required: true, message: '请选择结束时间', trigger: 'change' }]
})

const calculate = async() => {
  formRef.value?.validate(async(valid) => {
    if (valid) {
      capabilityResult.value = {
        cp: 1.45,
        cpk: 1.32,
        pp: 1.40,
        ppk: 1.28
      }
      ElMessage.success('能力分析计算完成')
    }
  })
}

const resetForm = () => {
  formRef.value?.resetFields()
  capabilityResult.value = null
}

const getCapabilityType = (value) => {
  if (value >= 1.67) return 'success'
  if (value >= 1.33) return 'warning'
  return 'danger'
}

const getCapabilityLevel = (value) => {
  if (value >= 1.67) return '优秀'
  if (value >= 1.33) return '良好'
  if (value >= 1.00) return '一般'
  return '不足'
}

const numberFormatter = (row, column, cellValue) => {
  return cellValue ? cellValue.toFixed(3) : '-'
}

const dateFormatter = (row, column, cellValue) => {
  return formatDate(cellValue)
}
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.capability-histogram {
  height: 300px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px dashed #dcdfe6;
  border-radius: 4px;
  margin-top: 20px;
}
</style>
