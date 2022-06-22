package presentation

import (
	"net/http"
	"rest-api-clean-arch/features/books"
	_requestBook "rest-api-clean-arch/features/books/presentation/request"
	_responseBook "rest-api-clean-arch/features/books/presentation/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	BookBusiness books.Business
}

func NewBookHandler(business books.Business) *BookHandler {
	return &BookHandler{
		BookBusiness: business,
	}
}

func (h *BookHandler) GetAll(c echo.Context) error {

	result, err := h.BookBusiness.GetAllData()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseBook.FromCoreList(result),
	})

}

func (h *BookHandler) InsertBook(c echo.Context) error {
	var newBook _requestBook.Book

	errBind := c.Bind(&newBook)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}

	dtBook := _requestBook.ToCore(newBook)
	result, err := h.BookBusiness.InsertData(dtBook)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseBook.FormCore(result),
	})

}

func (h *BookHandler) DetailBook(c echo.Context) error {

	id := c.Param("id")

	idBook, _ := strconv.Atoi(id)

	result, err := h.BookBusiness.DetailDataBook(idBook)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to show detail data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseBook.FormCore(result),
	})

}

func (h *BookHandler) DestroyBook(c echo.Context) error {

	id := c.Param("id")

	idBook, _ := strconv.Atoi(id)

	_, err := h.BookBusiness.DeleteDataBook(idBook)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete detail data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})

}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	var editBook _requestBook.Book

	id := c.Param("id")

	idBook, _ := strconv.Atoi(id)

	errBind := c.Bind(&editBook)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}

	dtBook := _requestBook.ToCore(editBook)

	result, err := h.BookBusiness.UpdateDataBook(dtBook, idBook)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseBook.FormCore(result),
	})

}
