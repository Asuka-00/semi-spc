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

// 腔室管理
export const getChamberList = (data) => {
  return service({
    url: '/spc/getChamberList',
    method: 'get',
    params: data
  })
}

export const createChamber = (data) => {
  return service({
    url: '/spc/createChamber',
    method: 'post',
    data
  })
}

export const updateChamber = (data) => {
  return service({
    url: '/spc/updateChamber',
    method: 'put',
    data
  })
}

export const deleteChamber = (data) => {
  return service({
    url: '/spc/deleteChamber',
    method: 'delete',
    data
  })
}

// 工艺管理
export const getTechnologyList = (data) => {
  return service({
    url: '/spc/getTechnologyList',
    method: 'get',
    params: data
  })
}

export const createTechnology = (data) => {
  return service({
    url: '/spc/createTechnology',
    method: 'post',
    data
  })
}

export const updateTechnology = (data) => {
  return service({
    url: '/spc/updateTechnology',
    method: 'put',
    data
  })
}

export const deleteTechnology = (data) => {
  return service({
    url: '/spc/deleteTechnology',
    method: 'delete',
    data
  })
}

// 产品管理
export const getProductList = (data) => {
  return service({
    url: '/spc/getProductList',
    method: 'get',
    params: data
  })
}

export const createProduct = (data) => {
  return service({
    url: '/spc/createProduct',
    method: 'post',
    data
  })
}

export const updateProduct = (data) => {
  return service({
    url: '/spc/updateProduct',
    method: 'put',
    data
  })
}

export const deleteProduct = (data) => {
  return service({
    url: '/spc/deleteProduct',
    method: 'delete',
    data
  })
}

// 工艺步骤管理
export const getProcessStepList = (data) => {
  return service({
    url: '/spc/getProcessStepList',
    method: 'get',
    params: data
  })
}

export const createProcessStep = (data) => {
  return service({
    url: '/spc/createProcessStep',
    method: 'post',
    data
  })
}

export const updateProcessStep = (data) => {
  return service({
    url: '/spc/updateProcessStep',
    method: 'put',
    data
  })
}

export const deleteProcessStep = (data) => {
  return service({
    url: '/spc/deleteProcessStep',
    method: 'delete',
    data
  })
}

// 配方管理
export const getRecipeList = (data) => {
  return service({
    url: '/spc/getRecipeList',
    method: 'get',
    params: data
  })
}

export const createRecipe = (data) => {
  return service({
    url: '/spc/createRecipe',
    method: 'post',
    data
  })
}

export const updateRecipe = (data) => {
  return service({
    url: '/spc/updateRecipe',
    method: 'put',
    data
  })
}

export const deleteRecipe = (data) => {
  return service({
    url: '/spc/deleteRecipe',
    method: 'delete',
    data
  })
}
