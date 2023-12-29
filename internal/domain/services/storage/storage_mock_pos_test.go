package storage_test

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/services/storage"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

type storageMockPos struct{}

func newMockStoragePos() *storageMockPos {
	return &storageMockPos{}
}

func newServicePos() *storage.Service {
	storageMock := newMockStoragePos()
	logger := logrus.StandardLogger()
	logger.SetOutput(io.Discard)
	return storage.NewStorageService(storageMock, logger, 0)
}

func (m *storageMockPos) AddUser(_ interfaces.UserDTO) error {
	return nil
}

func (m *storageMockPos) GetUser(request entities.AuthRequest) (*entities.User, error) {
	return &entities.User{
		ID:               0,
		Name:             request.Username,
		Password:         request.Password,
		Role:             entities.AdminRole,
		RegistrationDate: time.Now().Format(time.DateTime),
	}, nil
}

func (m *storageMockPos) GetAllModels(_ string) ([]entities.Model, error) {
	return []entities.Model{
		{ID: 0, Name: "0"},
		{ID: 1, Name: "1"},
		{ID: 2, Name: "2"},
	}, nil
}

func (m *storageMockPos) SaveModels(_ string, models []string) ([]int, error) {
	res := make([]int, 0, len(models))
	for i := range models {
		res = append(res, i)
	}
	return res, nil
}

func (m *storageMockPos) GetAllModelElements(_ string) ([]entities.ModelElement, error) {
	return []entities.ModelElement{
		{ID: 0, Name: "0"},
		{ID: 1, Name: "1"},
		{ID: 2, Name: "2"},
	}, nil
}

func (m *storageMockPos) SaveModelElements(_ string, modelElements []string) ([]int, error) {
	res := make([]int, 0, len(modelElements))
	for i := range modelElements {
		res = append(res, i)
	}
	return res, nil
}

func (m *storageMockPos) GetAllProperties() ([]entities.Property, error) {
	return []entities.Property{
		{ID: 0, Name: "0"},
		{ID: 1, Name: "1"},
		{ID: 2, Name: "2"},
	}, nil
}

func (m *storageMockPos) GetPropertiesByUnit(_ string, _ interfaces.SearchUnitDTO) ([]entities.Property, error) {
	return []entities.Property{
		{ID: 0, Name: "0"},
		{ID: 1, Name: "1"},
		{ID: 2, Name: "2"},
	}, nil
}

func (m *storageMockPos) SaveProperties(properties []string) ([]int, error) {
	res := make([]int, 0, len(properties))
	for i := range properties {
		res = append(res, i)
	}
	return res, nil
}

func (m *storageMockPos) GetAllUnits(_ string) (interfaces.OutputUnitsDTO, error) {
	return interfaces.OutputUnitsDTO{
		Units: []map[string]interfaces.OutputUnitDTO{
			{
				"ru": interfaces.OutputUnitDTO{ModelID: 0, RegDate: time.Now().Format(time.DateTime), Text: "ru1", PropertiesID: []int{0, 2}, ContextsID: []int{0, 1}},
				"en": interfaces.OutputUnitDTO{ModelID: 1, RegDate: time.Now().Format(time.DateTime), Text: "en1", PropertiesID: []int{0, 2}, ContextsID: []int{0, 2}},
			},
			{
				"ru": interfaces.OutputUnitDTO{ModelID: 0, RegDate: time.Now().Format(time.DateTime), Text: "ru2", PropertiesID: []int{1, 3}, ContextsID: []int{0, 1}},
				"en": interfaces.OutputUnitDTO{ModelID: 3, RegDate: time.Now().Format(time.DateTime), Text: "en2", PropertiesID: []int{1, 4}, ContextsID: []int{0, 2}},
			},
		},
		Contexts: []interfaces.ContextDTO{
			{ID: 0, Text: "0"},
			{ID: 1, Text: "1"},
			{ID: 2, Text: "2"},
		},
	}, nil
}

func (m *storageMockPos) GetUnitsByModels(_ string, modelsID []int) (interfaces.OutputUnitsDTO, error) {
	return interfaces.OutputUnitsDTO{
		Units: []map[string]interfaces.OutputUnitDTO{
			{
				"ru": interfaces.OutputUnitDTO{ModelID: modelsID[0], RegDate: time.Now().Format(time.DateTime), Text: "ru1", PropertiesID: []int{0, 2}, ContextsID: []int{0, 1}},
				"en": interfaces.OutputUnitDTO{ModelID: modelsID[1], RegDate: time.Now().Format(time.DateTime), Text: "en1", PropertiesID: []int{0, 2}, ContextsID: []int{0, 2}},
			},
			{
				"ru": interfaces.OutputUnitDTO{ModelID: modelsID[2], RegDate: time.Now().Format(time.DateTime), Text: "ru2", PropertiesID: []int{1, 3}, ContextsID: []int{0, 1}},
				"en": interfaces.OutputUnitDTO{ModelID: modelsID[3], RegDate: time.Now().Format(time.DateTime), Text: "en2", PropertiesID: []int{1, 4}, ContextsID: []int{0, 2}},
			},
		},
		Contexts: []interfaces.ContextDTO{
			{ID: 0, Text: "0"},
			{ID: 1, Text: "1"},
			{ID: 2, Text: "2"},
		},
	}, nil
}

func (m *storageMockPos) GetUnitsByProperties(_ string, propertiesID []int) (interfaces.OutputUnitsDTO, error) {
	return interfaces.OutputUnitsDTO{
		Units: []map[string]interfaces.OutputUnitDTO{
			{
				"ru": interfaces.OutputUnitDTO{ModelID: 0, RegDate: time.Now().Format(time.DateTime), Text: "ru1", PropertiesID: []int{propertiesID[0], propertiesID[1]}, ContextsID: []int{0, 1}},
				"en": interfaces.OutputUnitDTO{ModelID: 1, RegDate: time.Now().Format(time.DateTime), Text: "en1", PropertiesID: []int{propertiesID[0], propertiesID[1]}, ContextsID: []int{0, 2}},
			},
			{
				"ru": interfaces.OutputUnitDTO{ModelID: 0, RegDate: time.Now().Format(time.DateTime), Text: "ru2", PropertiesID: []int{propertiesID[2], propertiesID[3]}, ContextsID: []int{0, 1}},
				"en": interfaces.OutputUnitDTO{ModelID: 3, RegDate: time.Now().Format(time.DateTime), Text: "en2", PropertiesID: []int{propertiesID[0], propertiesID[3]}, ContextsID: []int{0, 2}},
			},
		},
		Contexts: []interfaces.ContextDTO{
			{ID: 0, Text: "0"},
			{ID: 1, Text: "1"},
			{ID: 2, Text: "2"},
		},
	}, nil
}

func (m *storageMockPos) SaveUnits(_ string, _ interfaces.SaveUnitsDTO) error {
	return nil
}

func (m *storageMockPos) RenameUnit(_, _, _, _ string) error {
	return nil
}

func (m *storageMockPos) SetUnitProperties(_, _, _ string, _ []int) error {
	return nil
}

func (m *storageMockPos) DeleteUnit(_, _, _ string) error {
	return nil
}

func (m *storageMockPos) LayerExist(_ string) (bool, error) {
	return true, nil
}

func (m *storageMockPos) GetAllLayers() ([]string, error) {
	return []string{"1", "2", "3"}, nil
}

func (m *storageMockPos) SaveLayer(_ string) error {
	return nil
}
