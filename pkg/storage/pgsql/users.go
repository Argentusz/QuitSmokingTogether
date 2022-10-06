package pgsql

import "QuitSmokingTogether/pkg/models"

func (s *Storage) GetUsers() ([]models.User, error) {
	// TODO
	return nil, nil
}

func (s *Storage) NewUser(data models.User) (int, error) {
	// TODO
	return 0, nil
}

func (s *Storage) PatchUser(data models.User) (int, error) {
	// TODO
	return 0, nil
}

func (s *Storage) DeleteUser(req models.Request) error {
	return nil
}

func (s *Storage) AuthUser(req models.Request) (int, error) {
	return 0, nil
}
