<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchFormRef" :inline="true" :model="searchInfo">
        <el-form-item label="工艺">
          <el-select v-model="searchInfo.technologyId" placeholder="选择工艺" clearable filterable style="width: 200px">
            <el-option
              v-for="tech in technologyList"
              :key="tech.ID"
              :label="`${tech.code} - ${tech.name}`"
              :value="tech.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="产品代码">
          <el-input v-model="searchInfo.code" placeholder="请输入产品代码" clearable />
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
        <el-table-column label="产品代码" prop="code" min-width="120" />
        <el-table-column label="产品名称" prop="name" min-width="150" />
        <el-table-column label="工艺" prop="technologyId" min-width="120">
          <template #default="{ row }">
            {{ getTechnologyName(row.technologyId) }}
          </template>
        </el-table-column>
        <el-table-column label="层数" prop="layers" min-width="80" />
        <el-table-column label="状态" prop="status" min-width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="备注" prop="remark" min-width="200" show-overflow-tooltip />
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

    <el-dialog v-model="dialogFormVisible" :title="dialogTitle" width="50%">
      <el-form ref="formRef" :model="formData" :rules="rules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="产品代码" prop="code">
              <el-input v-model="formData.code" placeholder="如：PROD001" clearable />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="产品名称" prop="name">
              <el-input v-model="formData.name" placeholder="请输入产品名称" clearable />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="工艺" prop="technologyId">
              <el-select v-model="formData.technologyId" placeholder="选择工艺" filterable style="width: 100%">
                <el-option
                  v-for="tech in technologyList"
                  :key="tech.ID"
                  :label="`${tech.code} - ${tech.name}`"
                  :value="tech.ID"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="层数" prop="layers">
              <el-input-number v-model="formData.layers" :min="1" style="width: 100%" />
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
import { getProductList, createProduct, updateProduct, deleteProduct, getTechnologyList } from '@/api/spc/master'

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = reactive({})
const technologyList = ref([])

const dialogFormVisible = ref(false)
const dialogTitle = ref('')
const formRef = ref(null)
const formData = ref({
  code: '',
  name: '',
  technologyId: null,
  layers: 1,
  status: 1,
  remark: ''
})

const rules = reactive({
  code: [{ required: true, message: '请输入产品代码', trigger: 'blur' }],
  name: [{ required: true, message: '请输入产品名称', trigger: 'blur' }],
  technologyId: [{ required: true, message: '请选择工艺', trigger: 'change' }]
})

const loadTechnologyList = async () => {
  try {
    const res = await getTechnologyList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      technologyList.value = res.data.list || []
    }
  } catch (error) {
    console.error('加载工艺列表失败', error)
  }
}

const getTechnologyName = (id) => {
  const tech = technologyList.value.find(t => t.ID === id)
  return tech ? `${tech.code} - ${tech.name}` : id
}

const getTableData = async() => {
  const table = await getProductList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

onMounted(() => {
  loadTechnologyList()
  getTableData()
})

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const onReset = () => {
  searchInfo.technologyId = null
  searchInfo.code = ''
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
    dialogTitle.value = '新增产品'
    formData.value = {
      code: '',
      name: '',
      technologyId: null,
      layers: 1,
      status: 1,
      remark: ''
    }
  } else {
    dialogTitle.value = '编辑产品'
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
        res = await updateProduct(formData.value)
      } else {
        res = await createProduct(formData.value)
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
    const res = await deleteProduct({ ID: row.ID })
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
