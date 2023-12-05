package display

import (
	"booking_go_with_text/interfaces"
	"fmt"
	"strconv"
)

type TextBookDisplay struct {
	service interfaces.TextBookService
}

func NewTextBookDisplay(service interfaces.TextBookService) *TextBookDisplay {
	return &TextBookDisplay{
		service: service,
	}
}

func (t *TextBookDisplay) AllBooks() {
	allbooks := t.service.GetAllBooks()

	if len(allbooks) > 0 {
		fmt.Println("\nDaftar Buku: ")

		for idx, book := range allbooks {
			fmt.Printf("%d. %s\n", idx+1, book)
		}
	} else {
		fmt.Println("Tidak ada buku")
	}
}

func (t *TextBookDisplay) CreateBook() {
	var title, author, publishYear, isbn string

	fmt.Print("Masukkan judul buku: ")
	fmt.Scanln(&title)
	fmt.Print("Masukkan nama pengarang: ")
	fmt.Println(&author)
	fmt.Print("Masukkan tahun terbit: ")
	fmt.Scanln(&publishYear)
	fmt.Print("Masukkan ISBN Buku: ")
	fmt.Scanln(&isbn)

	t.service.CreateBook(title, author, publishYear, isbn)

	fmt.Println("Buku berhasil ditambahkan.")
}

func (t *TextBookDisplay) UpdateBook() {
	t.AllBooks()

	var bookID, newTitle, newAuthor, newPublishYear, newISBN string

	fmt.Print("Pilih ID buku yang ingin diperbarui: ")

	fmt.Scanln(&bookID)

	bookIdx, err := strconv.Atoi(bookID)

	if err != nil || bookIdx < 1 || bookIdx > len(t.service.GetAllBooks()) {
		fmt.Println("ID buku tidak valid.")

		return
	}

	fmt.Print("Masukkan judul baru (kosongkan jika tidak ingin diubah): ")
	fmt.Scanln(&newTitle)
	fmt.Print("Masukkan pengarang baru (kosongkan jika tidak ingin diubah): ")
	fmt.Scanln(&newAuthor)
	fmt.Print("Masukkan tahun terbit baru (kosongkan jika tidak ingin diubah): ")
	fmt.Scanln(&newPublishYear)
	fmt.Print("Masukkan ISBN baru (kosongkan jika tidak ingin diubah): ")
	fmt.Scanln(&newISBN)

	t.service.UpdateBook(bookIdx-1, newTitle, newAuthor, newPublishYear, newISBN)

	fmt.Println("Buku berhasil diperbarui.")
}

func (t *TextBookDisplay) DeleteBook() {
	t.AllBooks()

	var bookID string

	fmt.Print("Pilih ID pengguna buku yang ingin dihapus: ")

	fmt.Scanln(&bookID)

	bookIdx, err := strconv.Atoi(bookID)

	if err != nil || bookIdx < 1 || bookIdx > len(t.service.GetAllBooks()) {
		fmt.Println("ID buku tidak valid.")
		return
	}

	t.service.DeleteBook(bookIdx - 1)

	fmt.Println("Buku telah dihapus.")
}
