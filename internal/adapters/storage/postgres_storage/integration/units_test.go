package storage_integration_test

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/migrations"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/test"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/services/storage"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/sirupsen/logrus"
	"github.com/testcontainers/testcontainers-go"
	"io"
	"testing"
	"time"
)

type UnitsServiceSuite struct {
	suite.Suite
	dbContainer testcontainers.Container
	db          *sqlx.DB
	migration   *migrations.Migration
}

func (s *UnitsServiceSuite) BeforeAll(t provider.T) {
	t.Epic("Units")
	t.WithNewStep("open DB connection", func(sCtx provider.StepCtx) {
		var err error
		s.dbContainer, s.db, err = test.SetupTestDatabase()
		sCtx.Require().NoError(err)
		s.migration = migrations.NewMigration(s.db)
	})
}

func (s *UnitsServiceSuite) AfterAll(t provider.T) {
	t.WithNewStep("close DB connection", func(sCtx provider.StepCtx) {
		if s.dbContainer != nil {
			err := test.TeardownTestDatabase(s.dbContainer, s.db)
			sCtx.Require().NoError(err)
		} else {
			t.FailNow()
		}
	})
}

func (s *UnitsServiceSuite) BeforeEach(t provider.T) {
	t.WithNewStep("migrate up", func(sCtx provider.StepCtx) {
		err := s.migration.UpTest()
		sCtx.Require().NoError(err)
	})
}

func (s *UnitsServiceSuite) AfterEach(t provider.T) {
	t.WithNewStep("migrate down", func(sCtx provider.StepCtx) {
		err := s.migration.DownTest()
		sCtx.Require().NoError(err)
	})
}

func (s *UnitsServiceSuite) TestGetUnits(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		units, err := service.GetUnits("test")
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(25, len(units.Units))
		sCtx.Assert().Equal(5, len(units.Contexts))
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsServiceSuite) TestGetUnitsByModels(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		units, err := service.GetUnitsByModels("test", interfaces.ModelsIdDTO{Models: []int{3}})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(6, len(units.Units))
		sCtx.Assert().Equal(1, len(units.Contexts))
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsServiceSuite) TestGetUnitsByProperties(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		units, err := service.GetUnitsByProperties("test", interfaces.PropertiesIdDTO{Properties: []int{4}})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(3, len(units.Units))
		sCtx.Assert().Equal(3, len(units.Contexts))
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsServiceSuite) TestSaveUnits(t provider.T) {
	t.XSkip()

	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		err := service.SaveUnits("test", interfaces.SaveUnitsDTO{
			Units: []map[string]interfaces.SaveUnitDTO{
				{
					"ru": interfaces.SaveUnitDTO{ModelID: 2, Text: "ru1", PropertiesID: []int{1, 2}},
					"en": interfaces.SaveUnitDTO{ModelID: 1, Text: "en1", PropertiesID: []int{1, 2}},
				},
				{
					"ru": interfaces.SaveUnitDTO{ModelID: 2, Text: "ru2", PropertiesID: []int{1, 3}},
					"en": interfaces.SaveUnitDTO{ModelID: 3, Text: "en2", PropertiesID: []int{1, 4}},
				},
			},
			Contexts: map[string]string{
				"ru": "абоба1",
				"en": "aboba2",
			},
		})
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsServiceSuite) TestUpdateUnits(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 5)
		newText1 := "abama1"
		newText2 := "abama2"
		newText3 := "abama3"
		// act
		err := service.UpdateUnits("test", interfaces.UpdateUnitsDTO{Units: []interfaces.UpdateUnitDTO{
			{Lang: "ru", OldText: "абоба1", NewText: &newText1, PropertiesID: []int{5, 1}},
			{Lang: "en", OldText: "aboba1", NewText: &newText2, PropertiesID: []int{1, 3}},
			{Lang: "ru", OldText: "абоба2", NewText: &newText3, PropertiesID: []int{2, 3}},
		}})
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsServiceSuite) TestDeleteUnit(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		err := service.DeleteUnit("test", interfaces.SearchUnitDTO{
			Lang: "ru",
			Text: "абоба1",
		})
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))
}

func TestUnits(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(UnitsServiceSuite))
}
