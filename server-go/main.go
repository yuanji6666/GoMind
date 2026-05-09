package main

import (
	"fmt"
	"log"

	"github.com/yuanji6666/gopherAI/common/mysql"
	"github.com/yuanji6666/gopherAI/common/redis"
	"github.com/yuanji6666/gopherAI/config"
	"github.com/yuanji6666/gopherAI/domain/ai_session"
	"github.com/yuanji6666/gopherAI/domain/user"
	"github.com/yuanji6666/gopherAI/router"
)

func main() {
	config.InitConfig()
	
	fmt.Println(config.GetConfig().MainConfig.Host)

	if err := mysql.InitMysql(); err != nil {
		log.Panic("mysql Init error", err)
	}

	// 在 InitMysql 建立 DB 连接后，由 main 提供 domain 的模型给迁移函数，避免循环依赖
	if err := mysql.RunMigration(
		new(user.User),
		new(ai_session.Message),
		new(ai_session.Session),
	); err != nil {
		log.Panic("database migration error", err)
	}

	redis.InitRedis()

	r := router.InitRouter()

	err := r.Run(config.GetConfig().MainConfig.Host + ":" + config.GetConfig().MainConfig.Port)
	if err != nil {
		return
	}
}
