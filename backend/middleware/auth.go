package middleware

import (
	"net/http"

	"github.com/Eizeed/vibe_gogo/models"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token, err := c.Cookie("jwt_token");
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, no token"})
            return;
        }

        var jwt models.JWT;

        claims, err := jwt.DecodeToken(token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, invalid token"})
            return;
        }
        
        c.Set("token", claims);

        c.Next();
    }
}
