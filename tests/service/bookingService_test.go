package service

import (
	mocks "booking_go_with_text/mocks/interfaces"
	"booking_go_with_text/service"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTextBookingService(t *testing.T) {
	mockRepo := &mocks.TextBookingRepository{}
	service := service.NewTextBookingService(mockRepo)

	// Test GetAllBookings
	mockRepo.On("ReadAllBookings").Return([]string{"booking1", "booking2"})
	bookings := service.GetAllBookings()

	assert.Equal(t, []string{"booking1", "booking2"}, bookings)

	// Test CreateBooking
	mockRepo.On("CreateBookingWithDetails", "tglPinjam", "userID", "tglKembali", "tglPengembalian", "status", "totalDenda")
	service.CreateBooking("tglPinjam", "userID", "tglKembali", "tglPengembalian", "status", "totalDenda")

	// Test UpdateBooking
	mockRepo.On("UpdateBookingWithDetails", 1, "tglPinjam", "userID", "tglKembali", "tglPengembalian", "status", "totalDenda")
	service.UpdateBooking(1, "tglPinjam", "userID", "tglKembali", "tglPengembalian", "status", "totalDenda")

	// Test DeleteBooking
	mockRepo.On("DeleteBooking", 1)
	service.DeleteBooking(1)

	mockRepo.AssertExpectations(t)
}
