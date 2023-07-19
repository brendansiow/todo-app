package helper

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func getJwtKey() string {
	key, exists := os.LookupEnv("JWT_KEY")
	if !exists {
		log.Fatal("JWT Key not found")
	}

	return key
}

func GenerateJwtToken(id uint) string {
	key := getJwtKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		log.Fatal("Error signing token: ", err)
		return ""
	}
	return tokenString
}

func IsJwtTokenValid(c *gin.Context) error {
	tokenString := getTokenFromHeader(c)
	if tokenString == "" {
		return errors.New("token is empty")
	}
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getJwtKey()), nil
	})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func getTokenFromHeader(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken != "" && len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func GetID(c *gin.Context) uint {
	tokenString := getTokenFromHeader(c)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getJwtKey()), nil
	})
	if err != nil {
		log.Fatal(err)
		return 0
	}

	if token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["sub"]), 10, 32)
		if err != nil {
			log.Fatal(err)
			return 0
		}
		return uint(uid)
	}
	return 0
}
