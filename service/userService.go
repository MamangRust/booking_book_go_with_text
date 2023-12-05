package service

import (
	"booking_go_with_text/interfaces"
)

type TextUsersService struct {
	UsersRepository interfaces.TextUserRepository
}

func NewTextUserService(repo interfaces.TextUserRepository) *TextUsersService {
	return &TextUsersService{
		UsersRepository: repo,
	}
}

func (t *TextUsersService) CreateUser(name string, email string, password string) {
	t.UsersRepository.CreateUserDetails(name, email, password)
}

func (t *TextUsersService) GetAllUsers() []string {
	return t.UsersRepository.ReadAllUsers()
}

func (t *TextUsersService) UpdateUser(user_id int, name string, email string, password string) {
	t.UsersRepository.UpdateUserDetails(user_id, name, email, password)

}

func (t *TextUsersService) DeleteUser(user_id int) {
	t.UsersRepository.DeleteUser(user_id)
}
