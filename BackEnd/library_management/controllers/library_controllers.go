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

func newBook(ID int, title string, author string) *models.Book {
	book := models.Book{
		ID:     ID,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	return &book
}
func newMember(ID int, name string) *models.Member {
	member := models.Member{
		ID:            ID,
		Name:          name,
		BorrowedBooks: []models.Book{},
	}
	return &member
}
func AddBook() {
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

	library.AddBook(*newBook(id, title, author))
}
func AddMember() {
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

	library.AddMember(*newMember(id, name))
}

func Start() {
	library.AddBook(*newBook(9781982103750, "Becoming Bulletproof", "Evy Poumpouras"))
	library.AddBook(*newBook(9781934255155, "The Art of War", "Sun Tzu & Lionel Giles & Shawn Conners"))
	library.AddBook(*newBook(9781476788654, "How Champions Think", "Bob Rotella"))
	fmt.Println("Wellcome to Library\nChoose from the options:-")
	fmt.Println("0.see available books")
	fmt.Println("----------------------------------")
	S, _ := input.ReadString('\n')
	fmt.Println("----------------------------------")
	S = strings.Trim(S, "\n")

	choice, err := strconv.Atoi(S)

	if err != nil || choice > 10 {
		fmt.Println("Make sure the U Entered Number:")
		return
	}
	if choice == 0 {
		for _,k := range library.ListAvailableBooks() {
			fmt.Println(k,"\t")
		}
	}

}
