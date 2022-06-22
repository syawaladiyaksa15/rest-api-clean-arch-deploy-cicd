package data

import (
	"rest-api-clean-arch/features/books"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	// gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Title     string `json:"title" form:"title" gorm:"not null; type:varchar(200)"`
	Author    string `json:"author" form:"author" gorm:"not null; type:varchar(100)"`
	Publisher string `json:"publisher" form:"publisher" gorm:"not null; type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (data *Book) toCore() books.Core {
	return books.Core{
		ID:        int(data.ID),
		Title:     data.Title,
		Author:    data.Author,
		Publisher: data.Publisher,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func toCoreList(data []Book) []books.Core {
	result := []books.Core{}

	for key := range data {
		result = append(result, data[key].toCore())
	}

	return result
}

func formCore(core books.Core) Book {
	return Book{
		Title:     core.Title,
		Author:    core.Author,
		Publisher: core.Publisher,
	}
}
