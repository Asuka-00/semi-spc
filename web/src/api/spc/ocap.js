import service from '@/utils/request'

// OCAP API
export const createOcap = (data) => {
  return service({
    url: '/spc/createOcap',
    method: 'post',
    data
  })
}

export const updateOcap = (data) => {
  return service({
    url: '/spc/updateOcap',
    method: 'put',
    data
  })
}

export const deleteOcap = (data) => {
  return service({
    url: '/spc/deleteOcap',
    method: 'delete',
    data
  })
}

export const findOcap = (params) => {
  return service({
    url: '/spc/findOcap',
    method: 'get',
    params
  })
}

export const getOcapList = (params) => {
  return service({
    url: '/spc/getOcapList',
    method: 'get',
    params
  })
}

export const startOcap = (data) => {
  return service({
    url: '/spc/startOcap',
    method: 'post',
    data
  })
}

// OCAP Execution API
export const updateOcapExecution = (data) => {
  return service({
    url: '/spc/updateOcapExecution',
    method: 'put',
    data
  })
}

export const getOcapExecutionList = (params) => {
  return service({
    url: '/spc/getOcapExecutionList',
    method: 'get',
    params
  })
}
