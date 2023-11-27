package main

import (
	"log"
)

func main() {
	db, err := NewPostgresdb()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":2000", db)
	server.Run()
}
