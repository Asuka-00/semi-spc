<template>
  <!-- SPC实时监控 - 控制图展示 -->
  <div>
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>控制图实时监控</span>
          <el-select v-model="selectedChartId" placeholder="选择控制图" style="width: 200px" @change="loadChartData">
            <el-option label="温度监控图" value="1" />
            <el-option label="压力监控图" value="2" />
            <el-option label="流量监控图" value="3" />
          </el-select>
        </div>
      </template>

      <div class="chart-container">
        <div class="chart-placeholder">
          <el-empty description="控制图展示区域 - 待集成ECharts">
            <template #image>
              <el-icon :size="100"><data-analysis /></el-icon>
            </template>
            <template #description>
              <p>此处将展示:</p>
              <ul style="text-align: left; display: inline-block">
                <li>X̄ 图 (均值图)</li>
                <li>R 或 S 图 (极差/标准差图)</li>
                <li>控制限线 (UCL, CL, LCL)</li>
                <li>规格限线 (USL, LSL)</li>
                <li>OOC/OOS 点标记</li>
              </ul>
            </template>
          </el-empty>
        </div>
      </div>
    </el-card>

    <el-card shadow="hover" style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>数据点列表</span>
        </div>
      </template>
      <el-table :data="sampleData" style="width: 100%">
        <el-table-column prop="subgroup_no" label="子组号" width="100" />
        <el-table-column prop="sample_time" label="采样时间" width="180" :formatter="dateFormatter" />
        <el-table-column prop="mean" label="均值" width="100" />
        <el-table-column prop="range" label="极差" width="100" />
        <el-table-column prop="stddev" label="标准差" width="100" />
        <el-table-column prop="is_oos" label="OOS" width="80">
          <template #default="scope">
            <el-tag v-if="scope.row.is_oos" type="danger" size="small">是</el-tag>
            <el-tag v-else type="success" size="small">否</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="is_ooc" label="OOC" width="80">
          <template #default="scope">
            <el-tag v-if="scope.row.is_ooc" type="warning" size="small">是</el-tag>
            <el-tag v-else type="success" size="small">否</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="ooc_rules" label="违规规则" min-width="200">
          <template #default="scope">
            <span v-if="scope.row.ooc_rules">{{ scope.row.ooc_rules }}</span>
            <span v-else style="color: #909399">-</span>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { formatDate } from '@/utils/format'

const selectedChartId = ref('1')
const sampleData = ref([
  { subgroup_no: 45, sample_time: '2026-08-27T10:00:00Z', mean: 650.18, range: 0.7, stddev: 0.28, is_oos: false, is_ooc: false, ooc_rules: null },
  { subgroup_no: 46, sample_time: '2026-08-27T11:00:00Z', mean: 650.25, range: 0.8, stddev: 0.31, is_oos: false, is_ooc: false, ooc_rules: null },
  { subgroup_no: 47, sample_time: '2026-08-27T12:00:00Z', mean: 653.50, range: 1.2, stddev: 0.45, is_oos: true, is_ooc: true, ooc_rules: 'WE1: 1个点超出3σ控制限' }
])

const dateFormatter = (row, column, cellValue) => {
  return formatDate(cellValue)
}

const loadChartData = () => {
  console.log('加载控制图:', selectedChartId.value)
}
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.chart-container {
  width: 100%;
  height: 400px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.chart-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px dashed #dcdfe6;
  border-radius: 4px;
}
</style>
