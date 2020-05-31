package testing

import (
	"fmt"
	"testing"
)

type MockBookSuggester struct {
	SuggestBookOut Book
	SuggestErr     error
}

func (m MockBookSuggester) Suggest(_ string) (Book, error) {
	return m.SuggestBookOut, m.SuggestErr
}

func TestSuggestABook(t *testing.T) {
	tests := []struct {
		name      string
		suggester BookSuggester
		isbnIn    string
		wantBook  Book
		wantErr   bool
	}{
		{
			name: "happy path",
			suggester: MockBookSuggester{
				SuggestBookOut: Book{ISBN: "978-1633430075", Title: "Go in Practice"},
				SuggestErr:     nil,
			},
			isbnIn:   "978-0134190440",
			wantBook: Book{ISBN: "978-1633430075", Title: "Go in Practice"},
			wantErr:  false,
		},
		{
			name: "error: suggestor error is propagated",
			suggester: MockBookSuggester{
				SuggestBookOut: Book{},
				SuggestErr:     fmt.Errorf("can't suggest: 💥"),
			},
			isbnIn:   "978-0134190440",
			wantBook: Book{},
			wantErr:  true,
		},
		// Spies (here be overengineering)
		// {
		// 	name: "happy path with spies",
		// 	suggester: SpyBookSuggester{
		// 		SuggestBookOut: Book{ISBN: "978-1633430075", Title: "Go in Practice"},
		// 		SuggestErr:     nil,
		// 		SuggestAssertion: func(t *testing.T, gotISBN string) {
		// 			wantISBN := "9781633430075"
		// 			if wantISBN != gotISBN {
		// 				t.Errorf("want: %+v\ngot: %+v", wantISBN, gotISBN)
		// 			}
		// 		},
		// 	},
		// 	isbnIn:   "978-1633430075",
		// 	wantBook: Book{ISBN: "978-1633430075", Title: "Go in Practice"},
		// 	wantErr:  false,
		// },
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockServer := BookServer{suggester: test.suggester}
			book, err := mockServer.SuggestABook(test.isbnIn)
			if book != test.wantBook {
				t.Errorf("want: %+v\ngot: %+v", test.wantBook, book)
			}
			if (err != nil) != test.wantErr {
				t.Errorf("BookServer.SuggestABook() error = %v, wantErr %v", err, test.wantErr)
				return
			}
		})
	}
}

// should be at the top of the file,
// down here to lessen the mental overloading
type SpyBookSuggester struct {
	T                *testing.T
	SuggestBookOut   Book
	SuggestErr       error
	SuggestAssertion func(t *testing.T, got string)
}

func (m SpyBookSuggester) Suggest(isbn string) (Book, error) {
	if m.SuggestAssertion != nil {
		m.SuggestAssertion(m.T, isbn)
	}
	return m.SuggestBookOut, m.SuggestErr
}
