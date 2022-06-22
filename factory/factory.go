package factory

import (
	// user
	_userBusiness "rest-api-clean-arch/features/users/business"
	_userData "rest-api-clean-arch/features/users/data"
	_userPresentation "rest-api-clean-arch/features/users/presentation"

	// book
	_bookBusiness "rest-api-clean-arch/features/books/business"
	_bookData "rest-api-clean-arch/features/books/data"
	_bookPresentation "rest-api-clean-arch/features/books/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter *_userPresentation.UserHandler
	BookPresenter *_bookPresentation.BookHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// user
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	// book
	bookData := _bookData.NewBookRepository(dbConn)
	bookBusiness := _bookBusiness.NewBookBusiness(bookData)
	bookPresentation := _bookPresentation.NewBookHandler(bookBusiness)

	return Presenter{
		UserPresenter: userPresentation,
		BookPresenter: bookPresentation,
	}
}
