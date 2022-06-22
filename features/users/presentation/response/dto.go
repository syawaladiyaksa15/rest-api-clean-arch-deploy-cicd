package response

import (
	"rest-api-clean-arch/features/users"
	"time"
)

type User struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

func FormCore(data users.Core) User {
	return User{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
	}
}

func FromCoreList(data []users.Core) []User {
	result := []User{}

	for k, _ := range data {
		result = append(result, FormCore(data[k]))
	}

	return result
}
