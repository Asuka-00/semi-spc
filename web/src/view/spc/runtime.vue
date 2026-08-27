<template>
  <div>
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span>SPC控制图实时监控</span>
          <div class="header-controls">
            <el-select
              v-model="selectedChartId"
              placeholder="选择控制图"
              style="width: 250px; margin-right: 10px"
              filterable
              @change="loadChartData"
            >
              <el-option
                v-for="chart in chartList"
                :key="chart.ID"
                :label="`${chart.code} - ${chart.name}`"
                :value="chart.ID"
              />
            </el-select>
            <el-date-picker
              v-model="timeRange"
              type="datetimerange"
              range-separator="至"
              start-placeholder="开始时间"
              end-placeholder="结束时间"
              style="width: 360px; margin-right: 10px"
              @change="loadChartData"
            />
            <el-button type="primary" :loading="loading" @click="loadChartData">刷新</el-button>
          </div>
        </div>
      </template>

      <div v-if="chartData" class="chart-info">
        <el-descriptions :column="4" size="small" border>
          <el-descriptions-item label="图表类型">{{ chartData.chart?.chartType }}</el-descriptions-item>
          <el-descriptions-item label="子组大小">{{ chartData.chart?.subgroupSize }}</el-descriptions-item>
          <el-descriptions-item label="控制限方法">{{ chartData.chart?.limitMethod }}</el-descriptions-item>
          <el-descriptions-item label="规则集">{{ chartData.chart?.ruleset }}</el-descriptions-item>
        </el-descriptions>
      </div>

      <div ref="chartContainer" class="chart-container" v-loading="loading"></div>
    </el-card>

    <el-card shadow="hover" style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>数据点列表 ({{ chartData?.total || 0 }} 个样本)</span>
        </div>
      </template>
      <el-table :data="chartData?.samples || []" style="width: 100%" @row-click="handleRowClick">
        <el-table-column prop="subgroupNo" label="子组号" width="100" />
        <el-table-column prop="sampleTime" label="采样时间" width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.sampleTime) }}
          </template>
        </el-table-column>
        <el-table-column prop="meanVal" label="均值 X̄" width="120" />
        <el-table-column prop="rangeVal" label="极差 R" width="120" />
        <el-table-column prop="stdVal" label="标准差 S" width="120" />
        <el-table-column prop="oosFlag" label="OOS" width="80">
          <template #default="{ row }">
            <el-tag v-if="row.oosFlag" type="danger" size="small">超规格</el-tag>
            <el-tag v-else type="success" size="small">正常</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="oocFlag" label="OOC" width="80">
          <template #default="{ row }">
            <el-tag v-if="row.oocFlag" type="warning" size="small">失控</el-tag>
            <el-tag v-else type="success" size="small">受控</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120">
          <template #default="{ row }">
            <el-button type="primary" link size="small" @click.stop="viewMeasurements(row)">
              查看测量值
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 测量值详情对话框 -->
    <el-dialog v-model="measurementDialog" title="测量值详情" width="600px">
      <el-descriptions :column="2" border size="small" v-if="currentSample">
        <el-descriptions-item label="子组号">{{ currentSample.subgroupNo }}</el-descriptions-item>
        <el-descriptions-item label="采样时间">{{ formatDateTime(currentSample.sampleTime) }}</el-descriptions-item>
        <el-descriptions-item label="均值">{{ currentSample.meanVal }}</el-descriptions-item>
        <el-descriptions-item label="测量点数">{{ currentSample.n }}</el-descriptions-item>
      </el-descriptions>
      <el-divider />
      <el-table :data="measurements" style="width: 100%" size="small">
        <el-table-column prop="seq" label="序号" width="80" />
        <el-table-column prop="value" label="测量值" />
      </el-table>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import * as echarts from 'echarts'
import { getChartList } from '@/api/spc/chart'
import { getChartRuntime, getMeasurementList } from '@/api/spc/collect'

const loading = ref(false)
const chartList = ref([])
const selectedChartId = ref(null)
const timeRange = ref([])
const chartData = ref(null)
const chartContainer = ref(null)
let chartInstance = null

const measurementDialog = ref(false)
const currentSample = ref(null)
const measurements = ref([])

const formatDateTime = (dateStr) => {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleString('zh-CN')
}

// 加载图表列表
const loadChartList = async () => {
  try {
    const res = await getChartList({ page: 1, pageSize: 100 })
    if (res.code === 0 && res.data.list) {
      chartList.value = res.data.list
      if (chartList.value.length > 0) {
        selectedChartId.value = chartList.value[0].ID
        await loadChartData()
      }
    }
  } catch (error) {
    ElMessage.error('加载图表列表失败')
  }
}

// 加载图表运行时数据
const loadChartData = async () => {
  if (!selectedChartId.value) return

  loading.value = true
  try {
    const params = {
      chartId: selectedChartId.value,
      page: 1,
      pageSize: 100
    }

    if (timeRange.value && timeRange.value.length === 2) {
      params.from = timeRange.value[0].toISOString()
      params.to = timeRange.value[1].toISOString()
    }

    const res = await getChartRuntime(params)
    if (res.code === 0) {
      chartData.value = res.data
      await nextTick()
      renderChart()
    }
  } catch (error) {
    ElMessage.error(error.message || '加载图表数据失败')
  } finally {
    loading.value = false
  }
}

