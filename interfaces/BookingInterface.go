package interfaces

type TextBookingRepository interface {
	ReadAllBookings() []string
	CreateBookingWithDetails(tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string)
	UpdateBookingWithDetails(bookingID int, tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string)
	DeleteBooking(bookingID int)
}

type TextBookingService interface {
	GetAllBookings() []string
	CreateBooking(tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string)
	UpdateBooking(bookingID int, tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string)
	DeleteBooking(bookingID int)
}
