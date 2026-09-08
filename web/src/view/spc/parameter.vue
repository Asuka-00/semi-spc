<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.parameter.title')" :subtitle="$t('spc.parameter.subtitle')" />
    <el-tabs v-model="activeTab" class="spc-panel !px-4 !pt-2">
      <el-tab-pane :label="$t('spc.parameter.paramTab')" name="parameter">
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openParam('add')">{{ $t('spc.parameter.addParam') }}</el-button></div>
        <el-table v-loading="param.loading" :data="param.tableData" row-key="ID">
          <el-table-column :label="$t('common.code')" prop="code" min-width="120" />
          <el-table-column :label="$t('common.name')" prop="name" min-width="150" />
          <el-table-column :label="$t('spc.parameter.unit')" prop="unit" width="90" />
          <el-table-column :label="$t('spc.parameter.dataType')" min-width="110"><template #default="{ row }">{{ $t(`spc.option.dataType.${row.dataType}`, row.dataType) }}</template></el-table-column>
          <el-table-column :label="$t('spc.parameter.sampleLevel')" min-width="110"><template #default="{ row }">{{ $t(`spc.option.sampleLevel.${row.sampleLevel}`, row.sampleLevel) }}</template></el-table-column>
          <el-table-column :label="$t('common.status')" width="90"><template #default="{ row }"><el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag></template></el-table-column>
          <el-table-column :label="$t('common.action')" width="160"><template #default="{ row }"><el-button type="primary" link @click="openParam('edit', row)">{{ $t('common.edit') }}</el-button><el-button type="danger" link @click="param.deleteRow(row, $t('spc.parameter.paramTab'))">{{ $t('common.delete') }}</el-button></template></el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="param.page" :page-size="param.pageSize" :total="param.total" layout="total, prev, pager, next" @current-change="param.handleCurrentChange" @size-change="param.handleSizeChange" /></div>
      </el-tab-pane>
      <el-tab-pane :label="$t('spc.parameter.specTab')" name="spec">
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openSpec('add')">{{ $t('spc.parameter.addSpec') }}</el-button></div>
        <el-table v-loading="spec.loading" :data="spec.tableData" row-key="ID">
          <el-table-column :label="$t('spc.parameter.parameter')" min-width="140"><template #default="{ row }">{{ row.parameter?.name || row.parameterId }}</template></el-table-column>
          <el-table-column :label="$t('spc.parameter.product')" min-width="120"><template #default="{ row }">{{ row.product?.name || row.productId }}</template></el-table-column>
          <el-table-column :label="$t('spc.parameter.version')" prop="version" width="90" />
          <el-table-column label="USL" width="90"><template #default="{ row }">{{ num(row.usl) }}</template></el-table-column>
          <el-table-column :label="$t('spc.parameter.target')" width="90"><template #default="{ row }">{{ num(row.target) }}</template></el-table-column>
          <el-table-column label="LSL" width="90"><template #default="{ row }">{{ num(row.lsl) }}</template></el-table-column>
          <el-table-column :label="$t('common.status')" width="90"><template #default="{ row }"><el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag></template></el-table-column>
          <el-table-column :label="$t('common.action')" width="160"><template #default="{ row }"><el-button type="primary" link @click="openSpec('edit', row)">{{ $t('common.edit') }}</el-button><el-button type="danger" link @click="spec.deleteRow(row, $t('spc.parameter.specTab'))">{{ $t('common.delete') }}</el-button></template></el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="spec.page" :page-size="spec.pageSize" :total="spec.total" layout="total, prev, pager, next" @current-change="spec.handleCurrentChange" @size-change="spec.handleSizeChange" /></div>
      </el-tab-pane>
    </el-tabs>

    <el-drawer v-model="param.drawerVisible" :title="param.drawerTitle" size="500px" destroy-on-close>
      <el-form :ref="bindFormRef(param)" :model="param.formData" :rules="paramRules" label-width="120px">
        <el-form-item :label="$t('spc.parameter.paramCode')" prop="code"><el-input v-model="param.formData.code" /></el-form-item>
        <el-form-item :label="$t('spc.parameter.paramName')" prop="name"><el-input v-model="param.formData.name" /></el-form-item>
        <el-form-item :label="$t('spc.parameter.unit')"><el-input v-model="param.formData.unit" /></el-form-item>
        <el-form-item :label="$t('spc.parameter.dataType')"><el-select v-model="param.formData.dataType" class="w-full"><el-option v-for="item in dataTypes" :key="item.value" :label="item.label" :value="item.value" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.parameter.sampleLevel')"><el-select v-model="param.formData.sampleLevel" class="w-full"><el-option v-for="item in sampleLevels" :key="item.value" :label="item.label" :value="item.value" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.parameter.decimals')"><el-input-number v-model="param.formData.decimalPlaces" :min="0" :max="6" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('common.status')"><el-radio-group v-model="param.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="param.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="param.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="param.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
    <el-drawer v-model="spec.drawerVisible" :title="spec.drawerTitle" size="520px" destroy-on-close>
      <el-form :ref="bindFormRef(spec)" :model="spec.formData" :rules="specRules" label-width="130px">
        <el-form-item :label="$t('spc.parameter.parameter')" prop="parameterId"><el-select v-model="spec.formData.parameterId" class="w-full" filterable><el-option v-for="item in params" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.parameter.product')" prop="productId"><el-select v-model="spec.formData.productId" class="w-full" filterable><el-option v-for="item in products" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.parameter.step')" prop="processStepId"><el-select v-model="spec.formData.processStepId" class="w-full" filterable><el-option v-for="item in steps" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.parameter.equipment')"><el-select v-model="spec.formData.equipmentId" class="w-full" clearable filterable><el-option v-for="item in eqs" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.parameter.version')"><el-input v-model="spec.formData.version" /></el-form-item>
        <el-form-item :label="$t('spc.parameter.usl')"><el-input-number v-model="spec.formData.usl" :precision="4" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('spc.parameter.target')"><el-input-number v-model="spec.formData.target" :precision="4" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('spc.parameter.lsl')"><el-input-number v-model="spec.formData.lsl" :precision="4" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('spc.parameter.effectiveFrom')"><el-date-picker v-model="spec.formData.effectiveFrom" type="datetime" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('spc.parameter.effectiveTo')"><el-date-picker v-model="spec.formData.effectiveTo" type="datetime" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('common.status')"><el-radio-group v-model="spec.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="spec.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="spec.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="spec.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { equipmentApi, parameterApi, processStepApi, productApi, specApi } from '@/api/spc'
