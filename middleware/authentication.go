package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ivohutasoit/alira/model/domain"
	"github.com/ivohutasoit/alira/util"
)

func SessionHeaderRequired(args ...interface{}) gin.HandlerFunc {
	/*if 1 > len(args) {
		panic("authentication uri must be provided")
	}*/
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

		url, err := util.GenerateUrl(c.Request.TLS, c.Request.Host, currentPath, true)
		if err != nil {
			fmt.Println(err)
		}
		redirect := fmt.Sprintf("%s?redirect=%s", os.Getenv("URL_LOGIN"), url)

		session := sessions.Default(c)
		accessToken := session.Get("access_token")
		if accessToken == nil && !opt {
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
				c.Redirect(http.StatusMovedPermanently, redirect)
				c.Abort()
				return
			}

			data := map[string]string{
				"type":  "Bearer",
				"token": accessToken.(string),
			}
			// https://tutorialedge.net/golang/consuming-restful-api-with-go/
			payload, _ := json.Marshal(data)
			resp, err := http.Post(os.Getenv("URL_AUTH"), "application/json", bytes.NewBuffer(payload))
			if err != nil && !opt {
				c.Redirect(http.StatusMovedPermanently, redirect)
				c.Abort()
				return
			}
			respData, err := ioutil.ReadAll(resp.Body)
			if err != nil && !opt {
				c.Redirect(http.StatusMovedPermanently, redirect)
				c.Abort()
				return
			}
			var response domain.Response
			json.Unmarshal(respData, &response)
			if response.Code != http.StatusOK && !opt {
				c.Redirect(http.StatusMovedPermanently, redirect)
				c.Abort()
				return
			}
			c.Set("userid", response.Data["userid"])
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
		if tokens[0] == "Refresh" {
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
		}

		data := map[string]string{
			"type":  tokens[0],
			"token": tokenString,
		}
		// https://tutorialedge.net/golang/consuming-restful-api-with-go/
		payload, _ := json.Marshal(data)
		resp, err := http.Post(os.Getenv("URL_AUTH"), "application/json", bytes.NewBuffer(payload))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":   http.StatusUnauthorized,
				"status": http.StatusText(http.StatusUnauthorized),
				"error":  "unable to verify token",
			})
			return
		}
		respData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":   http.StatusUnauthorized,
				"status": http.StatusText(http.StatusUnauthorized),
				"error":  "unable to read detail token",
			})
			return
		}
		var response domain.Response
		json.Unmarshal(respData, &response)
		if response.Code != http.StatusOK {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":   http.StatusUnauthorized,
				"status": http.StatusText(http.StatusUnauthorized),
				"error":  "unable to apply detail token",
			})
			return
		}
		c.Set("userid", response.Data["userid"])
		if tokens[0] == "Refresh" {
			c.Set("sub", claims.(*domain.RefreshTokenClaims).Sub)
		}
		c.Next()
	}
}
