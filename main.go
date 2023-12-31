package main

import (
	"flag"
	"fmt"
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

	fmt.Println("new Account =>", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Atanda", "Nafiu", "kolapo", "ethereumsolana")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	db, err := NewPostgresdb()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the db")
		seedAccounts(db)
	}

	server := NewAPIServer(":2000", db)
	server.Run()
}
