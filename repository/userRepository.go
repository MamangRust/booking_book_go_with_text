package repository

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TextUsersRepository struct {
	FilePath string
}

func NewTextUsersRepository(file_path string) *TextUsersRepository {
	return &TextUsersRepository{
		FilePath: file_path,
	}
}

func (t *TextUsersRepository) CreateUser(user_info string) {
	file, err := os.OpenFile(t.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)

		return
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, user_info)

	if err != nil {
		fmt.Println(err)
	}
}

func (t *TextUsersRepository) ReadAllUsers() []string {
	file, err := os.Open(t.FilePath)

	if err != nil {
		fmt.Println(err)

		return nil
	}

	defer file.Close()

	var users []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		users = append(users, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)

		return nil
	}

	return users
}

func (t *TextUsersRepository) CreateUserDetails(name string, email string, password string) {
	user_info := fmt.Sprintf("Name: %s, Email: %s, Password: %s", name, email, password)

	t.CreateUser(user_info)
}

func (t *TextUsersRepository) UpdateUserDetails(user_id int, name string, email string, password string) {
	users := t.ReadAllUsers()

	if 0 <= user_id && user_id < len(users) {
		user_info := strings.Split(strings.TrimSpace(users[user_id]), ", ")

		updated_info := map[string]string{
			"Name":     name,
			"Email":    email,
			"Password": password,
		}

		for key, value := range updated_info {
			if value == "" {
				updated_info[key] = strings.Split(user_info[getIndexUser(key)], ": ")[1]

			}
		}

		updated_user_info := fmt.Sprintf("Name: %s, Email: %s, Password: %s\n", updated_info["Name"], updated_info["Email"], updated_info["Password"])

		users[user_id] = updated_user_info

		file, err := os.OpenFile(t.FilePath, os.O_WRONLY|os.O_TRUNC, 0644)

		if err != nil {
			fmt.Println(err)
		}

		defer file.Close()

		for _, user := range users {
			fmt.Fprintln(file, user)
		}
	} else {
		fmt.Println("User ID out of range")
	}
}

func (t *TextUsersRepository) DeleteUser(user_id int) {
	users := t.ReadAllUsers()

	if 0 <= user_id && user_id < len(users) {
		users = append(users[:user_id], users[user_id+1:]...)

		file, err := os.OpenFile(t.FilePath, os.O_WRONLY|os.O_TRUNC, 0644)

		if err != nil {
			fmt.Println(err)

			return
		}

		defer file.Close()

		for _, user := range users {
			fmt.Fprintln(file, user)
		}
	} else {
		fmt.Println("User ID out of range")
	}
}

func getIndexUser(key string) int {
	switch key {
	case "Name":
		return 0
	case "Email":
		return 1
	case "Password":
		return 2
	default:
		return -1
	}
}
