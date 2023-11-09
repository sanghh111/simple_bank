package main

import (
	"database/sql"
	"log"

	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/uti"

	_ "github.com/lib/pq"
)

func main() {
	uti.LoadConfig(".")
	config, err := uti.GetConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Fatal(err)
	log.Fatal(config)

	conn, err := sql.Open(config.DbDriver, config.Uri_db)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	server.Start(config.ServerSource)

}
