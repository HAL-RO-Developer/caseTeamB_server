package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(obj interface{}, c *gin.Context) {
	c.JSON(http.StatusBadRequest, obj)
}

func Json(obj interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, obj)
}
