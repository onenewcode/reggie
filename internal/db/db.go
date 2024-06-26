package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"reggie/internal/config"
	"time"
)

var DBEngine *gorm.DB
var EmpDao empI = &employeeDao{}
var CatDao catI = &categoryDao{}
var DisDao dishI = &dishDao{}
var UserDao userI = &userDao{}
var DishFDao dishFI = &dishFDao{}
var MealDishDao meal_dishI = &mealDishDao{}
var MealDao mealI = &mealDao{}
var ShopCartDao shoppingcartI = &shoppingcartDao{}
var AddressDA0 addressI = &addressDao{}
var OrderDao orderI = &orderDao{}
var OrderDetailDao order_detialI = &order_detialDao{}

func InitDB() {
	// 创建一个新的logger实例，设置为verbose模式以输出详细日志，包括SQL语句
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:      logger.Info, // 设置日志级别
			SlowThreshold: time.Second, // 慢查询阈值
			Colorful:      true,        // 是否使用彩色日志
		},
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DatabaseSetting.Url, // DSN data source name
		DefaultStringSize:         256,                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		// 取消生成外键。
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   dbLogger,
	})
	if err != nil {
		panic("数据库链接失败")
	}
	DBEngine = db
}
