package response

import (
	"rest-api-clean-arch/features/books"
	"time"
)

type Book struct {
	ID        int       `json:"id" form:"id"`
	Title     string    `json:"title" form:"title"`
	Author    string    `json:"author" form:"author"`
	Publisher string    `json:"publisher" form:"publisher"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

func FormCore(data books.Core) Book {
	return Book{
		ID:        data.ID,
		Title:     data.Title,
		Author:    data.Author,
		Publisher: data.Publisher,
		CreatedAt: data.CreatedAt,
	}
}

func FromCoreList(data []books.Core) []Book {
	result := []Book{}

	for k, _ := range data {
		result = append(result, FormCore(data[k]))
	}

	return result
}
