<template>
  <!-- SPC Dashboard 仪表板 -->
  <div>
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="今日告警" :value="statistics.todayAlarms">
            <template #suffix>
              <el-icon><warning /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="未处理告警" :value="statistics.openAlarms">
            <template #suffix>
              <el-icon><bell /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="OOC率" :value="statistics.oocRate" :precision="2" suffix="%">
            <template #suffix>
              <el-icon><data-line /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <el-statistic title="活跃控制图" :value="statistics.activeCharts">
            <template #suffix>
              <el-icon><data-analysis /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <span>最新告警</span>
              <el-button type="text" @click="gotoAlarmCenter">查看全部</el-button>
            </div>
          </template>
          <el-table :data="recentAlarms" style="width: 100%">
            <el-table-column prop="ID" label="ID" width="60" />
            <el-table-column prop="alarmType" label="类型" width="80">
              <template #default="scope">
                <el-tag :type="scope.row.alarmType === 'OOC' ? 'warning' : 'danger'" size="small">
                  {{ scope.row.alarmType }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="rule_code" label="规则" width="80" />
            <el-table-column prop="status" label="状态" width="80">
              <template #default="scope">
                <el-tag :type="scope.row.status === 'OPEN' ? 'danger' : 'success'" size="small">
                  {{ scope.row.status }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="CreatedAt" label="时间" :formatter="dateFormatter" />
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <span>告警趋势（近7天）</span>
            </div>
          </template>
          <div style="height: 300px; display: flex; align-items: center; justify-content: center">
            <el-empty description="图表功能待集成ECharts" />
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" style="margin-top: 20px">
      <el-col :span="24">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <span>TOP设备告警统计</span>
            </div>
          </template>
          <el-table :data="topEquipment" style="width: 100%">
            <el-table-column prop="equipment_name" label="设备名称" />
            <el-table-column prop="alarm_count" label="告警次数" sortable />
            <el-table-column prop="ooc_count" label="OOC次数" sortable />
            <el-table-column prop="oos_count" label="OOS次数" sortable />
            <el-table-column label="操作" width="150">
              <template #default="scope">
                <el-button type="text" size="small" @click="viewEquipmentDetail(scope.row)">查看详情</el-button>
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
import { useRouter } from 'vue-router'
import { getAlarmList, getAlarmStatistics } from '@/api/spc/collect'
import { formatDate } from '@/utils/format'

const router = useRouter()

const statistics = ref({
  todayAlarms: 0,
  openAlarms: 0,
  oocRate: 0,
  activeCharts: 0
})

const recentAlarms = ref([])
const topEquipment = ref([
  { equipment_name: '扩散炉01', alarm_count: 15, ooc_count: 10, oos_count: 5 },
  { equipment_name: '刻蚀机02', alarm_count: 12, ooc_count: 8, oos_count: 4 },
  { equipment_name: 'CVD设备03', alarm_count: 8, ooc_count: 5, oos_count: 3 }
])

const dateFormatter = (row, column, cellValue) => {
  return formatDate(cellValue)
}

const loadStatistics = async() => {
  const res = await getAlarmStatistics({ days: 7 })
  if (res.code === 0) {
    statistics.value = {
      todayAlarms: res.data.total_count || 0,
      openAlarms: res.data.open_count || 0,
      oocRate: res.data.ooc_count > 0 ? (res.data.ooc_count / res.data.total_count * 100) : 0,
      activeCharts: 10
    }
  }
}

const loadRecentAlarms = async() => {
  const res = await getAlarmList({ page: 1, pageSize: 5 })
  if (res.code === 0) {
    recentAlarms.value = res.data.list || []
  }
}

const gotoAlarmCenter = () => {
  router.push('/spc/alarm')
}

const viewEquipmentDetail = (row) => {
  console.log('查看设备详情:', row)
}

onMounted(() => {
  loadStatistics()
  loadRecentAlarms()
})
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
