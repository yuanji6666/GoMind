package main

import (
	"log"

	"github.com/yuanji6666/gopherAI/common/mysql"
	"github.com/yuanji6666/gopherAI/common/redis"
	"github.com/yuanji6666/gopherAI/config"
	"github.com/yuanji6666/gopherAI/router"
)

func main() {
	config.InitConfig()

	if err := mysql.InitMysql(); err != nil {
		log.Panic("mysql Init error", err)
	}
	
	redis.InitRedis()
	
	
	/*
	err := readDataFromDB()
	if err != nil {
		return
	}
	*/

	r := router.InitRouter()

	err := r.Run(config.GetConfig().MainConfig.Host + ":" + config.GetConfig().MainConfig.Port)
	if err != nil {
		return
	}
}

/*
func readDataFromDB() error {
	messages, err := message.GetAllMessages()

	if err != nil {

		return err
	}

	for i := range messages {
		msg := messages[i]
		modelType := "1"
		config := map[string]interface{}{}
		helper, err := manager.GetOrCreateAIHelper(msg.UserName, msg.SessionID, modelType, config)
		if err != nil {
			log.Printf("[readDataFromDB] failed to create helper for user=%s session=%s: %v", msg.UserName, msg.SessionID, err)
			return err
		}
		log.Println("readDataFromDB init:  ", helper.SessionID)
		helper.AddMessage(msg.UserName, msg.Content, msg.IsUser, false)
	}

	log.Println("AIHelperManager init success ")
	return nil

}
*/