package gorm

import (
	"fmt"
	"gorm.io/gorm"
)

type AutoTable struct {
}

func (auto AutoTable) CreateTableAuto(db *gorm.DB, model interface{}, tableName string, tableComment string) {
	db.Table(tableName).Set("gorm:table_options", fmt.Sprintf("COMMENT='%s'", tableComment)).AutoMigrate(model)
}
