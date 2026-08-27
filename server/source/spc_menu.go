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
			Path:      "equipment",
			Name:      "spcEquipment",
			Hidden:    false,
			Component: "view/spc/equipment.vue",
			Sort:      2,
			Meta: system.Meta{
				Title: "设备管理",
				Icon:  "platform",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "parameter",
			Name:      "spcParameter",
			Hidden:    false,
			Component: "view/spc/parameter.vue",
			Sort:      3,
			Meta: system.Meta{
				Title: "参数规格",
				Icon:  "data-line",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "chart",
			Name:      "spcChart",
			Hidden:    false,
			Component: "view/spc/chart.vue",
			Sort:      4,
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
			Sort:      5,
			Meta: system.Meta{
				Title: "实时监控",
				Icon:  "monitor",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "collect",
			Name:      "spcCollect",
			Hidden:    false,
			Component: "view/spc/collect.vue",
			Sort:      6,
			Meta: system.Meta{
				Title: "数据采集",
				Icon:  "upload",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "alarm",
			Name:      "spcAlarm",
			Hidden:    false,
			Component: "view/spc/alarm.vue",
			Sort:      7,
			Meta: system.Meta{
				Title: "告警中心",
				Icon:  "bell",
			},
		},
		{
			ParentId:  spcMenu.ID,
			Path:      "capability",
			Name:      "spcCapability",
			Hidden:    false,
			Component: "view/spc/capability.vue",
			Sort:      8,
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
			Sort:      9,
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

	// 注册API
	apis := []system.SysApi{
		{Path: "/spc/getSiteList", Description: "获取厂区列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/site", Description: "创建厂区", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/site", Description: "更新厂区", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/site", Description: "删除厂区", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/site", Description: "获取厂区", ApiGroup: "spc", Method: "GET"},
		
		{Path: "/spc/getAreaList", Description: "获取区域列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createArea", Description: "创建区域", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateArea", Description: "更新区域", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteArea", Description: "删除区域", ApiGroup: "spc", Method: "DELETE"},
		
		{Path: "/spc/getEquipmentList", Description: "获取设备列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createEquipment", Description: "创建设备", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateEquipment", Description: "更新设备", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteEquipment", Description: "删除设备", ApiGroup: "spc", Method: "DELETE"},
		
		{Path: "/spc/getChartList", Description: "获取控制图列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/createChart", Description: "创建控制图", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/updateChart", Description: "更新控制图", ApiGroup: "spc", Method: "PUT"},
		{Path: "/spc/deleteChart", Description: "删除控制图", ApiGroup: "spc", Method: "DELETE"},
		{Path: "/spc/findChart", Description: "查询控制图", ApiGroup: "spc", Method: "GET"},
		
		{Path: "/spc/collect", Description: "数据采集", ApiGroup: "spc", Method: "POST"},
		
		{Path: "/spc/getAlarmList", Description: "获取告警列表", ApiGroup: "spc", Method: "GET"},
		{Path: "/spc/acknowledgeAlarm", Description: "确认告警", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/closeAlarm", Description: "关闭告警", ApiGroup: "spc", Method: "POST"},
		{Path: "/spc/getAlarmStatistics", Description: "告警统计", ApiGroup: "spc", Method: "GET"},
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
