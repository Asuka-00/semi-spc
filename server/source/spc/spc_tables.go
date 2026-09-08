package spc

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/model/spc"
	sysService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"gorm.io/gorm"
)

type initSpcTables struct{}

func init() {
	sysService.RegisterInit(sysService.InitOrderInternal+40, &initSpcTables{})
}

func (i *initSpcTables) InitializerName() string {
	return "spc_tables"
}

func spcModels() []interface{} {
	return []interface{}{
		&spc.SpcSite{},
		&spc.SpcArea{},
		&spc.SpcEquipment{},
		&spc.SpcChamber{},
		&spc.SpcTechnology{},
		&spc.SpcProduct{},
		&spc.SpcProcessStep{},
		&spc.SpcRecipe{},
		&spc.SpcLot{},
		&spc.SpcWafer{},
		&spc.SpcParameter{},
		&spc.SpcSpec{},
		&spc.SpcChart{},
		&spc.SpcControlLimit{},
		&spc.SpcRule{},
		&spc.SpcSample{},
		&spc.SpcMeasurement{},
		&spc.SpcAlarm{},
		&spc.SpcOcap{},
		&spc.SpcOcapExecution{},
		&spc.SpcCapability{},
	}
}

func (i *initSpcTables) MigrateTable(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, sysService.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(spcModels()...)
}

func (i *initSpcTables) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	for _, model := range spcModels() {
		if !db.Migrator().HasTable(model) {
			return false
		}
	}
	return true
}

func (i *initSpcTables) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (i *initSpcTables) DataInserted(ctx context.Context) bool {
	return true
}
