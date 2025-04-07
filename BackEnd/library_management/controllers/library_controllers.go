package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

var input = bufio.NewReader(os.Stdin)
var library = services.NewLibrary()

type LibraryController struct {
	service services.LibraryManager
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{service: service}
}
func (lbcon *LibraryController) newBook(ID int, title string, author string) *models.Book {
	book := models.Book{
		ID:     ID,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	return &book
}
func (lbcon *LibraryController) newMember(ID int, name string) *models.Member {
	member := models.Member{
		ID:            ID,
		Name:          name,
		BorrowedBooks: []models.Book{},
	}
	return &member
}
func (lbcon *LibraryController) AddBook() {
	fmt.Println("Enter ID: ")
	idStr, _ := input.ReadString('\n')
	idStr = strings.Trim(idStr, "\n")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Make sure the U Enter an Integer:")
		return
	}

	fmt.Println("Enter Title: ")
	title, _ := input.ReadString('\n')
	title = strings.Trim(title, "\n")

	fmt.Println("Enter Author: ")
	author, _ := input.ReadString('\n')
	author = strings.Trim(author, "\n")

	library.AddBook(*lbcon.newBook(id, title, author))
}
func (lbcon *LibraryController) AddMember() {
	fmt.Println("Enter ID: ")
	idStr, _ := input.ReadString('\n')
	idStr = strings.Trim(idStr, "\n")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Make sure the U Enter an Integer:")
		return
	}

	fmt.Println("Enter Name: ")
	name, _ := input.ReadString('\n')
	name = strings.Trim(name, "\n")

	library.AddMember(*lbcon.newMember(id, name))
}
