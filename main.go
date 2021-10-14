package main

import (
	"github.com/lumencode/livechat/config"
	ctrl "github.com/lumencode/livechat/controllers"

	"github.com/gin-gonic/gin"
)

func main () {

	r := gin.Default()

	r.POST("/create-account", ctrl.CreateAccount)

	r.Run(config.Config().Server.Addr)
}
