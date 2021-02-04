package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"account-service/src/v1/usecase"
)

func CheckToken(tokenType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			err   error
		)

		isValid := true

		token = c.Request.Header.Get("Authorization")
		token, err = usecase.CleanUpToken(token)

		if err != nil {
			isValid = false
		}

		if isValid {
			t, valid, e := usecase.VerifyToken(token)

			if valid && e == nil {
				if tokenType == t {
					c.Next()
				} else {
					isValid = false
				}
			} else {
				isValid = false
			}
		}

		if !isValid {
			if err == nil {
				err = errors.New("user not permitted")
			}

			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   err.Error(),
				"code":    http.StatusUnauthorized,
				"message": "User not permitted",
			})
			c.Abort()
			return
		}
	}
}
