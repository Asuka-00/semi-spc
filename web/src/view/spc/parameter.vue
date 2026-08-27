<template>
  <!-- SPC参数规格管理 -->
  <div>
    <el-tabs v-model="activeTab">
      <el-tab-pane label="参数管理" name="parameter">
        <div class="gva-table-box">
          <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openParamDialog('add')">新增参数</el-button>
          </div>
          <el-table :data="paramList" row-key="ID">
            <el-table-column label="ID" prop="ID" width="80" />
            <el-table-column label="参数代码" prop="code" min-width="120" />
            <el-table-column label="参数名称" prop="name" min-width="150" />
            <el-table-column label="单位" prop="unit" width="100" />
            <el-table-column label="类型" prop="param_type" width="100" />
            <el-table-column label="状态" prop="status" width="100">
              <template #default="scope">
                <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
                  {{ scope.row.status === 1 ? '启用' : '禁用' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-button type="primary" link icon="edit" @click="openParamDialog('edit', scope.row)">编辑</el-button>
                <el-button type="primary" link icon="delete" @click="deleteParam(scope.row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-tab-pane>

      <el-tab-pane label="规格管理" name="spec">
        <div class="gva-table-box">
          <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openSpecDialog('add')">新增规格</el-button>
          </div>
          <el-table :data="specList" row-key="ID">
            <el-table-column label="ID" prop="ID" width="80" />
            <el-table-column label="参数" prop="parameter_id" width="100" />
            <el-table-column label="产品" prop="product_id" width="100" />
            <el-table-column label="USL" prop="usl" width="100" />
            <el-table-column label="目标值" prop="target" width="100" />
            <el-table-column label="LSL" prop="lsl" width="100" />
            <el-table-column label="状态" prop="status" width="100">
              <template #default="scope">
                <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
                  {{ scope.row.status === 1 ? '启用' : '禁用' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-button type="primary" link icon="edit" @click="openSpecDialog('edit', scope.row)">编辑</el-button>
                <el-button type="primary" link icon="delete" @click="deleteSpec(scope.row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-tab-pane>
    </el-tabs>

    <el-dialog v-model="paramDialogVisible" :title="paramDialogTitle" width="50%">
      <el-form ref="paramFormRef" :model="paramForm" label-width="100px">
        <el-form-item label="参数代码" prop="code">
          <el-input v-model="paramForm.code" placeholder="请输入参数代码" />
        </el-form-item>
        <el-form-item label="参数名称" prop="name">
          <el-input v-model="paramForm.name" placeholder="请输入参数名称" />
        </el-form-item>
        <el-form-item label="单位" prop="unit">
          <el-input v-model="paramForm.unit" placeholder="如: nm, ℃, Ω" />
        </el-form-item>
        <el-form-item label="参数类型" prop="param_type">
          <el-select v-model="paramForm.param_type">
            <el-option label="连续型" value="CONTINUOUS" />
            <el-option label="离散型" value="DISCRETE" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="paramForm.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="paramDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveParam">确定</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="specDialogVisible" :title="specDialogTitle" width="50%">
      <el-form ref="specFormRef" :model="specForm" label-width="100px">
        <el-form-item label="参数ID">
          <el-input-number v-model="specForm.parameter_id" :min="1" style="width: 100%" />
        </el-form-item>
        <el-form-item label="产品ID">
          <el-input-number v-model="specForm.product_id" :min="1" style="width: 100%" />
        </el-form-item>
        <el-form-item label="上规格限(USL)">
          <el-input-number v-model="specForm.usl" style="width: 100%" />
        </el-form-item>
        <el-form-item label="目标值">
          <el-input-number v-model="specForm.target" style="width: 100%" />
        </el-form-item>
        <el-form-item label="下规格限(LSL)">
          <el-input-number v-model="specForm.lsl" style="width: 100%" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="specForm.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="specDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveSpec">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

const activeTab = ref('parameter')
const paramList = ref([])
const specList = ref([])

const paramDialogVisible = ref(false)
const paramDialogTitle = ref('')
const paramFormRef = ref(null)
const paramForm = ref({
  code: '',
  name: '',
  unit: '',
  param_type: 'CONTINUOUS',
  status: 1
})

const specDialogVisible = ref(false)
const specDialogTitle = ref('')
const specFormRef = ref(null)
const specForm = ref({
  parameter_id: null,
  product_id: null,
  usl: null,
  target: null,
  lsl: null,
  status: 1
})

const openParamDialog = (type, row) => {
  paramDialogVisible.value = true
  paramDialogTitle.value = type === 'add' ? '新增参数' : '编辑参数'
  if (type === 'edit') {
    paramForm.value = { ...row }
  } else {
    paramForm.value = { code: '', name: '', unit: '', param_type: 'CONTINUOUS', status: 1 }
  }
}

const openSpecDialog = (type, row) => {
  specDialogVisible.value = true
  specDialogTitle.value = type === 'add' ? '新增规格' : '编辑规格'
  if (type === 'edit') {
    specForm.value = { ...row }
  } else {
    specForm.value = { parameter_id: null, product_id: null, usl: null, target: null, lsl: null, status: 1 }
  }
}

const saveParam = () => {
  ElMessage.success('参数保存成功')
  paramDialogVisible.value = false
}

const saveSpec = () => {
  ElMessage.success('规格保存成功')
  specDialogVisible.value = false
}

const deleteParam = (row) => {
  ElMessage.success('参数删除成功')
}

const deleteSpec = (row) => {
  ElMessage.success('规格删除成功')
}
</script>

<style scoped>
</style>
