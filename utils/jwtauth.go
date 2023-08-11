package util

import (
	"fmt"
	"strconv"
	"strings"

	"container/list"
	"time"

	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

var whitelistTokens = list.New()

func GenerateToken(uid string) (string, error) {
	token_lifetime, err := strconv.Atoi(GetEnv("TOKEN_HOUR_LIFETIME"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uid"] = uid
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifetime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetEnv("JWT_KEY")))
}

func BlockToken(tokenString string) error {
	if whitelistTokens.Front() != nil {
		pointer := whitelistTokens.Front()
		for {
			if pointer.Value == tokenString {
				return fmt.Errorf("token has already blocked")
			}
			if pointer.Next() == nil {
				break
			} else {
				pointer.Next()
			}
		}
	}
	whitelistTokens.PushBack(tokenString)
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (string, error) {
	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetEnv("JWT_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["uid"].(string), nil
	}
	return "", nil
}
