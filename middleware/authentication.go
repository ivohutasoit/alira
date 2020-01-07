package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ivohutasoit/alira/model"
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
		excepts := strings.Split(os.Getenv("EXCEPT_WEB"), ";")
		optionals := strings.Split(os.Getenv("OPTIONAL_WEB"), ";")

		currentPath := c.Request.URL.Path
		for _, value := range excepts {
			if strings.TrimSpace(value) == currentPath {
				c.Next()
				return
			}
		}

		optional := false
		for _, value := range optionals {
			if strings.TrimSpace(value) == currentPath {
				optional = true
				break
			}
		}

		url, err := util.GenerateUrl(c.Request.TLS, c.Request.Host, c.Request.URL.Path, true)
		if err != nil {
			fmt.Println(err)
			return
		}

		session := sessions.Default(c)
		accessToken := session.Get("access_token")
		if accessToken == nil {
			redirect := fmt.Sprintf("%s?redirect=%s", args[0].(string), url)
			c.Redirect(http.StatusMovedPermanently, redirect)
			c.Abort()
			return
		}
		claims := &domain.AccessTokenClaims{}
		token, err := jwt.ParseWithClaims(accessToken.(string), claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			redirect := fmt.Sprintf("%s?redirect=%s", args[0].(string), url)
			c.Redirect(http.StatusMovedPermanently, redirect)
			c.Abort()
			return
		}

		var user *domain.User
		model.GetDatabase().First(&user, "user_id = ? AND active = ? AND deleted_at IS NULL",
			claims.UserID, true)
		if user == nil && !optional {
			redirect := fmt.Sprintf("%s?redirect=%s", args[0].(string), url)
			c.Redirect(http.StatusMovedPermanently, redirect)
			c.Abort()
			return
		}
		c.Set("userid", user.ID)

		c.Next()
	}
}

func TokenHeaderRequired(args ...interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		excepts := strings.Split(os.Getenv("EXCEPT_API"), ";")

		currentPath := c.Request.URL.Path

		for _, value := range excepts {
			if strings.TrimSpace(value) == currentPath {
				c.Next()
				return
			}
		}

		authorization := c.Request.Header.Get("Authorization")

		if authorization == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   401,
				"status": "Unauthorized",
				"error":  "missing authorization token",
			})
			c.Abort()
			return
		}

		tokens := strings.Split(authorization, " ")
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

		var claims jwt.Claims
		if tokens[0] == "Bearer" {
			claims = &domain.AccessTokenClaims{}
		} else if tokens[0] == "Refresh" {
			if currentPath != "/api/alpha/auth/refresh" {
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":   401,
					"status": "Unauthorized",
					"error":  "invalid refresh uri",
				})
				c.Abort()
				return
			}
			claims = &domain.RefreshTokenClaims{}
		} else {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   401,
				"status": "Unauthorized",
				"error":  "invalid token indentifier",
			})
			c.Abort()
			return
		}

		tokenString := tokens[1]
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   401,
				"status": "Unauthorized",
				"error":  err.Error(),
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

		var user *domain.User
		var userID string
		var sub int
		if tokens[0] == "Bearer" {
			userID = claims.(*domain.AccessTokenClaims).UserID
		} else if tokens[0] == "Refresh" {
			if claims.(*domain.RefreshTokenClaims).Sub != 1 {
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":   401,
					"status": "Unauthorized",
					"error":  "invalid refresh token",
				})
				c.Abort()
				return
			}
			userID = claims.(*domain.RefreshTokenClaims).UserID
			sub = claims.(*domain.RefreshTokenClaims).Sub
		}

		model.GetDatabase().First(&user, "user_id = ? AND active = ? AND deleted_at IS NULL",
			userID, true)
		if user == nil {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   401,
				"status": "Unauthorized",
				"error":  "invalid token user",
			})
			c.Abort()
			return
		}
		c.Set("userid", user.ID)
		c.Set("sub", sub)

		c.Next()
	}
}
