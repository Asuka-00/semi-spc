import service from '@/utils/request'

// 运行时数据API
export const getChartRuntime = (params) => {
  return service({
    url: '/spc/getChartRuntime',
    method: 'get',
    params
  })
}

// 能力分析API
export const getCapability = (params) => {
  return service({
    url: '/spc/getCapability',
    method: 'get',
    params
  })
}

export const getCapabilityHistory = (params) => {
  return service({
    url: '/spc/getCapabilityList',
    method: 'get',
    params
  })
}

// 控制限计算API
export const calculateLimits = (data) => {
  return service({
    url: '/spc/calculateLimits',
    method: 'post',
    data
  })
}

// Dashboard API
export const getDashboard = () => {
  return service({
    url: '/spc/getDashboard',
    method: 'get'
  })
}
