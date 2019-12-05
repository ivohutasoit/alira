package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
			fmt.Println(strings.Trim(url))
			url, err := util.Encrypt(strings.Trim(url), os.Getenv("SECRET_KEY"))
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
