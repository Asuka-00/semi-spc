<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.process.title')" :subtitle="$t('spc.process.subtitle')" />
    <el-tabs v-model="activeTab" class="spc-panel !px-4 !pt-2">
      <el-tab-pane :label="$t('spc.process.tech')" name="tech">
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openTech('add')">{{ $t('spc.process.addTech') }}</el-button></div>
        <el-table v-loading="tech.loading" :data="tech.tableData" row-key="ID">
          <el-table-column :label="$t('common.code')" prop="code" />
          <el-table-column :label="$t('common.name')" prop="name" />
          <el-table-column :label="$t('spc.process.nodeNm')" prop="nodeNm" width="120" />
          <el-table-column :label="$t('common.status')" width="90"><template #default="{ row }"><el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag></template></el-table-column>
          <el-table-column :label="$t('common.action')" width="160"><template #default="{ row }"><el-button type="primary" link @click="openTech('edit', row)">{{ $t('common.edit') }}</el-button><el-button type="danger" link @click="tech.deleteRow(row, $t('spc.process.tech'))">{{ $t('common.delete') }}</el-button></template></el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="tech.page" :page-size="tech.pageSize" :total="tech.total" layout="total, prev, pager, next" @current-change="tech.handleCurrentChange" @size-change="tech.handleSizeChange" /></div>
      </el-tab-pane>
      <el-tab-pane :label="$t('spc.process.product')" name="product">
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openProduct('add')">{{ $t('spc.process.addProduct') }}</el-button></div>
        <el-table v-loading="product.loading" :data="product.tableData" row-key="ID">
          <el-table-column :label="$t('common.code')" prop="code" />
          <el-table-column :label="$t('common.name')" prop="name" />
          <el-table-column :label="$t('spc.process.tech')"><template #default="{ row }">{{ row.technology?.name || row.technologyId }}</template></el-table-column>
          <el-table-column :label="$t('common.status')" width="90"><template #default="{ row }"><el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag></template></el-table-column>
          <el-table-column :label="$t('common.action')" width="160"><template #default="{ row }"><el-button type="primary" link @click="openProduct('edit', row)">{{ $t('common.edit') }}</el-button><el-button type="danger" link @click="product.deleteRow(row, $t('spc.process.product'))">{{ $t('common.delete') }}</el-button></template></el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="product.page" :page-size="product.pageSize" :total="product.total" layout="total, prev, pager, next" @current-change="product.handleCurrentChange" @size-change="product.handleSizeChange" /></div>
      </el-tab-pane>
      <el-tab-pane :label="$t('spc.process.step')" name="step">
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openStep('add')">{{ $t('spc.process.addStep') }}</el-button></div>
        <el-table v-loading="step.loading" :data="step.tableData" row-key="ID">
          <el-table-column :label="$t('common.code')" prop="code" />
          <el-table-column :label="$t('common.name')" prop="name" />
          <el-table-column :label="$t('spc.process.stepType')" prop="stepType" />
          <el-table-column :label="$t('common.status')" width="90"><template #default="{ row }"><el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag></template></el-table-column>
          <el-table-column :label="$t('common.action')" width="160"><template #default="{ row }"><el-button type="primary" link @click="openStep('edit', row)">{{ $t('common.edit') }}</el-button><el-button type="danger" link @click="step.deleteRow(row, $t('spc.process.step'))">{{ $t('common.delete') }}</el-button></template></el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="step.page" :page-size="step.pageSize" :total="step.total" layout="total, prev, pager, next" @current-change="step.handleCurrentChange" @size-change="step.handleSizeChange" /></div>
      </el-tab-pane>
      <el-tab-pane :label="$t('spc.process.recipe')" name="recipe">
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openRecipe('add')">{{ $t('spc.process.addRecipe') }}</el-button></div>
        <el-table v-loading="recipe.loading" :data="recipe.tableData" row-key="ID">
          <el-table-column :label="$t('common.code')" prop="code" />
          <el-table-column :label="$t('common.name')" prop="name" />
          <el-table-column :label="$t('spc.process.version')" prop="version" width="100" />
          <el-table-column :label="$t('spc.equipment.equipment')"><template #default="{ row }">{{ row.equipment?.name || row.equipmentId }}</template></el-table-column>
          <el-table-column :label="$t('spc.process.step')"><template #default="{ row }">{{ row.processStep?.name || row.processStepId }}</template></el-table-column>
          <el-table-column :label="$t('common.action')" width="160"><template #default="{ row }"><el-button type="primary" link @click="openRecipe('edit', row)">{{ $t('common.edit') }}</el-button><el-button type="danger" link @click="recipe.deleteRow(row, $t('spc.process.recipe'))">{{ $t('common.delete') }}</el-button></template></el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="recipe.page" :page-size="recipe.pageSize" :total="recipe.total" layout="total, prev, pager, next" @current-change="recipe.handleCurrentChange" @size-change="recipe.handleSizeChange" /></div>
      </el-tab-pane>
    </el-tabs>

    <el-drawer v-model="tech.drawerVisible" :title="tech.drawerTitle" size="460px" destroy-on-close>
      <el-form :ref="bindFormRef(tech)" :model="tech.formData" :rules="codeNameRules" label-width="120px">
        <el-form-item :label="$t('spc.process.techCode')" prop="code"><el-input v-model="tech.formData.code" /></el-form-item>
        <el-form-item :label="$t('spc.process.techName')" prop="name"><el-input v-model="tech.formData.name" /></el-form-item>
        <el-form-item :label="$t('spc.process.nodeNm')"><el-input-number v-model="tech.formData.nodeNm" :min="0" :precision="2" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('common.status')"><el-radio-group v-model="tech.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="tech.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="tech.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="tech.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
    <el-drawer v-model="product.drawerVisible" :title="product.drawerTitle" size="460px" destroy-on-close>
      <el-form :ref="bindFormRef(product)" :model="product.formData" :rules="productRules" label-width="120px">
        <el-form-item :label="$t('spc.process.tech')" prop="technologyId"><el-select v-model="product.formData.technologyId" class="w-full" filterable><el-option v-for="item in techs" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.process.productCode')" prop="code"><el-input v-model="product.formData.code" /></el-form-item>
        <el-form-item :label="$t('spc.process.productName')" prop="name"><el-input v-model="product.formData.name" /></el-form-item>
        <el-form-item :label="$t('common.status')"><el-radio-group v-model="product.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="product.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="product.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="product.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
    <el-drawer v-model="step.drawerVisible" :title="step.drawerTitle" size="460px" destroy-on-close>
      <el-form :ref="bindFormRef(step)" :model="step.formData" :rules="codeNameRules" label-width="120px">
        <el-form-item :label="$t('spc.process.stepCode')" prop="code"><el-input v-model="step.formData.code" /></el-form-item>
        <el-form-item :label="$t('spc.process.stepName')" prop="name"><el-input v-model="step.formData.name" /></el-form-item>
        <el-form-item :label="$t('spc.process.stepType')"><el-input v-model="step.formData.stepType" /></el-form-item>
        <el-form-item :label="$t('common.status')"><el-radio-group v-model="step.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="step.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="step.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="step.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
    <el-drawer v-model="recipe.drawerVisible" :title="recipe.drawerTitle" size="520px" destroy-on-close>
      <el-form :ref="bindFormRef(recipe)" :model="recipe.formData" :rules="recipeRules" label-width="120px">
        <el-form-item :label="$t('spc.equipment.equipment')" prop="equipmentId"><el-select v-model="recipe.formData.equipmentId" class="w-full" filterable><el-option v-for="item in eqs" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.process.step')" prop="processStepId"><el-select v-model="recipe.formData.processStepId" class="w-full" filterable><el-option v-for="item in steps" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.process.recipeCode')" prop="code"><el-input v-model="recipe.formData.code" /></el-form-item>
        <el-form-item :label="$t('spc.process.recipeName')" prop="name"><el-input v-model="recipe.formData.name" /></el-form-item>
        <el-form-item :label="$t('spc.process.version')"><el-input v-model="recipe.formData.version" /></el-form-item>
        <el-form-item :label="$t('common.status')"><el-radio-group v-model="recipe.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="recipe.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="recipe.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="recipe.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { equipmentApi, processStepApi, productApi, recipeApi, technologyApi } from '@/api/spc'
