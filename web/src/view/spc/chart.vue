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
        <el-table-column label="操作" fixed="right" min-width="200">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="openDialog('edit', scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteFunc(scope.row)">删除</el-button>
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
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getChartList, createChart, updateChart, deleteChart } from '@/api/spc/chart'
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
    }
  })
}
</script>

<style scoped>
</style>
