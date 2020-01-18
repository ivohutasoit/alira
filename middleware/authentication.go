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

func SessionHeaderRequired(args ...interface{}) gin.HandlerFunc {
	if 1 > len(args) {
		panic("authentication uri must be provided")
	}
	return func(c *gin.Context) {
		currentPath := c.Request.URL.Path
		except := os.Getenv("WEB_EXCEPT")
		if except != "" {
			excepts := strings.Split(except, ";")
			for _, value := range excepts {
				if currentPath == strings.TrimSpace(value) {
					c.Next()
					return
				}
			}
		}

		opt := false
		optional := os.Getenv("WEB_OPTIONAL")
		if optional != "" {
			optionals := strings.Split(optional, ";")
			for _, value := range optionals {
				if value == "/" && (currentPath == "" || currentPath == "/") {
					opt = true
					break
				} else if currentPath == strings.TrimSpace(value) {
					opt = true
					break
				}
			}
		}

		url, err := util.GenerateUrl(c.Request.TLS, c.Request.Host, c.Request.URL.Path, true)
		if err != nil {
			fmt.Println(err)
			return
		}

		session := sessions.Default(c)
		accessToken := session.Get("access_token")
		if accessToken == nil && !opt {
			redirect := fmt.Sprintf("%s?redirect=%s", args[0].(string), url)
			c.Redirect(http.StatusMovedPermanently, redirect)
			c.Abort()
			return
		}
		if accessToken != nil {
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

			sessionToken := &domain.Token{}
			model.GetDatabase().First(sessionToken, "access_token = ? AND valid = ?",
				strings.TrimSpace(accessToken.(string)), true)

			if sessionToken == nil && !opt {
				redirect := fmt.Sprintf("%s?redirect=%s", args[0].(string), url)
				c.Redirect(http.StatusMovedPermanently, redirect)
				c.Abort()
				return
			}
			user := &domain.User{}
			model.GetDatabase().First(user, "id = ? AND active = ? AND deleted_at IS NULL",
				sessionToken.UserID, true)

			if user == nil && !opt {
				redirect := fmt.Sprintf("%s?redirect=%s", args[0].(string), url)
				c.Redirect(http.StatusMovedPermanently, redirect)
				c.Abort()
				return
			}
			c.Set("userid", user.ID)
		}

		c.Next()
	}
}

func TokenHeaderRequired(args ...interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentPath := c.Request.URL.Path
		except := os.Getenv("API_EXCEPT")
		if except != "" {
			excepts := strings.Split(except, ";")

			for _, value := range excepts {
				if currentPath == strings.TrimSpace(value) {
					c.Next()
					return
				}
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

		user := &domain.User{}
		sessionToken := &domain.Token{}
		var userID string
		var sub int
		if tokens[0] == "Bearer" {
			// https://tutorialedge.net/golang/consuming-restful-api-with-go/
			/*jsonData := map[string]string{"token": tokenString}
			jsonValue, _ := json.Marshal(jsonData)
			response, err := http.Post("http://localhost:9000/api/alpha/token/detail/", "application/json", bytes.NewBuffer(jsonValue))
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, _ := ioutil.ReadAll(response.Body)
				fmt.Println(string(data))
			}*/
			model.GetDatabase().First(sessionToken, "access_token = ? AND valid = ?",
				tokenString, true)
			if sessionToken == nil {
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":   401,
					"status": "Unauthorized",
					"error":  "invalid token",
				})
				c.Abort()
				return
			}
			userID = sessionToken.UserID
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
			model.GetDatabase().First(&sessionToken, "refresh_token = ? AND valid = ?",
				tokenString, true)

			if sessionToken == nil {
				c.Header("Content-Type", "application/json")
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":   401,
					"status": "Unauthorized",
					"error":  "invalid refresh token",
				})
				c.Abort()
				return
			}
			userID = sessionToken.UserID
			sub = claims.(*domain.RefreshTokenClaims).Sub
		}

		model.GetDatabase().First(user, "id = ? AND active = ? AND deleted_at IS NULL",
			userID, true)
		if user == nil {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":   401,
				"status": "Unauthorized",
				"error":  "invalid user",
			})
			c.Abort()
			return
		}
		c.Set("userid", user.ID)
		c.Set("sub", sub)

		c.Next()
	}
}