import { requiredRule, selectRule, statusTag, statusText } from './constants'
import { bindFormRef, loadOptions, useSpcCrud } from './composables/useSpcCrud'
import SpcPageHeader from './components/SpcPageHeader.vue'

const { t } = useI18n()
const activeTab = ref('tech')
const techs = ref([])
const steps = ref([])
const eqs = ref([])
const tech = useSpcCrud({ listApi: technologyApi.getList, createApi: technologyApi.create, updateApi: technologyApi.update, deleteApi: technologyApi.remove })
const product = useSpcCrud({ listApi: productApi.getList, createApi: productApi.create, updateApi: productApi.update, deleteApi: productApi.remove })
const step = useSpcCrud({ listApi: processStepApi.getList, createApi: processStepApi.create, updateApi: processStepApi.update, deleteApi: processStepApi.remove })
const recipe = useSpcCrud({ listApi: recipeApi.getList, createApi: recipeApi.create, updateApi: recipeApi.update, deleteApi: recipeApi.remove })
const codeNameRules = computed(() => ({ code: requiredRule(t), name: requiredRule(t) }))
const productRules = computed(() => ({ technologyId: selectRule(t), code: requiredRule(t), name: requiredRule(t) }))
const recipeRules = computed(() => ({ equipmentId: selectRule(t), processStepId: selectRule(t), code: requiredRule(t), name: requiredRule(t) }))

const openTech = (type, row) => tech.openDrawer(type === 'add' ? t('spc.process.addTech') : t('spc.process.editTech'), type === 'add' ? { code: '', name: '', nodeNm: 28, status: 1, remark: '' } : row)
const openProduct = (type, row) => product.openDrawer(type === 'add' ? t('spc.process.addProduct') : t('spc.process.editProduct'), type === 'add' ? { technologyId: techs.value[0]?.ID, code: '', name: '', status: 1, remark: '' } : row)
const openStep = (type, row) => step.openDrawer(type === 'add' ? t('spc.process.addStep') : t('spc.process.editStep'), type === 'add' ? { code: '', name: '', stepType: '', status: 1, remark: '' } : row)
const openRecipe = (type, row) => recipe.openDrawer(type === 'add' ? t('spc.process.addRecipe') : t('spc.process.editRecipe'), type === 'add' ? { equipmentId: eqs.value[0]?.ID, processStepId: steps.value[0]?.ID, code: '', name: '', version: '1.0', status: 1, remark: '' } : row)

onMounted(async () => {
  tech.getTableData(); product.getTableData(); step.getTableData(); recipe.getTableData()
  techs.value = await loadOptions(technologyApi.getList)
  steps.value = await loadOptions(processStepApi.getList)
  eqs.value = await loadOptions(equipmentApi.getList)
})
</script>
