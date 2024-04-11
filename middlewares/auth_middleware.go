package middlewares

import (
	"Application/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		token := session.Get("token")
		if token == nil {
			context.Redirect(http.StatusTemporaryRedirect, "/network-management-solutions")
		} else {
			validateToken, validateTokenError := services.ValidateToken(token.(string))
			if validateTokenError != nil && !validateToken.Valid {
				context.Redirect(http.StatusTemporaryRedirect, "/network-management-solutions")
			} else {
				claims := validateToken.Claims.(jwt.MapClaims)
				context.Set("user_id", claims["user_id"])
				context.Set("name", claims["name"])
				context.Set("role", claims["role"])
				context.Set("email", claims["email"])
			}
		}
	}
}
