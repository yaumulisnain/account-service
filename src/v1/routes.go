package v1

import (
	"github.com/gin-gonic/gin"

	"account-service/src/v1/delivery"
)

func Route(route *gin.Engine) *gin.RouterGroup {

	api := route.Group("/v1")
	{
		api.POST("/sign-up", delivery.SignUp)
		api.POST("/sign-in", delivery.SignIn)
		// api.POST("/refresh-token", delivery.RefreshToken)
	}

	// auth := route.Group("/v1")
	// auth.Use(middleware.CheckToken())
	// {
	// 	auth.GET("/chart", deliver.GetMusicChart)
	// 	auth.GET("/fav", deliver.GetUserFavMusic)
	// }

	return api
}
