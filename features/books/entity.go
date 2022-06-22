package books

import (
	"time"
)

type Core struct {
	ID        int
	Title     string
	Author    string
	Publisher string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	GetAllData() (data []Core, err error)
	InsertData(newBook Core) (data Core, err error)
	DetailDataBook(id int) (data Core, err error)
	DeleteDataBook(id int) (result int, err error)
	UpdateDataBook(editBook Core, id int) (data Core, err error)
}

type Data interface {
	SelectData() (data []Core, err error)
	CreateData(newBook Core) (data Core, err error)
	FirstData(id int) (data Core, err error)
	DestroyData(id int) (result int, err error)
	EditData(editBook Core, id int) (data Core, err error)
}
