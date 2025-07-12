package library

import (
	"fmt"
	"map/utils"
	"time"
)

func AddBook(lib *Library) error {
	id := utils.GenerateId()
	title := utils.GetNoneEmptyStr("Nhập tiêu đề :")
	author := utils.GetNoneEmptyStr("Nhập tên tác giả :")

	if err := lib.AddBookStore(id, title, author); err != nil {
		return err
	}

	fmt.Println("Thêm sách thành công!")
	err := ListBook(lib)
	if err != nil {
		return err
	}
	return nil
}
func ListBook(lib *Library) error {
	books := lib.ListBooksStore()
	if len(books) == 0 {
		fmt.Println("Thư viện chưa có sách! Vui lòng thêm sách")
		return nil
	}

	for _, book := range books {
		status := "Còn"
		if book.IsBorrowed {
			status = "Đã mượn"
		}
		fmt.Printf("Id : %s | Tiêu đề : %s | Tác giả : %s | Trạng thái : %s \n", book.Id, book.Title, book.Author, status)
	}
	return nil
}
func ListBorrower(lib *Library) error {
	borrowers := lib.ListBorrowerStore()
	if len(borrowers) == 0 {
		fmt.Println("Thư viện chưa có người mượn! Vui lòng thêm người mượn!")
		return nil
	}

	for _, borrower := range borrowers {

		fmt.Printf("Id : %s | Tên người mượn : %s | Email : %s \n", borrower.Id, borrower.Name, borrower.Email)
	}
	return nil
}

func AddBorrower(lib *Library) error {
	id := utils.GenerateId()
	name := utils.GetNoneEmptyStr("Nhập tên người mượn :")
	email := utils.GetNoneEmptyStr("Nhập email :")

	if err := lib.AddBorrowerStore(id, name, email); err != nil {
		return err
	}

	fmt.Println("Thêm người mượn thành công!")
	err := ListBorrower(lib)
	if err != nil {
		return err
	}
	return nil
}

func BorrowBook(lib *Library) error {
	id := utils.GenerateId()
	bookId := utils.GetNoneEmptyStr("Nhập ID sách :")
	borrowerId := utils.GetNoneEmptyStr("Nhập ID người mượn :")
	borrowerDate := time.Now()

	if err := lib.AddBorrowBookStore(id, bookId, borrowerId, borrowerDate); err != nil {
		return err
	}
	return nil
}
func HistoryBorrowerBook(lib *Library) error {
	borrowerId := utils.GetNoneEmptyStr("Nhap ID cua nguoi muon :")
	history := lib.HistoryBorrower(borrowerId)
	fmt.Printf("Lich su nguoi muon voi ID: %s \n", borrowerId)
	for _, trans := range history {
		var returnDate = "Chua tra"
		if trans.ReturnDate.IsZero() {
			returnDate = trans.ReturnDate.Format("2006-01-02")
		}
		fmt.Printf("\nId : %s |Id sach : %s |Ngay muon :%s |Ngay tra :%s ", trans.Id, lib.GetTitleBookStore(trans.BookId), trans.BorrowerDate.Format("2006-01-02"), returnDate)
		fmt.Println()
	}
	return nil
}
func ReturnBook(lib *Library) error {
	transactionId := utils.GetNoneEmptyStr("Nhap ID giao dich :")
	if err := lib.ReturnBookStore(transactionId); err != nil {
		return err
	}
	fmt.Println("Tra sach thanh cong!")
	return nil
}

func FindBook(lib *Library) error {
	query := utils.GetNoneEmptyStr("Nhap tieu de hay ten tac gia :")
	books := lib.FindBookStore(query)
	if len(books) == 0 {
		fmt.Println("Thu vien khong co sach!")
		return nil
	}
	for _, book := range books {
		status := "Còn"
		if book.IsBorrowed {
			status = "Đã mượn"
		}
		fmt.Printf("Id : %s | Tiêu đề : %s | Tác giả : %s | Trạng thái : %s \n", book.Id, book.Title, book.Author, status)
	}
	return nil
}
