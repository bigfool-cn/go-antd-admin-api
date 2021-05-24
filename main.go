package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"wechat-bot-api/configs"
	"wechat-bot-api/routers"
)

func main()  {
	if configs.Conf.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := routers.InitRouter()

	if err := r.Run(configs.Conf.App.Host + ":" + configs.Conf.App.Port);err != nil {
		log.Fatalf("run server error: %v",err)
	}
}
