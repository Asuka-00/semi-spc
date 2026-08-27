<template>
  <div>
    <!-- 统计卡片 -->
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="今日告警" :value="stats.todayAlarms">
            <template #suffix>
              <el-icon color="#F56C6C"><warning /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="未处理告警" :value="stats.openAlarms">
            <template #suffix>
              <el-icon color="#E6A23C"><bell /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="OOC率 (%)" :value="stats.oocRate" :precision="2">
            <template #suffix>
              <el-icon color="#409EFF"><trend-charts /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="进行中OCAP" :value="stats.openOcap">
            <template #suffix>
              <el-icon color="#67C23A"><document /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
    </el-row>

    <!-- TOP设备告警 -->
    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <span>TOP设备（按告警数）</span>
              <span class="header-sub">近7天</span>
            </div>
          </template>
          <el-table :data="stats.topEquipment || []" style="width: 100%">
            <el-table-column prop="equipment_id" label="设备ID" width="100" />
            <el-table-column prop="alarm_count" label="告警数" width="100">
              <template #default="{ row }">
                <el-tag type="danger">{{ row.alarm_count }}</el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <span>TOP参数（按告警数）</span>
              <span class="header-sub">近7天</span>
            </div>
          </template>
          <el-table :data="stats.topParameter || []" style="width: 100%">
            <el-table-column prop="parameter_id" label="参数ID" width="100" />
            <el-table-column prop="alarm_count" label="告警数" width="100">
              <template #default="{ row }">
                <el-tag type="warning">{{ row.alarm_count }}</el-tag>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { Warning, Bell, TrendCharts, Document } from '@element-plus/icons-vue'
import { getDashboard } from '@/api/spc/runtime'

const stats = ref({
  todayAlarms: 0,
  openAlarms: 0,
  oocRate: 0,
  openOcap: 0,
  topEquipment: [],
  topParameter: []
})

const loadDashboard = async () => {
  try {
    const res = await getDashboard()
    if (res.code === 0) {
      stats.value = res.data
    }
  } catch (error) {
    ElMessage.error('加载Dashboard失败')
  }
}

onMounted(() => {
  loadDashboard()
  // 每30秒刷新一次
  setInterval(loadDashboard, 30000)
})
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-sub {
  font-size: 12px;
  color: #909399;
}
</style>
