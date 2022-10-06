package pgsql

import "QuitSmokingTogether/pkg/models"

func (s *Storage) GetGroups() ([]models.User, error) {
	// TODO
	return nil, nil
}

func (s *Storage) NewGroup(data models.User) (int, error) {
	// TODO
	return 0, nil
}

func (s *Storage) PatchGroup(data models.User) (int, error) {
	// TODO
	return 0, nil
}

func (s *Storage) DeleteGroup(req models.Request) error {
	return nil
}
