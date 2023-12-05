package display

import (
	"booking_go_with_text/interfaces"
	"fmt"
	"strconv"
	"time"
)

type TextBookingDisplay struct {
	serviceBooking interfaces.TextBookingService
	serviceUser    interfaces.TextUserService
	serviceBook    interfaces.TextBookService
}

func NewTextBookingDisplay(serviceBooking interfaces.TextBookingService, serviceUser interfaces.TextUserService, serviceBook interfaces.TextBookService) *TextBookingDisplay {
	return &TextBookingDisplay{
		serviceBooking: serviceBooking,
		serviceUser:    serviceUser,
		serviceBook:    serviceBook,
	}
}

func (t *TextBookingDisplay) AllBookings() {
	allBookings := t.serviceBooking.GetAllBookings()

	if len(allBookings) > 0 {
		fmt.Println("\nDaftar Booking Buku")

		for idx, booking := range allBookings {
			fmt.Printf("%d. %s\n", idx+1, booking)
		}
	} else {
		fmt.Println("Tidak ada booking buku")
	}
}

func (bs *TextBookingDisplay) CreateBooking() {
	// Display lists of users and books
	fmt.Println("\n Daftar User")
	fmt.Println(bs.serviceUser.GetAllUsers())

	fmt.Println("\n Daftar Book")
	fmt.Println(bs.serviceBook.GetAllBooks())

	// Collect booking details from user input
	var tgl_pinjam, tgl_kembali, tgl_pengembalian time.Time
	var user_id, status string
	var total_denda int
	fmt.Print("Masukkan Tanggal Peminjaman buku (YYYY-MM-DD): ")
	fmt.Scanln(&tgl_pinjam)
	fmt.Print("Masukkan User id: ")
	fmt.Scanln(&user_id)
	fmt.Print("Masukkan Tanggal Kembali Buku (YYYY-MM-DD): ")
	fmt.Scanln(&tgl_kembali)
	fmt.Print("Masukkan Tanggal Pengembalian Buku (YYYY-MM-DD): ")
	fmt.Scanln(&tgl_pengembalian)
	fmt.Print("Masukkan Status (Kembali/Pinjam): ")
	fmt.Scanln(&status)
	fmt.Print("Masukkan Denda buku (Jika terjadi keterlambatan kembali buku): ")
	fmt.Scanln(&total_denda)

	// Convert input dates to datetime objects
	var err error
	if tgl_pinjam, err = convertToTime(tgl_pinjam.String()); err != nil {
		fmt.Println("Format tanggal salah. Gunakan format YYYY-MM-DD.")
		return
	}
	if tgl_kembali, err = convertToTime(tgl_kembali.String()); err != nil {
		fmt.Println("Format tanggal salah. Gunakan format YYYY-MM-DD.")
		return
	}
	if tgl_pengembalian, err = convertToTime(tgl_pengembalian.String()); err != nil {
		fmt.Println("Format tanggal salah. Gunakan format YYYY-MM-DD.")
		return
	}

	bs.serviceBooking.CreateBooking(
		tgl_pinjam.String(), user_id, tgl_kembali.String(), tgl_pengembalian.String(), status, strconv.Itoa(total_denda),
	)
}

func (t *TextBookingDisplay) UpdateBooking() {
	t.AllBookings()

	var booking_id int
	fmt.Print("Pilih ID booking buku yang ingin diperbarui: ")
	fmt.Scanln(&booking_id)
	booking_id -= 1

	if 0 <= booking_id && booking_id < len(t.serviceBooking.GetAllBookings()) {
		// Display lists of users and books for reference
		fmt.Println("\n Daftar User")
		fmt.Println(t.serviceUser.GetAllUsers())

		fmt.Println("\n Daftar Book")
		fmt.Println(t.serviceBook.GetAllBooks())

		// Collect updated booking details from user input
		var tgl_pinjam, tgl_kembali, tgl_pengembalian time.Time
		var user_id, status string
		var total_denda int
		fmt.Print("Masukkan Tanggal Peminjaman buku (YYYY-MM-DD): ")
		fmt.Scanln(&tgl_pinjam)
		fmt.Print("Masukkan User id: ")
		fmt.Scanln(&user_id)
		fmt.Print("Masukkan Tanggal Kembali Buku (YYYY-MM-DD): ")
		fmt.Scanln(&tgl_kembali)
		fmt.Print("Masukkan Tanggal Pengembalian Buku (YYYY-MM-DD): ")
		fmt.Scanln(&tgl_pengembalian)
		fmt.Print("Masukkan Status (Kembali/Pinjam): ")
		fmt.Scanln(&status)
		fmt.Print("Masukkan Denda buku (Jika terjadi keterlambatan kembali buku): ")
		fmt.Scanln(&total_denda)

		// Convert input dates to datetime objects
		var err error
		if tgl_pinjam, err = time.Parse("2006-01-02", tgl_pinjam.String()); err != nil {
			fmt.Println("Format tanggal salah. Gunakan format YYYY-MM-DD.")
			return
		}
		if tgl_kembali, err = time.Parse("2006-01-02", tgl_kembali.String()); err != nil {
			fmt.Println("Format tanggal salah. Gunakan format YYYY-MM-DD.")
			return
		}
		if tgl_pengembalian, err = time.Parse("2006-01-02", tgl_pengembalian.String()); err != nil {
			fmt.Println("Format tanggal salah. Gunakan format YYYY-MM-DD.")
			return
		}

		// Call the service to update the chosen booking
		t.serviceBooking.UpdateBooking(
			booking_id,
			tgl_pinjam.String(),
			user_id,
			tgl_kembali.String(),
			tgl_pengembalian.String(),
			status,
			strconv.Itoa(total_denda),
		)

		fmt.Println("Informasi booking buku telah diperbarui.")
	} else {
		fmt.Println("ID booking buku tidak valid.")
	}
}

func (t *TextBookingDisplay) DeleteBooking() {
	t.AllBookings()

	var bookingID string

	fmt.Print("Pilih ID Booking buku yang ingin dihapus: ")

	fmt.Scanln(&bookingID)

	bookingIdx, err := strconv.Atoi(bookingID)

	if err != nil || bookingIdx < 1 || bookingIdx > len(t.serviceBooking.GetAllBookings()) {
		fmt.Println("ID buku tidak valid.")
		return
	}

	t.serviceBooking.DeleteBooking(bookingIdx - 1)

	fmt.Println("Buku telah dihapus.")
}

func convertToTime(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}
