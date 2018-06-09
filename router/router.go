package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors)

	r.GET("/log", log)
	api := r.Group("/api")
	apiRouter(api)
	return r

}

func cors(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
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
