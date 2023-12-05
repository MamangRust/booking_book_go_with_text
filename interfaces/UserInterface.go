package interfaces

type TextUserRepository interface {
	CreateUserDetails(name string, email string, password string)
	ReadAllUsers() []string
	UpdateUserDetails(user_id int, name string, email string, password string)
	DeleteUser(user_id int)
}

type TextUserService interface {
	CreateUser(name string, email string, password string)
	GetAllUsers() []string
	UpdateUser(user_id int, name string, email string, password string)
	DeleteUser(user_id int)
}
