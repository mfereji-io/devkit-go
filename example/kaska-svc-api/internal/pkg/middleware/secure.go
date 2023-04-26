package middleware

import (
	"os"
	"strings"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
)

func Secure() gin.HandlerFunc {
	return secure.Secure(secure.Options{
		SSLRedirect:          strings.ToLower(os.Getenv("FORCE_SSL")) == "true",
		SSLProxyHeaders:      map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:           315360000,
		STSIncludeSubdomains: true,
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
	})
}
