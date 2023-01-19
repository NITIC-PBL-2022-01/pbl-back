package middleware

import (
	"log"
	"nitic-pbl-2022-01/pbl-back/src/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
      c.JSON(401, map[string]string{ "message": "Unauthorized" })
      return
    }

    email, err := auth.GetUserEmail(authHeader)
    if err != nil {
      log.Println(err)
      c.JSON(401, map[string]string{ "message": "Unauthorized" })
      return
    }

    c.Set("email", email)
    c.Next()
  }
}
