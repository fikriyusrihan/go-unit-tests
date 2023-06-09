package router

import (
	"github.com/gin-gonic/gin"
	"go-product/controllers"
	"go-product/infrastructure/http/handler"
	"go-product/infrastructure/http/middleware"
)

func NewRouter(ctr controllers.AppController) *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	{
		users.POST("/login", middleware.AuthRequestValidator(), handler.PostUserLogin(ctr))
		users.POST("/register", middleware.UserRequestValidator(), handler.PostUserRegister(ctr))
	}

	products := router.Group("/products", middleware.Authentication())
	{
		products.POST("/", middleware.ProductRequestValidator(), handler.PostProduct(ctr))
		products.GET("/:productId", middleware.ProductAuthorization(), handler.GetProductById(ctr))

		adminAccess := products.Group("/", middleware.AdminAuthorization())
		{
			adminAccess.PUT(":productId", middleware.ProductRequestValidator(), handler.PutProduct(ctr))
			adminAccess.DELETE(":productId", handler.DeleteProduct(ctr))
			adminAccess.GET("/", handler.GetProducts(ctr))
		}
	}

	return router
}
