package storage

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/errors"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
)

func (s *Service) GetLayers() (interfaces.LayersDTO, error) {
	layers, err := s.storage.GetAllLayers()
	if err != nil {
		s.logger.Error(err)
		return interfaces.LayersDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.LayersDTO{Layers: layers}, nil
}

func (s *Service) SaveLayer(layer string) error {
	err := s.storage.SaveLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}

	return nil
}
