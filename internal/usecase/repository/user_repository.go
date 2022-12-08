package repository

import "github.com/juniperfags/gql-golang/internal/domain/models"

type IUserRepository interface {
	FindUser(id string) *models.User
	FindAll() *[]models.User
	CreateUser() *models.User
}
