package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	contexthelper "github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/pkg/utils/contextHelper"
)

const userAgentHeaderKey = "user-agent"
const xForwardedFor = "x-forwarded-for"
const xEnvoyExternalAddress = "x-envoy-external-address"

func SetAppContext() gin.HandlerFunc {
	return func(c *gin.Context) {

		ipAddress := c.Request.Header.Get(xEnvoyExternalAddress)

		if ipAddress == "" {
			ipsList := strings.Split(c.Request.Header.Get(xForwardedFor), ",")
			ipAddress = strings.TrimSpace(ipsList[0])
		}

		ctx := contexthelper.WithIpAddress(c.Request.Context(), ipAddress)

		userAgent := c.Request.Header.Get(userAgentHeaderKey)
		ctx = contexthelper.WithUserAgent(ctx, userAgent)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
