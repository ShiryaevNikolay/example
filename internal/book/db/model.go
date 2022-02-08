package db

import (
	"database/sql"

	"github.com/ShiryaevNikolay/example/internal/author"
	"github.com/ShiryaevNikolay/example/internal/book"
)

type Book struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Age     sql.NullInt32   `json:"age"`
	Authors []author.Author `json:"authors"`
}

func (m *Book) ToDomain() book.Book {
	b := book.Book{
		ID:   m.ID,
		Name: m.Name,
	}
	if m.Age.Valid {
		b.Age = int(m.Age.Int32)
	}
	return b
}
