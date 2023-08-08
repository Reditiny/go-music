package common

import (
	"fmt"
	"go-music/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局mysql数据库变量
var DB *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
		config.Conf.Mysql.Charset,
		config.Conf.Mysql.Collation,
		config.Conf.Mysql.Query,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Printf("初始化mysql数据库异常: %v\n", err)
		panic(fmt.Errorf("初始化mysql数据库异常: %v", err))
	}

	// 开启mysql日志
	if config.Conf.Mysql.LogMode {
		db.Debug()
	}
	DB = db
	fmt.Println("mysql数据库初始化完成")
}
