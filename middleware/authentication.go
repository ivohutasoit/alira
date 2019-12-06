package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ivohutasoit/alira/model/domain"
	"github.com/ivohutasoit/alira/util"
)

func AuthenticationRequired(args ...interface{}) gin.HandlerFunc {
	if 1 > len(args) {
		panic("Redirect authentication url must be provided")
	}

	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")
		if token == nil {
			url := fmt.Sprintf("%s%s", c.Request.Host, c.Request.RequestURI)
			fmt.Println(strings.TrimSpace(url))
			url, err := util.Encrypt(strings.TrimSpace(url), os.Getenv("SECRET_KEY"))
			if err != nil {
				fmt.Println(err)
			}
			redirect := fmt.Sprintf("%s%s", args[0].(string), url)
			c.Redirect(http.StatusMovedPermanently, redirect)
			c.Abort()
			return
		}
		c.Next()
	}
}

func SessionHeaderRequired(args ...interface{}) gin.HandlerFunc {
	if 1 > len(args) {
		panic("authentication uri must be provided")
	}
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")
		if token == nil {
			url := fmt.Sprintf("%s%s", c.Request.Host, c.Request.URL.Path)
			fmt.Println(strings.TrimSpace(url))
			url, err := util.Encrypt(strings.TrimSpace(url), os.Getenv("SECRET_KEY"))
			if err != nil {
				fmt.Println(err)
			}
			redirect := fmt.Sprintf("%s?source=%s", args[0].(string), url)
			c.Redirect(http.StatusMovedPermanently, redirect)
			c.Abort()
			return
		}
		c.Next()
	}
}

func TokenHeaderRequired(args ...interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		excepts := []string{"/api/alpha/account/register", "/api/alpha/account/login"}

		currentPath := c.Request.URL.Path

		for _, value := range excepts {
			if strings.TrimSpace(value) == currentPath {
				c.Next()
				return
			}
		}

		header := c.Request.Header.Get("Authorization")

		if header == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   401,
				"status": "Unauthorized",
				"error":  "missing authorization token",
			})
			c.Abort()
			return
		}

		tokens := strings.Split(header, " ")
		if len(tokens) != 2 {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   401,
				"status": "Unauthorized",
				"error":  "invalid token",
			})
			c.Abort()
			return
		}

		tokenString := tokens[1]
		userToken := &domain.UserToken{}

		token, err := jwt.ParseWithClaims(tokenString, userToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   401,
				"status": "Unauthorized",
				"error":  "malformed authentication token",
			})
			c.Abort()
			return
		}

		if !token.Valid {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   401,
				"status": "Unauthorized",
				"error":  "invalid token",
			})
			c.Abort()
			return
		}

		fmt.Printf("User id %s", userToken.UserID)
		c.Set("user_id", userToken.UserID)
		c.Next()
	}
}
