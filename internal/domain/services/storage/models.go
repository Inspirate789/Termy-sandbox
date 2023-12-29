package storage

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/errors"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
)

func (s *Service) GetModels(layer string) (interfaces.OutputModelsDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputModelsDTO{}, err
	}

	models, err := s.storage.GetAllModels(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputModelsDTO{}, errors.IdentifyStorageError(err)
	}

	modelsDTO := make([]interfaces.OutputModelDTO, 0, len(models))
	for i := range models {
		modelsDTO = append(modelsDTO, interfaces.OutputModelDTO(models[i]))
	}

	return interfaces.OutputModelsDTO{Models: modelsDTO}, nil
}

func (s *Service) SaveModels(layer string, modelsDTO interfaces.ModelNamesDTO) (interfaces.ModelsIdDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.ModelsIdDTO{}, err
	}

	for _, modelName := range modelsDTO.Models {
		model := entities.Model{Name: modelName}
		err = model.IsValidName()
		if err != nil {
			s.logger.Error(err)
			return interfaces.ModelsIdDTO{}, err
		}
	}

	modelsID, err := s.storage.SaveModels(layer, modelsDTO.Models)
	if err != nil {
		s.logger.Error(err)
		return interfaces.ModelsIdDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.ModelsIdDTO{Models: modelsID}, nil
}
