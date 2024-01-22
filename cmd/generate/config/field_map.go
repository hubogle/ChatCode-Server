package config

import (
	"gorm.io/gorm"
)

var DataMap = map[string]func(gorm.ColumnType) (dataType string){
	// int mapping
	"int": func(columnType gorm.ColumnType) (dataType string) {
		if n, ok := columnType.Nullable(); ok && n {
			return "*int64"
		}
		return "int64"
	},
}
