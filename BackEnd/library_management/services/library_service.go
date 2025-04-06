package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	members map[int]models.Member
	books   map[int]models.Book
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}
func (l Library) AddBook(book models.Book) {
	l.books[book.ID] = book
}
func (l Library) RemoveBook(book models.Book) {
	delete(l.books, book.ID)
}
func (l Library) BorrowBook(memberID int, bookID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book doesn't exist")
	}
	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member doesn't exist")
	}
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}
