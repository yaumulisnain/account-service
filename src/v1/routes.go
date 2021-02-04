package v1

import (
	"github.com/gin-gonic/gin"

	"account-service/src/middleware"
	"account-service/src/v1/delivery"
)

func Route(route *gin.Engine) *gin.RouterGroup {

	api := route.Group("/v1")
	{
		api.POST("/sign-up", delivery.SignUp)
		api.POST("/sign-in", delivery.SignIn)
	}

	apiWithToken := route.Group("/v1")
	apiWithToken.Use(middleware.CheckToken("token"))
	{
		apiWithToken.POST("/sign-out", delivery.SignOut)
		apiWithToken.GET("/chart", delivery.GetMusicChart)
		apiWithToken.GET("/fav", delivery.GetUserFavMusic)
	}

	apiWithRefresh := route.Group("/v1")
	apiWithRefresh.Use(middleware.CheckToken("refresh"))
	{
		apiWithRefresh.POST("/refresh-token", delivery.RefreshToken)
	}

	return api
}
