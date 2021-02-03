package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	h "account-service/src/helper"
	"account-service/src/v1/model"
	"account-service/src/v1/usecase"
)

func SignUp(c *gin.Context) {
	var (
		user model.User
		err  error
	)

	v := validator.New()
	err = c.ShouldBindJSON(&user)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed when parsing request"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	if err = v.Struct(user); err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Validation error"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	err = usecase.Registration(&user)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed to register user"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}
