package service

import "booking_go_with_text/interfaces"

type TextBookService struct {
	repository interfaces.TextBookRepository
}

func NewTextBookService(repository interfaces.TextBookRepository) *TextBookService {
	return &TextBookService{
		repository: repository,
	}
}

func (t *TextBookService) GetAllBooks() []string {
	return t.repository.ReadAllBooks()
}

func (t *TextBookService) CreateBook(title string, author string, publish_year string, isbn string) {
	t.repository.CreateBookWithDetails(title, author, publish_year, isbn)
}

func (bs *TextBookService) UpdateBook(book_id int, title string, author string, publish_year string, isbn string) {
	bs.repository.UpdateBookWithDetails(book_id, title, author, publish_year, isbn)
}

func (bs *TextBookService) DeleteBook(book_id int) {
	bs.repository.DeleteBook(book_id)
}
