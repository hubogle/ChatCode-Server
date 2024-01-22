package main

import (
	"fmt"

	"github.com/hubogle/chatcode-server/cmd/generate/config"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/rawsql"
)

func generateModel() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/dal/query",
		// ModelPkgPath: outPath,
		FieldNullable:     true,  // if you want the nullable field generation property to be pointer type
		FieldCoverable:    false, // if you want to assign field which has a default value in the `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		FieldSignable:     true,  // if you want to generate field with unsigned integer type
		FieldWithIndexTag: false, // if you want to generate index tags from database
		FieldWithTypeTag:  false, // if you want to generate type tags from database
		WithUnitTest:      false, // if you need unit tests for query code, set WithUnitTest true
		Mode:              gen.WithQueryInterface | gen.WithDefaultQuery | gen.WithoutContext,
	})

	db, _ := gorm.Open(rawsql.New(rawsql.Config{
		// SQL:      rawsql,                      //create table sql
		FilePath: []string{
			"./database/init.sql", // create table sql file
		},
	}))

	g.UseDB(db)

	// 去除掉 gorm tag 中的 comment
	g.WithOpts(gen.FieldGORMTagReg(".", func(tag field.GormTag) field.GormTag {
		tag.Remove("comment")
		return tag
	}))

	// GenerateModel 生成表的 CURD
	// GenerateModelAs 自定义某个表生成特性
	g.ApplyBasic(
		g.GenerateModel("group"),
	)
	g.GenerateAllTable()

	// model生成自定义方法，可以用ApplyInterface，第一个参数是方法接口
	commonMethod := config.CommonMethod{}
	// g.ApplyInterface(
	// 	// func(config.UserMethod, config.Querier) {},
	// 	g.GenerateModel("user",
	// 		gen.WithMethod(commonMethod),
	// 		gen.FieldType("ext_info", "json.RawMessage"),
	// 	), // 可以绑定一个方法，也可以绑定全部方法
	// )

	// 过滤不需要生成的表
	filterTable := map[string]struct{}{}

	// opts := []gen.ModelOpt{}
	tableList, err := db.Migrator().GetTables()
	if err != nil {
		panic(fmt.Errorf("get all tables fail: %w", err))
	}
	tableModels := make([]interface{}, 0, len(tableList))
	for _, tableName := range tableList {
		if _, ok := filterTable[tableName]; !ok {
			tableModels = append(tableModels, g.GenerateModel(tableName, gen.WithMethod(commonMethod)))
		}
	}

	// 导出所有的表结构，使用通用方法
	g.ApplyBasic(
		tableModels...,
	)

	// 指定字段类型和数据库列类型之间的数据类型映射
	g.WithDataTypeMap(config.DataMap)
	g.Execute()
}

func main() {
	generateModel()
}
