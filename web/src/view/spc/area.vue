<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-search-box">
        <el-form :inline="true">
          <el-form-item label="厂区">
            <el-select v-model="searchInfo.siteId" placeholder="选择厂区" clearable @change="getList">
              <el-option
                v-for="site in siteList"
                :key="site.ID"
                :label="site.name"
                :value="site.ID"
              />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="getList">查询</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="gva-btn-list">
        <el-button type="primary" @click="handleAdd">新增区域</el-button>
      </div>
      <el-table :data="tableData" style="width: 100%" v-loading="loading">
        <el-table-column prop="code" label="区域代码" width="150" />
        <el-table-column prop="name" label="区域名称" width="200" />
        <el-table-column prop="siteCode" label="厂区" width="100" />
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'info'">
              {{ row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="remark" label="备注" show-overflow-tooltip />
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" link @click="handleDelete(row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="40%">
      <el-form :model="form" label-width="100px">
        <el-form-item label="区域代码">
          <el-input v-model="form.code" placeholder="请输入区域代码" />
        </el-form-item>
        <el-form-item label="区域名称">
          <el-input v-model="form.name" placeholder="请输入区域名称" />
        </el-form-item>
        <el-form-item label="所属厂区">
          <el-select v-model="form.siteId" placeholder="请选择厂区" clearable>
            <el-option
              v-for="site in siteList"
              :key="site.ID"
              :label="site.name"
              :value="site.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
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
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getAreaList,
  createArea,
  updateArea,
  deleteArea,
  getSiteList
} from '@/api/spc/master'

const loading = ref(false)
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const searchInfo = ref({
  siteId: null
})

const siteList = ref([])

const dialogVisible = ref(false)
const dialogTitle = ref('新增区域')
const form = ref({
  code: '',
  name: '',
  siteId: null,
  status: 1,
  remark: ''
})

const getList = async () => {
  loading.value = true
  try {
    const params = {
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    }
    const res = await getAreaList(params)
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
  dialogTitle.value = '新增区域'
  form.value = {
    code: '',
    name: '',
    siteId: null,
    status: 1,
    remark: ''
  }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogTitle.value = '编辑区域'
  form.value = { ...row }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (form.value.ID) {
      await updateArea(form.value)
      ElMessage.success('更新成功')
    } else {
      await createArea(form.value)
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
    await ElMessageBox.confirm('确定要删除该区域吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await deleteArea({ ID: row.ID })
    ElMessage.success('删除成功')
    await getList()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
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

const loadSites = async () => {
  try {
    const res = await getSiteList({ page: 1, pageSize: 100, status: 1 })
    if (res.code === 0) {
      siteList.value = res.data.list || []
    }
  } catch (error) {
    console.error('加载厂区失败', error)
  }
}

onMounted(() => {
  getList()
  loadSites()
})
</script>

<style scoped>
.gva-table-box {
  padding: 20px;
}
.gva-search-box {
  margin-bottom: 15px;
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
