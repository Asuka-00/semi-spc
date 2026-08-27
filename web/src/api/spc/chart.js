import service from '@/utils/request'

// 控制图管理
export const getChartList = (data) => {
  return service({
    url: '/spc/getChartList',
    method: 'get',
    params: data
  })
}

export const createChart = (data) => {
  return service({
    url: '/spc/createChart',
    method: 'post',
    data
  })
}

export const updateChart = (data) => {
  return service({
    url: '/spc/updateChart',
    method: 'put',
    data
  })
}

export const deleteChart = (data) => {
  return service({
    url: '/spc/deleteChart',
    method: 'delete',
    data
  })
}

export const findChart = (params) => {
  return service({
    url: '/spc/findChart',
    method: 'get',
    params
  })
}
