package business

import (
	"errors"
	"rest-api-clean-arch/features/users"
)

type userUseCase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUseCase{
		userData: usrData,
	}
}

func (uc *userUseCase) AuthLogin(dataLogin users.Core) (response users.Core, err error) {
	response, err = uc.userData.LoginAuthData(dataLogin)

	return response, err
}

func (uc *userUseCase) GetAllData(limit, offset int) (response []users.Core, err error) {
	response, err = uc.userData.SelectData()

	return response, err
}

func (uc *userUseCase) InsertData(newUser users.Core) (response users.Core, err error) {
	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		return users.Core{}, errors.New("all input data must be filled")
	}

	response, err = uc.userData.CreateData(newUser)

	return response, err
}

func (uc *userUseCase) DetailDataUser(id int) (response users.Core, err error) {
	response, err = uc.userData.FirstData(id)

	return response, err
}

func (uc *userUseCase) DeleteDataUser(id int) (response int, err error) {
	response, err = uc.userData.DestroyData(id)

	return response, err
}

func (uc *userUseCase) UpdateDataUser(editUser users.Core, id int) (response users.Core, err error) {
	response, err = uc.userData.EditData(editUser, id)

	return response, err
}