import { DATA_TYPE_VALUES, SAMPLE_LEVEL_VALUES, num, optionOf, requiredRule, selectRule, statusTag, statusText } from './constants'
import { bindFormRef, loadOptions, useSpcCrud } from './composables/useSpcCrud'
import SpcPageHeader from './components/SpcPageHeader.vue'

const { t } = useI18n()
const activeTab = ref('parameter')
const params = ref([])
const products = ref([])
const steps = ref([])
const eqs = ref([])
const dataTypes = computed(() => optionOf(t, 'spc.option.dataType', DATA_TYPE_VALUES))
const sampleLevels = computed(() => optionOf(t, 'spc.option.sampleLevel', SAMPLE_LEVEL_VALUES))
const param = useSpcCrud({ listApi: parameterApi.getList, createApi: parameterApi.create, updateApi: parameterApi.update, deleteApi: parameterApi.remove })
const spec = useSpcCrud({ listApi: specApi.getList, createApi: specApi.create, updateApi: specApi.update, deleteApi: specApi.remove })
const paramRules = computed(() => ({ code: requiredRule(t), name: requiredRule(t) }))
const specRules = computed(() => ({ parameterId: selectRule(t), productId: selectRule(t), processStepId: selectRule(t) }))
const openParam = (type, row) => param.openDrawer(type === 'add' ? t('spc.parameter.addParam') : t('spc.parameter.editParam'), type === 'add' ? { code: '', name: '', unit: 'nm', dataType: 'VARIABLE', sampleLevel: 'WAFER', decimalPlaces: 3, status: 1, remark: '' } : row)
const openSpec = (type, row) => spec.openDrawer(type === 'add' ? t('spc.parameter.addSpec') : t('spc.parameter.editSpec'), type === 'add' ? { parameterId: params.value[0]?.ID, productId: products.value[0]?.ID, processStepId: steps.value[0]?.ID, equipmentId: undefined, version: 'v1', usl: undefined, target: undefined, lsl: undefined, status: 1, remark: '' } : row)
onMounted(async () => {
  param.getTableData(); spec.getTableData()
  params.value = await loadOptions(parameterApi.getList)
  products.value = await loadOptions(productApi.getList)
  steps.value = await loadOptions(processStepApi.getList)
  eqs.value = await loadOptions(equipmentApi.getList)
})
</script>
