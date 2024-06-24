package middleware

import (
	"github.com/doddeeph/todo-app/domain"
	"github.com/doddeeph/todo-app/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := util.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := util.ExtractIDFromToken(authToken, secret)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					ctx.Abort()
					return
				}
				ctx.Set("x-user-id", userID)
				ctx.Next()
				return
			}
			ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
		ctx.Abort()
	}
}
