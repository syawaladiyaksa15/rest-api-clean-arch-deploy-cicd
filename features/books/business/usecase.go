package business

import (
	"errors"
	"rest-api-clean-arch/features/books"
)

type bookUseCase struct {
	bookData books.Data
}

func NewBookBusiness(bkData books.Data) books.Business {
	return &bookUseCase{
		bookData: bkData,
	}
}

func (uc *bookUseCase) GetAllData() (response []books.Core, err error) {
	response, err = uc.bookData.SelectData()

	return response, err
}

func (uc *bookUseCase) InsertData(newBook books.Core) (response books.Core, err error) {
	if newBook.Title == "" || newBook.Author == "" || newBook.Publisher == "" {
		return books.Core{}, errors.New("all input data must be filled")
	}

	response, err = uc.bookData.CreateData(newBook)

	return response, err
}

func (uc *bookUseCase) DetailDataBook(id int) (response books.Core, err error) {
	response, err = uc.bookData.FirstData(id)

	return response, err
}

func (uc *bookUseCase) DeleteDataBook(id int) (response int, err error) {
	response, err = uc.bookData.DestroyData(id)

	return response, err
}

func (uc *bookUseCase) UpdateDataBook(editBook books.Core, id int) (response books.Core, err error) {
	response, err = uc.bookData.EditData(editBook, id)

	return response, err
}
