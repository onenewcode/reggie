package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reggie/internal/global"
)

// 通过
//
//	func buildDSN() *string {
//		//var dsnbuilder strings.Builder
//		value := reflect.ValueOf(global.DatabaseSetting).Elem()
//		for i := 0; i < value.NumField(); i++ {
//			fieldValue := value.Field(i)
//			fmt.Println(fieldValue)
//		}
//		return nil
//	}
func InitDB() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       global.DatabaseSetting.Url, // DSN data source name
		DefaultStringSize:         256,                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("数据库链接失败")
	}
	global.DBEngine = db
}
