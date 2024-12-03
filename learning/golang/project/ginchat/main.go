package main

import (
	"ginchat/router"
	"ginchat/utils"
)
func main() {
	utils.InitConfig()
	utils.InitMysql()
	router.Router().Run(":8081")
}
