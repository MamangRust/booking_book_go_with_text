package main

import (
	"booking_go_with_text/display"
	"booking_go_with_text/repository"
	"booking_go_with_text/service"
	"fmt"
)

func displayMenu() {
	fmt.Println("===== Book Management System =====")
	fmt.Println("1. Tampilkan Display Buku")
	fmt.Println("2. Tampilkan Display Pengguna")
	fmt.Println("3. Tampilkan Display Peminjaman")

	fmt.Println("0. Keluar")
	fmt.Println("=================================")

}

func displayMenuBooks() {
	fmt.Println("\n===== Menu Buku =====")
	fmt.Println("1. Tampilkan Semua Buku")
	fmt.Println("2. Tambah Buku")
	fmt.Println("3. Perbarui Informasi Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("0. Kembali")
	fmt.Println("======================")
}

func displayMenuUsers() {
	fmt.Println("\n===== Menu Pengguna =====")
	fmt.Println("1. Tampilkan Pengguna Buku")
	fmt.Println("2. Tambah Pengguna Buku")
	fmt.Println("3. Perbarui Informasi Pengguna Buku")
	fmt.Println("4. Hapus Pengguna Buku")
	fmt.Println("0. Kembali")
	fmt.Println("======================")
}

func displayMenuBookings() {
	fmt.Println("\n===== Menu Peminjam =====")
	fmt.Println("1. Tampilkan Pinjam Buku")
	fmt.Println("2. Tambah Pinjam Buku")
	fmt.Println("3. Perbarui Informasi Pinjam Buku")
	fmt.Println("4. Hapus Pinjam Buku")
	fmt.Println("0. Kembali")
	fmt.Println("======================")
}

func main() {
	book_repository := repository.NewTextBookRepository("books.txt")
	users_repository := repository.NewTextUsersRepository("users.txt")
	booking_repository := repository.NewTextBookingRepository("bookings.txt")

	book_service := service.NewTextBookService(book_repository)
	user_service := service.NewTextUserService(users_repository)
	booking_service := service.NewTextBookingService(booking_repository)

	book_repository.CreateBookWithDetails(
		"The Great Gatsby", "F. Scott Fitzgerald", "1925", "9780743273565",
	)
	book_repository.CreateBookWithDetails(
		"To Kill a Mockingbird", "Harper Lee", "1960", "9780061120084",
	)
	book_repository.CreateBookWithDetails(
		"1984", "George Orwell", "1949", "9780451524935",
	)

	users_repository.CreateUserDetails(
		"John Doe", "john@example.com", "password123",
	)
	users_repository.CreateUserDetails(
		"Alice Smith", "alice@example.com", "securepwd456",
	)
	users_repository.CreateUserDetails(
		"Bob Johnson", "bob@example.com", "mysecretp@ssword",
	)

	booking_repository.CreateBookingWithDetails(
		"2023-11-01", "1", "2023-11-08", "2023-11-10", "Pinjam", "0",
	)
	booking_repository.CreateBookingWithDetails(
		"2023-11-03", "2", "2023-11-10", "2023-11-15", "Pinjam", "0",
	)
	booking_repository.CreateBookingWithDetails(
		"2023-11-05", "3", "2023-11-12", "2023-11-18", "Pinjam", "0",
	)

	for {
		displayMenu()
		var choice string
		fmt.Print("Pilih operasi yang ingin Anda lakukan: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			displayMenuBooks()
			displayBuku := display.NewTextBookDisplay(book_service)

			var choiceBuku string
			fmt.Print("Pilih operasi untuk buku: ")
			fmt.Scanln(&choiceBuku)

			switch choiceBuku {
			case "1":
				displayBuku.AllBooks()
			case "2":
				displayBuku.CreateBook()
			case "3":
				displayBuku.UpdateBook()
			case "4":
				displayBuku.DeleteBook()
			case "0":
				continue
			default:
				fmt.Println("Pilihan tidak valid. Silakan pilih operasi yang sesuai.")
			}

		case "2":
			displayMenuUsers()
			displayUser := display.NewTextUserDisplay(user_service)

			var choicePengguna string
			fmt.Print("Pilih operasi untuk pengguna buku: ")
			fmt.Scanln(&choicePengguna)

			switch choicePengguna {
			case "1":
				displayUser.AllUsers()
			case "2":
				displayUser.CreateUser()
			case "3":
				displayUser.UpdateUser()
			case "4":
				displayUser.DeleteUser()
			case "0":
				break
			default:
				fmt.Println("Pilihan tidak valid. Silakan pilih operasi yang sesuai.")
			}

		case "3":
			displayMenuBookings()
			displayBooking := display.NewTextBookingDisplay(booking_service, user_service, book_service)

			var choiceBooking string
			fmt.Print("Pilih operasi untuk booking buku: ")
			fmt.Scanln(&choiceBooking)

			switch choiceBooking {
			case "1":
				displayBooking.AllBookings()
			case "2":
				displayBooking.CreateBooking()
			case "3":
				displayBooking.UpdateBooking()
			case "4":
				displayBooking.DeleteBooking()
			case "0":
				break
			default:
				fmt.Println("Pilihan tidak valid. Silakan pilih operasi yang sesuai.")
			}

		case "0":
			fmt.Println("Terima kasih! Sampai jumpa!")
			break
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih operasi yang sesuai.")
		}
	}
}
