package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	GetAccountByNumber(int) (*Account, error)
}

type Postgresdb struct {
	db *sql.DB
}

func NewPostgresdb() (*Postgresdb, error) {
	conn := "user=root dbname=fintechdb password=ethereumsolana sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Postgresdb{
		db: db,
	}, nil
}

func (d *Postgresdb) Init() error {
	return d.CreateAccountTable()
}

func (d *Postgresdb) CreateAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(200),
		last_name varchar(200),
		other_name varchar(200),
		number serial,
		encrypted_password varchar(200),
		balance serial, 
		created_at timestamp
	)`
	_, err := d.db.Exec(query)
	return err
}

func (d *Postgresdb) CreateAccount(acc *Account) error {
	query := `insert into account (first_name, last_name, other_name, number, encrypted_password, balance, created_at)
	values ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := d.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.OtherName,
		acc.Number,
		acc.EncryptedPassword,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (d *Postgresdb) UpdateAccount(*Account) error {
	return nil
}

func (d *Postgresdb) DeleteAccount(id int) error {
	_, err := d.db.Query("select from account where id = $1", id)
	return err
}

func (d *Postgresdb) GetAccountByNumber(number int) (*Account, error) {
	rows, err := d.db.Query("select * from account where number = $1", number)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account with number [%d] not found", number)
}

func (d *Postgresdb) GetAccountByID(id int) (*Account, error) {
	rows, err := d.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

func (d *Postgresdb) GetAccounts() ([]*Account, error) {
	rows, err := d.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil

}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.OtherName,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
