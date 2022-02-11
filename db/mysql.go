package db

import (
	"defaultProject/config"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type MysqlConnect struct{}

func (m *MysqlConnect) Init() error {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Cfg.GetString("Mysql.Host"),
		DefaultStringSize:         1024, // string 类型字段的默认长度
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		logrus.Error(err)
		return err
	}

	sqldb, err := db.DB()
	if err != nil {
		logrus.Error(err)
		return err
	}
	sqldb.SetConnMaxIdleTime(10)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(time.Second * 180)

	DB = db
	return nil
}
