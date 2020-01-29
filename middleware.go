package alira

import "github.com/gin-gonic/gin"

type AuthenticationMiddleware interface {
	SessionRequired(args ...interface{}) gin.HandlerFunc
	TokenRequired(args ...interface{}) gin.HandlerFunc
}

type AccessMiddleware interface {
	AppAdminRequired(args ...interface{}) gin.HandlerFunc
	CustomerAdminRequired(args ...interface{}) gin.HandlerFunc
}
