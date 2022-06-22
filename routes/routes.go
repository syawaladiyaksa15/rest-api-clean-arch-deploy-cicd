package routes

import (
	"rest-api-clean-arch/factory"
	"rest-api-clean-arch/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	// presenter := factory.InitFactory()
	e := echo.New()
	//test
	// login
	e.POST("/login", presenter.UserPresenter.AuthLogin)

	e.GET("/users", presenter.UserPresenter.GetAll, middlewares.JWTMiddleware())
	e.POST("/users", presenter.UserPresenter.InsertUser)
	e.GET("/users/:id", presenter.UserPresenter.DetailUser, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", presenter.UserPresenter.DestroyUser, middlewares.JWTMiddleware())
	e.PUT("/users/:id", presenter.UserPresenter.UpdateUser, middlewares.JWTMiddleware())

	e.GET("/books", presenter.BookPresenter.GetAll)
	e.POST("/books", presenter.BookPresenter.InsertBook, middlewares.JWTMiddleware())
	e.GET("/books/:id", presenter.BookPresenter.DetailBook, middlewares.JWTMiddleware())
	e.DELETE("/books/:id", presenter.BookPresenter.DestroyBook, middlewares.JWTMiddleware())
	e.PUT("/books/:id", presenter.BookPresenter.UpdateBook, middlewares.JWTMiddleware())

	return e
}
