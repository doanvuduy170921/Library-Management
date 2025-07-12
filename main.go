package main

import (
	"fmt"
	"map/library"
	"map/utils"
)

func main() {
	lib := library.NewLibrary()

	for {
		utils.ClearScreen()
		fmt.Println("CHƯƠNG TRÌNH QUẢN LÝ THƯ VIỆN")
		fmt.Println("1.THÊM SÁCH")
		fmt.Println("2.XEM DANH SÁCH SÁCH")
		fmt.Println("3.THÊM NGƯỜI MƯỢN")
		fmt.Println("4.XEM DANH SÁCH NGƯỜI MƯỢN")
		fmt.Println("5.MƯỢN SÁCH")
		fmt.Println("6.XEM LỊCH SỬ MƯỢN")
		fmt.Println("7.TRẢ SÁCH")
		fmt.Println("8.TÌM KIẾM SÁCH")
		fmt.Println("9.THOÁT")

		choise := utils.GetPositiveInt("CHỌN CHỨC NĂNG : ")
		switch choise {
		case 1:
			fmt.Println("=========Thêm sách=========")
			if err := library.AddBook(lib); err != nil {
				fmt.Printf("Loi khi them sach :%s", err)
			}
		case 2:
			fmt.Println("=========Xem danh sách sách=========")
			if err := library.ListBook(lib); err != nil {
				fmt.Printf("Loi khi xem danh sach sach :%s", err)
			}
		case 3:
			fmt.Println("=========Thêm người mượn=========")
			if err := library.AddBorrower(lib); err != nil {
				fmt.Printf("Lỗi khi thêm người mượn :%s", err)
			}
		case 4:
			fmt.Println("=========Xem danh sách người mượn=========")
			if err := library.ListBorrower(lib); err != nil {
				fmt.Printf("Lỗi khi xem danh sách người mượn:%s", err)
			}
		case 5:
			fmt.Println("=========Mượn sách=========")
			if err := library.BorrowBook(lib); err != nil {
				fmt.Printf("Lỗi khi mượn sách: %v\n", err)
			}
		case 6:
			fmt.Println("=========Lịch sử mượn sách=========")
			if err := library.HistoryBorrowerBook(lib); err != nil {
				fmt.Printf("Lỗi khi xem lịch sử mượn sách :%v", err)
			}
		case 7:
			fmt.Println("=========Trả sách=========")
			if err := library.ReturnBook(lib); err != nil {
				fmt.Printf("Lỗi khi trả sách :%s", err)
			}
		case 8:
			fmt.Println("=========Tìm kiếm sách=========")
			if err := library.FindBook(lib); err != nil {
				fmt.Printf("Lỗi khi tìm sách :%s", err)
			}
		case 9:
			return
		default:
			fmt.Println("Vui lòng nhập giá trị hợp lệ!")
		}
		utils.ReadInput("Nhấn Enter để tiếp tục...")
	}
}
