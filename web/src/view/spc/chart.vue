<template>
  <!-- SPC控制图配置 -->
  <div>
    <div class="gva-search-box">
      <el-form ref="searchFormRef" :inline="true" :model="searchInfo">
        <el-form-item label="图表代码">
          <el-input v-model="searchInfo.code" placeholder="请输入图表代码" clearable />
        </el-form-item>
        <el-form-item label="图表名称">
          <el-input v-model="searchInfo.name" placeholder="请输入图表名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog('add')">新增控制图</el-button>
      </div>
      <el-table :data="tableData" ref="multipleTable" row-key="ID">
        <el-table-column label="ID" prop="ID" min-width="60" />
        <el-table-column label="图表代码" prop="code" min-width="120" />
        <el-table-column label="图表名称" prop="name" min-width="150" />
        <el-table-column label="图表类型" prop="chart_type" min-width="100" />
        <el-table-column label="子组大小" prop="subgroup_size" min-width="100" />
        <el-table-column label="状态" prop="status" min-width="80">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" prop="CreatedAt" min-width="160">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" min-width="340">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="openDialog('edit', scope.row)">编辑</el-button>
            <el-button type="success" link icon="setting" @click="openRuleDrawer(scope.row)">规则</el-button>
            <el-button type="info" link icon="data-line" @click="openLimitDrawer(scope.row)">控制限</el-button>
            <el-button type="warning" link icon="refresh" @click="recalculateLimits(scope.row)">重算</el-button>
            <el-button type="danger" link icon="delete" @click="deleteFunc(scope.row)">删除</el-button>
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

    <el-dialog v-model="dialogFormVisible" :title="dialogTitle" width="60%">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="120px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="图表代码" prop="code">
              <el-input v-model="formData.code" placeholder="请输入图表代码" clearable />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="图表名称" prop="name">
              <el-input v-model="formData.name" placeholder="请输入图表名称" clearable />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="图表类型" prop="chart_type">
              <el-select v-model="formData.chart_type" placeholder="请选择图表类型" clearable>
                <el-option label="均值-极差图(XBAR-R)" value="XBAR_R" />
                <el-option label="均值-标准差图(XBAR-S)" value="XBAR_S" />
                <el-option label="单值-移动极差图(I-MR)" value="I_MR" />
                <el-option label="P图" value="P" />
                <el-option label="NP图" value="NP" />
                <el-option label="C图" value="C" />
                <el-option label="U图" value="U" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="子组大小" prop="subgroup_size">
              <el-input-number v-model="formData.subgroup_size" :min="1" :max="25" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="状态" prop="status">
              <el-radio-group v-model="formData.status">
                <el-radio :label="1">启用</el-radio>
                <el-radio :label="0">禁用</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="formData.remark" type="textarea" :rows="3" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取消</el-button>
          <el-button type="primary" @click="enterDialog">确定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 规则配置抽屉 -->
    <el-drawer v-model="ruleDrawerVisible" :title="`规则配置 - ${currentChart?.name || ''}`" size="50%">
      <div style="padding: 0 20px">
        <el-form label-width="180px">
          <el-form-item label="启用Western Electric规则">
            <el-checkbox-group v-model="selectedRules">
              <el-checkbox label="WE1">WE1: 点超出3σ控制限</el-checkbox>
              <el-checkbox label="WE2">WE2: 连续9点同侧</el-checkbox>
              <el-checkbox label="WE3">WE3: 连续6点递增/递减</el-checkbox>
              <el-checkbox label="WE4">WE4: 连续14点交替</el-checkbox>
            </el-checkbox-group>
          </el-form-item>
          <el-form-item label="启用Nelson规则">
            <el-checkbox-group v-model="selectedRules">
              <el-checkbox label="NELSON1">Nelson1: 1点超3σ</el-checkbox>
              <el-checkbox label="NELSON2">Nelson2: 9点连续同侧</el-checkbox>
              <el-checkbox label="NELSON3">Nelson3: 6点递增/递减</el-checkbox>
              <el-checkbox label="NELSON4">Nelson4: 14点交替</el-checkbox>
              <el-checkbox label="NELSON5">Nelson5: 3点中2点超2σ</el-checkbox>
              <el-checkbox label="NELSON6">Nelson6: 5点中4点超1σ</el-checkbox>
              <el-checkbox label="NELSON7">Nelson7: 15点在1σ内</el-checkbox>
              <el-checkbox label="NELSON8">Nelson8: 8点超1σ</el-checkbox>
            </el-checkbox-group>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="saveRules">保存规则配置</el-button>
          </el-form-item>
        </el-form>
      </div>
    </el-drawer>

    <!-- 控制限历史抽屉 -->
    <el-drawer v-model="limitDrawerVisible" :title="`控制限历史 - ${currentChart?.name || ''}`" size="70%">
      <div style="padding: 0 20px">
        <el-table :data="limitData" style="width: 100%">
          <el-table-column label="ID" prop="ID" width="60" />
          <el-table-column label="来源" prop="source" width="100">
            <template #default="scope">
              <el-tag :type="scope.row.source === 'CALC' ? 'success' : 'info'">
                {{ scope.row.source === 'CALC' ? '自动计算' : '手工录入' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="UCL" prop="ucl" width="100">
            <template #default="scope">{{ scope.row.ucl?.toFixed(3) || '-' }}</template>
          </el-table-column>
          <el-table-column label="CL" prop="cl" width="100">
            <template #default="scope">{{ scope.row.cl?.toFixed(3) || '-' }}</template>
          </el-table-column>
          <el-table-column label="LCL" prop="lcl" width="100">
            <template #default="scope">{{ scope.row.lcl?.toFixed(3) || '-' }}</template>
          </el-table-column>
          <el-table-column label="样本数" prop="calcN" width="100" />
          <el-table-column label="生效时间" prop="effectiveFrom" min-width="160">
            <template #default="scope">{{ formatDate(scope.row.effectiveFrom) }}</template>
          </el-table-column>
          <el-table-column label="备注" prop="remark" min-width="150" show-overflow-tooltip />
        </el-table>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getChartList, createChart, updateChart, deleteChart } from '@/api/spc/chart'
import { calculateLimits } from '@/api/spc/runtime'
import { getRuleList, createRule, deleteRule } from '@/api/spc/control'
import { getControlLimitList } from '@/api/spc/control'
import { formatDate } from '@/utils/format'

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = reactive({})

const dialogFormVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref(null)
const formData = ref({
  code: '',
  name: '',
  chart_type: 'XBAR_R',
  subgroup_size: 5,
  status: 1,
  remark: ''
})

const rules = reactive({
  code: [{ required: true, message: '请输入图表代码', trigger: 'blur' }],
  name: [{ required: true, message: '请输入图表名称', trigger: 'blur' }],
  chart_type: [{ required: true, message: '请选择图表类型', trigger: 'change' }],
  subgroup_size: [{ required: true, message: '请输入子组大小', trigger: 'blur' }]
})

const getTableData = async() => {
  const table = await getChartList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.code = ''
  searchInfo.name = ''
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

const openDialog = (type, row) => {
  dialogFormVisible.value = true
  if (type === 'add') {
    dialogTitle.value = '新增控制图'
    formData.value = {
      code: '',
      name: '',
      chart_type: 'XBAR_R',
      subgroup_size: 5,
      status: 1,
      remark: ''
    }
  } else {
    dialogTitle.value = '编辑控制图'
    formData.value = { ...row }
  }
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formRef.value?.resetFields()
}

const enterDialog = async() => {
  formRef.value?.validate(async(valid) => {
    if (valid) {
      let res
      if (formData.value.ID) {
        res = await updateChart(formData.value)
      } else {
        res = await createChart(formData.value)
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: formData.value.ID ? '编辑成功' : '创建成功'
        })
        closeDialog()
        getTableData()
      }
    }
  })
}

