package middleware

import (
	"errors"
	"fmt"
	"majoo/helpers"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type CustomClaim struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func GenerateToken(userId int) (string, error) {
	claim := CustomClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    os.Getenv("APP_NAME"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func ValidateToken(token string) (*jwt.Token, error) {
	plainToken := strings.Replace(token, "Bearer ", "", 1)
	return jwt.Parse(plainToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}

func JWTauth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer ") {
			res := helpers.SendResponse(http.StatusUnauthorized, "Unauthorized", errors.New("Token Invalid"), nil)
			c.AbortWithStatusJSON(res.StatusCode, res)
			return
		}

		validatedToken, err := ValidateToken(authHeader)
		
		claims, ok := validatedToken.Claims.(jwt.MapClaims)
		if !ok {
			res := helpers.SendResponse(http.StatusUnauthorized, "Unauthorized", err, nil)
			c.AbortWithStatusJSON(res.StatusCode, res)
			return
		}

		fmt.Println("ID:", claims["user_id"])
		if err != nil {
			res := helpers.SendResponse(http.StatusUnauthorized, "Unauthorized", err, nil)
			c.AbortWithStatusJSON(res.StatusCode, res)
			return
		}

		c.Set("user_id", int(claims["user_id"].(float64)))		
	}
}