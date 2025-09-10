package middlewares

import (
	"errors"
	"net/http"

	"app/repositories"
	"app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HTTPError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if last := c.Errors.Last(); last == nil {
			return
		} else {
			err := last.Err

			if errors.Is(err, repositories.ErrNotFoundOrNotOwner) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "not allowed to modify this post"})
				return
			}
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "data not found"})
				return
			}
			var api utils.APIError
			if errors.As(err, &api) {
				c.AbortWithStatusJSON(api.Code, gin.H{"error": api.Message})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}
