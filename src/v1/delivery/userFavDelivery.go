package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"

	h "account-service/src/helper"
	"account-service/src/v1/model"
	"account-service/src/v1/usecase"
)

func GetUserFavMusic(c *gin.Context) {
	var (
		music []model.Music
		token string
		err   error
	)

	token = c.Request.Header.Get("Authorization")
	token, err = usecase.CleanUpToken(token)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Token is not valid"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	music, err = usecase.ListUserFavorite(token)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed to get data"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    music,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})

	return
}
