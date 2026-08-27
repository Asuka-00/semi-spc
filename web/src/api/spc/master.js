import service from '@/utils/request'

// 厂区管理
export const getSiteList = (data) => {
  return service({
    url: '/spc/getSiteList',
    method: 'get',
    params: data
  })
}

export const createSite = (data) => {
  return service({
    url: '/spc/site',
    method: 'post',
    data
  })
}

export const updateSite = (data) => {
  return service({
    url: '/spc/site',
    method: 'put',
    data
  })
}

export const deleteSite = (data) => {
  return service({
    url: '/spc/site',
    method: 'delete',
    data
  })
}

export const getSite = (params) => {
  return service({
    url: '/spc/site',
    method: 'get',
    params
  })
}

// 区域管理
export const getAreaList = (data) => {
  return service({
    url: '/spc/getAreaList',
    method: 'get',
    params: data
  })
}

export const createArea = (data) => {
  return service({
    url: '/spc/createArea',
    method: 'post',
    data
  })
}

export const updateArea = (data) => {
  return service({
    url: '/spc/updateArea',
    method: 'put',
    data
  })
}

export const deleteArea = (data) => {
  return service({
    url: '/spc/deleteArea',
    method: 'delete',
    data
  })
}

// 设备管理
export const getEquipmentList = (data) => {
  return service({
    url: '/spc/getEquipmentList',
    method: 'get',
    params: data
  })
}

export const createEquipment = (data) => {
  return service({
    url: '/spc/createEquipment',
    method: 'post',
    data
  })
}

export const updateEquipment = (data) => {
  return service({
    url: '/spc/updateEquipment',
    method: 'put',
    data
  })
}

export const deleteEquipment = (data) => {
  return service({
    url: '/spc/deleteEquipment',
    method: 'delete',
    data
  })
}
