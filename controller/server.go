package controller

import "github.com/gin-gonic/gin"

func StartServer() {
	// ReleaseMode  ช่วยให้เร็วขึ้น แต่จะ debug ไม่ได้
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"massage": "pong!!!!!!!",
	// 	})
	// })

	// Load  Controllers
	DemoController(router)
	PersonController(router)
	router.Run()
}
