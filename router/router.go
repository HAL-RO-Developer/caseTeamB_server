package router

import (
	"github.com/gin-gonic/gin"
	"github.com/HAL-RO-Developer/caseTeamB_server/middleware"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	apiRouter(api)
	auth := api.Group("")
	auth.Use(middleware.Jwt("hogehoge", 3600*24*365))
	authApiRouter(auth)
	return r

}
