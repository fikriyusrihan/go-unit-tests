package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-product/domain"
	"go-product/infrastructure/db"
	"net/http"
	"strconv"
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

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := db.GetDBInstance()
		claim := c.MustGet("claim").(jwt.MapClaims)

		uid := uint(claim["id"].(float64))
		pid, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, domain.ApiResponse{
				Code:    http.StatusBadRequest,
				Status:  "BAD REQUEST",
				Message: "Invalid product id",
			})
			return
		}

		var user domain.User
		err = database.Select("is_admin").First(&user, uid).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, domain.ApiResponse{
				Code:    http.StatusNotFound,
				Status:  "NOT FOUND",
				Message: "User not found",
			})
			return
		}

		if !user.IsAdmin {
			var product domain.Product
			err = database.Select("user_id").First(&product, pid).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, domain.ApiResponse{
					Code:    http.StatusNotFound,
					Status:  "NOT FOUND",
					Message: "Product not found",
				})
				return
			}

			if product.UserID != uid {
				c.AbortWithStatusJSON(http.StatusForbidden, domain.ApiResponse{
					Code:    http.StatusForbidden,
					Status:  "FORBIDDEN",
					Message: "You are not authorized to access this product",
				})
				return
			}
		}

		c.Next()
	}
}
