package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdatedAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountID(int) (*Account, error)
	GetAccountByNumber(int)(*Account,error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table  if not exists account(
	id serial primary key,
	first_name varchar(50),
	last_name varchar(50),
	number serial,
	encrypted_password varchar(150),
	balance serial,
	created_at timestamp
	);`

	_, err := s.db.Exec(query)
	return err
}

// func → Declares a function.
// (s *PostgresStore) → This is a receiver. It means this function belongs to PostgresStore.
// CreateAccount → This is the function name.
// (acc *Account) → This is the parameter (a pointer to an Account struct).
// error → This is the return type.

// This is a method attached to PostgresStore.
// The function name is still CreateAccount, but now it can only be called on a PostgresStore instance.

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query :=
		`INSERT INTO account (first_name,last_name,number,encrypted_password,balance,created_at)
	VALUES($1,$2,$3,$4,$5,$6)`
	_, err := s.db.Query(query, acc.FirstName, acc.LastName, acc.Number,acc.EncryptedPassword, acc.Balance, acc.CreatedAt)
	if err != nil {
		return err
	}
	

	return nil
}

func (s *PostgresStore) UpdatedAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	result, err := s.db.Exec("DELETE FROM account WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("account %d not found", id)
	}

	return nil

}

func (s *PostgresStore) GetAccountID(id int) (*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanintoAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account ")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanintoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func scanintoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.EncryptedPassword,
		&account.Balance,
		&account.CreatedAt)

	return account, err

}

func (s *PostgresStore) GetAccountByNumber(number int)(*Account ,error){
	rows, err := s.db.Query("SELECT * FROM account WHERE number = $1", number)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanintoAccount(rows)
	}
	return nil, fmt.Errorf("account %d not found", number)
}
