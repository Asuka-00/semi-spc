<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchFormRef" :inline="true" :model="searchInfo">
        <el-form-item label="参数">
          <el-select v-model="searchInfo.parameterId" placeholder="选择参数" clearable filterable style="width: 200px">
            <el-option
              v-for="param in parameterList"
              :key="param.ID"
              :label="`${param.code} - ${param.name}`"
              :value="param.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="版本">
          <el-input v-model="searchInfo.version" placeholder="请输入版本" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog('add')">新增</el-button>
      </div>
      <el-table :data="tableData" ref="multipleTable" row-key="ID">
        <el-table-column label="ID" prop="ID" min-width="80" />
        <el-table-column label="参数" prop="parameterId" min-width="150">
          <template #default="{ row }">
            {{ getParameterName(row.parameterId) }}
          </template>
        </el-table-column>
        <el-table-column label="版本" prop="version" min-width="100" />
        <el-table-column label="规格上限 USL" prop="usl" min-width="120">
          <template #default="{ row }">
            {{ row.usl != null ? row.usl : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="目标值 Target" prop="target" min-width="120">
          <template #default="{ row }">
            {{ row.target != null ? row.target : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="规格下限 LSL" prop="lsl" min-width="120">
          <template #default="{ row }">
            {{ row.lsl != null ? row.lsl : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="生效起" prop="effectiveFrom" min-width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.effectiveFrom) }}
          </template>
        </el-table-column>
        <el-table-column label="生效止" prop="effectiveTo" min-width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.effectiveTo) }}
          </template>
        </el-table-column>
        <el-table-column label="状态" prop="status" min-width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
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
            <el-form-item label="参数" prop="parameterId">
              <el-select v-model="formData.parameterId" placeholder="选择参数" filterable style="width: 100%">
                <el-option
                  v-for="param in parameterList"
                  :key="param.ID"
                  :label="`${param.code} - ${param.name}`"
                  :value="param.ID"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="版本" prop="version">
              <el-input v-model="formData.version" placeholder="如：V1.0" clearable />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="规格上限 USL" prop="usl">
              <el-input-number
                v-model="formData.usl"
                :controls="false"
                placeholder="规格上限"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="目标值 Target" prop="target">
              <el-input-number
                v-model="formData.target"
                :controls="false"
                placeholder="目标值"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="规格下限 LSL" prop="lsl">
              <el-input-number
                v-model="formData.lsl"
                :controls="false"
                placeholder="规格下限"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-alert
          v-if="formData.usl != null && formData.lsl != null && formData.usl <= formData.lsl"
          title="规格上限必须大于下限！"
          type="error"
          :closable="false"
          show-icon
          style="margin-bottom: 15px"
        />
        <el-alert
          v-if="formData.target != null && formData.usl != null && formData.target > formData.usl"
          title="目标值必须在规格范围内！"
          type="warning"
          :closable="false"
          show-icon
          style="margin-bottom: 15px"
        />
        <el-alert
          v-if="formData.target != null && formData.lsl != null && formData.target < formData.lsl"
          title="目标值必须在规格范围内！"
          type="warning"
          :closable="false"
          show-icon
          style="margin-bottom: 15px"
        />
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="生效起" prop="effectiveFrom">
              <el-date-picker
                v-model="formData.effectiveFrom"
                type="datetime"
                placeholder="选择生效起始时间"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="生效止" prop="effectiveTo">
              <el-date-picker
                v-model="formData.effectiveTo"
                type="datetime"
                placeholder="选择生效结束时间"
                style="width: 100%"
              />
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
          <el-input v-model="formData.remark" type="textarea" placeholder="请输入备注" />
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
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getSpecList, createSpec, updateSpec, deleteSpec, getParameterList } from '@/api/spc/material'

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = reactive({})
const parameterList = ref([])

const dialogFormVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref(null)
const formData = ref({
  parameterId: null,
  version: '',
  usl: null,
  target: null,
  lsl: null,
  effectiveFrom: null,
  effectiveTo: null,
  status: 1,
  remark: ''
})

const rules = reactive({
  parameterId: [{ required: true, message: '请选择参数', trigger: 'change' }],
  version: [{ required: true, message: '请输入版本', trigger: 'blur' }]
})

const formatDateTime = (dateStr) => {
  if (!dateStr) return '-'
  return new Date(dateStr).toLocaleString('zh-CN')
}

const loadParameterList = async () => {
  try {
    const res = await getParameterList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      parameterList.value = res.data.list || []
    }
  } catch (error) {
    console.error('加载参数列表失败', error)
  }
}

const getParameterName = (id) => {
  const param = parameterList.value.find(p => p.ID === id)
  return param ? `${param.code} - ${param.name}` : id
}

const getTableData = async() => {
  const table = await getSpecList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

onMounted(() => {
  loadParameterList()
  getTableData()
})

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const onReset = () => {
  searchInfo.parameterId = null
  searchInfo.version = ''
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
    dialogTitle.value = '新增规格'
    formData.value = {
      parameterId: null,
      version: '',
      usl: null,
      target: null,
      lsl: null,
      effectiveFrom: null,
      effectiveTo: null,
      status: 1,
      remark: ''
    }
  } else {
    dialogTitle.value = '编辑规格'
    formData.value = {
      ...row,
      effectiveFrom: row.effectiveFrom ? new Date(row.effectiveFrom) : null,
      effectiveTo: row.effectiveTo ? new Date(row.effectiveTo) : null
    }
  }
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formRef.value?.resetFields()
}

const enterDialog = async() => {
  formRef.value?.validate(async(valid) => {
    if (valid) {
      // 前端验证
      if (formData.value.usl != null && formData.value.lsl != null && formData.value.usl <= formData.value.lsl) {
        ElMessage.error('规格上限必须大于下限')
        return
      }
      if (formData.value.target != null) {
        if (formData.value.usl != null && formData.value.target > formData.value.usl) {
          ElMessage.error('目标值必须在规格范围内')
          return
        }
        if (formData.value.lsl != null && formData.value.target < formData.value.lsl) {
          ElMessage.error('目标值必须在规格范围内')
          return
        }
      }

      const submitData = {
        ...formData.value,
        effectiveFrom: formData.value.effectiveFrom ? formData.value.effectiveFrom.toISOString() : null,
        effectiveTo: formData.value.effectiveTo ? formData.value.effectiveTo.toISOString() : null
      }

      let res
      if (formData.value.ID) {
        res = await updateSpec(submitData)
      } else {
        res = await createSpec(submitData)
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: formData.value.ID ? '编辑成功' : '创建成功'
        })
        closeDialog()
        getTableData()
      } else {
        ElMessage.error(res.msg || '操作失败')
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
    const res = await deleteSpec({ ID: row.ID })
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
</script>

<style scoped>
</style>