// 渲染ECharts图表
const renderChart = () => {
  if (!chartContainer.value || !chartData.value) return

  if (!chartInstance) {
    chartInstance = echarts.init(chartContainer.value)
  }

  const samples = chartData.value.samples || []
  const chartType = chartData.value.chart?.chartType || ''
  const currentLimit = chartData.value.currentLimit || {}
  const currentSpec = chartData.value.currentSpec || {}

  // 准备数据
  const xData = samples.map(s => s.subgroupNo)
  const meanData = samples.map(s => s.meanVal)
  const rangeData = samples.map(s => s.rangeVal)
  const stdData = samples.map(s => s.stdVal)

  // 标记OOC/OOS点
  const oocPoints = []
  const oosPoints = []
  samples.forEach((s, idx) => {
    if (s.oocFlag) oocPoints.push([idx, s.meanVal])
    if (s.oosFlag) oosPoints.push([idx, s.meanVal])
  })

  const series = [
    {
      name: '均值 X̄',
      type: 'line',
      data: meanData,
      symbol: 'circle',
      symbolSize: 6,
      itemStyle: { color: '#409EFF' }
    }
  ]

  // 添加控制限线
  if (currentLimit.CL != null) {
    series.push({
      name: '中心线 CL',
      type: 'line',
      data: Array(xData.length).fill(currentLimit.CL),
      lineStyle: { type: 'solid', color: '#67C23A', width: 2 },
      symbol: 'none'
    })
  }
  if (currentLimit.UCL != null) {
    series.push({
      name: '上控制限 UCL',
      type: 'line',
      data: Array(xData.length).fill(currentLimit.UCL),
      lineStyle: { type: 'dashed', color: '#E6A23C', width: 2 },
      symbol: 'none'
    })
  }
  if (currentLimit.LCL != null) {
    series.push({
      name: '下控制限 LCL',
      type: 'line',
      data: Array(xData.length).fill(currentLimit.LCL),
      lineStyle: { type: 'dashed', color: '#E6A23C', width: 2 },
      symbol: 'none'
    })
  }

  // 添加规格限线
  if (currentSpec.usl != null) {
    series.push({
      name: '规格上限 USL',
      type: 'line',
      data: Array(xData.length).fill(currentSpec.usl),
      lineStyle: { type: 'dashed', color: '#F56C6C', width: 2 },
      symbol: 'none'
    })
  }
  if (currentSpec.lsl != null) {
    series.push({
      name: '规格下限 LSL',
      type: 'line',
      data: Array(xData.length).fill(currentSpec.lsl),
      lineStyle: { type: 'dashed', color: '#F56C6C', width: 2 },
      symbol: 'none'
    })
  }

  // 标记OOC点
  if (oocPoints.length > 0) {
    series.push({
      name: 'OOC点',
      type: 'scatter',
      data: oocPoints,
      symbolSize: 12,
      itemStyle: { color: '#E6A23C', borderColor: '#fff', borderWidth: 2 }
    })
  }

  // 标记OOS点
  if (oosPoints.length > 0) {
    series.push({
      name: 'OOS点',
      type: 'scatter',
      data: oosPoints,
      symbolSize: 12,
      itemStyle: { color: '#F56C6C', borderColor: '#fff', borderWidth: 2 }
    })
  }

  const option = {
    title: {
      text: chartData.value.chart?.name || '控制图',
      left: 'center'
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: { type: 'cross' }
    },
    legend: {
      data: series.map(s => s.name),
      top: 30
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '10%',
      top: 80,
      containLabel: true
    },
    xAxis: {
      type: 'category',
      data: xData,
      name: '子组号',
      nameLocation: 'middle',
      nameGap: 30
    },
    yAxis: {
      type: 'value',
      name: '测量值',
      nameLocation: 'middle',
      nameGap: 50
    },
    series: series
  }

  chartInstance.setOption(option, true)

  // 点击事件
  chartInstance.off('click')
  chartInstance.on('click', (params) => {
    if (params.componentType === 'series' && params.seriesName === '均值 X̄') {
      const sample = samples[params.dataIndex]
      if (sample) {
        viewMeasurements(sample)
      }
    }
  })
}

// 查看测量值
const viewMeasurements = async (sample) => {
  currentSample.value = sample
  measurementDialog.value = true

  try {
    const res = await getMeasurementList({ sampleId: sample.ID })
    if (res.code === 0 && res.data) {
      measurements.value = res.data.map((m, idx) => ({
        seq: idx + 1,
        value: m.value
      }))
    }
  } catch (error) {
    ElMessage.error('获取测量值失败')
  }
}

const handleRowClick = (row) => {
  viewMeasurements(row)
}

onMounted(() => {
  loadChartList()

  // 窗口resize时重绘图表
  window.addEventListener('resize', () => {
    if (chartInstance) {
      chartInstance.resize()
    }
  })
})
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-controls {
  display: flex;
  align-items: center;
}

.chart-info {
  margin-bottom: 20px;
}

.chart-container {
  width: 100%;
  height: 500px;
}
</style>
