package middleware

import (
	"github.com/gin-gonic/gin"
	"go-product/domain"
	"go-product/helpers"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ApiResponse{
				Code:    http.StatusUnauthorized,
				Status:  "UNAUTHORIZED",
				Message: "Invalid access token. Please login to get a new access token",
			})
			return
		}

		c.Set("claim", claim)
		c.Next()
	}
}
