package repository

import (
	mocks "booking_go_with_text/mocks/interfaces"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTextBookRepository(t *testing.T) {
	mockRepo := &mocks.TextBookRepository{}
	books := []string{"book1", "book2"}

	// Test ReadAllBooks
	mockRepo.On("ReadAllBooks").Return(books)
	assert.Equal(t, books, mockRepo.ReadAllBooks())

	// Test CreateBookWithDetails
	mockRepo.On("CreateBookWithDetails", "title", "author", "publish_year", "isbn")
	mockRepo.CreateBookWithDetails("title", "author", "publish_year", "isbn")

	// Test UpdateBookWithDetails
	mockRepo.On("UpdateBookWithDetails", 1, "title", "author", "publish_year", "isbn")
	mockRepo.UpdateBookWithDetails(1, "title", "author", "publish_year", "isbn")

	// Test DeleteBook
	mockRepo.On("DeleteBook", 1)
	mockRepo.DeleteBook(1)

	mockRepo.AssertExpectations(t)
}
