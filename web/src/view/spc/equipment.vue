<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.equipment.title')" :subtitle="$t('spc.equipment.subtitle')" />
    <el-tabs v-model="activeTab" class="spc-panel !px-4 !pt-2">
      <el-tab-pane :label="$t('spc.equipment.eqpTab')" name="eqp">
        <div class="gva-search-box !bg-transparent !p-0 !my-3">
          <el-form :inline="true" :model="eqp.searchInfo">
            <el-form-item :label="$t('common.keyword')"><el-input v-model="eqp.searchInfo.keyword" :placeholder="`${$t('common.code')} / ${$t('common.name')}`" clearable /></el-form-item>
            <el-form-item :label="$t('spc.equipment.site')">
              <el-select v-model="eqp.searchInfo.siteId" clearable filterable :placeholder="$t('common.all')">
                <el-option v-for="item in sites" :key="item.ID" :label="item.name" :value="item.ID" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="search" @click="eqp.onSubmit">{{ $t('common.query') }}</el-button>
              <el-button icon="refresh" @click="eqp.onReset()">{{ $t('common.reset') }}</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="gva-btn-list">
          <el-button type="primary" icon="plus" @click="openEqp('add')">{{ $t('spc.equipment.addEqp') }}</el-button>
        </div>
        <el-table v-loading="eqp.loading" :data="eqp.tableData" row-key="ID">
          <el-table-column :label="$t('common.code')" prop="code" min-width="120" />
          <el-table-column :label="$t('common.name')" prop="name" min-width="150" />
          <el-table-column :label="$t('spc.equipment.type')" min-width="110">
            <template #default="{ row }">{{ $t(`spc.option.eqpType.${row.eqpType}`, row.eqpType) }}</template>
          </el-table-column>
          <el-table-column :label="$t('spc.equipment.area')" min-width="120"><template #default="{ row }">{{ row.area?.name || '-' }}</template></el-table-column>
          <el-table-column :label="$t('spc.equipment.vendor')" prop="vendor" min-width="120" />
          <el-table-column :label="$t('common.status')" width="90">
            <template #default="{ row }"><el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag></template>
          </el-table-column>
          <el-table-column :label="$t('common.action')" fixed="right" width="160">
            <template #default="{ row }">
              <el-button type="primary" link @click="openEqp('edit', row)">{{ $t('common.edit') }}</el-button>
              <el-button type="danger" link @click="eqp.deleteRow(row, $t('spc.equipment.equipment'))">{{ $t('common.delete') }}</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination mt-4">
          <el-pagination :current-page="eqp.page" :page-size="eqp.pageSize" :total="eqp.total" layout="total, prev, pager, next" @current-change="eqp.handleCurrentChange" @size-change="eqp.handleSizeChange" />
        </div>
      </el-tab-pane>
      <el-tab-pane :label="$t('spc.equipment.chTab')" name="ch">
        <div class="gva-search-box !bg-transparent !p-0 !my-3">
          <el-form :inline="true">
            <el-form-item :label="$t('spc.equipment.equipment')">
              <el-select v-model="ch.searchInfo.equipmentId" clearable filterable :placeholder="$t('spc.equipment.allEqp')">
                <el-option v-for="item in equipments" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="search" @click="ch.onSubmit">{{ $t('common.query') }}</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="gva-btn-list"><el-button type="primary" icon="plus" @click="openCh('add')">{{ $t('spc.equipment.addCh') }}</el-button></div>
        <el-table v-loading="ch.loading" :data="ch.tableData" row-key="ID">
          <el-table-column :label="$t('common.code')" prop="code" />
          <el-table-column :label="$t('common.name')" prop="name" />
          <el-table-column :label="$t('spc.equipment.equipment')"><template #default="{ row }">{{ row.equipment?.name || row.equipmentId }}</template></el-table-column>
          <el-table-column :label="$t('common.status')" width="90">
            <template #default="{ row }"><el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag></template>
          </el-table-column>
          <el-table-column :label="$t('common.action')" width="160">
            <template #default="{ row }">
              <el-button type="primary" link @click="openCh('edit', row)">{{ $t('common.edit') }}</el-button>
              <el-button type="danger" link @click="ch.deleteRow(row, $t('spc.equipment.chTab'))">{{ $t('common.delete') }}</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination mt-4">
          <el-pagination :current-page="ch.page" :page-size="ch.pageSize" :total="ch.total" layout="total, prev, pager, next" @current-change="ch.handleCurrentChange" @size-change="ch.handleSizeChange" />
        </div>
      </el-tab-pane>
    </el-tabs>

    <el-drawer v-model="eqp.drawerVisible" :title="eqp.drawerTitle" size="520px" destroy-on-close>
      <el-form :ref="bindFormRef(eqp)" :model="eqp.formData" :rules="eqpRules" label-width="110px">
        <el-form-item :label="$t('spc.equipment.site')" prop="siteId">
          <el-select v-model="eqp.formData.siteId" class="w-full" filterable @change="onSiteChange">
            <el-option v-for="item in sites" :key="item.ID" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('spc.equipment.area')" prop="areaId">
          <el-select v-model="eqp.formData.areaId" class="w-full" filterable>
            <el-option v-for="item in filteredAreas" :key="item.ID" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('spc.equipment.eqpCode')" prop="code"><el-input v-model="eqp.formData.code" /></el-form-item>
        <el-form-item :label="$t('spc.equipment.eqpName')" prop="name"><el-input v-model="eqp.formData.name" /></el-form-item>
        <el-form-item :label="$t('spc.equipment.type')" prop="eqpType">
          <el-select v-model="eqp.formData.eqpType" class="w-full">
            <el-option v-for="item in eqpTypes" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('spc.equipment.vendor')"><el-input v-model="eqp.formData.vendor" /></el-form-item>
        <el-form-item :label="$t('common.status')">
          <el-radio-group v-model="eqp.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group>
        </el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="eqp.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="eqp.closeDrawer">{{ $t('common.cancel') }}</el-button>
        <el-button type="primary" :loading="eqp.submitting" @click="eqp.enterDrawer">{{ $t('common.save') }}</el-button>
      </template>
    </el-drawer>

    <el-drawer v-model="ch.drawerVisible" :title="ch.drawerTitle" size="480px" destroy-on-close>
      <el-form :ref="bindFormRef(ch)" :model="ch.formData" :rules="chRules" label-width="110px">
        <el-form-item :label="$t('spc.equipment.equipment')" prop="equipmentId">
          <el-select v-model="ch.formData.equipmentId" class="w-full" filterable>
            <el-option v-for="item in equipments" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('spc.equipment.chamberCode')" prop="code"><el-input v-model="ch.formData.code" /></el-form-item>
        <el-form-item :label="$t('spc.equipment.chamberName')" prop="name"><el-input v-model="ch.formData.name" /></el-form-item>
        <el-form-item :label="$t('common.status')">
          <el-radio-group v-model="ch.formData.status"><el-radio :label="1">{{ $t('common.enabled') }}</el-radio><el-radio :label="0">{{ $t('common.disabled') }}</el-radio></el-radio-group>
        </el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="ch.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="ch.closeDrawer">{{ $t('common.cancel') }}</el-button>
        <el-button type="primary" @click="ch.enterDrawer">{{ $t('common.save') }}</el-button>
      </template>
    </el-drawer>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { areaApi, chamberApi, equipmentApi, siteApi } from '@/api/spc'
