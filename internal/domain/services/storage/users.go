package storage

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/errors"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
)

func (s *Service) AddUser(user interfaces.UserDTO) error {
	err := s.storage.AddUser(user)
	if err != nil {
		s.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}
	s.logger.Infof("storage user (name: %s, role: %s) inserted to database", user.Name, user.Role)

	return nil
}

func (s *Service) GetUser(request entities.AuthRequest) (*entities.User, error) {
	user, err := s.storage.GetUser(request)
	if err != nil {
		s.logger.Error(err)
		return nil, errors.IdentifyStorageError(err)
	}

	return user, nil
}
