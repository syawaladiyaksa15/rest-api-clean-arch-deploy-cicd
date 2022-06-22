package data

import (
	"rest-api-clean-arch/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" form:"name" gorm:"not null; type:varchar(100)"`
	Email     string `json:"email" form:"email" gorm:"not null; type:varchar(100); unique"`
	Password  string `json:"password" form:"password" gorm:"not null; type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func toCoreList(data []User) []users.Core {
	result := []users.Core{}

	for key := range data {
		result = append(result, data[key].toCore())
	}

	return result
}

// func toCoreNoSlice(data User) users.Core {
// 	result := users.Core{
// 		ID:        int(data.ID),
// 		Name:      data.Name,
// 		Email:     data.Email,
// 		Password:  data.Password,
// 		CreatedAt: data.CreatedAt,
// 		UpdatedAt: data.UpdatedAt,
// 	}

// 	return result

// }

func formCore(core users.Core) User {
	return User{
		Name:     core.Name,
		Email:    core.Email,
		Password: core.Password,
	}
}
