package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/techschool/simplebank/api/schema"
	securityJWT "github.com/techschool/simplebank/api/security/jwt"
	db "github.com/techschool/simplebank/db/sqlc"
)

// Server server HTTP requests for our banking service
type Server struct {
	store     db.Store
	route     *gin.Engine
	jwtMarker securityJWT.Maker
}

func NewServer(sotre db.Store, marker securityJWT.Maker) *Server {

	server := &Server{store: sotre, jwtMarker: marker}
	router := gin.Default()

	router.POST("/accounts/", server.createAccount)
	router.POST("/user/", server.createUser)
	router.GET("/account/:id", server.getAccountById)
	router.GET("/accounts/", server.getListAccount)
	router.POST("/login/", server.login)
	router.GET("/info/", server.getInfo)

	router.POST("/transfertx/", server.transferMoney)

	server.route = router
	return server
}

// Start run the HTTP server a specific address
func (server *Server) Start(address string) error {
	return server.route.Run(address)
}

func errResponse(err error, requestId string, requestTime time.Time) gin.H {
	return gin.H{"error": schema.GetResponseError(err, requestId, requestTime)}
}
