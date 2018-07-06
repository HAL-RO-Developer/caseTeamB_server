package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors)

	r.GET("/log", log)

	// アプリ側処理
	api := r.Group("/")
	apiRouter(api)

	// デバイス側処理
	thing := r.Group("/thing")
	thingRouter(thing)
	return r

}

func cors(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", headers)
	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	}
}

func log(c *gin.Context) {
	c2 := *c
	s, _ := c2.GetRawData()
	fmt.Println(string(s))
}
