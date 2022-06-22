package users

import (
	"time"
)

type Core struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	GetAllData(limit int, offset int) (data []Core, err error)
	InsertData(newUser Core) (data Core, err error)
	DetailDataUser(id int) (data Core, err error)
	DeleteDataUser(id int) (result int, err error)
	UpdateDataUser(editUser Core, id int) (data Core, err error)
	AuthLogin(dataLogin Core) (data Core, err error)
}

type Data interface {
	SelectData() (data []Core, err error)
	CreateData(newUser Core) (data Core, err error)
	FirstData(id int) (data Core, err error)
	DestroyData(id int) (result int, err error)
	EditData(editUser Core, id int) (data Core, err error)
	LoginAuthData(dataLogin Core) (data Core, err error)
}
