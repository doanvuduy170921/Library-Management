package library

import (
	"fmt"
	"map/models"
	"strings"
	"time"
)

type Library struct {
	book        map[string]models.Book
	borrower    map[string]models.Borrower
	transaction map[string]models.Transaction
}

func NewLibrary() *Library {
	return &Library{
		book:        make(map[string]models.Book),
		borrower:    make(map[string]models.Borrower),
		transaction: make(map[string]models.Transaction),
	}
}

func (lib *Library) AddBookStore(id, title, author string) error {
	if _, exists := lib.book[id]; exists {
		return fmt.Errorf("Sách với ID %s đã tồn tại \n", id)
	}

	lib.book[id] = models.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}
	return nil
}
func (lib *Library) AddBorrowerStore(id, name, email string) error {
	if _, exists := lib.borrower[id]; exists {
		return fmt.Errorf("Người mượn với ID %s đã tồn tại! \n", id)
	}

	lib.borrower[id] = models.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return nil
}

func (lib *Library) ListBooksStore() []models.Book {
	books := make([]models.Book, 0, len(lib.book))

	for _, book := range lib.book {
		books = append(books, book)
	}
	return books
}

func (lib *Library) ListBorrowerStore() []models.Borrower {
	borrowers := make([]models.Borrower, 0, len(lib.borrower))

	for _, borrower := range lib.borrower {
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func (lib *Library) AddBorrowBookStore(id, bookId, borrowerId string, borrowerDate time.Time) error {
	book, exists := lib.book[bookId]
	if !exists {
		return fmt.Errorf("sách với ID: %s không tồn tại", bookId)
	}
	if book.IsBorrowed {
		return fmt.Errorf("sách %s đã được mượn", book.Title)
	}
	_, ok := lib.borrower[borrowerId]
	if !ok {
		return fmt.Errorf("người mượn với ID: %s không tồn tại", bookId)
	}
	_, transactionExists := lib.transaction[id]
	if transactionExists {
		return fmt.Errorf("Giao dịch không tồn tại với ID:%s ", id)
	}
	book.IsBorrowed = true
	lib.book[bookId] = book

	lib.transaction[id] = models.Transaction{
		Id:           id,
		BookId:       bookId,
		BorrowerId:   borrowerId,
		BorrowerDate: borrowerDate,
	}
	fmt.Println("Mượn sách thành công!")

	transactions := lib.transaction
	for _, transaction := range transactions {
		tempDate := "Chưa trả"
		if transaction.ReturnDate.IsZero() {
			tempDate = transaction.ReturnDate.Format("2006-01-02")
		}
		fmt.Printf("Id : %s | Tên sách : %s | BorrowerId : %s |Ngày muợn : %s |Ngày trả : %s \n", transaction.Id, lib.book[bookId].Title, transaction.BorrowerId, transaction.BorrowerDate, tempDate)
	}
	return nil
}

func (lib *Library) HistoryBorrower(borrowerId string) []models.Transaction {
	if _, borrowerExists := lib.borrower[borrowerId]; !borrowerExists {
		return nil
	}
	var history []models.Transaction
	for _, trans := range lib.transaction {
		if trans.BorrowerId == borrowerId {
			history = append(history, trans)
		}
	}
	return history
}

func (lib *Library) GetTitleBookStore(bookId string) string {
	book := lib.book[bookId]
	return book.Title
}

func (lib *Library) ReturnBookStore(transactionId string) error {
	transaction, exists := lib.transaction[transactionId]
	if !exists {
		return fmt.Errorf("giao dich voi id : %s khong ton tai", transactionId)
	}
	if !lib.transaction[transactionId].ReturnDate.IsZero() {
		return fmt.Errorf("giao dich voi id : %s da tra sach", transactionId)
	}
	book, exists := lib.book[transaction.BookId]
	if !exists {
		return fmt.Errorf("sach voi id : %s khong ton tai", transaction.BookId)
	}
	book.IsBorrowed = false
	lib.book[transaction.BookId] = book

	transaction.ReturnDate = time.Now()
	lib.transaction[transactionId] = transaction
	return nil
}
func (lib *Library) FindBookStore(query string) []models.Book {
	books := make([]models.Book, 0)
	for _, book := range lib.book {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(query)) || strings.Contains(strings.ToLower(book.Author), strings.ToLower(query)) {
			books = append(books, book)
		}
	}
	return books
}
