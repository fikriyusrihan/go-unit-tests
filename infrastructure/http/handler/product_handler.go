package handler

import (
	"github.com/gin-gonic/gin"
	"go-product/controllers"
)

func PostProduct(ctr controllers.AppController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctr.HandleCreateProduct(ctx)
	}
}

func PutProduct(ctr controllers.AppController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctr.HandleUpdateProduct(ctx)
	}
}

func DeleteProduct(ctr controllers.AppController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctr.HandleDeleteProduct(ctx)
	}
}
