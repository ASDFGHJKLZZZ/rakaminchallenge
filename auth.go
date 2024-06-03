package middlewares

import (
    "myapp/helpers"
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
            c.Abort()
            return
        }

        tokenString := strings.Split(authHeader, "Bearer ")[1]
        _, err := helpers.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}