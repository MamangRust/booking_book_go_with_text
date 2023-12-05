package service

import (
	"booking_go_with_text/interfaces"
)

type TextBookingService struct {
	BookingRepository interfaces.TextBookingRepository
}

func NewTextBookingService(repo interfaces.TextBookingRepository) *TextBookingService {
	return &TextBookingService{BookingRepository: repo}
}

func (t *TextBookingService) GetAllBookings() []string {
	return t.BookingRepository.ReadAllBookings()
}

func (t *TextBookingService) CreateBooking(tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string) {
	t.BookingRepository.CreateBookingWithDetails(tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda)
}

func (t *TextBookingService) UpdateBooking(bookingID int, tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda string) {
	t.BookingRepository.UpdateBookingWithDetails(bookingID, tglPinjam, userID, tglKembali, tglPengembalian, status, totalDenda)
}

func (t *TextBookingService) DeleteBooking(bookingID int) {
	t.BookingRepository.DeleteBooking(bookingID)
}
