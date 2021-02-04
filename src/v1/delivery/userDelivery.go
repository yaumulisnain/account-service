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

func SignIn(c *gin.Context) {
	var (
		userLogin model.UserLogin
		token     model.TokenResponse
		err       error
	)

	v := validator.New()
	err = c.ShouldBindJSON(&userLogin)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed when parsing request"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	if err = v.Struct(userLogin); err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Validation error"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	token, err = usecase.Login(userLogin)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed to login"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    token,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}

func RefreshToken(c *gin.Context) {
	var (
		tokenResponse model.TokenResponse
		token         string
		err           error
	)

	token = c.Request.Header.Get("Authorization")
	token, err = usecase.CleanUpToken(token)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Token is not valid"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	tokenResponse, err = usecase.RefreshToken(token)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed create token"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    tokenResponse,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}

func SignOut(c *gin.Context) {
	var (
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

	err = usecase.Logout(token)

	if err != nil {
		errResp := &h.ErrorResp{Error: err.Error(), Code: http.StatusBadRequest, Message: "Failed to log out"}
		c.AbortWithStatusJSON(errResp.Code, errResp)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    true,
		"code":    http.StatusOK,
		"message": http.StatusText(http.StatusOK),
	})
}
