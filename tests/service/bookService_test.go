package service

import (
	mocks "booking_go_with_text/mocks/interfaces"
	"booking_go_with_text/service"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTextBookService(t *testing.T) {
	mockRepo := &mocks.TextBookRepository{}
	service := service.NewTextBookService(mockRepo)

	// Test GetAllBooks
	mockRepo.On("ReadAllBooks").Return([]string{"book1", "book2"})
	books := service.GetAllBooks()
	assert.Equal(t, []string{"book1", "book2"}, books)

	// Test CreateBook
	mockRepo.On("CreateBookWithDetails", "title", "author", "publish_year", "isbn")
	service.CreateBook("title", "author", "publish_year", "isbn")

	// Test UpdateBook
	mockRepo.On("UpdateBookWithDetails", 1, "title", "author", "publish_year", "isbn")
	service.UpdateBook(1, "title", "author", "publish_year", "isbn")

	// Test DeleteBook
	mockRepo.On("DeleteBook", 1)
	service.DeleteBook(1)

	mockRepo.AssertExpectations(t)
}
