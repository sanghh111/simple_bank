package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/techschool/simplebank/api"
	securityJWT "github.com/techschool/simplebank/api/security/jwt"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/uti"

	_ "github.com/lib/pq"
)

func main() {
	uti.LoadConfig(".")
	config, err := uti.GetConfig()
	if err != nil {
		log.Fatal("cannot connect to config: ", err)
		return
	}

	conn, err := sql.Open(config.DB_DRIVER, config.URI_DB)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
		return
	}
	d, err := time.ParseDuration(config.DURATION)
	if err != nil {
		log.Fatal("Cannot parse Duration: ", err)
		return
	}
	optionJwt := securityJWT.OptionJwt{
		TimeDuration: d,
	}
	marker, err := securityJWT.NewJWTMaker(config.APIKEY, optionJwt)
	if err != nil {
		log.Fatal("cannot load newJWT Maker: ", err)
		return
	}

	store := db.NewStore(conn)
	server := api.NewServer(store, marker)

	server.Start(config.SERVER_SOURCE)

}
