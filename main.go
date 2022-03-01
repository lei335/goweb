package main

import (
	godb "go-web/db"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := godb.ConnectPQ()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
