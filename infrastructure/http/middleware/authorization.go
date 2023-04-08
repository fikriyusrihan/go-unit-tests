package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-product/domain"
	"go-product/infrastructure/db"
	"net/http"
)

func AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := db.GetDBInstance()
		claim := c.MustGet("claim").(jwt.MapClaims)
		uid := uint(claim["id"].(float64))

		var user domain.User
		err := database.Select("is_admin").First(&user, uid).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, domain.ApiResponse{
				Code:    http.StatusNotFound,
				Status:  "NOT FOUND",
				Message: "User not found",
			})
			return
		}

		if !user.IsAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, domain.ApiResponse{
				Code:    http.StatusForbidden,
				Status:  "FORBIDDEN",
				Message: "You are not authorized to access this resource",
			})
			return
		}

		c.Next()
	}
}
