import { reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n'

export function useSpcCrud({ listApi, createApi, updateApi, deleteApi, extraParams } = {}) {
  const { t } = useI18n()
  const state = reactive({
    page: 1,
    pageSize: 10,
    total: 0,
    tableData: [],
    loading: false,
    searchInfo: {},
    drawerVisible: false,
    drawerTitle: '',
    formRef: null,
    formData: {},
    submitting: false
  })

  const mergedParams = () => ({
    page: state.page,
    pageSize: state.pageSize,
    ...state.searchInfo,
    ...(typeof extraParams === 'function' ? extraParams() : extraParams || {})
  })

  const getTableData = async () => {
    if (!listApi) return
    state.loading = true
    try {
      const res = await listApi(mergedParams())
      if (res.code === 0) {
        state.tableData = res.data?.list || []
        state.total = res.data?.total || 0
        state.page = res.data?.page || state.page
        state.pageSize = res.data?.pageSize || state.pageSize
      }
    } finally {
      state.loading = false
    }
  }

  const onSubmit = () => {
    state.page = 1
    getTableData()
  }

  const onReset = (fields = []) => {
    if (!fields.length) {
      Object.keys(state.searchInfo).forEach((key) => {
        state.searchInfo[key] = undefined
      })
    } else {
      fields.forEach((key) => {
        state.searchInfo[key] = undefined
      })
    }
    onSubmit()
  }

  const handleSizeChange = (val) => {
    state.pageSize = val
    getTableData()
  }

  const handleCurrentChange = (val) => {
    state.page = val
    getTableData()
  }

  const openDrawer = (title, data) => {
    state.drawerTitle = title
    state.formData = { ...data }
    state.drawerVisible = true
  }

  const closeDrawer = () => {
    state.drawerVisible = false
  }

  const enterDrawer = async () => {
    const run = async () => {
      state.submitting = true
      try {
        const isEdit = Boolean(state.formData.ID)
        const res = isEdit ? await updateApi(state.formData) : await createApi(state.formData)
        if (res.code === 0) {
          ElMessage.success(isEdit ? t('common.saved') : t('common.created'))
          closeDrawer()
          getTableData()
        }
      } finally {
        state.submitting = false
      }
    }
    if (state.formRef?.validate) {
      await state.formRef.validate(async (valid) => {
        if (valid) await run()
      })
      return
    }
    await run()
  }

  const deleteRow = async (row, name) => {
    await ElMessageBox.confirm(t('common.deleteConfirm', { name: name || t('common.keyword') }), t('common.deleteTitle'), {
      type: 'warning',
      confirmButtonText: t('common.delete'),
      cancelButtonText: t('common.cancel')
    })
    const res = await deleteApi({ ID: row.ID, id: row.ID })
    if (res.code === 0) {
      ElMessage.success(t('common.deleted'))
      getTableData()
    }
  }

  return Object.assign(state, {
    getTableData,
    onSubmit,
    onReset,
    handleSizeChange,
    handleCurrentChange,
    openDrawer,
    closeDrawer,
    enterDrawer,
    deleteRow
  })
}

export async function loadOptions(api, params = { page: 1, pageSize: 200 }) {
  const res = await api(params)
  if (res.code === 0) return res.data?.list || []
  return []
}

export const bindFormRef = (state) => (el) => {
  state.formRef = el
}
