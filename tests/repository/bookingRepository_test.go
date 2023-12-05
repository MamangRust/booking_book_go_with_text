package repository

import (
	mocks "booking_go_with_text/mocks/interfaces"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTextBookingRepository(t *testing.T) {
	mockRepo := &mocks.TextBookingRepository{}
	bookings := []string{"booking1", "booking2"}

	// Test ReadAllBookings
	mockRepo.On("ReadAllBookings").Return(bookings)
	assert.Equal(t, bookings, mockRepo.ReadAllBookings())

	// Test CreateBookingWithDetails
	mockRepo.On("CreateBookingWithDetails", "tglPinjam", "userID", "tglKembali", "tglPengembalian", "status", "totalDenda")
	mockRepo.CreateBookingWithDetails("tglPinjam", "userID", "tglKembali", "tglPengembalian", "status", "totalDenda")

	// Test UpdateBookingWithDetails
	mockRepo.On("UpdateBookingWithDetails", 1, "tglPinjam", "userID", "tglKembali", "tglPengembalian", "status", "totalDenda")
	mockRepo.UpdateBookingWithDetails(1, "tglPinjam", "userID", "tglKembali", "tglPengembalian", "status", "totalDenda")

	// Test DeleteBooking
	mockRepo.On("DeleteBooking", 1)
	mockRepo.DeleteBooking(1)

	mockRepo.AssertExpectations(t)
}
