package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// TODO: answer here
		//Get session_token cookie
		cookie, err := ctx.Request.Cookie("session_token")
		if err != nil || cookie.Value == "" {
			//No session_token cookie, return 401 or redirect to login page
			if ctx.GetHeader("Content-Type") == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			} else {
				ctx.Redirect(http.StatusSeeOther, "/client/login")
			}
			return
		}

		//Parse JWT token from session_token cookie
		tokenString := cookie.Value
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		//Failed to parse token, return 401
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		}

		//Invalid token, return 401
		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		//Store email value from claims in context
		ctx.Set("email", claims.Email)

		// Call the next handler
		ctx.Next()
	})
}
