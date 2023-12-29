package storage

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/errors"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
)

func (s *Service) GetUnits(layer string) (interfaces.OutputUnitsDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, err
	}

	units, err := s.storage.GetAllUnits(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}

	return units, nil
}

func (s *Service) GetUnitsByModels(layer string, modelsDTO interfaces.ModelsIdDTO) (interfaces.OutputUnitsDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, err
	}

	if len(modelsDTO.Models) == 0 {
		return interfaces.OutputUnitsDTO{
			Units:    make([]map[string]interfaces.OutputUnitDTO, 0),
			Contexts: make([]interfaces.ContextDTO, 0),
		}, nil
	}

	units, err := s.storage.GetUnitsByModels(layer, modelsDTO.Models)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}

	return units, nil
}

func (s *Service) GetUnitsByProperties(layer string, propertiesDTO interfaces.PropertiesIdDTO) (interfaces.OutputUnitsDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, err
	}

	if len(propertiesDTO.Properties) == 0 {
		return interfaces.OutputUnitsDTO{
			Units:    make([]map[string]interfaces.OutputUnitDTO, 0),
			Contexts: make([]interfaces.ContextDTO, 0),
		}, nil
	}

	units, err := s.storage.GetUnitsByProperties(layer, propertiesDTO.Properties)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}

	return units, nil
}

func (s *Service) SaveUnits(layer string, unitsDTO interfaces.SaveUnitsDTO) error {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	for i := 0; i < s.repeatOnFailure+1; i++ {
		err = s.storage.SaveUnits(layer, unitsDTO)
		if err == nil || errors.IdentifyStorageError(err) != errors.ErrConnectDatabase {
			break
		}
	}
	if err != nil {
		s.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}

	return nil
}

func (s *Service) UpdateUnits(layer string, unitsDTO interfaces.UpdateUnitsDTO) error {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	for _, unit := range unitsDTO.Units {
		name := unit.OldText

		if unit.NewText != nil {
			err = s.storage.RenameUnit(layer, unit.Lang, unit.OldText, *unit.NewText)
			if err != nil {
				s.logger.Error(err)
				return errors.IdentifyStorageError(err)
			}
			name = *unit.NewText
		}

		if len(unit.PropertiesID) != 0 {
			err = s.storage.SetUnitProperties(layer, unit.Lang, name, unit.PropertiesID)
			if err != nil {
				s.logger.Error(err)
				return errors.IdentifyStorageError(err)
			}
		}
	}

	return nil
}

func (s *Service) DeleteUnit(layer string, unitDTO interfaces.SearchUnitDTO) error {
	err := s.storage.DeleteUnit(layer, unitDTO.Lang, unitDTO.Text)
	if err != nil {
		s.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}
	return nil
}
