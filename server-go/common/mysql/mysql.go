// Package mysql 提供mysql数据库访问基础设施
package mysql

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB //全局实例

// InitMysql 初始化Mysql数据库连接
// 根据Gin的运行模式设置日志等级
// 配置数据库连接池
func InitMysql() error {

	//初始化Mysql数据库连接
	host := config.GetConfig().MysqlHost
	port := config.GetConfig().MysqlPort
	user := config.GetConfig().MysqlUser
	password := config.GetConfig().MysqlPassword
	DBname := config.GetConfig().MysqlDBName
	charSet := config.GetConfig().MysqlCharset

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", user, password, host, port, DBname, charSet)

	var log logger.Interface

	// 根据Gin的运行模式设置日志等级
	if gin.Mode() == "debug" {
		log = logger.Default.LogMode(logger.Info)
	} else {
		log = logger.Default
	}

	// 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: log,
	})

	if err != nil {
		return err
	}

	// 获取底层sqlDB对象，设置连接池
	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)           //最大空闲连接
	sqlDB.SetMaxOpenConns(100)          //最大打开连接
	sqlDB.SetConnMaxLifetime(time.Hour) //连接最大声明周期

	DB = db

	return nil

}

// RunMigration 运行数据库迁移，接受模型列表以避免与 domain 包产生循环依赖
func RunMigration(models ...interface{}) error {
    return DB.AutoMigrate(models...)
}
