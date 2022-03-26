package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	httpError "go-ddd-blog/internal/errors"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func Error(c *gin.Context, err error) {
	var httpErr httpError.HTTPError
	if ok := errors.As(err, &httpErr); ok {
		c.AbortWithStatusJSON(httpErr.StatusCode(), gin.H{
			"success": false,
			"message": httpErr.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"success": false,
		"message": "internal error",
	})
}
