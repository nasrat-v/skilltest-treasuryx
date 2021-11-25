package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Basic Authentication middleware
func BasicAuth(context *gin.Context) {
	user, password, hasAuth := context.Request.BasicAuth()

	if !hasAuth || user != "nasraty" || password != "xxxx" { // default credentials
		context.Abort()
		context.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad credentials",
		})
	}
}