import { EQP_TYPE_VALUES, optionOf, requiredRule, selectRule, statusTag, statusText } from './constants'
import { bindFormRef, loadOptions, useSpcCrud } from './composables/useSpcCrud'
import SpcPageHeader from './components/SpcPageHeader.vue'

const { t } = useI18n()
const activeTab = ref('eqp')
const sites = ref([])
const areas = ref([])
const equipments = ref([])
const eqpTypes = computed(() => optionOf(t, 'spc.option.eqpType', EQP_TYPE_VALUES))
const eqp = useSpcCrud({ listApi: equipmentApi.getList, createApi: equipmentApi.create, updateApi: equipmentApi.update, deleteApi: equipmentApi.remove })
const ch = useSpcCrud({ listApi: chamberApi.getList, createApi: chamberApi.create, updateApi: chamberApi.update, deleteApi: chamberApi.remove })
const filteredAreas = computed(() => areas.value.filter((item) => !eqp.formData.siteId || item.siteId === eqp.formData.siteId))
const eqpRules = computed(() => ({ siteId: selectRule(t), areaId: selectRule(t), code: requiredRule(t), name: requiredRule(t) }))
const chRules = computed(() => ({ equipmentId: selectRule(t), code: requiredRule(t), name: requiredRule(t) }))
const onSiteChange = () => { eqp.formData.areaId = undefined }
const openEqp = (type, row) => {
  eqp.openDrawer(type === 'add' ? t('spc.equipment.addEqp') : t('spc.equipment.editEqp'), type === 'add'
    ? { siteId: sites.value[0]?.ID, areaId: undefined, code: '', name: '', eqpType: 'METROLOGY', vendor: '', status: 1, remark: '' }
    : row)
}
const openCh = (type, row) => {
  ch.openDrawer(type === 'add' ? t('spc.equipment.addCh') : t('spc.equipment.editCh'), type === 'add'
    ? { equipmentId: equipments.value[0]?.ID, code: '', name: '', status: 1, remark: '' }
    : row)
}
onMounted(async () => {
  eqp.getTableData()
  ch.getTableData()
  sites.value = await loadOptions(siteApi.getList)
  areas.value = await loadOptions(areaApi.getList)
  equipments.value = await loadOptions(equipmentApi.getList)
})
</script>
