package data

import (
	"fmt"
	"rest-api-clean-arch/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) LoginAuthData(dataLogin users.Core) (response users.Core, err error) {
	dtLogin := formCore(dataLogin)

	result := repo.db.Where("email = ? AND password = ?", dtLogin.Email, dtLogin.Password).First(&dtLogin)

	if result.Error != nil {
		return users.Core{}, result.Error
	}

	if result.RowsAffected != 1 {
		return users.Core{}, fmt.Errorf("failed to login user")
	}

	return dtLogin.toCore(), nil
}

func (repo *mysqlUserRepository) SelectData() (response []users.Core, err error) {
	var dataUser []User

	result := repo.db.Find(&dataUser)

	if result.Error != nil {
		return []users.Core{}, result.Error
	}

	return toCoreList(dataUser), nil
}

func (repo *mysqlUserRepository) CreateData(newUser users.Core) (response users.Core, err error) {

	// user := User{
	// 	Name:     newUser.Name,
	// 	Email:    newUser.Email,
	// 	Password: newUser.Password,
	// }

	user := formCore(newUser)

	result := repo.db.Create(&user)

	if result.Error != nil {
		return users.Core{}, result.Error
	}

	if result.RowsAffected != 1 {
		return users.Core{}, fmt.Errorf("failed to insert user")
	}

	return user.toCore(), nil
}

func (repo *mysqlUserRepository) FirstData(id int) (response users.Core, err error) {
	var dataUser User

	result := repo.db.Find(&dataUser, id)

	if result.RowsAffected != 1 {
		return users.Core{}, fmt.Errorf("user not found")
	}

	if result.Error != nil {
		return users.Core{}, result.Error
	}

	return dataUser.toCore(), nil
}

func (repo *mysqlUserRepository) DestroyData(id int) (response int, err error) {
	var dataUser User

	result := repo.db.Delete(&dataUser, id)

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("user not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlUserRepository) EditData(editUser users.Core, id int) (response users.Core, err error) {
	user := formCore(editUser)

	// user := User{
	// 	Name:     editUser.Name,
	// 	Email:    editUser.Email,
	// 	Password: editUser.Password,
	// }

	result := repo.db.Model(User{}).Where("id = ?", id).Updates(&user)

	if result.RowsAffected != 1 {
		return users.Core{}, fmt.Errorf("user not found")
	}

	if result.Error != nil {
		return users.Core{}, result.Error
	}

	user.ID = uint(id)

	return user.toCore(), nil
}
