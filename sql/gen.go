package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// 根据数据库生成表结构
func main() {
	g := gen.NewGenerator(gen.Config{
		//  设置输出路径
		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery,
		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
	})
	//  建立数据库连接
	gormdb, _ := gorm.Open(mysql.Open("root:root@(121.37.143.160:3306)/dish?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // 选择数据库连接
	g.ApplyBasic(
		// 从当前数据库的所有表生成结构
		g.GenerateAllTable()...,
	)
	// 生成代码
	g.Execute()
}
