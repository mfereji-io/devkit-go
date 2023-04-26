package middleware

//TODO: implement or depend on istio/envoy+jaeger
import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	contexthelper "github.com/mfereji-io/devkit-go/example/kaska-svc-api/internal/pkg/utils/contextHelper"
)

const requestIdHeaderKey = "x-request-id"

func SetRequestId() gin.HandlerFunc {

	return func(c *gin.Context) {

		requestId := uuid.New().String()
		ctx := contexthelper.WithRequestId(c.Request.Context(), requestId)
		c.Request = c.Request.WithContext(ctx)
		c.Header(requestIdHeaderKey, requestId)
		c.Next()

	}
}
