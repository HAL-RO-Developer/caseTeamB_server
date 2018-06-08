package router

import (
	"fmt"

	"github.com/HAL-RO-Developer/caseTeamB_server/middleware"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(preflightRequest)

	api := r.Group("/api")
	apiRouter(api)
	auth := api.Group("")
	auth.Use(middleware.Jwt("hogehoge", 3600*24*365))
	authApiRouter(auth)
	return r

}

//func cros(c *gin.Context) {
//	c.Header("Access-Control-Allow-Origin", "*")
//	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
//}

/*func cros(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", headers)
	if c.Request.Method == "OPTIONS" {
		c.Status(200)
		c.Abort()
	}

	c.Set("start_time", time.Now())
	c.Next()
}*/

func preflightRequest(c *gin.Context) {
	headers := c.Request.Header.Get("Access-Control-Request-Headers")

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", headers)
	if c.Request.Method == "OPTIONS" {
		fmt.Println("in!!!!")
		c.Status(200)
		c.Abort()
	}
}
