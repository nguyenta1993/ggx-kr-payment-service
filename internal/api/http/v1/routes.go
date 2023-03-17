package v1

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("./templates/*")
	// Serve frontend static files
	router.Use(static.Serve("/static", static.LocalFile("./static", true)))
	// setup client side templates
	router.GET("/", IndexHandler)
	router.GET("/preview/:type", PreviewHandler)
	router.GET("/checkout/:type", CheckoutHandler)
	router.GET("/result/:status", ResultHandler)

	// Setup route group and routes for the API
	api := router.Group("/api")

	api.POST("/sessions", SessionsHandler)
	api.POST("/webhooks/notifications", WebhookHandler)

	// handle redirects
	api.GET("/handleShopperRedirect", RedirectHandler)
	api.POST("/handleShopperRedirect", RedirectHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
