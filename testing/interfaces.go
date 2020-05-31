package testing

import "strings"

// https://deliveroo.engineering/2019/05/17/testing-go-services-using-interfaces.html

// A Book is a book
type Book struct {
	ISBN  string
	Title string
}

// BookSuggester suggests a book
type BookSuggester interface {
	// Suggest a book, based off of another book
	// ISBN must not include spaces or hyphens.
	Suggest(ISBN string) (Book, error)
}

// A BookServer is a server
type BookServer struct {
	suggester BookSuggester
}

// ommitted http server code for brevity

// SuggestABook given an ISBN suggest a book to read
// Accepts hyphens and spaces in the ISBN string
func (s *BookServer) SuggestABook(ISBN string) (Book, error) {
	ISBN = strings.ReplaceAll(ISBN, " ", "")
	ISBN = strings.ReplaceAll(ISBN, "-", "")
	return s.suggester.Suggest(ISBN)
}
