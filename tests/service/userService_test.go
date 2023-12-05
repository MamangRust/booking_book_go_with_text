package service

import (
	mocks "booking_go_with_text/mocks/interfaces"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestTextUserRepository(t *testing.T) {
	mockRepo := &mocks.TextUserRepository{}
	users := []string{"user1", "user2"}

	// Test ReadAllUsers
	mockRepo.On("ReadAllUsers").Return(users)
	assert.Equal(t, users, mockRepo.ReadAllUsers())

	// Test CreateUserDetails
	mockRepo.On("CreateUserDetails", "name", "email", "password")
	mockRepo.CreateUserDetails("name", "email", "password")

	// Test UpdateUserDetails
	mockRepo.On("UpdateUserDetails", 1, "name", "email", "password")
	mockRepo.UpdateUserDetails(1, "name", "email", "password")

	// Test DeleteUser
	mockRepo.On("DeleteUser", 1)
	mockRepo.DeleteUser(1)

	mockRepo.AssertExpectations(t)
}
