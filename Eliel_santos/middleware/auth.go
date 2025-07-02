package middleware

import (
    "net/http"
    "os"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

const patientKey = "patientId"

func PatientID(c *gin.Context) uint {
    if id, ok := c.Get(patientKey); ok {
        return id.(uint)
    }
    return 0
}

func Auth() gin.HandlerFunc {
    secret := []byte(os.Getenv("JWT_SECRET"))
    return func(c *gin.Context) {
        header := c.GetHeader("Authorization")
        if !strings.HasPrefix(header, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "message": "Token inválido ou expirado",
                "code":    401,
            })
            return
        }
        tokenString := strings.TrimPrefix(header, "Bearer ")

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrTokenInvalidType
            }
            return secret, nil
        })
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "message": "Token inválido ou expirado",
                "code":    401,
            })
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok {
            if idFloat, ok := claims["id"].(float64); ok {
                c.Set(patientKey, uint(idFloat))
            }
        }

        c.Next()
    }
}
