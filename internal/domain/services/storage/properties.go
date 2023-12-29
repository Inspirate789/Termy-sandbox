package storage

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/errors"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
)

func (s *Service) GetProperties() (interfaces.OutputPropertiesDTO, error) {
	properties, err := s.storage.GetAllProperties()
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputPropertiesDTO{}, errors.IdentifyStorageError(err)
	}

	propertiesDTO := make([]interfaces.OutputPropertyDTO, 0, len(properties))
	for i := range properties {
		propertiesDTO = append(propertiesDTO, interfaces.OutputPropertyDTO(properties[i]))
	}

	return interfaces.OutputPropertiesDTO{Properties: propertiesDTO}, nil
}

func (s *Service) GetPropertiesByUnit(layer string, unit interfaces.SearchUnitDTO) (interfaces.OutputPropertiesDTO, error) {
	properties, err := s.storage.GetPropertiesByUnit(layer, unit)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputPropertiesDTO{}, errors.IdentifyStorageError(err)
	}

	propertiesDTO := make([]interfaces.OutputPropertyDTO, 0, len(properties))
	for i := range properties {
		propertiesDTO = append(propertiesDTO, interfaces.OutputPropertyDTO(properties[i]))
	}

	return interfaces.OutputPropertiesDTO{Properties: propertiesDTO}, nil
}

func (s *Service) SaveProperties(propertiesDTO interfaces.PropertyNamesDTO) (interfaces.PropertiesIdDTO, error) {
	propertiesID, err := s.storage.SaveProperties(propertiesDTO.Properties)
	if err != nil {
		s.logger.Error(err)
		return interfaces.PropertiesIdDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.PropertiesIdDTO{Properties: propertiesID}, nil
}
