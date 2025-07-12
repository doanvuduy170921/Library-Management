package models

import "time"

type Book struct {
	Id         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Borrower struct {
	Id    string
	Name  string
	Email string
}

type Transaction struct {
	Id           string
	BookId       string
	BorrowerId   string
	BorrowerDate time.Time
	ReturnDate   time.Time
}
