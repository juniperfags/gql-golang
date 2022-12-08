package output

import "github.com/juniperfags/gql-golang/internal/domain/models"

type UserOutput struct {
}

type IFindUserOutput interface {
	FindUser(id string) *models.User
}
type IFindAllUserOutput interface {
	FindAll() *[]models.User
}
