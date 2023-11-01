package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"moduledemo/model"
)

// 全局mysql数据库变量
var DB *gorm.DB

// 初始化mysql数据库
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		"root",
		"986203",
		"112.74.125.238",
		3306,
		"neshield",
		"utf8",
		"utf8mb4_general_ci",
		"parseTime=True&loc=Local&timeout=10000ms",
	)
	// Log.Info("数据库连接DSN: ", showDsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		//// 指定表前缀
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: config.Conf.Mysql.TablePrefix + "_",
		//},
	})
	if err != nil {
		log.Panicf("初始化mysql数据库异常: %v", err)
		panic(fmt.Errorf("初始化mysql数据库异常: %v", err))
	}

	// 开启mysql日志
	db.Debug()
	// 自动生成建表语句
	err = db.AutoMigrate(&model.TNesDevice{}, &model.User{})
	if err != nil {
		log.Println(err)
	}

	// 全局DB赋值
	DB = db
	log.Println("初始化mysql数据库完成! dsn: %s", dsn)
}
