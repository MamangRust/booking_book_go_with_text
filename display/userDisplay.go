package display

import (
	"booking_go_with_text/interfaces"
	"fmt"
	"strconv"
)

type TextUserDisplay struct {
	service interfaces.TextUserService
}

func NewTextUserDisplay(service interfaces.TextUserService) *TextUserDisplay {
	return &TextUserDisplay{
		service: service,
	}
}

func (t *TextUserDisplay) AllUsers() {
	allUsers := t.service.GetAllUsers()

	if len(allUsers) > 0 {
		fmt.Println("\nDaftar Pengguna Buku: ")

		for idx, user := range allUsers {
			fmt.Printf("%d. %s\n", idx+1, user)
		}
	} else {
		fmt.Println("Tidak ada pengguna buku")
	}
}

func (t *TextUserDisplay) CreateUser() {
	var name, email, password string

	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan alamat email: ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan kata sandi: ")
	fmt.Scanln(&password)

	t.service.CreateUser(name, email, password)
	fmt.Println("Pengguna buku telah ditambahkan.")
}

func (t *TextUserDisplay) UpdateUser() {
	t.AllUsers()

	var userID, newName, newEmail, newPassword string

	fmt.Print("Pilih ID pengguna buku yang ingin diperbarui: ")
	fmt.Scanln(&userID)

	userIdx, err := strconv.Atoi(userID)
	if err != nil || userIdx < 1 || userIdx > len(t.service.GetAllUsers()) {
		fmt.Println("ID pengguna buku tidak valid.")
		return
	}

	fmt.Print("Masukkan nama baru (biarkan kosong jika tidak ingin mengubah): ")
	fmt.Scanln(&newName)
	fmt.Print("Masukkan email baru (biarkan kosong jika tidak ingin mengubah): ")
	fmt.Scanln(&newEmail)
	fmt.Print("Masukkan kata sandi baru (biarkan kosong jika tidak ingin mengubah): ")
	fmt.Scanln(&newPassword)

	t.service.UpdateUser(userIdx-1, newName, newEmail, newPassword)
	fmt.Println("Informasi pengguna buku telah diperbarui.")
}

func (t *TextUserDisplay) DeleteUser() {
	t.AllUsers()

	var userID string

	fmt.Print("Pilih ID pengguna buku yang ingin dihapus: ")
	fmt.Scanln(&userID)

	userIdx, err := strconv.Atoi(userID)

	if err != nil || userIdx < 1 || userIdx > len(t.service.GetAllUsers()) {
		fmt.Println("ID pengguna buku tidak valid.")
		return
	}

	t.service.DeleteUser(userIdx - 1)
	fmt.Println("Pengguna buku telah dihapus.")

}
