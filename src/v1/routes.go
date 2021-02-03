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
		api.POST("/refresh-token", delivery.RefreshToken)

		api.GET("/chart", delivery.GetMusicChart)
		api.GET("/fav", delivery.GetUserFavMusic)
	}

	return api
}
