<template>
  <div>
    <!-- 手动采集卡片 -->
    <el-card class="box-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <span>手动数据采集</span>
        </div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="控制图" prop="chartCode">
              <el-select
                v-model="formData.chartCode"
                placeholder="选择控制图"
                filterable
                style="width: 100%"
                @change="onChartChange"
              >
                <el-option
                  v-for="chart in chartList"
                  :key="chart.ID"
                  :label="`${chart.code} - ${chart.name}`"
                  :value="chart.code"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="批次" prop="lotId">
              <el-select v-model="formData.lotId" placeholder="选择批次" filterable style="width: 100%">
                <el-option
                  v-for="lot in lotList"
                  :key="lot.ID"
                  :label="lot.code"
                  :value="lot.code"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="设备" prop="equipmentId">
              <el-select v-model="formData.equipmentId" placeholder="选择设备" filterable style="width: 100%">
                <el-option
                  v-for="eqp in equipmentList"
                  :key="eqp.ID"
                  :label="`${eqp.code} - ${eqp.name}`"
                  :value="eqp.ID"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="采样时间" prop="sampleTime">
              <el-date-picker
                v-model="formData.sampleTime"
                type="datetime"
                placeholder="选择日期时间"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="测量值" prop="values">
          <el-input
            v-model="measurementsText"
            type="textarea"
            :rows="3"
            :placeholder="`请输入测量值，多个值用逗号分隔（需${selectedChart?.subgroupSize || 'N'}个值）`"
          />
          <div class="help-text">
            <span v-if="selectedChart">当前图表类型：{{ selectedChart.chartType }}，子组大小：{{ selectedChart.subgroupSize }}</span>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitData" :loading="submitting">提交数据</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- CSV批量上传卡片 -->
    <el-card class="box-card" shadow="hover" style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>CSV批量采集</span>
          <el-button type="success" size="small" @click="downloadTemplate">下载CSV模板</el-button>
        </div>
      </template>
      <el-upload
        class="upload-demo"
        drag
        action="#"
        :auto-upload="false"
        :on-change="handleFileChange"
        :file-list="fileList"
        accept=".csv"
      >
        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
        <div class="el-upload__text">将CSV文件拖到此处，或<em>点击上传</em></div>
        <template #tip>
          <div class="el-upload__tip">
            仅支持.csv格式文件。格式：chart_id,lot_id,wafer_id,equipment_id,sample_time,measurements
          </div>
        </template>
      </el-upload>
      <div style="margin-top: 15px">
        <el-button type="primary" @click="uploadCsv" :loading="uploading" :disabled="!selectedFile">
          开始上传
        </el-button>
      </div>

      <!-- CSV上传结果 -->
      <div v-if="csvResult" style="margin-top: 20px">
        <el-alert
          :title="`上传完成：成功 ${csvResult.success} 条，失败 ${csvResult.fail} 条`"
          :type="csvResult.fail > 0 ? 'warning' : 'success'"
          :closable="false"
          show-icon
        />
        <el-table
          v-if="csvResult.errors && csvResult.errors.length > 0"
          :data="csvResult.errors"
          style="width: 100%; margin-top: 15px"
          max-height="300"
        >
          <el-table-column prop="row" label="行号" width="100" />
          <el-table-column prop="msg" label="错误信息" />
        </el-table>
      </div>
    </el-card>

    <!-- 采集结果卡片 -->
    <el-card class="box-card" shadow="hover" style="margin-top: 20px" v-if="collectResult">
      <template #header>
        <div class="card-header">
          <span>采集结果</span>
        </div>
      </template>
      <el-descriptions :column="2" border>
        <el-descriptions-item label="样本ID">{{ collectResult.sampleId }}</el-descriptions-item>
        <el-descriptions-item label="OOC">
          <el-tag :type="collectResult.oocFlag ? 'warning' : 'success'">
            {{ collectResult.oocFlag ? '失控' : '受控' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="OOS">
          <el-tag :type="collectResult.oosFlag ? 'danger' : 'success'">
            {{ collectResult.oosFlag ? '超规格' : '正常' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="违规规则" :span="2">
          <div v-if="collectResult.violations && collectResult.violations.length > 0">
            <el-tag
              v-for="(v, idx) in collectResult.violations"
              :key="idx"
              type="warning"
              style="margin: 2px"
            >
              {{ v.ruleCode }}: {{ v.message }}
            </el-tag>
          </div>
          <span v-else>无</span>
        </el-descriptions-item>
        <el-descriptions-item label="告警" :span="2">
          <div v-if="collectResult.alarms && collectResult.alarms.length > 0">
            <el-tag v-for="alarmId in collectResult.alarms" :key="alarmId" type="danger" style="margin: 2px">
              告警ID: {{ alarmId }}
            </el-tag>
          </div>
          <span v-else>无</span>
        </el-descriptions-item>
      </el-descriptions>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { UploadFilled } from '@element-plus/icons-vue'
import { getChartList } from '@/api/spc/chart'
import { getLotList } from '@/api/spc/material'
import { getEquipmentList } from '@/api/spc/master'
import { collectData, collectCsv } from '@/api/spc/collect'

const formRef = ref(null)
const formData = ref({
  chartCode: '',
  lotId: '',
  equipmentId: null,
  sampleTime: new Date(),
  values: []
})

const measurementsText = ref('')
const collectResult = ref(null)
const submitting = ref(false)

const chartList = ref([])
const lotList = ref([])
const equipmentList = ref([])
const selectedChart = ref(null)

const fileList = ref([])
const selectedFile = ref(null)
const uploading = ref(false)
const csvResult = ref(null)

const rules = reactive({
  chartCode: [{ required: true, message: '请选择控制图', trigger: 'change' }],
  lotId: [{ required: true, message: '请选择批次', trigger: 'change' }],
  equipmentId: [{ required: true, message: '请选择设备', trigger: 'change' }],
  sampleTime: [{ required: true, message: '请选择采样时间', trigger: 'change' }]
})

// 加载下拉列表数据
const loadData = async () => {
  try {
    const [chartsRes, lotsRes, eqpRes] = await Promise.all([
      getChartList({ page: 1, pageSize: 100 }),
      getLotList({ page: 1, pageSize: 100 }),
      getEquipmentList({ page: 1, pageSize: 100 })
    ])

    if (chartsRes.code === 0) chartList.value = chartsRes.data.list || []
    if (lotsRes.code === 0) lotList.value = lotsRes.data.list || []
    if (eqpRes.code === 0) equipmentList.value = eqpRes.data.list || []
  } catch (error) {
    ElMessage.error('加载数据失败')
  }
}

const onChartChange = () => {
  selectedChart.value = chartList.value.find(c => c.code === formData.value.chartCode)
}

const submitData = async () => {
  formRef.value?.validate(async (valid) => {
    if (valid) {
      if (!measurementsText.value) {
        ElMessage.error('请输入测量值')
        return
      }

      const values = measurementsText.value
        .split(',')
        .map(v => parseFloat(v.trim()))
        .filter(v => !isNaN(v))

      if (values.length === 0) {
        ElMessage.error('测量值格式不正确')
        return
      }

      submitting.value = true
      try {
        const data = {
          chartCode: formData.value.chartCode,
          lotId: formData.value.lotId,
          equipmentId: formData.value.equipmentId,
          sampleTime: formData.value.sampleTime.toISOString(),
          values
        }

        const res = await collectData(data)
        if (res.code === 0) {
          collectResult.value = res.data
          ElMessage.success(res.data.message || '数据采集成功')
          if (res.data.alarms && res.data.alarms.length > 0) {
            ElMessage.warning('检测到异常，已创建告警')
          }
        }
      } catch (error) {
        ElMessage.error(error.message || '数据采集失败')
      } finally {
        submitting.value = false
      }
    }
  })
}

const resetForm = () => {
  formRef.value?.resetFields()
  measurementsText.value = ''
  collectResult.value = null
}

// CSV模板下载
const downloadTemplate = () => {
  const template = `chart_id,lot_id,wafer_id,equipment_id,sample_time,measurements
1,LOT001,W001,1,2024-08-27T10:00:00Z,"100.1,100.2,100.3,100.4,100.5"
2,LOT001,W002,1,2024-08-27T11:00:00Z,"99.8,100.0,100.1"`

  const blob = new Blob([template], { type: 'text/csv;charset=utf-8;' })
  const link = document.createElement('a')
  link.href = URL.createObjectURL(blob)
  link.download = 'spc_collect_template.csv'
  link.click()
  URL.revokeObjectURL(link.href)
}

const handleFileChange = (file) => {
  selectedFile.value = file.raw
  fileList.value = [file]
}

const uploadCsv = async () => {
  if (!selectedFile.value) {
    ElMessage.error('请选择文件')
    return
  }

  uploading.value = true
  csvResult.value = null

  try {
    const formData = new FormData()
    formData.append('file', selectedFile.value)

    const res = await collectCsv(formData)
    if (res.code === 0) {
      csvResult.value = res.data
      if (res.data.fail === 0) {
        ElMessage.success(`批量上传成功！共上传 ${res.data.success} 条数据`)
      } else {
        ElMessage.warning(`上传完成：成功 ${res.data.success} 条，失败 ${res.data.fail} 条`)
      }
    }
  } catch (error) {
    ElMessage.error(error.message || 'CSV上传失败')
  } finally {
    uploading.value = false
  }
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.help-text {
  font-size: 12px;
  color: #909399;
  margin-top: 5px;
}

.upload-demo {
  width: 100%;
}
</style>
