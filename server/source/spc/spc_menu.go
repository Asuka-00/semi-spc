package spc

import (
	"context"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	sysService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type initSpcMenu struct{}

func init() {
	sysService.RegisterInit(sysService.InitOrderInternal+50, &initSpcMenu{})
}

func (i *initSpcMenu) InitializerName() string {
	return "spc_menu"
}

func (i *initSpcMenu) MigrateTable(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (i *initSpcMenu) TableCreated(ctx context.Context) bool {
	return true
}

func (i *initSpcMenu) InitializeData(ctx context.Context) (next context.Context, err error) {
	db := global.GVA_DB
	adminAuthorityID := uint(888)
	adminID := strconv.Itoa(int(adminAuthorityID))

	var parent system.SysBaseMenu
	err = db.Where("name = ?", "spc").First(&parent).Error
	if err != nil {
		parent = system.SysBaseMenu{
			ParentId:  0,
			Path:      "spc",
			Name:      "spc",
			Hidden:    false,
			Component: "view/spc/index.vue",
			Sort:      90,
			Meta: system.Meta{
				Title: "SPC统计过程控制",
				Icon:  "monitor",
			},
		}
		if err = db.Create(&parent).Error; err != nil {
			return ctx, err
		}
	}

	subMenus := []system.SysBaseMenu{
		{Path: "dashboard", Name: "spcDashboard", Component: "view/spc/dashboard.vue", Sort: 1, Meta: system.Meta{Title: "过程总览", Icon: "data-board"}},
		{Path: "site", Name: "spcSite", Component: "view/spc/site.vue", Sort: 2, Meta: system.Meta{Title: "厂区与区域", Icon: "office-building"}},
		{Path: "equipment", Name: "spcEquipment", Component: "view/spc/equipment.vue", Sort: 3, Meta: system.Meta{Title: "设备与腔室", Icon: "platform"}},
		{Path: "process", Name: "spcProcess", Component: "view/spc/process.vue", Sort: 4, Meta: system.Meta{Title: "工艺与产品", Icon: "cpu"}},
		{Path: "material", Name: "spcMaterial", Component: "view/spc/material.vue", Sort: 5, Meta: system.Meta{Title: "批次物料", Icon: "box"}},
		{Path: "parameter", Name: "spcParameter", Component: "view/spc/parameter.vue", Sort: 6, Meta: system.Meta{Title: "参数规格", Icon: "data-line"}},
		{Path: "chart", Name: "spcChart", Component: "view/spc/chart.vue", Sort: 7, Meta: system.Meta{Title: "控制图配置", Icon: "data-analysis"}},
		{Path: "runtime", Name: "spcRuntime", Component: "view/spc/runtime.vue", Sort: 8, Meta: system.Meta{Title: "实时监控", Icon: "odometer"}},
		{Path: "collect", Name: "spcCollect", Component: "view/spc/collect.vue", Sort: 9, Meta: system.Meta{Title: "数据采集", Icon: "upload"}},
		{Path: "alarm", Name: "spcAlarm", Component: "view/spc/alarm.vue", Sort: 10, Meta: system.Meta{Title: "告警中心", Icon: "bell"}},
		{Path: "ocap", Name: "spcOcap", Component: "view/spc/ocap.vue", Sort: 11, Meta: system.Meta{Title: "OCAP行动", Icon: "finished"}},
		{Path: "capability", Name: "spcCapability", Component: "view/spc/capability.vue", Sort: 12, Meta: system.Meta{Title: "能力分析", Icon: "pie-chart"}},
	}

	menuIDs := []uint{parent.ID}
	for _, item := range subMenus {
		var exist system.SysBaseMenu
		e := db.Where("name = ?", item.Name).First(&exist).Error
		if e != nil {
			item.ParentId = parent.ID
			item.Hidden = false
			if err = db.Create(&item).Error; err != nil {
				return ctx, err
			}
			menuIDs = append(menuIDs, item.ID)
			continue
		}
		_ = db.Model(&exist).Updates(map[string]interface{}{
			"parent_id": parent.ID,
			"path":      item.Path,
			"component": item.Component,
			"sort":      item.Sort,
			"title":     item.Meta.Title,
			"icon":      item.Meta.Icon,
			"hidden":    false,
		}).Error
		menuIDs = append(menuIDs, exist.ID)
	}

	for _, menuID := range menuIDs {
		var count int64
		db.Model(&system.SysAuthorityMenu{}).
			Where("sys_base_menu_id = ? AND sys_authority_authority_id = ?", strconv.Itoa(int(menuID)), adminID).
			Count(&count)
		if count == 0 {
			if err = db.Create(&system.SysAuthorityMenu{
				MenuId:      strconv.Itoa(int(menuID)),
				AuthorityId: adminID,
			}).Error; err != nil {
				return ctx, err
			}
		}
	}

	apis := spcAPIList()
	for _, apiItem := range apis {
		var exist system.SysApi
		e := db.Where("path = ? AND method = ?", apiItem.Path, apiItem.Method).First(&exist).Error
		if e != nil {
			if err = db.Create(&apiItem).Error; err != nil {
				return ctx, err
			}
		}
	}

	if enforcer := utils.GetCasbin(); enforcer != nil {
		rules := make([][]string, 0, len(apis))
		for _, apiItem := range apis {
			rules = append(rules, []string{adminID, apiItem.Path, apiItem.Method})
		}
		_, _ = enforcer.AddPolicies(rules)
	}

	global.GVA_LOG.Info("SPC菜单和权限初始化完成")
	return ctx, nil
}

func (i *initSpcMenu) DataInserted(ctx context.Context) bool {
	db := global.GVA_DB
	if db == nil {
		db, _ = ctx.Value("db").(*gorm.DB)
	}
	if db == nil {
		return false
	}
	required := []string{"spc", "spcDashboard", "spcProcess", "spcMaterial", "spcOcap", "spcRuntime"}
	for _, name := range required {
		var count int64
		db.Model(&system.SysBaseMenu{}).Where("name = ?", name).Count(&count)
		if count == 0 {
			return false
		}
	}
	var apiCount int64
	db.Model(&system.SysApi{}).Where("path = ? AND method = ?", "/spc/getDashboardOverview", "GET").Count(&apiCount)
	return apiCount > 0
}

func spcAPIList() []system.SysApi {
	type kv struct{ p, d, m string }
	items := []kv{
		{"/spc/getSiteList", "获取厂区列表", "GET"},
		{"/spc/site", "创建厂区", "POST"},
		{"/spc/site", "更新厂区", "PUT"},
		{"/spc/site", "删除厂区", "DELETE"},
		{"/spc/site", "获取厂区", "GET"},
		{"/spc/getAreaList", "获取区域列表", "GET"},
		{"/spc/createArea", "创建区域", "POST"},
		{"/spc/updateArea", "更新区域", "PUT"},
		{"/spc/deleteArea", "删除区域", "DELETE"},
		{"/spc/findArea", "查询区域", "GET"},
		{"/spc/getEquipmentList", "获取设备列表", "GET"},
		{"/spc/createEquipment", "创建设备", "POST"},
		{"/spc/updateEquipment", "更新设备", "PUT"},
		{"/spc/deleteEquipment", "删除设备", "DELETE"},
		{"/spc/findEquipment", "查询设备", "GET"},
		{"/spc/getChamberList", "获取腔室列表", "GET"},
		{"/spc/createChamber", "创建腔室", "POST"},
		{"/spc/updateChamber", "更新腔室", "PUT"},
		{"/spc/deleteChamber", "删除腔室", "DELETE"},
		{"/spc/findChamber", "查询腔室", "GET"},
		{"/spc/getTechnologyList", "获取技术节点列表", "GET"},
		{"/spc/createTechnology", "创建技术节点", "POST"},
		{"/spc/updateTechnology", "更新技术节点", "PUT"},
		{"/spc/deleteTechnology", "删除技术节点", "DELETE"},
		{"/spc/findTechnology", "查询技术节点", "GET"},
		{"/spc/getProductList", "获取产品列表", "GET"},
		{"/spc/createProduct", "创建产品", "POST"},
		{"/spc/updateProduct", "更新产品", "PUT"},
		{"/spc/deleteProduct", "删除产品", "DELETE"},
		{"/spc/findProduct", "查询产品", "GET"},
		{"/spc/getProcessStepList", "获取工艺步骤列表", "GET"},
		{"/spc/createProcessStep", "创建工艺步骤", "POST"},
		{"/spc/updateProcessStep", "更新工艺步骤", "PUT"},
		{"/spc/deleteProcessStep", "删除工艺步骤", "DELETE"},
		{"/spc/findProcessStep", "查询工艺步骤", "GET"},
		{"/spc/getRecipeList", "获取配方列表", "GET"},
		{"/spc/createRecipe", "创建配方", "POST"},
		{"/spc/updateRecipe", "更新配方", "PUT"},
		{"/spc/deleteRecipe", "删除配方", "DELETE"},
		{"/spc/findRecipe", "查询配方", "GET"},
		{"/spc/getLotList", "获取批次列表", "GET"},
		{"/spc/createLot", "创建批次", "POST"},
		{"/spc/updateLot", "更新批次", "PUT"},
		{"/spc/deleteLot", "删除批次", "DELETE"},
		{"/spc/findLot", "查询批次", "GET"},
		{"/spc/getWaferList", "获取晶圆列表", "GET"},
		{"/spc/createWafer", "创建晶圆", "POST"},
		{"/spc/updateWafer", "更新晶圆", "PUT"},
		{"/spc/deleteWafer", "删除晶圆", "DELETE"},
		{"/spc/findWafer", "查询晶圆", "GET"},
		{"/spc/getParameterList", "获取参数列表", "GET"},
		{"/spc/createParameter", "创建参数", "POST"},
		{"/spc/updateParameter", "更新参数", "PUT"},
		{"/spc/deleteParameter", "删除参数", "DELETE"},
		{"/spc/findParameter", "查询参数", "GET"},
		{"/spc/getSpecList", "获取规格列表", "GET"},
		{"/spc/createSpec", "创建规格", "POST"},
		{"/spc/updateSpec", "更新规格", "PUT"},
		{"/spc/deleteSpec", "删除规格", "DELETE"},
		{"/spc/findSpec", "查询规格", "GET"},
		{"/spc/getChartList", "获取控制图列表", "GET"},
		{"/spc/createChart", "创建控制图", "POST"},
		{"/spc/updateChart", "更新控制图", "PUT"},
		{"/spc/deleteChart", "删除控制图", "DELETE"},
		{"/spc/findChart", "查询控制图", "GET"},
		{"/spc/getControlLimitList", "获取控制限列表", "GET"},
		{"/spc/createControlLimit", "创建控制限", "POST"},
		{"/spc/updateControlLimit", "更新控制限", "PUT"},
		{"/spc/deleteControlLimit", "删除控制限", "DELETE"},
		{"/spc/findControlLimit", "查询控制限", "GET"},
		{"/spc/getRuleList", "获取规则列表", "GET"},
		{"/spc/createRule", "创建规则", "POST"},
		{"/spc/updateRule", "更新规则", "PUT"},
		{"/spc/deleteRule", "删除规则", "DELETE"},
		{"/spc/findRule", "查询规则", "GET"},
		{"/spc/collect", "数据采集", "POST"},
		{"/spc/getSampleList", "获取样本列表", "GET"},
		{"/spc/findSample", "查询样本", "GET"},
		{"/spc/getMeasurementList", "获取测量值", "GET"},
		{"/spc/getAlarmList", "获取告警列表", "GET"},
		{"/spc/acknowledgeAlarm", "确认告警", "POST"},
		{"/spc/closeAlarm", "关闭告警", "POST"},
		{"/spc/getAlarmStatistics", "告警统计", "GET"},
		{"/spc/getOcapList", "获取OCAP列表", "GET"},
		{"/spc/createOcap", "创建OCAP", "POST"},
		{"/spc/updateOcap", "更新OCAP", "PUT"},
		{"/spc/deleteOcap", "删除OCAP", "DELETE"},
		{"/spc/findOcap", "查询OCAP", "GET"},
		{"/spc/getOcapExecutionList", "获取OCAP执行列表", "GET"},
		{"/spc/createOcapExecution", "创建OCAP执行", "POST"},
		{"/spc/updateOcapExecution", "更新OCAP执行", "PUT"},
		{"/spc/findOcapExecution", "查询OCAP执行", "GET"},
		{"/spc/startOcapExecution", "开始OCAP执行", "POST"},
		{"/spc/completeOcapExecution", "完成OCAP执行", "POST"},
		{"/spc/calculateCapability", "计算过程能力", "POST"},
		{"/spc/getCapabilityList", "获取能力分析列表", "GET"},
		{"/spc/getCapabilityHistory", "获取能力分析历史", "GET"},
		{"/spc/getDashboardOverview", "仪表盘总览", "GET"},
		{"/spc/getChartRuntime", "控制图运行时数据", "GET"},
		{"/spc/getRuleCatalog", "规则目录", "GET"},
		{"/spc/saveChartRules", "保存控制图规则", "POST"},
	}
	apis := make([]system.SysApi, 0, len(items))
	for _, it := range items {
		apis = append(apis, system.SysApi{Path: it.p, Description: it.d, ApiGroup: "spc", Method: it.m})
	}
	return apis
}
