package spc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	"github.com/gin-gonic/gin"
)

type ChamberApi struct{}
type TechnologyApi struct{}
type ProductApi struct{}
type ProcessStepApi struct{}
type RecipeApi struct{}
type LotApi struct{}
type WaferApi struct{}
type ParameterApi struct{}
type SpecApi struct{}
type ControlLimitApi struct{}
type RuleApi struct{}
type OcapApi struct{}
type OcapExecutionApi struct{}

func createJSON[T any](c *gin.Context, create func(*T) error) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := create(&item); err != nil {
		failCreate(c, err)
		return
	}
	response.OkWithData(item, c)
}

func updateJSON[T any](c *gin.Context, update func(*T) error) {
	var item T
	if err := c.ShouldBindJSON(&item); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := update(&item); err != nil {
		failUpdate(c, err)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func deleteByID(c *gin.Context, del func(uint) error) {
	id, err := bindJSONID(c)
	if err != nil || id == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}
	if err = del(id); err != nil {
		failDelete(c, err)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// ---- Chamber ----

// CreateSpcChamber
// @Tags SpcChamber @Summary 创建腔室 @Security ApiKeyAuth @Router /spc/createChamber [post]
func (a *ChamberApi) CreateSpcChamber(c *gin.Context) {
	createJSON(c, chamberService.CreateSpcChamber)
}

func (a *ChamberApi) DeleteSpcChamber(c *gin.Context) {
	deleteByID(c, chamberService.DeleteSpcChamber)
}

func (a *ChamberApi) UpdateSpcChamber(c *gin.Context) {
	updateJSON(c, chamberService.UpdateSpcChamber)
}

func (a *ChamberApi) FindSpcChamber(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := chamberService.GetSpcChamber(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}

func (a *ChamberApi) GetSpcChamberList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := chamberService.GetSpcChamberList(page, parseUintQuery(c, "equipmentId"))
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- Technology ----

func (a *TechnologyApi) CreateSpcTechnology(c *gin.Context) {
	createJSON(c, technologyService.CreateSpcTechnology)
}
func (a *TechnologyApi) DeleteSpcTechnology(c *gin.Context) {
	deleteByID(c, technologyService.DeleteSpcTechnology)
}
func (a *TechnologyApi) UpdateSpcTechnology(c *gin.Context) {
	updateJSON(c, technologyService.UpdateSpcTechnology)
}
func (a *TechnologyApi) FindSpcTechnology(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := technologyService.GetSpcTechnology(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *TechnologyApi) GetSpcTechnologyList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := technologyService.GetSpcTechnologyList(page)
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- Product ----

func (a *ProductApi) CreateSpcProduct(c *gin.Context) {
	createJSON(c, productService.CreateSpcProduct)
}
func (a *ProductApi) DeleteSpcProduct(c *gin.Context) {
	deleteByID(c, productService.DeleteSpcProduct)
}
func (a *ProductApi) UpdateSpcProduct(c *gin.Context) {
	updateJSON(c, productService.UpdateSpcProduct)
}
func (a *ProductApi) FindSpcProduct(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := productService.GetSpcProduct(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *ProductApi) GetSpcProductList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := productService.GetSpcProductList(page, parseUintQuery(c, "technologyId"))
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- ProcessStep ----

func (a *ProcessStepApi) CreateSpcProcessStep(c *gin.Context) {
	createJSON(c, processStepService.CreateSpcProcessStep)
}
func (a *ProcessStepApi) DeleteSpcProcessStep(c *gin.Context) {
	deleteByID(c, processStepService.DeleteSpcProcessStep)
}
func (a *ProcessStepApi) UpdateSpcProcessStep(c *gin.Context) {
	updateJSON(c, processStepService.UpdateSpcProcessStep)
}
func (a *ProcessStepApi) FindSpcProcessStep(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := processStepService.GetSpcProcessStep(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *ProcessStepApi) GetSpcProcessStepList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := processStepService.GetSpcProcessStepList(page)
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- Recipe ----

func (a *RecipeApi) CreateSpcRecipe(c *gin.Context) {
	createJSON(c, recipeService.CreateSpcRecipe)
}
func (a *RecipeApi) DeleteSpcRecipe(c *gin.Context) {
	deleteByID(c, recipeService.DeleteSpcRecipe)
}
func (a *RecipeApi) UpdateSpcRecipe(c *gin.Context) {
	updateJSON(c, recipeService.UpdateSpcRecipe)
}
func (a *RecipeApi) FindSpcRecipe(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := recipeService.GetSpcRecipe(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *RecipeApi) GetSpcRecipeList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := recipeService.GetSpcRecipeList(page, parseUintQuery(c, "equipmentId"), parseUintQuery(c, "processStepId"))
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- Lot ----

func (a *LotApi) CreateSpcLot(c *gin.Context) {
	createJSON(c, lotService.CreateSpcLot)
}
func (a *LotApi) DeleteSpcLot(c *gin.Context) {
	deleteByID(c, lotService.DeleteSpcLot)
}
func (a *LotApi) UpdateSpcLot(c *gin.Context) {
	updateJSON(c, lotService.UpdateSpcLot)
}
func (a *LotApi) FindSpcLot(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := lotService.GetSpcLot(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *LotApi) GetSpcLotList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := lotService.GetSpcLotList(page, parseUintQuery(c, "siteId"), parseUintQuery(c, "productId"))
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- Wafer ----

func (a *WaferApi) CreateSpcWafer(c *gin.Context) {
	createJSON(c, waferService.CreateSpcWafer)
}
func (a *WaferApi) DeleteSpcWafer(c *gin.Context) {
	deleteByID(c, waferService.DeleteSpcWafer)
}
func (a *WaferApi) UpdateSpcWafer(c *gin.Context) {
	updateJSON(c, waferService.UpdateSpcWafer)
}
func (a *WaferApi) FindSpcWafer(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := waferService.GetSpcWafer(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *WaferApi) GetSpcWaferList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := waferService.GetSpcWaferList(page, parseUintQuery(c, "lotId"))
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- Parameter ----

func (a *ParameterApi) CreateSpcParameter(c *gin.Context) {
	createJSON(c, parameterService.CreateSpcParameter)
}
func (a *ParameterApi) DeleteSpcParameter(c *gin.Context) {
	deleteByID(c, parameterService.DeleteSpcParameter)
}
func (a *ParameterApi) UpdateSpcParameter(c *gin.Context) {
	updateJSON(c, parameterService.UpdateSpcParameter)
}
func (a *ParameterApi) FindSpcParameter(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := parameterService.GetSpcParameter(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *ParameterApi) GetSpcParameterList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := parameterService.GetSpcParameterList(page)
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- Spec ----

func (a *SpecApi) CreateSpcSpec(c *gin.Context) {
	createJSON(c, specService.CreateSpcSpec)
}
func (a *SpecApi) DeleteSpcSpec(c *gin.Context) {
	deleteByID(c, specService.DeleteSpcSpec)
}
func (a *SpecApi) UpdateSpcSpec(c *gin.Context) {
	updateJSON(c, specService.UpdateSpcSpec)
}
func (a *SpecApi) FindSpcSpec(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := specService.GetSpcSpec(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *SpecApi) GetSpcSpecList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := specService.GetSpcSpecList(page, parseUintQuery(c, "parameterId"), parseUintQuery(c, "productId"))
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- ControlLimit ----

func (a *ControlLimitApi) CreateSpcControlLimit(c *gin.Context) {
	createJSON(c, controlLimitService.CreateSpcControlLimit)
}
func (a *ControlLimitApi) DeleteSpcControlLimit(c *gin.Context) {
	deleteByID(c, controlLimitService.DeleteSpcControlLimit)
}
func (a *ControlLimitApi) UpdateSpcControlLimit(c *gin.Context) {
	updateJSON(c, controlLimitService.UpdateSpcControlLimit)
}
func (a *ControlLimitApi) FindSpcControlLimit(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := controlLimitService.GetSpcControlLimit(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *ControlLimitApi) GetSpcControlLimitList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := controlLimitService.GetSpcControlLimitList(page, parseUintQuery(c, "chartId"))
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- Rule ----

func (a *RuleApi) CreateSpcRule(c *gin.Context) {
	createJSON(c, ruleService.CreateSpcRule)
}
func (a *RuleApi) DeleteSpcRule(c *gin.Context) {
	deleteByID(c, ruleService.DeleteSpcRule)
}
func (a *RuleApi) UpdateSpcRule(c *gin.Context) {
	updateJSON(c, ruleService.UpdateSpcRule)
}
func (a *RuleApi) FindSpcRule(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := ruleService.GetSpcRule(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *RuleApi) GetSpcRuleList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := ruleService.GetSpcRuleList(page, parseUintQuery(c, "chartId"))
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

// ---- OCAP ----

func (a *OcapApi) CreateSpcOcap(c *gin.Context) {
	createJSON(c, ocapService.CreateSpcOcap)
}
func (a *OcapApi) DeleteSpcOcap(c *gin.Context) {
	deleteByID(c, ocapService.DeleteSpcOcap)
}
func (a *OcapApi) UpdateSpcOcap(c *gin.Context) {
	updateJSON(c, ocapService.UpdateSpcOcap)
}
func (a *OcapApi) FindSpcOcap(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := ocapService.GetSpcOcap(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *OcapApi) GetSpcOcapList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := ocapService.GetSpcOcapList(page, parseUintQuery(c, "chartId"))
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

func (a *OcapExecutionApi) CreateSpcOcapExecution(c *gin.Context) {
	createJSON(c, ocapExecutionService.CreateSpcOcapExecution)
}
func (a *OcapExecutionApi) UpdateSpcOcapExecution(c *gin.Context) {
	updateJSON(c, ocapExecutionService.UpdateSpcOcapExecution)
}
func (a *OcapExecutionApi) FindSpcOcapExecution(c *gin.Context) {
	id, _ := bindQueryID(c)
	item, err := ocapExecutionService.GetSpcOcapExecution(id)
	if err != nil {
		failGet(c, err)
		return
	}
	response.OkWithData(item, c)
}
func (a *OcapExecutionApi) GetSpcOcapExecutionList(c *gin.Context) {
	page := parsePage(c)
	list, total, err := ocapExecutionService.GetSpcOcapExecutionList(page, parseUintQuery(c, "alarmId"))
	if err != nil {
		failGet(c, err)
		return
	}
	okPage(c, list, total, page)
}

func (a *OcapExecutionApi) StartSpcOcapExecution(c *gin.Context) {
	var req struct {
		ID    int    `json:"ID"`
		Id    int    `json:"id"`
		Owner string `json:"owner"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	id := uint(req.ID)
	if id == 0 {
		id = uint(req.Id)
	}
	if id == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}
	if err := ocapExecutionService.StartSpcOcapExecution(id, req.Owner); err != nil {
		failUpdate(c, err)
		return
	}
	response.OkWithMessage("已开始执行", c)
}

func (a *OcapExecutionApi) CompleteSpcOcapExecution(c *gin.Context) {
	var req struct {
		ID      int    `json:"ID"`
		Id      int    `json:"id"`
		Comment string `json:"comment"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	id := uint(req.ID)
	if id == 0 {
		id = uint(req.Id)
	}
	if id == 0 {
		response.FailWithMessage("ID不能为空", c)
		return
	}
	if err := ocapExecutionService.CompleteSpcOcapExecution(id, req.Comment); err != nil {
		failUpdate(c, err)
		return
	}
	response.OkWithMessage("已完成", c)
}

var _ = spc.SpcChamber{}
