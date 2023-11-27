package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	OtherName string `json:"other_name"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_Name"`
	LastName  string    `json:"last_Name"`
	OtherName string    `json:"other_Name"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName, otherName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		OtherName: otherName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}
