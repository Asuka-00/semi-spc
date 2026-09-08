<template>
  <div class="space-y-4">
    <SpcPageHeader :title="$t('spc.site.title')" :subtitle="$t('spc.site.subtitle')" />
    <el-tabs v-model="activeTab" class="spc-panel !px-4 !pt-2">
      <el-tab-pane :label="$t('spc.site.siteTab')" name="site">
        <div class="gva-search-box !bg-transparent !p-0 !my-3">
          <el-form :inline="true" :model="site.searchInfo">
            <el-form-item :label="$t('common.keyword')">
              <el-input v-model="site.searchInfo.keyword" :placeholder="`${$t('common.code')} / ${$t('common.name')}`" clearable @keyup.enter="site.onSubmit" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="search" @click="site.onSubmit">{{ $t('common.query') }}</el-button>
              <el-button icon="refresh" @click="site.onReset()">{{ $t('common.reset') }}</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="gva-btn-list">
          <el-button type="primary" icon="plus" @click="openSite('add')">{{ $t('spc.site.addSite') }}</el-button>
        </div>
        <el-table v-loading="site.loading" :data="site.tableData" row-key="ID">
          <el-table-column :label="$t('common.code')" prop="code" min-width="120" />
          <el-table-column :label="$t('common.name')" prop="name" min-width="160" />
          <el-table-column :label="$t('spc.site.timezone')" prop="timezone" min-width="140" />
          <el-table-column :label="$t('common.status')" width="100">
            <template #default="{ row }">
              <el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column :label="$t('common.remark')" prop="remark" min-width="160" show-overflow-tooltip />
          <el-table-column :label="$t('common.action')" fixed="right" width="160">
            <template #default="{ row }">
              <el-button type="primary" link icon="edit" @click="openSite('edit', row)">{{ $t('common.edit') }}</el-button>
              <el-button type="danger" link icon="delete" @click="site.deleteRow(row, $t('spc.site.siteTab'))">{{ $t('common.delete') }}</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination mt-4">
          <el-pagination :current-page="site.page" :page-size="site.pageSize" :page-sizes="[10, 30, 50, 100]" :total="site.total" layout="total, sizes, prev, pager, next" @current-change="site.handleCurrentChange" @size-change="site.handleSizeChange" />
        </div>
      </el-tab-pane>

      <el-tab-pane :label="$t('spc.site.areaTab')" name="area">
        <div class="gva-search-box !bg-transparent !p-0 !my-3">
          <el-form :inline="true" :model="area.searchInfo">
            <el-form-item :label="$t('spc.site.parentSite')">
              <el-select v-model="area.searchInfo.siteId" :placeholder="$t('spc.site.allSites')" clearable filterable>
                <el-option v-for="item in siteOptions" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="search" @click="area.onSubmit">{{ $t('common.query') }}</el-button>
              <el-button icon="refresh" @click="area.onReset(['siteId'])">{{ $t('common.reset') }}</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="gva-btn-list">
          <el-button type="primary" icon="plus" @click="openArea('add')">{{ $t('spc.site.addArea') }}</el-button>
        </div>
        <el-table v-loading="area.loading" :data="area.tableData" row-key="ID">
          <el-table-column :label="$t('common.code')" prop="code" min-width="120" />
          <el-table-column :label="$t('common.name')" prop="name" min-width="160" />
          <el-table-column :label="$t('spc.site.parentSite')" min-width="140">
            <template #default="{ row }">{{ row.site?.name || row.siteId }}</template>
          </el-table-column>
          <el-table-column :label="$t('common.status')" width="100">
            <template #default="{ row }">
              <el-tag :type="statusTag(row.status)" round>{{ statusText(row.status, t) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column :label="$t('common.action')" fixed="right" width="160">
            <template #default="{ row }">
              <el-button type="primary" link icon="edit" @click="openArea('edit', row)">{{ $t('common.edit') }}</el-button>
              <el-button type="danger" link icon="delete" @click="area.deleteRow(row, $t('spc.site.areaTab'))">{{ $t('common.delete') }}</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination mt-4">
          <el-pagination :current-page="area.page" :page-size="area.pageSize" :page-sizes="[10, 30, 50, 100]" :total="area.total" layout="total, sizes, prev, pager, next" @current-change="area.handleCurrentChange" @size-change="area.handleSizeChange" />
        </div>
      </el-tab-pane>
    </el-tabs>

    <el-drawer v-model="site.drawerVisible" :title="site.drawerTitle" size="480px" destroy-on-close>
      <el-form :ref="bindFormRef(site)" :model="site.formData" :rules="siteRules" label-width="110px">
        <el-form-item :label="$t('spc.site.siteCode')" prop="code"><el-input v-model="site.formData.code" /></el-form-item>
        <el-form-item :label="$t('spc.site.siteName')" prop="name"><el-input v-model="site.formData.name" /></el-form-item>
        <el-form-item :label="$t('spc.site.timezone')" prop="timezone">
          <el-select v-model="site.formData.timezone" class="w-full">
            <el-option v-for="tz in TIMEZONES" :key="tz.value" :label="tz.label" :value="tz.value" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('common.status')">
          <el-radio-group v-model="site.formData.status">
            <el-radio :label="1">{{ $t('common.enabled') }}</el-radio>
            <el-radio :label="0">{{ $t('common.disabled') }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="site.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="site.closeDrawer">{{ $t('common.cancel') }}</el-button>
        <el-button type="primary" :loading="site.submitting" @click="site.enterDrawer">{{ $t('common.save') }}</el-button>
      </template>
    </el-drawer>

    <el-drawer v-model="area.drawerVisible" :title="area.drawerTitle" size="480px" destroy-on-close>
      <el-form :ref="bindFormRef(area)" :model="area.formData" :rules="areaRules" label-width="110px">
        <el-form-item :label="$t('spc.site.parentSite')" prop="siteId">
          <el-select v-model="area.formData.siteId" class="w-full" filterable>
            <el-option v-for="item in siteOptions" :key="item.ID" :label="`${item.code} · ${item.name}`" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('spc.site.areaCode')" prop="code"><el-input v-model="area.formData.code" /></el-form-item>
        <el-form-item :label="$t('spc.site.areaName')" prop="name"><el-input v-model="area.formData.name" /></el-form-item>
        <el-form-item :label="$t('common.status')">
          <el-radio-group v-model="area.formData.status">
            <el-radio :label="1">{{ $t('common.enabled') }}</el-radio>
            <el-radio :label="0">{{ $t('common.disabled') }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item :label="$t('common.remark')"><el-input v-model="area.formData.remark" type="textarea" :rows="3" /></el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="area.closeDrawer">{{ $t('common.cancel') }}</el-button>
        <el-button type="primary" :loading="area.submitting" @click="area.enterDrawer">{{ $t('common.save') }}</el-button>
      </template>
    </el-drawer>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { siteApi, areaApi } from '@/api/spc'
import { TIMEZONES, requiredRule, selectRule, statusTag, statusText } from './constants'
import { loadOptions, useSpcCrud, bindFormRef } from './composables/useSpcCrud'
import SpcPageHeader from './components/SpcPageHeader.vue'

const { t } = useI18n()
const activeTab = ref('site')
const siteOptions = ref([])
const site = useSpcCrud({ listApi: siteApi.getList, createApi: siteApi.create, updateApi: siteApi.update, deleteApi: siteApi.remove })
const area = useSpcCrud({ listApi: areaApi.getList, createApi: areaApi.create, updateApi: areaApi.update, deleteApi: areaApi.remove })
const siteRules = computed(() => ({ code: requiredRule(t), name: requiredRule(t), timezone: selectRule(t) }))
const areaRules = computed(() => ({ siteId: selectRule(t), code: requiredRule(t), name: requiredRule(t) }))

const openSite = (type, row) => {
  site.openDrawer(type === 'add' ? t('spc.site.addSite') : t('spc.site.editSite'), type === 'add'
    ? { code: '', name: '', timezone: 'Asia/Shanghai', status: 1, remark: '' }
    : row)
}
const openArea = (type, row) => {
  area.openDrawer(type === 'add' ? t('spc.site.addArea') : t('spc.site.editArea'), type === 'add'
    ? { siteId: siteOptions.value[0]?.ID, code: '', name: '', status: 1, remark: '' }
    : row)
}

onMounted(async () => {
  site.getTableData()
  area.getTableData()
  siteOptions.value = await loadOptions(siteApi.getList)
})
</script>
