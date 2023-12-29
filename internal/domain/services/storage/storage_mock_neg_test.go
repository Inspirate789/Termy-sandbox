package storage_test

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/services/storage"
	"github.com/sirupsen/logrus"
	"io"
)

type storageMockNeg struct {
	error error
}

func newMockStorageNeg(err error) *storageMockNeg {
	return &storageMockNeg{
		error: err,
	}
}

func newServiceNeg(err error) *storage.Service {
	storageMock := newMockStorageNeg(err)
	logger := logrus.StandardLogger()
	logger.SetOutput(io.Discard)
	return storage.NewStorageService(storageMock, logger, 0)
}

func (m *storageMockNeg) AddUser(_ interfaces.UserDTO) error {
	return m.error
}

func (m *storageMockNeg) GetUser(request entities.AuthRequest) (*entities.User, error) {
	return nil, m.error
}

func (m *storageMockNeg) GetAllModels(layer string) ([]entities.Model, error) {
	return nil, m.error
}

func (m *storageMockNeg) SaveModels(layer string, models []string) ([]int, error) {
	return nil, m.error
}

func (m *storageMockNeg) GetAllModelElements(layer string) ([]entities.ModelElement, error) {
	return nil, m.error
}

func (m *storageMockNeg) SaveModelElements(layer string, modelElements []string) ([]int, error) {
	return nil, m.error
}

func (m *storageMockNeg) GetAllProperties() ([]entities.Property, error) {
	return nil, m.error
}

func (m *storageMockNeg) GetPropertiesByUnit(layer string, unit interfaces.SearchUnitDTO) ([]entities.Property, error) {
	return nil, m.error
}

func (m *storageMockNeg) SaveProperties(properties []string) ([]int, error) {
	return nil, m.error
}

func (m *storageMockNeg) GetAllUnits(layer string) (interfaces.OutputUnitsDTO, error) {
	return interfaces.OutputUnitsDTO{}, m.error
}

func (m *storageMockNeg) GetUnitsByModels(layer string, modelsID []int) (interfaces.OutputUnitsDTO, error) {
	return interfaces.OutputUnitsDTO{}, m.error
}

func (m *storageMockNeg) GetUnitsByProperties(layer string, propertiesID []int) (interfaces.OutputUnitsDTO, error) {
	return interfaces.OutputUnitsDTO{}, m.error
}

func (m *storageMockNeg) SaveUnits(layer string, data interfaces.SaveUnitsDTO) error {
	return m.error
}

func (m *storageMockNeg) RenameUnit(layer, lang, oldName, newName string) error {
	return m.error
}

func (m *storageMockNeg) SetUnitProperties(layer, lang, unitName string, propertiesID []int) error {
	return m.error
}

func (m *storageMockNeg) DeleteUnit(layer, lang, unitName string) error {
	return m.error
}

func (m *storageMockNeg) LayerExist(layer string) (bool, error) {
	return false, m.error
}

func (m *storageMockNeg) GetAllLayers() ([]string, error) {
	return nil, m.error
}

func (m *storageMockNeg) SaveLayer(name string) error {
	return m.error
}
