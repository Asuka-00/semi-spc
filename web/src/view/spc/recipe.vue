<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchFormRef" :inline="true" :model="searchInfo">
        <el-form-item label="设备">
          <el-select v-model="searchInfo.equipmentId" placeholder="选择设备" clearable filterable style="width: 200px">
            <el-option
              v-for="eqp in equipmentList"
              :key="eqp.ID"
              :label="`${eqp.code} - ${eqp.name}`"
              :value="eqp.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="配方代码">
          <el-input v-model="searchInfo.code" placeholder="请输入配方代码" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    
    <el-tabs v-model="activeTab" type="border-card">
      <el-tab-pane label="配方管理" name="recipe">
        <div class="gva-table-box">
          <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openRecipeDialog('add')">新增配方</el-button>
          </div>
          <el-table :data="recipeData" ref="multipleTable" row-key="ID">
            <el-table-column label="ID" prop="ID" min-width="80" />
            <el-table-column label="配方代码" prop="code" min-width="120" />
            <el-table-column label="配方名称" prop="name" min-width="150" />
            <el-table-column label="设备" prop="equipmentId" min-width="150">
              <template #default="{ row }">
                {{ getEquipmentName(row.equipmentId) }}
              </template>
            </el-table-column>
            <el-table-column label="工艺步骤" prop="processStepId" min-width="150">
              <template #default="{ row }">
                {{ getProcessStepName(row.processStepId) }}
              </template>
            </el-table-column>
            <el-table-column label="版本" prop="version" min-width="80" />
            <el-table-column label="状态" prop="status" min-width="100">
              <template #default="scope">
                <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
                  {{ scope.row.status === 1 ? '启用' : '禁用' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column label="操作" fixed="right" min-width="200">
              <template #default="scope">
                <el-button type="primary" link icon="edit" @click="openRecipeDialog('edit', scope.row)">编辑</el-button>
                <el-button type="primary" link icon="delete" @click="deleteRecipe(scope.row)">删除</el-button>
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
      </el-tab-pane>

      <el-tab-pane label="工艺步骤管理" name="processStep">
        <div class="gva-table-box">
          <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openStepDialog('add')">新增工艺步骤</el-button>
          </div>
          <el-table :data="stepData" ref="stepTable" row-key="ID">
            <el-table-column label="ID" prop="ID" min-width="80" />
            <el-table-column label="步骤代码" prop="code" min-width="120" />
            <el-table-column label="步骤名称" prop="name" min-width="150" />
            <el-table-column label="步骤序号" prop="sequence" min-width="100" />
            <el-table-column label="状态" prop="status" min-width="100">
              <template #default="scope">
                <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
                  {{ scope.row.status === 1 ? '启用' : '禁用' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column label="操作" fixed="right" min-width="200">
              <template #default="scope">
                <el-button type="primary" link icon="edit" @click="openStepDialog('edit', scope.row)">编辑</el-button>
                <el-button type="primary" link icon="delete" @click="deleteStep(scope.row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
          <div class="gva-pagination">
            <el-pagination
              :current-page="stepPage"
              :page-size="stepPageSize"
              :page-sizes="[10, 30, 50, 100]"
              :total="stepTotal"
              layout="total, sizes, prev, pager, next, jumper"
              @current-change="handleStepCurrentChange"
              @size-change="handleStepSizeChange"
            />
          </div>
        </div>
      </el-tab-pane>
    </el-tabs>

    <!-- 配方对话框 -->
    <el-dialog v-model="recipeDialogVisible" :title="recipeDialogTitle" width="60%">
      <el-form ref="recipeFormRef" :model="recipeFormData" :rules="recipeRules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="配方代码" prop="code">
              <el-input v-model="recipeFormData.code" placeholder="如：RECIPE001" clearable />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="配方名称" prop="name">
              <el-input v-model="recipeFormData.name" placeholder="请输入配方名称" clearable />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="设备" prop="equipmentId">
              <el-select v-model="recipeFormData.equipmentId" placeholder="选择设备" filterable style="width: 100%">
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
            <el-form-item label="工艺步骤" prop="processStepId">
              <el-select v-model="recipeFormData.processStepId" placeholder="选择工艺步骤" filterable style="width: 100%">
                <el-option
                  v-for="step in processStepList"
                  :key="step.ID"
                  :label="`${step.code} - ${step.name}`"
                  :value="step.ID"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="版本" prop="version">
              <el-input v-model="recipeFormData.version" placeholder="如：V1.0" clearable />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="状态" prop="status">
              <el-radio-group v-model="recipeFormData.status">
                <el-radio :label="1">启用</el-radio>
                <el-radio :label="0">禁用</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="recipeFormData.remark" type="textarea" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeRecipeDialog">取消</el-button>
          <el-button type="primary" @click="enterRecipeDialog">确定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 工艺步骤对话框 -->
    <el-dialog v-model="stepDialogVisible" :title="stepDialogTitle" width="50%">
      <el-form ref="stepFormRef" :model="stepFormData" :rules="stepRules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="步骤代码" prop="code">
              <el-input v-model="stepFormData.code" placeholder="如：STEP001" clearable />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="步骤名称" prop="name">
              <el-input v-model="stepFormData.name" placeholder="请输入步骤名称" clearable />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="步骤序号" prop="sequence">
              <el-input-number v-model="stepFormData.sequence" :min="1" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="状态" prop="status">
              <el-radio-group v-model="stepFormData.status">
                <el-radio :label="1">启用</el-radio>
                <el-radio :label="0">禁用</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="stepFormData.remark" type="textarea" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeStepDialog">取消</el-button>
          <el-button type="primary" @click="enterStepDialog">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getRecipeList,
  createRecipe,
  updateRecipe,
  deleteRecipe as apiDeleteRecipe,
  getProcessStepList,
  createProcessStep,
  updateProcessStep,
  deleteProcessStep,
  getEquipmentList
} from '@/api/spc/master'

const activeTab = ref('recipe')
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const recipeData = ref([])
const searchInfo = reactive({})

const stepPage = ref(1)
const stepPageSize = ref(10)
const stepTotal = ref(0)
const stepData = ref([])

const equipmentList = ref([])
const processStepList = ref([])

// Recipe dialog
const recipeDialogVisible = ref(false)
const recipeDialogTitle = ref('')
const recipeFormRef = ref(null)
const recipeFormData = ref({
  code: '',
  name: '',
  equipmentId: null,
  processStepId: null,
  version: '',
  status: 1,
  remark: ''
})

const recipeRules = reactive({
  code: [{ required: true, message: '请输入配方代码', trigger: 'blur' }],
  name: [{ required: true, message: '请输入配方名称', trigger: 'blur' }],
  equipmentId: [{ required: true, message: '请选择设备', trigger: 'change' }]
})

// Process Step dialog
const stepDialogVisible = ref(false)
const stepDialogTitle = ref('')
const stepFormRef = ref(null)
const stepFormData = ref({
  code: '',
  name: '',
  sequence: 1,
  status: 1,
  remark: ''
})

const stepRules = reactive({
  code: [{ required: true, message: '请输入步骤代码', trigger: 'blur' }],
  name: [{ required: true, message: '请输入步骤名称', trigger: 'blur' }]
})

const loadEquipmentList = async () => {
  try {
    const res = await getEquipmentList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      equipmentList.value = res.data.list || []
    }
  } catch (error) {
    console.error('加载设备列表失败', error)
  }
}

const loadProcessStepList = async () => {
  try {
    const res = await getProcessStepList({ page: 1, pageSize: 100 })
    if (res.code === 0) {
      processStepList.value = res.data.list || []
    }
  } catch (error) {
    console.error('加载工艺步骤列表失败', error)
  }
}

const getEquipmentName = (id) => {
  const eqp = equipmentList.value.find(e => e.ID === id)
  return eqp ? `${eqp.code} - ${eqp.name}` : id
}

const getProcessStepName = (id) => {
  const step = processStepList.value.find(s => s.ID === id)
  return step ? `${step.code} - ${step.name}` : id || '-'
}

const getRecipeData = async() => {
  const table = await getRecipeList({ page: page.value, pageSize: pageSize.value, ...searchInfo })
  if (table.code === 0) {
    recipeData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const getStepData = async() => {
  const table = await getProcessStepList({ page: stepPage.value, pageSize: stepPageSize.value })
  if (table.code === 0) {
    stepData.value = table.data.list
    stepTotal.value = table.data.total
    stepPage.value = table.data.page
    stepPageSize.value = table.data.pageSize
  }
}

onMounted(() => {
  loadEquipmentList()
  loadProcessStepList()
  getRecipeData()
  getStepData()
})

const onSubmit = () => {
  page.value = 1
  getRecipeData()
}

const onReset = () => {
  searchInfo.equipmentId = null
  searchInfo.code = ''
  onSubmit()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getRecipeData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getRecipeData()
}

const handleStepSizeChange = (val) => {
  stepPageSize.value = val
  getStepData()
}

const handleStepCurrentChange = (val) => {
  stepPage.value = val
  getStepData()
}

// Recipe CRUD
const openRecipeDialog = (type, row) => {
  recipeDialogVisible.value = true
  if (type === 'add') {
    recipeDialogTitle.value = '新增配方'
    recipeFormData.value = {
      code: '',
      name: '',
      equipmentId: null,
      processStepId: null,
      version: '',
      status: 1,
      remark: ''
    }
  } else {
    recipeDialogTitle.value = '编辑配方'
    recipeFormData.value = { ...row }
  }
}

const closeRecipeDialog = () => {
  recipeDialogVisible.value = false
  recipeFormRef.value?.resetFields()
}

const enterRecipeDialog = async() => {
  recipeFormRef.value?.validate(async(valid) => {
    if (valid) {
      let res
      if (recipeFormData.value.ID) {
        res = await updateRecipe(recipeFormData.value)
      } else {
        res = await createRecipe(recipeFormData.value)
      }
      if (res.code === 0) {
        ElMessage.success(recipeFormData.value.ID ? '编辑成功' : '创建成功')
        closeRecipeDialog()
        getRecipeData()
      } else {
        ElMessage.error(res.msg || '操作失败')
      }
    }
  })
}

const deleteRecipe = async(row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await apiDeleteRecipe({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getRecipeData()
    } else {
      ElMessage.error(res.msg || '删除失败')
    }
  })
}

// ProcessStep CRUD
const openStepDialog = (type, row) => {
  stepDialogVisible.value = true
  if (type === 'add') {
    stepDialogTitle.value = '新增工艺步骤'
    stepFormData.value = {
      code: '',
      name: '',
      sequence: 1,
      status: 1,
      remark: ''
    }
  } else {
    stepDialogTitle.value = '编辑工艺步骤'
    stepFormData.value = { ...row }
  }
}

const closeStepDialog = () => {
  stepDialogVisible.value = false
  stepFormRef.value?.resetFields()
}

const enterStepDialog = async() => {
  stepFormRef.value?.validate(async(valid) => {
    if (valid) {
      let res
      if (stepFormData.value.ID) {
        res = await updateProcessStep(stepFormData.value)
      } else {
        res = await createProcessStep(stepFormData.value)
      }
      if (res.code === 0) {
        ElMessage.success(stepFormData.value.ID ? '编辑成功' : '创建成功')
        closeStepDialog()
        getStepData()
        loadProcessStepList() // Reload for recipe dropdown
      } else {
        ElMessage.error(res.msg || '操作失败')
      }
    }
  })
}

const deleteStep = async(row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteProcessStep({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getStepData()
      loadProcessStepList()
    } else {
      ElMessage.error(res.msg || '删除失败')
    }
  })
}
</script>

<style scoped>
</style>
