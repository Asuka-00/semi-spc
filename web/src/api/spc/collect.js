import service from '@/utils/request'

// 运行时API
export const getChartRuntime = (params) => {
  return service({
    url: '/spc/getChartRuntime',
    method: 'get',
    params
  })
}

export const getMeasurementList = (params) => {
  return service({
    url: '/spc/getMeasurementList',
    method: 'get',
    params
  })
}

// 数据采集
export const collectData = (data) => {
  return service({
    url: '/spc/collect',
    method: 'post',
    data
  })
}

export const collectCsv = (data) => {
  return service({
    url: '/spc/collectCsv',
    method: 'post',
    data,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

// 告警API
export const getAlarmList = (params) => {
  return service({
    url: '/spc/getAlarmList',
    method: 'get',
    params
  })
}

export const acknowledgeAlarm = (data) => {
  return service({
    url: '/spc/acknowledgeAlarm',
    method: 'post',
    data
  })
}

export const closeAlarm = (data) => {
  return service({
    url: '/spc/closeAlarm',
    method: 'post',
    data
  })
}

export const getAlarmStatistics = (params) => {
  return service({
    url: '/spc/getAlarmStatistics',
    method: 'get',
    params
  })
}
