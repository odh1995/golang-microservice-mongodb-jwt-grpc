package main

import (
	"flag"
	"golang_microservice_mongodb_kub_jwt_grpc/db"
	"log"

	"github.com/joho/godotenv"
)

var (
	local bool
)

func init() {
	flag.BoolVar(&local, "local", true, "run service local")
	flag.Parse()
}

func main() {

	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panic(err)
		}
	}
	cfg := db.NewConfig()

	// get Client, Context, CancelFunc and err from connect method.
	conn, err := db.NewConnection(cfg)
	if err != nil {
		panic(err)
	}
	// Release resource when main function is returned.
	defer conn.Close()
}