const deleteFunc = async(row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteChart({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      getTableData()
    } else {
      ElMessage.error(res.msg || '删除失败')
    }
  })
}

// 规则配置
const ruleDrawerVisible = ref(false)
const currentChart = ref(null)
const selectedRules = ref([])

const openRuleDrawer = async(chart) => {
  currentChart.value = chart
  ruleDrawerVisible.value = true
  await loadChartRules()
}

const loadChartRules = async() => {
  if (!currentChart.value) return
  try {
    const res = await getRuleList({ page: 1, pageSize: 100, chartId: currentChart.value.ID })
    if (res.code === 0 && res.data.list) {
      selectedRules.value = res.data.list
        .filter(r => r.enabled)
        .map(r => r.ruleType)
    }
  } catch (error) {
    console.error('加载规则失败', error)
  }
}

const saveRules = async() => {
  if (!currentChart.value) return
  try {
    // 简化实现：先删除现有规则，再创建新规则
    const existingRes = await getRuleList({ page: 1, pageSize: 100, chartId: currentChart.value.ID })
    if (existingRes.code === 0 && existingRes.data.list) {
      for (const rule of existingRes.data.list) {
        await deleteRule({ ID: rule.ID })
      }
    }
    
    // 创建选中的规则
    for (const ruleType of selectedRules.value) {
      await createRule({
        chartId: currentChart.value.ID,
        ruleType: ruleType,
        enabled: 1,
        severity: ruleType.includes('1') ? 'CRITICAL' : 'WARNING'
      })
    }
    
    ElMessage.success('规则配置已保存')
    ruleDrawerVisible.value = false
  } catch (error) {
    ElMessage.error('保存规则失败')
  }
}

// 控制限历史
const limitDrawerVisible = ref(false)
const limitData = ref([])

const openLimitDrawer = async(chart) => {
  currentChart.value = chart
  limitDrawerVisible.value = true
  await loadControlLimits()
}

const loadControlLimits = async() => {
  if (!currentChart.value) return
  try {
    const res = await getControlLimitList({ page: 1, pageSize: 100, chartId: currentChart.value.ID })
    if (res.code === 0) {
      limitData.value = res.data.list || []
    }
  } catch (error) {
    ElMessage.error('加载控制限失败')
  }
}

// 重算控制限
const recalculateLimits = async(chart) => {
  try {
    await ElMessageBox.confirm(
      `确定要重新计算控制限吗？将使用最近30个样本计算。`,
      '重算控制限',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    
    const res = await calculateLimits({ chartId: chart.ID, window: 30 })
    if (res.code === 0) {
      ElMessage.success('控制限重算成功')
    } else {
      ElMessage.error(res.msg || '重算失败')
    }
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('重算失败')
    }
  }
}
</script>

<style scoped>
</style>
