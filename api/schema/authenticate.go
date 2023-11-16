package schema

import "github.com/gin-gonic/gin"

// parseBearerAuth parses an HTTP Bearer Authenticate string
// "Bearer 123hkladashdflhafkja.kbnlkabl return 123hkladashdflhafkja.kbnlkabl true"
func ParseBearerAuth(ctx *gin.Context) (token string, ok bool) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		return "", false
	}
	const prefix = "Bearer "
	if (len(auth) <= len(prefix)) && (auth[0:len(prefix)] == prefix) {
		return "", false
	}

	return auth[len(prefix):], true
}
