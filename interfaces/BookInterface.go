package interfaces

type TextBookRepository interface {
	CreateBook(book_info string)
	ReadAllBooks() []string
	CreateBookWithDetails(title string, author string, publish_year string, isbn string)
	UpdateBookWithDetails(book_id int, title string, author string, publish_year string, isbn string)
	DeleteBook(book_id int)
}

type TextBookService interface {
	GetAllBooks() []string
	CreateBook(title string, author string, publish_year string, isbn string)
	UpdateBook(book_id int, title string, author string, publish_year string, isbn string)
	DeleteBook(book_id int)
}
