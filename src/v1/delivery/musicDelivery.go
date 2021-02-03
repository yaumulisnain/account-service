package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"

	h "account-service/src/helper"
	"account-service/src/v1/model"
	"account-service/src/v1/usecase"
)

func GetMusicChart(c *gin.Context) {
	var (
		data []model.Music
		meta h.ResponseMeta
		err  error
	)

	data, meta, err = usecase.GetMusicChart(c.Request.URL.Query())

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed to get list users"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"meta":    meta,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}
