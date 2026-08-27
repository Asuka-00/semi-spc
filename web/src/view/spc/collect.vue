<template>
  <!-- SPC数据采集页面 -->
  <div>
    <el-card class="box-card" shadow="hover">
      <template #header>
        <div class="card-header">
          <span>手动数据采集</span>
        </div>
      </template>
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="控制图ID" prop="chart_id">
              <el-input-number v-model="formData.chart_id" :min="1" placeholder="请输入控制图ID" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="批次ID" prop="lot_id">
              <el-input-number v-model="formData.lot_id" :min="1" placeholder="请输入批次ID" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="设备ID" prop="equipment_id">
              <el-input-number v-model="formData.equipment_id" :min="1" placeholder="请输入设备ID" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="采样时间" prop="sample_time">
              <el-date-picker
                v-model="formData.sample_time"
                type="datetime"
                placeholder="选择日期时间"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="测量值" prop="measurements">
          <el-input
            v-model="measurementsText"
            type="textarea"
            :rows="3"
            placeholder="请输入测量值，多个值用逗号分隔，如：650.1, 650.3, 649.8, 650.5, 650.2"
          />
          <div class="help-text">提示：输入的值数量应与控制图的子组大小一致</div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitData">提交数据</el-button>
          <el-button @click="resetForm">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="box-card" shadow="hover" style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>采集结果</span>
        </div>
      </template>
      <div v-if="collectResult">
        <el-descriptions :column="2" border>
          <el-descriptions-item label="样本ID">{{ collectResult.sample_id }}</el-descriptions-item>
          <el-descriptions-item label="子组编号">{{ collectResult.subgroup_no }}</el-descriptions-item>
          <el-descriptions-item label="均值">{{ collectResult.mean?.toFixed(4) }}</el-descriptions-item>
          <el-descriptions-item label="极差">{{ collectResult.range?.toFixed(4) }}</el-descriptions-item>
          <el-descriptions-item label="标准差">{{ collectResult.stddev?.toFixed(4) }}</el-descriptions-item>
          <el-descriptions-item label="OOS数量">
            <el-tag :type="collectResult.oos_count > 0 ? 'danger' : 'success'">
              {{ collectResult.oos_count }}
            </el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="OOC违规">
            <div v-if="collectResult.ooc_violations && collectResult.ooc_violations.length > 0">
              <el-tag v-for="(v, idx) in collectResult.ooc_violations" :key="idx" type="warning" style="margin: 2px">
                {{ v.rule_code }}: {{ v.message }}
              </el-tag>
            </div>
            <el-tag v-else type="success">无</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="告警创建">
            <el-tag :type="collectResult.alarm_created ? 'danger' : 'success'">
              {{ collectResult.alarm_created ? '是 (ID:' + collectResult.alarm_id + ')' : '否' }}
            </el-tag>
          </el-descriptions-item>
        </el-descriptions>
      </div>
      <el-empty v-else description="暂无采集结果" />
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { collectData } from '@/api/spc/collect'

const formRef = ref(null)
const formData = ref({
  chart_id: null,
  lot_id: null,
  equipment_id: null,
  sample_time: new Date()
})

const measurementsText = ref('')
const collectResult = ref(null)

const rules = reactive({
  chart_id: [{ required: true, message: '请输入控制图ID', trigger: 'blur' }],
  lot_id: [{ required: true, message: '请输入批次ID', trigger: 'blur' }],
  equipment_id: [{ required: true, message: '请输入设备ID', trigger: 'blur' }],
  sample_time: [{ required: true, message: '请选择采样时间', trigger: 'change' }]
})

const submitData = async() => {
  formRef.value?.validate(async(valid) => {
    if (valid) {
      if (!measurementsText.value) {
        ElMessage.error('请输入测量值')
        return
      }

      const measurements = measurementsText.value
        .split(',')
        .map(v => parseFloat(v.trim()))
        .filter(v => !isNaN(v))

      if (measurements.length === 0) {
        ElMessage.error('测量值格式不正确')
        return
      }

      const data = {
        ...formData.value,
        sample_time: formData.value.sample_time.toISOString(),
        measurements
      }

      const res = await collectData(data)
      if (res.code === 0) {
        collectResult.value = res.data
        ElMessage({
          type: 'success',
          message: '数据采集成功'
        })
        if (res.data.alarm_created) {
          ElMessage({
            type: 'warning',
            message: '检测到异常，已创建告警',
            duration: 5000
          })
        }
      }
    }
  })
}

const resetForm = () => {
  formRef.value?.resetFields()
  measurementsText.value = ''
  collectResult.value = null
}
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
</style>
