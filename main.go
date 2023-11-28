package main

import (
	"log"
)

func seedAccount(store Storage, fname, lname, oname, pw string) *Account {
	acc, err := NewAccount(fname, lname, oname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Atanda", "Nafiu", "kolapo", "ethereumsolana")
}

func main() {
	db, err := NewPostgresdb()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	seedAccounts(db)

	server := NewAPIServer(":2000", db)
	server.Run()
}
