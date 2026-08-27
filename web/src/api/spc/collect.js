import service from '@/utils/request'

// 数据采集
export const collectData = (data) => {
  return service({
    url: '/spc/collect',
    method: 'post',
    data
  })
}

// 告警管理
export const getAlarmList = (data) => {
  return service({
    url: '/spc/getAlarmList',
    method: 'get',
    params: data
  })
}

export const acknowledgeAlarm = (data) => {
  return service({
    url: '/spc/acknowledgeAlarm',
    method: 'post',
    data,
    params: { remark: data.remark || '' }
  })
}

export const closeAlarm = (data) => {
  return service({
    url: '/spc/closeAlarm',
    method: 'post',
    data,
    params: { remark: data.remark || '' }
  })
}

export const getAlarmStatistics = (params) => {
  return service({
    url: '/spc/getAlarmStatistics',
    method: 'get',
    params
  })
}
