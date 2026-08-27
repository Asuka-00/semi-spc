<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" @click="handleAdd">新增批次</el-button>
      </div>
      <el-table :data="tableData" style="width: 100%" v-loading="loading">
        <el-table-column prop="code" label="批次号" width="150" />
        <el-table-column prop="siteCode" label="厂区" width="100" />
        <el-table-column prop="productCode" label="产品" width="120" />
        <el-table-column prop="waferCount" label="晶圆数" width="100" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 'HELD' ? 'danger' : 'success'">
              {{ row.status === 'HELD' ? 'HOLD' : row.status }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" show-overflow-tooltip />
        <el-table-column label="操作" width="280" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">编辑</el-button>
            <el-button type="info" link @click="openWaferDrawer(row)">晶圆</el-button>
            <el-button type="danger" link @click="handleDelete(row)">删除</el-button>
            <el-button v-if="row.status !== 'HELD'" type="warning" link @click="handleHold(row)">Hold</el-button>
            <el-button v-else type="success" link @click="handleRelease(row)">Release</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="50%">
      <el-form :model="form" label-width="120px">
        <el-form-item label="批次号">
          <el-input v-model="form.code" placeholder="请输入批次号" />
        </el-form-item>
        <el-form-item label="厂区">
          <el-select v-model="form.siteId" placeholder="请选择厂区">
            <el-option label="FAB1" :value="1" />
          </el-select>
        </el-form-item>
        <el-form-item label="产品">
          <el-select v-model="form.productId" placeholder="请选择产品">
            <el-option label="产品A" :value="1" />
          </el-select>
        </el-form-item>
        <el-form-item label="晶圆数">
          <el-input-number v-model="form.waferCount" :min="1" :max="25" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="form.remark" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <!-- Wafer晶圆抽屉 -->
    <el-drawer v-model="waferDrawerVisible" :title="`晶圆管理 - ${currentLot?.code || ''}`" size="60%">
      <div style="padding: 0 20px">
        <div style="margin-bottom: 15px">
          <el-button type="primary" icon="plus" @click="openWaferDialog('add')">新增晶圆</el-button>
        </div>
        <el-table :data="waferData" style="width: 100%">
          <el-table-column label="槽位号" prop="slotNo" width="100" />
          <el-table-column label="晶圆ID" prop="waferId" width="150" />
          <el-table-column label="状态" prop="status" width="100">
            <template #default="scope">
              <el-tag :type="scope.row.status === 'SCRAPPED' ? 'danger' : 'success'">
                {{ scope.row.status || 'NORMAL' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="备注" prop="remark" min-width="150" show-overflow-tooltip />
          <el-table-column label="操作" fixed="right" width="200">
            <template #default="scope">
              <el-button type="primary" link icon="edit" @click="openWaferDialog('edit', scope.row)">编辑</el-button>
              <el-button type="primary" link icon="delete" @click="deleteWafer(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-drawer>

    <!-- Wafer对话框 -->
    <el-dialog v-model="waferDialogVisible" :title="waferDialogTitle" width="50%">
      <el-form ref="waferFormRef" :model="waferFormData" :rules="waferRules" label-width="100px">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="槽位号" prop="slotNo">
              <el-input-number v-model="waferFormData.slotNo" :min="1" :max="25" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="晶圆ID" prop="waferId">
              <el-input v-model="waferFormData.waferId" placeholder="如：W001" clearable />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="waferFormData.remark" type="textarea" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeWaferDialog">取消</el-button>
          <el-button type="primary" @click="enterWaferDialog">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getLotList,
  createLot,
  updateLot,
  deleteLot,
  holdLot,
  releaseLot,
  getWaferList,
  createWafer,
  updateWafer,
  deleteWafer as apiDeleteWafer
} from '@/api/spc/material'

const loading = ref(false)
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

const dialogVisible = ref(false)
const dialogTitle = ref('新增批次')
const form = ref({
  code: '',
  siteId: null,
  productId: null,
  waferCount: 25,
  status: 'RELEASED',
  remark: ''
})

const getList = async () => {
  loading.value = true
  try {
    const res = await getLotList({ page: page.value, pageSize: pageSize.value })
    if (res.code === 0) {
      tableData.value = res.data.list || []
      total.value = res.data.total || 0
    }
  } catch (error) {
    ElMessage.error('获取列表失败')
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  dialogTitle.value = '新增批次'
  form.value = {
    code: '',
    siteId: null,
    productId: null,
    waferCount: 25,
    status: 'RELEASED',
    remark: ''
  }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogTitle.value = '编辑批次'
  form.value = { ...row }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (form.value.ID) {
      await updateLot(form.value)
      ElMessage.success('更新成功')
    } else {
      await createLot(form.value)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    await getList()
  } catch (error) {
    ElMessage.error(error.message || '操作失败')
  }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该批次吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await deleteLot({ ID: row.ID })
    ElMessage.success('删除成功')
    await getList()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

const handleHold = async (row) => {
  try {
    const result = await ElMessageBox.prompt('请输入Hold原因（必填）', 'Hold批次', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputPattern: /.+/,
      inputErrorMessage: 'Hold原因不能为空'
    })
    if (!result.value || result.value.trim() === '') {
      ElMessage.error('Hold原因不能为空')
      return
    }
    await holdLot({ ID: row.ID, comment: result.value.trim() })
    ElMessage.success('Hold成功')
    await getList()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || 'Hold失败')
    }
  }
}

const handleRelease = async (row) => {
  try {
    const result = await ElMessageBox.prompt('请输入Release原因（必填）', 'Release批次', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputPattern: /.+/,
      inputErrorMessage: 'Release原因不能为空'
    })
    if (!result.value || result.value.trim() === '') {
      ElMessage.error('Release原因不能为空')
      return
    }
    await releaseLot({ ID: row.ID, comment: result.value.trim() })
    ElMessage.success('Release成功')
    await getList()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error(error.message || 'Release失败')
    }
  }
}

