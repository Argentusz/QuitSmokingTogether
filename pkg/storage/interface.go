package storage

import "QuitSmokingTogether/pkg/models"

type DB interface {
	GetUsers() ([]models.User, error)
	NewUser(data models.User) (int, error)
	PatchUser(data models.User) (int, error)
	DeleteUser(req models.Request) error
	AuthUser(req models.Request) (int, error)

	GetGroups() ([]models.User, error)
	NewGroup(data models.User) (int, error)
	PatchGroup(data models.User) (int, error)
	DeleteGroup(req models.Request) error
}
