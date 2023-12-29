package storage

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/errors"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	storage         Storage
	logger          *log.Logger
	repeatOnFailure int
}

func NewStorageService(storage Storage, logger *log.Logger, repeatOnFailure int) *Service {
	return &Service{
		storage:         storage,
		logger:          logger,
		repeatOnFailure: repeatOnFailure,
	}
}

func (s *Service) checkLayer(layer string) error {
	exist, err := s.storage.LayerExist(layer)
	if err != nil {
		return errors.IdentifyStorageError(err)
	}
	if !exist {
		return errors.ErrUnknownLayerWrap(layer)
	}
	return nil
}