const handleCurrentChange = (val) => {
  page.value = val
  getList()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getList()
}

onMounted(() => {
  getList()
})

// Wafer管理
const waferDrawerVisible = ref(false)
const currentLot = ref(null)
const waferData = ref([])

const waferDialogVisible = ref(false)
const waferDialogTitle = ref('')
const waferFormRef = ref(null)
const waferFormData = ref({
  lotId: null,
  slotNo: 1,
  waferId: '',
  remark: ''
})

const waferRules = {
  slotNo: [{ required: true, message: '请输入槽位号', trigger: 'blur' }],
  waferId: [{ required: true, message: '请输入晶圆ID', trigger: 'blur' }]
}

const openWaferDrawer = async(lot) => {
  currentLot.value = lot
  waferDrawerVisible.value = true
  await loadWaferData()
}

const loadWaferData = async() => {
  if (!currentLot.value) return
  try {
    const res = await getWaferList({ page: 1, pageSize: 100, lotId: currentLot.value.ID })
    if (res.code === 0) {
      waferData.value = res.data.list || []
    }
  } catch (error) {
    ElMessage.error('加载晶圆列表失败')
  }
}

const openWaferDialog = (type, row) => {
  waferDialogVisible.value = true
  if (type === 'add') {
    waferDialogTitle.value = '新增晶圆'
    waferFormData.value = {
      lotId: currentLot.value.ID,
      slotNo: 1,
      waferId: '',
      remark: ''
    }
  } else {
    waferDialogTitle.value = '编辑晶圆'
    waferFormData.value = { ...row }
  }
}

const closeWaferDialog = () => {
  waferDialogVisible.value = false
  waferFormRef.value?.resetFields()
}

const enterWaferDialog = async() => {
  waferFormRef.value?.validate(async(valid) => {
    if (valid) {
      try {
        let res
        if (waferFormData.value.ID) {
          res = await updateWafer(waferFormData.value)
        } else {
          res = await createWafer(waferFormData.value)
        }
        if (res.code === 0) {
          ElMessage.success(waferFormData.value.ID ? '编辑成功' : '创建成功')
          closeWaferDialog()
          loadWaferData()
        } else {
          ElMessage.error(res.msg || '操作失败')
        }
      } catch (error) {
        ElMessage.error(error.message || '操作失败')
      }
    }
  })
}

const deleteWafer = async(row) => {
  try {
    await ElMessageBox.confirm('确定要删除该晶圆吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    const res = await apiDeleteWafer({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      loadWaferData()
    } else {
      ElMessage.error(res.msg || '删除失败')
    }
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}
</script>

<style scoped>
.gva-table-box {
  padding: 20px;
}
.gva-btn-list {
  margin-bottom: 15px;
}
.gva-pagination {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}
</style>
