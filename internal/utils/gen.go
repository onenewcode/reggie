package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/db",                                                      //同时输出结构体和查询方法路径
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"root",
		"121.37.143.160:3306",
		"sky_take_out")))
	g.UseDB(gormdb) // reuse your gorm db// reuse your gorm db

	// Generate basic type-safe DAO API for struct `models.User` following conventions

	g.ApplyBasic(
	// 根据 `user` 表生成结构 `User`
	//g.GenerateModel("users"),

	// 根据 `user` 表生成结构 `Employee` `
	//g.GenerateModelAs("users", "Employee"),

	// Generate struct `User` based on table `users` and generating options
	//g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),
	)
	g.ApplyBasic(
		//  从当前数据库中生成所有表的结构
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
