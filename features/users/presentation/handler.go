package presentation

import (
	"net/http"
	"rest-api-clean-arch/features/users"
	_requestUser "rest-api-clean-arch/features/users/presentation/request"
	_responseUser "rest-api-clean-arch/features/users/presentation/response"
	_middlewares "rest-api-clean-arch/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		UserBusiness: business,
	}
}

func (h *UserHandler) AuthLogin(c echo.Context) error {
	var loginUser _requestUser.User

	errBind := c.Bind(&loginUser)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}

	dtUser := _requestUser.ToCore(loginUser)
	result, err := h.UserBusiness.AuthLogin(dtUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to login data",
		})
	}

	token, errToken := _middlewares.CreateToken(int(result.ID))

	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to login data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    token,
	})

}

func (h *UserHandler) GetAll(c echo.Context) error {

	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	//
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	//
	result, err := h.UserBusiness.GetAllData(limitint, offsetint)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseUser.FromCoreList(result),
	})

}

func (h *UserHandler) InsertUser(c echo.Context) error {
	var newUser _requestUser.User

	// user := users.Core{
	// 	Name:     c.FormValue("name"),
	// 	Email:    c.FormValue("email"),
	// 	Password: c.FormValue("password"),
	// }

	errBind := c.Bind(&newUser)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}

	dtUser := _requestUser.ToCore(newUser)
	result, err := h.UserBusiness.InsertData(dtUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseUser.FormCore(result),
	})

}

func (h *UserHandler) DetailUser(c echo.Context) error {

	id := c.Param("id")

	idUser, _ := strconv.Atoi(id)

	result, err := h.UserBusiness.DetailDataUser(idUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to show detail data",
		})
	}

	idToken, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}

	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseUser.FormCore(result),
	})

}

func (h *UserHandler) DestroyUser(c echo.Context) error {

	id := c.Param("id")

	idUser, _ := strconv.Atoi(id)

	_, err := h.UserBusiness.DeleteDataUser(idUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete detail data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})

}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var editUser _requestUser.User

	id := c.Param("id")

	idUser, _ := strconv.Atoi(id)

	errBind := c.Bind(&editUser)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}

	dtUser := _requestUser.ToCore(editUser)

	// user := users.Core{
	// 	Name:     c.FormValue("name"),
	// 	Email:    c.FormValue("email"),
	// 	Password: c.FormValue("password"),
	// }

	result, err := h.UserBusiness.UpdateDataUser(dtUser, idUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseUser.FormCore(result),
	})

}
