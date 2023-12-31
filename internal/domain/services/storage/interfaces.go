package storage

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
)

type Storage interface {
	UsersRepository
	ModelsRepository
	ModelElementsRepository
	PropertiesRepository
	UnitsRepository
	LayersRepository
}

type UsersRepository interface {
	AddUser(user interfaces.UserDTO) error
	GetUser(entities.AuthRequest) (*entities.User, error)
}

type ModelsRepository interface {
	GetAllModels(layer string) ([]entities.Model, error)
	SaveModels(layer string, models []string) ([]int, error)
}

type ModelElementsRepository interface {
	GetAllModelElements(layer string) ([]entities.ModelElement, error)
	SaveModelElements(layer string, modelElements []string) ([]int, error)
}

type PropertiesRepository interface {
	GetAllProperties() ([]entities.Property, error)
	GetPropertiesByUnit(layer string, unit interfaces.SearchUnitDTO) ([]entities.Property, error)
	SaveProperties(properties []string) ([]int, error)
}

type UnitsRepository interface {
	GetAllUnits(layer string) (interfaces.OutputUnitsDTO, error)
	GetUnitsByModels(layer string, modelsID []int) (interfaces.OutputUnitsDTO, error)
	GetUnitsByProperties(layer string, propertiesID []int) (interfaces.OutputUnitsDTO, error)
	SaveUnits(layer string, data interfaces.SaveUnitsDTO) error
	RenameUnit(layer, lang, oldName, newName string) error
	SetUnitProperties(layer, lang, unitName string, propertiesID []int) error
	DeleteUnit(layer, lang, unitName string) error
}

type LayersRepository interface {
	LayerExist(layer string) (bool, error)
	GetAllLayers() ([]string, error)
	SaveLayer(name string) error
}
