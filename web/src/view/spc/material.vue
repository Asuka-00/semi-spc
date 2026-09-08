<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.material.title')" :subtitle="$t('spc.material.subtitle')" />
    <el-tabs v-model="activeTab" class="spc-panel !px-4 !pt-2">
      <el-tab-pane :label="$t('spc.material.lot')" name="lot">
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openLot('add')">{{ $t('spc.material.addLot') }}</el-button></div>
        <el-table v-loading="lot.loading" :data="lot.tableData" row-key="ID">
          <el-table-column :label="$t('spc.material.lotId')" prop="lotId" min-width="140" />
          <el-table-column :label="$t('spc.material.lotType')" width="110"><template #default="{ row }">{{ $t(`spc.option.lotType.${row.lotType}`, row.lotType) }}</template></el-table-column>
          <el-table-column :label="$t('spc.material.product')"><template #default="{ row }">{{ row.product?.name || row.productId }}</template></el-table-column>
          <el-table-column :label="$t('spc.material.site')"><template #default="{ row }">{{ row.site?.name || row.siteId }}</template></el-table-column>
          <el-table-column :label="$t('spc.material.qty')" prop="qty" width="80" />
          <el-table-column :label="$t('common.action')" width="160"><template #default="{ row }"><el-button type="primary" link @click="openLot('edit', row)">{{ $t('common.edit') }}</el-button><el-button type="danger" link @click="lot.deleteRow(row, $t('spc.material.lot'))">{{ $t('common.delete') }}</el-button></template></el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="lot.page" :page-size="lot.pageSize" :total="lot.total" layout="total, prev, pager, next" @current-change="lot.handleCurrentChange" @size-change="lot.handleSizeChange" /></div>
      </el-tab-pane>
      <el-tab-pane :label="$t('spc.material.wafer')" name="wafer">
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openWafer('add')">{{ $t('spc.material.addWafer') }}</el-button></div>
        <el-table v-loading="wafer.loading" :data="wafer.tableData" row-key="ID">
          <el-table-column :label="$t('spc.material.waferId')" prop="waferId" min-width="140" />
          <el-table-column :label="$t('spc.material.slot')" prop="slotNo" width="90" />
          <el-table-column :label="$t('spc.material.lot')"><template #default="{ row }">{{ row.lot?.lotId || row.lotId }}</template></el-table-column>
          <el-table-column :label="$t('common.action')" width="160"><template #default="{ row }"><el-button type="primary" link @click="openWafer('edit', row)">{{ $t('common.edit') }}</el-button><el-button type="danger" link @click="wafer.deleteRow(row, $t('spc.material.wafer'))">{{ $t('common.delete') }}</el-button></template></el-table-column>
        </el-table>
        <div class="gva-pagination mt-4"><el-pagination :current-page="wafer.page" :page-size="wafer.pageSize" :total="wafer.total" layout="total, prev, pager, next" @current-change="wafer.handleCurrentChange" @size-change="wafer.handleSizeChange" /></div>
      </el-tab-pane>
    </el-tabs>

    <el-drawer v-model="lot.drawerVisible" :title="lot.drawerTitle" size="480px" destroy-on-close>
      <el-form :ref="bindFormRef(lot)" :model="lot.formData" :rules="lotRules" label-width="110px">
        <el-form-item :label="$t('spc.material.site')" prop="siteId"><el-select v-model="lot.formData.siteId" class="w-full" filterable><el-option v-for="item in sites" :key="item.ID" :label="item.name" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.material.product')" prop="productId"><el-select v-model="lot.formData.productId" class="w-full" filterable><el-option v-for="item in products" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.material.lotId')" prop="lotId"><el-input v-model="lot.formData.lotId" /></el-form-item>
        <el-form-item :label="$t('spc.material.lotType')"><el-select v-model="lot.formData.lotType" class="w-full"><el-option v-for="item in lotTypes" :key="item.value" :label="item.label" :value="item.value" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.material.qty')"><el-input-number v-model="lot.formData.qty" :min="1" :max="25" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('common.status')"><el-radio-group v-model="lot.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="lot.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="lot.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="lot.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
    <el-drawer v-model="wafer.drawerVisible" :title="wafer.drawerTitle" size="480px" destroy-on-close>
      <el-form :ref="bindFormRef(wafer)" :model="wafer.formData" :rules="waferRules" label-width="110px">
        <el-form-item :label="$t('spc.material.lot')" prop="lotId"><el-select v-model="wafer.formData.lotId" class="w-full" filterable><el-option v-for="item in lots" :key="item.ID" :label="item.lotId" :value="item.ID" /></el-select></el-form-item>
        <el-form-item :label="$t('spc.material.waferId')" prop="waferId"><el-input v-model="wafer.formData.waferId" /></el-form-item>
        <el-form-item :label="$t('spc.material.slot')" prop="slotNo"><el-input-number v-model="wafer.formData.slotNo" :min="1" :max="25" class="!w-full" /></el-form-item>
        <el-form-item :label="$t('common.status')"><el-radio-group v-model="wafer.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group></el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="wafer.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="wafer.closeDrawer">{{ $t('common.cancel') }}</el-button><el-button type="primary" @click="wafer.enterDrawer">{{ $t('common.save') }}</el-button></template>
    </el-drawer>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { lotApi, productApi, siteApi, waferApi } from '@/api/spc'
import { LOT_TYPE_VALUES, optionOf, requiredRule, selectRule } from './constants'
import { bindFormRef, loadOptions, useSpcCrud } from './composables/useSpcCrud'
import SpcPageHeader from './components/SpcPageHeader.vue'

const { t } = useI18n()
const activeTab = ref('lot')
const sites = ref([])
const products = ref([])
const lots = ref([])
const lotTypes = computed(() => optionOf(t, 'spc.option.lotType', LOT_TYPE_VALUES))
const lot = useSpcCrud({ listApi: lotApi.getList, createApi: lotApi.create, updateApi: lotApi.update, deleteApi: lotApi.remove })
const wafer = useSpcCrud({ listApi: waferApi.getList, createApi: waferApi.create, updateApi: waferApi.update, deleteApi: waferApi.remove })
const lotRules = computed(() => ({ siteId: selectRule(t), productId: selectRule(t), lotId: requiredRule(t) }))
const waferRules = computed(() => ({ lotId: selectRule(t), waferId: requiredRule(t) }))
const openLot = (type, row) => lot.openDrawer(type === 'add' ? t('spc.material.addLot') : t('spc.material.editLot'), type === 'add' ? { siteId: sites.value[0]?.ID, productId: products.value[0]?.ID, lotId: '', lotType: 'PROD', qty: 25, status: 1, remark: '' } : row)
const openWafer = (type, row) => wafer.openDrawer(type === 'add' ? t('spc.material.addWafer') : t('spc.material.editWafer'), type === 'add' ? { lotId: lots.value[0]?.ID, waferId: '', slotNo: 1, status: 1, remark: '' } : row)
onMounted(async () => {
  lot.getTableData(); wafer.getTableData()
  sites.value = await loadOptions(siteApi.getList)
  products.value = await loadOptions(productApi.getList)
  lots.value = await loadOptions(lotApi.getList)
})
</script>
