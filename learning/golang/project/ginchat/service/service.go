package service

import (
	"github.com/gin-gonic/gin"
)
// GetIndex godoc
// @Tags         首页
// @Success      200  {string}  hello
// @Router       /index [get]
func GetIndex(c *gin.Context){
	c.JSON(200,gin.H{
		"message" : "hello",
	})
}


