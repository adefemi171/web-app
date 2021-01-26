package route

import (
	controller "github.com/adefemi171/webapp/pkg/controller"
	
	"github.com/gin-gonic/gin"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.Welcome)
	router.GET("/users", controller.UserHandler)
	router.GET("/header", gin.WrapF(controller.GetHeader))
	router.NoRoute(controller.NotFound)
}

// func PromRoutes(router *gin.Engine) {
// 	router.GET("/metrics", gin.WrapF(promhttp.Handler))
// }
