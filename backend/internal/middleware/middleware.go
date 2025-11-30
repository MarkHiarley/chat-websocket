package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markHiarley/projetinho/internal/auth"
)

func ValidateTokenJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		tokenString, err := auth.ExtractTokenFromHeader(authHeader)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token não fornecido",
			})
			ctx.Abort()
			return
		}

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token inválido ou expirado",
			})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)
		ctx.Next()

	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")

		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Usuário não autenticado",
			})
			ctx.Abort()
			return
		}

		if role != "admin" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "Acesso negado. Apenas administradores podem acessar esta rota.",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
