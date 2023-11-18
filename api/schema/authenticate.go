package schema

import (
	"github.com/gin-gonic/gin"
	securityJWT "github.com/techschool/simplebank/api/security/jwt"
	db "github.com/techschool/simplebank/db/sqlc"
)

// parseBearerAuth parses an HTTP Bearer Authenticate string
// "Bearer 123hkladashdflhafkja.kbnlkabl return 123hkladashdflhafkja.kbnlkabl true"
func parseBearerAuth(ctx *gin.Context) (token string, ok bool) {
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

func GetUserByBeareToken(ctx *gin.Context, jwtMaker securityJWT.Maker, store db.Store) (string, bool) {
	token, ok := parseBearerAuth(ctx)
	if !ok {
		return "", false
	}
	payload, err := jwtMaker.VerifyToken(token)
	if err != nil {
		return "", false
	}
	user, err := store.GetUser(ctx, payload.Username)
	if err != nil {
		return "", false
	}
	return user.Username, true
}
