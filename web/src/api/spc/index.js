import service from '@/utils/request'

const request = (url, method, data, params) =>
  service({ url, method, data, params })

const get = (url, params) => request(url, 'get', undefined, params)
const post = (url, data, params) => request(url, 'post', data, params)
const put = (url, data) => request(url, 'put', data)
const del = (url, data) => request(url, 'delete', data)

export const createSpcCrud = (name) => ({
  getList: (params) => get(`/spc/get${name}List`, params),
  create: (data) => post(`/spc/create${name}`, data),
  update: (data) => put(`/spc/update${name}`, data),
  remove: (data) => del(`/spc/delete${name}`, data),
  find: (params) => get(`/spc/find${name}`, { id: params?.id || params?.ID, ID: params?.ID || params?.id })
})

export const siteApi = {
  getList: (params) => get('/spc/getSiteList', params),
  create: (data) => post('/spc/site', data),
  update: (data) => put('/spc/site', data),
  remove: (data) => del('/spc/site', data),
  find: (params) => get('/spc/site', { id: params?.id || params?.ID, ID: params?.ID || params?.id })
}

export const areaApi = createSpcCrud('Area')
export const equipmentApi = createSpcCrud('Equipment')
export const chamberApi = createSpcCrud('Chamber')
export const technologyApi = createSpcCrud('Technology')
export const productApi = createSpcCrud('Product')
export const processStepApi = createSpcCrud('ProcessStep')
export const recipeApi = createSpcCrud('Recipe')
export const lotApi = createSpcCrud('Lot')
export const waferApi = createSpcCrud('Wafer')
export const parameterApi = createSpcCrud('Parameter')
export const specApi = createSpcCrud('Spec')
export const chartApi = createSpcCrud('Chart')
export const controlLimitApi = createSpcCrud('ControlLimit')
export const ruleApi = createSpcCrud('Rule')
export const ocapApi = createSpcCrud('Ocap')
export const ocapExecutionApi = {
  ...createSpcCrud('OcapExecution'),
  start: (data) => post('/spc/startOcapExecution', data),
  complete: (data) => post('/spc/completeOcapExecution', data)
}

export const collectData = (data) => post('/spc/collect', data)
export const getAlarmList = (params) => get('/spc/getAlarmList', params)
export const acknowledgeAlarm = (data) => post('/spc/acknowledgeAlarm', data)
export const closeAlarm = (data) => post('/spc/closeAlarm', data)
export const getAlarmStatistics = (params) => get('/spc/getAlarmStatistics', params)
export const getDashboardOverview = (params) => get('/spc/getDashboardOverview', params)
export const getChartRuntime = (params) => get('/spc/getChartRuntime', params)
export const getSampleList = (params) => get('/spc/getSampleList', params)
export const findSample = (params) => get('/spc/findSample', params)
export const getMeasurementList = (params) => get('/spc/getMeasurementList', params)
export const calculateCapability = (data) => post('/spc/calculateCapability', data)
export const getCapabilityList = (params) => get('/spc/getCapabilityList', params)
export const getCapabilityHistory = (params) => get('/spc/getCapabilityHistory', params)

export const getSiteList = siteApi.getList
export const createSite = siteApi.create
export const updateSite = siteApi.update
export const deleteSite = siteApi.remove
export const getSite = siteApi.find
export const getAreaList = areaApi.getList
export const createArea = areaApi.create
export const updateArea = areaApi.update
export const deleteArea = areaApi.remove
export const getEquipmentList = equipmentApi.getList
export const createEquipment = equipmentApi.create
export const updateEquipment = equipmentApi.update
export const deleteEquipment = equipmentApi.remove
export const getChartList = chartApi.getList
export const createChart = chartApi.create
export const updateChart = chartApi.update
export const deleteChart = chartApi.remove
export const getRuleCatalog = (params) => get('/spc/getRuleCatalog', params)
export const saveChartRules = (data) => post('/spc/saveChartRules', data)
export const calculateControlLimit = (data) => post('/spc/calculateControlLimit', data)
