package data

import (
	"fmt"
	"rest-api-clean-arch/features/books"

	"gorm.io/gorm"
)

type mysqlBookRepository struct {
	db *gorm.DB
}

func NewBookRepository(conn *gorm.DB) books.Data {
	return &mysqlBookRepository{
		db: conn,
	}
}

func (repo *mysqlBookRepository) SelectData() (response []books.Core, err error) {
	var dataBook []Book

	result := repo.db.Find(&dataBook)

	if result.Error != nil {
		return []books.Core{}, result.Error
	}

	return toCoreList(dataBook), nil
}

func (repo *mysqlBookRepository) CreateData(newBook books.Core) (response books.Core, err error) {

	book := formCore(newBook)

	result := repo.db.Create(&book)

	if result.Error != nil {
		return books.Core{}, result.Error
	}

	if result.RowsAffected != 1 {
		return books.Core{}, fmt.Errorf("failed to insert book")
	}

	return book.toCore(), nil
}

func (repo *mysqlBookRepository) FirstData(id int) (response books.Core, err error) {
	var dataBook Book

	result := repo.db.Find(&dataBook, id)

	if result.RowsAffected != 1 {
		return books.Core{}, fmt.Errorf("book not found")
	}

	if result.Error != nil {
		return books.Core{}, result.Error
	}

	return dataBook.toCore(), nil
}

func (repo *mysqlBookRepository) DestroyData(id int) (response int, err error) {
	var dataBook Book

	result := repo.db.Delete(&dataBook, id)

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("book not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlBookRepository) EditData(editBook books.Core, id int) (response books.Core, err error) {
	book := formCore(editBook)

	result := repo.db.Model(Book{}).Where("id = ?", id).Updates(&book)

	if result.RowsAffected != 1 {
		return books.Core{}, fmt.Errorf("book not found")
	}

	if result.Error != nil {
		return books.Core{}, result.Error
	}

	book.ID = uint(id)

	return book.toCore(), nil
}
