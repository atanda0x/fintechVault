package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_Name"`
	LastName  string `json:"last_Name"`
	OtherName string `json:"other_Name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastName, otherName string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		FirstName: firstName,
		LastName:  lastName,
		OtherName: otherName,
		Number:    int64(rand.Intn(1000000)),
	}
}
