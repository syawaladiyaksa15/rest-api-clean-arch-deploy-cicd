package request

import "rest-api-clean-arch/features/books"

type Book struct {
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}

func ToCore(req Book) books.Core {
	return books.Core{
		Title:     req.Title,
		Author:    req.Author,
		Publisher: req.Publisher,
	}
}
