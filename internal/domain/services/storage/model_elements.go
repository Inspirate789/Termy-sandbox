package storage

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/errors"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
)

func (s *Service) GetModelElements(layer string) (interfaces.OutputModelElementsDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputModelElementsDTO{}, err
	}

	modelElements, err := s.storage.GetAllModelElements(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputModelElementsDTO{}, errors.IdentifyStorageError(err)
	}

	modelElementsDTO := make([]interfaces.OutputModelElementDTO, 0, len(modelElements))
	for i := range modelElements {
		modelElementsDTO = append(modelElementsDTO, interfaces.OutputModelElementDTO(modelElements[i]))
	}

	return interfaces.OutputModelElementsDTO{Elements: modelElementsDTO}, nil
}

func (s *Service) SaveModelElements(layer string, modelElementsDTO interfaces.ModelElementNamesDTO) (interfaces.ModelElementsIdDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.ModelElementsIdDTO{}, err
	}

	modelElementsID, err := s.storage.SaveModelElements(layer, modelElementsDTO.ModelElements)
	if err != nil {
		s.logger.Error(err)
		return interfaces.ModelElementsIdDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.ModelElementsIdDTO{ModelElements: modelElementsID}, nil
}
