package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		appHeaders := []string{"X-APP-NAME", "X-APP-TOKEN", "X-APP-ID", "X-CLIENT-VERSION"}

		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "HEAD, OPTIONS, POST, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, "+strings.Join(appHeaders, ", "))
		c.Writer.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token, Authorization, X-Requested-With, "+strings.Join(appHeaders, ", "))

		//Obscure server
		c.Writer.Header().Set("Server", "Mfereji.io Devkit-go-server")

		if c.Request.Method == "OPTIONS" || c.Request.Method == "HEAD" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
