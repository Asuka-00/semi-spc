package source

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type initSpcMenu struct{}

// auto initialize data
func init() {
	RegisterInit(InitOrderMenu+1, &initSpcMenu{})
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

	// 检查是否已有SPC菜单
	var count int64
	db.Model(&system.SysBaseMenu{}).Where("name = ?", "spc").Count(&count)
	if count > 0 {
		return ctx, nil
	}

	adminAuthorityID := uint(888)

	// 一级菜单 - SPC统计过程控制
	spcMenu := system.SysBaseMenu{
		ParentId:  "0",
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
	if err = db.Create(&spcMenu).Error; err != nil {
		return ctx, err
	}

	// 二级菜单
	subMenus := []system.SysBaseMenu{
		// 主数据管理
		{
			ParentId:  spcMenu.ID,
			Path:      "site",
			Name:      "spcSite",
			Hidden:    false,
			Component: "view/spc/site.vue",
			Sort:      1,
			Meta: system.Meta{
				Title: "厂区管理",
				Icon:  "office-building",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "area",
			Name:      "spcArea",
			Hidden:    false,
			Component: "view/spc/area.vue",
			Sort:      2,
			Meta: system.Meta{
				Title: "区域管理",
				Icon:  "grid",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "equipment",
			Name:      "spcEquipment",
			Hidden:    false,
			Component: "view/spc/equipment.vue",
			Sort:      3,
			Meta: system.Meta{
				Title: "设备管理",
				Icon:  "platform",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "technology",
			Name:      "spcTechnology",
			Hidden:    false,
			Component: "view/spc/technology.vue",
			Sort:      4,
			Meta: system.Meta{
				Title: "工艺技术",
				Icon:  "cpu",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "product",
			Name:      "spcProduct",
			Hidden:    false,
			Component: "view/spc/product.vue",
			Sort:      5,
			Meta: system.Meta{
				Title: "产品管理",
				Icon:  "goods",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "recipe",
			Name:      "spcRecipe",
			Hidden:    false,
			Component: "view/spc/recipe.vue",
			Sort:      6,
			Meta: system.Meta{
				Title: "配方管理",
				Icon:  "document",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "lot",
			Name:      "spcLot",
			Hidden:    false,
			Component: "view/spc/lot.vue",
			Sort:      7,
			Meta: system.Meta{
				Title: "批次管理",
				Icon:  "tickets",
			},
		},
		// 参数与规格
		{
			ParentId:  spcMenu.ID,
			Path:      "parameter",
			Name:      "spcParameter",
			Hidden:    false,
			Component: "view/spc/parameter.vue",
			Sort:      10,
			Meta: system.Meta{
				Title: "参数管理",
				Icon:  "data-line",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "spec",
			Name:      "spcSpec",
			Hidden:    false,
			Component: "view/spc/spec.vue",
			Sort:      11,
			Meta: system.Meta{
				Title: "规格管理",
				Icon:  "files",
			},
		},
		// 控制图
		{
			ParentId:  spcMenu.ID,
			Path:      "chart",
			Name:      "spcChart",
			Hidden:    false,
			Component: "view/spc/chart.vue",
			Sort:      20,
			Meta: system.Meta{
				Title: "控制图配置",
				Icon:  "data-analysis",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "runtime",
			Name:      "spcRuntime",
			Hidden:    false,
			Component: "view/spc/runtime.vue",
			Sort:      21,
			Meta: system.Meta{
				Title: "实时监控",
				Icon:  "monitor",
			},
		},
		// 数据采集
		{
			ParentId:  spcMenu.ID,
			Path:      "collect",
			Name:      "spcCollect",
			Hidden:    false,
			Component: "view/spc/collect.vue",
			Sort:      30,
			Meta: system.Meta{
				Title: "数据采集",
				Icon:  "upload",
			},
		},
		// 告警与OCAP
		{
			ParentId:  spcMenu.ID,
			Path:      "alarm",
			Name:      "spcAlarm",
			Hidden:    false,
			Component: "view/spc/alarm.vue",
			Sort:      40,
			Meta: system.Meta{
				Title: "告警中心",
				Icon:  "bell",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "ocap",
			Name:      "spcOcap",
			Hidden:    false,
			Component: "view/spc/ocap.vue",
			Sort:      41,
			Meta: system.Meta{
				Title: "OCAP管理",
				Icon:  "document-checked",
			},
		},
		// 分析
		{
			ParentId:  spcMenu.ID,
			Path:      "capability",
			Name:      "spcCapability",
			Hidden:    false,
			Component: "view/spc/capability.vue",
			Sort:      50,
			Meta: system.Meta{
				Title: "能力分析",
				Icon:  "pie-chart",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "dashboard",
			Name:      "spcDashboard",
			Hidden:    false,
			Component: "view/spc/dashboard.vue",
			Sort:      51,
			Meta: system.Meta{
				Title: "Dashboard",
				Icon:  "data-board",
			},
		},
	}

	if err = db.Create(&subMenus).Error; err != nil {
		return ctx, err
	}

	// 为admin角色分配菜单
	var menuIDs []string
	menuIDs = append(menuIDs, spcMenu.ID)
	for _, menu := range subMenus {
		menuIDs = append(menuIDs, menu.ID)
	}

	// 创建菜单-角色关联
	for _, menuID := range menuIDs {
		if err = db.Create(&system.SysAuthorityMenu{
			MenuId:      menuID,
			AuthorityId: adminAuthorityID,
		}).Error; err != nil {
			return ctx, err
		}
	}

	// 注册API - 完整的SPC API列表
	apis := []system.SysApi{
		// Site API
		{Path: "/spc/getSiteList", Description: "获取厂区列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createSite", Description: "创建厂区", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateSite", Description: "更新厂区", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteSite", Description: "删除厂区", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findSite", Description: "查询厂区", ApiGroup: "spc", Method: "GET"},
		
		// Area API
		{Path: "/spc/getAreaList", Description: "获取区域列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createArea", Description: "创建区域", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateArea", Description: "更新区域", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteArea", Description: "删除区域", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findArea", Description: "查询区域", ApiGroup: "spc", Method: "GET"},
		
		// Equipment API
		{Path: "/spc/getEquipmentList", Description: "获取设备列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createEquipment", Description: "创建设备", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateEquipment", Description: "更新设备", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteEquipment", Description: "删除设备", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findEquipment", Description: "查询设备", ApiGroup: "spc", Method: "GET"},
		
		// Chamber API
		{Path: "/spc/getChamberList", Description: "获取腔室列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createChamber", Description: "创建腔室", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateChamber", Description: "更新腔室", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteChamber", Description: "删除腔室", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findChamber", Description: "查询腔室", ApiGroup: "spc", Method: "GET"},
		
		// Technology API
		{Path: "/spc/getTechnologyList", Description: "获取工艺列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createTechnology", Description: "创建工艺", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateTechnology", Description: "更新工艺", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteTechnology", Description: "删除工艺", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findTechnology", Description: "查询工艺", ApiGroup: "spc", Method: "GET"},
		
		// Product API
		{Path: "/spc/getProductList", Description: "获取产品列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createProduct", Description: "创建产品", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateProduct", Description: "更新产品", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteProduct", Description: "删除产品", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findProduct", Description: "查询产品", ApiGroup: "spc", Method: "GET"},
		
		// ProcessStep API
		{Path: "/spc/getProcessStepList", Description: "获取工艺步骤列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createProcessStep", Description: "创建工艺步骤", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateProcessStep", Description: "更新工艺步骤", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteProcessStep", Description: "删除工艺步骤", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findProcessStep", Description: "查询工艺步骤", ApiGroup: "spc", Method: "GET"},
		
		// Recipe API
		{Path: "/spc/getRecipeList", Description: "获取配方列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createRecipe", Description: "创建配方", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateRecipe", Description: "更新配方", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteRecipe", Description: "删除配方", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findRecipe", Description: "查询配方", ApiGroup: "spc", Method: "GET"},
		
		// Lot API
		{Path: "/spc/getLotList", Description: "获取批次列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createLot", Description: "创建批次", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateLot", Description: "更新批次", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteLot", Description: "删除批次", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findLot", Description: "查询批次", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/holdLot", Description: "冻结批次", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/releaseLot", Description: "释放批次", ApiGroup: "spc", Method: "POST"},
		
		// Wafer API
		{Path: "/spc/getWaferList", Description: "获取晶圆列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createWafer", Description: "创建晶圆", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateWafer", Description: "更新晶圆", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteWafer", Description: "删除晶圆", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findWafer", Description: "查询晶圆", ApiGroup: "spc", Method: "GET"},
		
		// Parameter API
		{Path: "/spc/getParameterList", Description: "获取参数列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createParameter", Description: "创建参数", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateParameter", Description: "更新参数", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteParameter", Description: "删除参数", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findParameter", Description: "查询参数", ApiGroup: "spc", Method: "GET"},
		
		// Spec API
		{Path: "/spc/getSpecList", Description: "获取规格列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createSpec", Description: "创建规格", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateSpec", Description: "更新规格", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteSpec", Description: "删除规格", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findSpec", Description: "查询规格", ApiGroup: "spc", Method: "GET"},
		
		// Chart API
		{Path: "/spc/getChartList", Description: "获取控制图列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createChart", Description: "创建控制图", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateChart", Description: "更新控制图", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteChart", Description: "删除控制图", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findChart", Description: "查询控制图", ApiGroup: "spc", Method: "GET"},
		
		// ControlLimit API
		{Path: "/spc/getControlLimitList", Description: "获取控制限列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createControlLimit", Description: "创建控制限", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateControlLimit", Description: "更新控制限", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteControlLimit", Description: "删除控制限", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findControlLimit", Description: "查询控制限", ApiGroup: "spc", Method: "GET"},
		
		// Rule API
		{Path: "/spc/getRuleList", Description: "获取规则列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createRule", Description: "创建规则", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateRule", Description: "更新规则", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteRule", Description: "删除规则", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findRule", Description: "查询规则", ApiGroup: "spc", Method: "GET"},
		
		// Sample API
		{Path: "/spc/getSampleList", Description: "获取样本列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createSample", Description: "创建样本", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateSample", Description: "更新样本", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteSample", Description: "删除样本", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findSample", Description: "查询样本", ApiGroup: "spc", Method: "GET"},
		
		// Measurement API
		{Path: "/spc/getMeasurementList", Description: "获取测量值列表", ApiGroup: "spc", Method: "GET"},
		
		// OCAP API
		{Path: "/spc/getOcapList", Description: "获取OCAP模板列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createOcap", Description: "创建OCAP模板", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateOcap", Description: "更新OCAP模板", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteOcap", Description: "删除OCAP模板", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findOcap", Description: "查询OCAP模板", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/startOcap", Description: "启动OCAP", ApiGroup: "spc", Method: "POST"},
		
		// OcapExecution API
		{Path: "/spc/getOcapExecutionList", Description: "获取OCAP执行列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/updateOcapExecution", Description: "更新OCAP执行", ApiGroup: "spc", Method: "PUT"},
		
		// Collect API
		{Path: "/spc/collect", Description: "数据采集", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/collectCsv", Description: "CSV批量采集", ApiGroup: "spc", Method: "POST"},
		
		// Alarm API
		{Path: "/spc/getAlarmList", Description: "获取告警列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/acknowledgeAlarm", Description: "确认告警", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/closeAlarm", Description: "关闭告警", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/getAlarmStatistics", Description: "告警统计", ApiGroup: "spc", Method: "GET"},
		
		// Runtime API
		{Path: "/spc/getChartRuntime", Description: "获取图表运行时数据", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/calculateLimits", Description: "计算控制限", ApiGroup: "spc", Method: "POST"},
		
		// Capability API
		{Path: "/spc/getCapability", Description: "计算过程能力", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/getCapabilityList", Description: "获取能力分析列表", ApiGroup: "spc", Method: "GET"},
		
		// Dashboard API
		{Path: "/spc/getDashboard", Description: "获取Dashboard数据", ApiGroup: "spc", Method: "GET"},
	}

	if err = db.Create(&apis).Error; err != nil {
		return ctx, err
	}

	// 为admin角色分配API权限
	var casbinRules []system.CasbinInfo
	for _, api := range apis {
		casbinRules = append(casbinRules, system.CasbinInfo{
			Path:   api.Path,
			Method: api.Method,
		})
	}

	if err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		rules := [][]string{}
		for _, rule := range casbinRules {
			rules = append(rules, []string{
				string(rune(adminAuthorityID)),
				rule.Path,
				rule.Method,
			})
		}
		return global.GVA_Casbin.AddPolicies(rules)
	}); err != nil {
		return ctx, err
	}

	global.GVA_LOG.Info("SPC菜单和权限初始化完成")
	return ctx, nil
}

func (i *initSpcMenu) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var count int64
	db.Model(&system.SysBaseMenu{}).Where("name = ?", "spc").Count(&count)
	return count > 0
}
