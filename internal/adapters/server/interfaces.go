package server

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
)

type StorageService interface {
	UnitsService
	ModelsService
	ModelElementsService
	PropertiesService
	LayersService
	UsersService
}

type UnitsService interface {
	GetUnits(layer string) (interfaces.OutputUnitsDTO, error)
	GetUnitsByModels(layer string, modelsDTO interfaces.ModelsIdDTO) (interfaces.OutputUnitsDTO, error)
	GetUnitsByProperties(layer string, propertiesDTO interfaces.PropertiesIdDTO) (interfaces.OutputUnitsDTO, error)
	SaveUnits(layer string, unitsDTO interfaces.SaveUnitsDTO) error
	UpdateUnits(layer string, unitsDTO interfaces.UpdateUnitsDTO) error
	DeleteUnit(layer string, unitDTO interfaces.SearchUnitDTO) error
}

type ModelsService interface {
	GetModels(layer string) (interfaces.OutputModelsDTO, error)
	SaveModels(layer string, modelsDTO interfaces.ModelNamesDTO) (interfaces.ModelsIdDTO, error)
}

type ModelElementsService interface {
	GetModelElements(layer string) (interfaces.OutputModelElementsDTO, error)
	SaveModelElements(layer string, modelElementsDTO interfaces.ModelElementNamesDTO) (interfaces.ModelElementsIdDTO, error)
}

type PropertiesService interface {
	GetProperties() (interfaces.OutputPropertiesDTO, error)
	GetPropertiesByUnit(layer string, unit interfaces.SearchUnitDTO) (interfaces.OutputPropertiesDTO, error)
	SaveProperties(propertiesDTO interfaces.PropertyNamesDTO) (interfaces.PropertiesIdDTO, error)
}

type LayersService interface {
	GetLayers() (interfaces.LayersDTO, error)
	SaveLayer(layer string) error
}

type UsersService interface {
	AddUser(user interfaces.UserDTO) error
	GetUser(entities.AuthRequest) (*entities.User, error)
}
