package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

func ERROR(c *gin.Context, statusCode int, err error) {
	if err != nil {
		JSON(c, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(c, http.StatusBadRequest, nil)
}
