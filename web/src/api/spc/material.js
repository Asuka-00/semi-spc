import service from '@/utils/request'

// Lot API
export const createLot = (data) => {
  return service({
    url: '/spc/createLot',
    method: 'post',
    data
  })
}

export const updateLot = (data) => {
  return service({
    url: '/spc/updateLot',
    method: 'put',
    data
  })
}

export const deleteLot = (data) => {
  return service({
    url: '/spc/deleteLot',
    method: 'delete',
    data
  })
}

export const findLot = (params) => {
  return service({
    url: '/spc/findLot',
    method: 'get',
    params
  })
}

export const getLotList = (params) => {
  return service({
    url: '/spc/getLotList',
    method: 'get',
    params
  })
}

export const holdLot = (data) => {
  return service({
    url: '/spc/holdLot',
    method: 'post',
    data
  })
}

export const releaseLot = (data) => {
  return service({
    url: '/spc/releaseLot',
    method: 'post',
    data
  })
}

// Wafer API
export const createWafer = (data) => {
  return service({
    url: '/spc/createWafer',
    method: 'post',
    data
  })
}

export const updateWafer = (data) => {
  return service({
    url: '/spc/updateWafer',
    method: 'put',
    data
  })
}

export const deleteWafer = (data) => {
  return service({
    url: '/spc/deleteWafer',
    method: 'delete',
    data
  })
}

export const findWafer = (params) => {
  return service({
    url: '/spc/findWafer',
    method: 'get',
    params
  })
}

export const getWaferList = (params) => {
  return service({
    url: '/spc/getWaferList',
    method: 'get',
    params
  })
}

// Parameter API
export const createParameter = (data) => {
  return service({
    url: '/spc/createParameter',
    method: 'post',
    data
  })
}

export const updateParameter = (data) => {
  return service({
    url: '/spc/updateParameter',
    method: 'put',
    data
  })
}

export const deleteParameter = (data) => {
  return service({
    url: '/spc/deleteParameter',
    method: 'delete',
    data
  })
}

export const findParameter = (params) => {
  return service({
    url: '/spc/findParameter',
    method: 'get',
    params
  })
}

export const getParameterList = (params) => {
  return service({
    url: '/spc/getParameterList',
    method: 'get',
    params
  })
}

// Spec API
export const createSpec = (data) => {
  return service({
    url: '/spc/createSpec',
    method: 'post',
    data
  })
}

export const updateSpec = (data) => {
  return service({
    url: '/spc/updateSpec',
    method: 'put',
    data
  })
}

export const deleteSpec = (data) => {
  return service({
    url: '/spc/deleteSpec',
    method: 'delete',
    data
  })
}

export const findSpec = (params) => {
  return service({
    url: '/spc/findSpec',
    method: 'get',
    params
  })
}

export const getSpecList = (params) => {
  return service({
    url: '/spc/getSpecList',
    method: 'get',
    params
  })
}
