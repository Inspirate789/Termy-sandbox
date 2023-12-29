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

type PropertiesServiceSuite struct {
	suite.Suite
	dbContainer testcontainers.Container
	db          *sqlx.DB
	migration   *migrations.Migration
}

func (s *PropertiesServiceSuite) BeforeAll(t provider.T) {
	t.Epic("Properties")
	t.WithNewStep("open DB connection", func(sCtx provider.StepCtx) {
		var err error
		s.dbContainer, s.db, err = test.SetupTestDatabase()
		sCtx.Require().NoError(err)
		s.migration = migrations.NewMigration(s.db)
	})
}

func (s *PropertiesServiceSuite) AfterAll(t provider.T) {
	t.WithNewStep("close DB connection", func(sCtx provider.StepCtx) {
		if s.dbContainer != nil {
			err := test.TeardownTestDatabase(s.dbContainer, s.db)
			sCtx.Require().NoError(err)
		} else {
			t.FailNow()
		}
	})
}

func (s *PropertiesServiceSuite) BeforeEach(t provider.T) {
	t.WithNewStep("migrate up", func(sCtx provider.StepCtx) {
		err := s.migration.UpTest()
		sCtx.Require().NoError(err)
	})
}

func (s *PropertiesServiceSuite) AfterEach(t provider.T) {
	t.WithNewStep("migrate down", func(sCtx provider.StepCtx) {
		err := s.migration.DownTest()
		sCtx.Require().NoError(err)
	})
}

func (s *PropertiesServiceSuite) TestGetProperties(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		properties, err := service.GetProperties()
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(properties.Properties, []interfaces.OutputPropertyDTO{
			{ID: 1, Name: "property1"},
			{ID: 2, Name: "property2"},
			{ID: 3, Name: "property3"},
			{ID: 4, Name: "property4"},
			{ID: 5, Name: "property5"},
		})
	}, allure.NewParameter("time", time.Now()))
}

func (s *PropertiesServiceSuite) TestGetPropertiesByUnit(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		properties, err := service.GetPropertiesByUnit("test", interfaces.SearchUnitDTO{
			Lang: "ru",
			Text: "абоба1",
		})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(properties.Properties, []interfaces.OutputPropertyDTO{
			{ID: 1, Name: "property1"},
			{ID: 3, Name: "property3"},
		})
	}, allure.NewParameter("time", time.Now()))
}

func (s *PropertiesServiceSuite) TestSaveProperties(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		id, err := service.SaveProperties(interfaces.PropertyNamesDTO{Properties: []string{"a", "b", "c", "d", "e"}})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(id.Properties, []int{6, 7, 8, 9, 10})
	}, allure.NewParameter("time", time.Now()))
}

func TestProperties(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(PropertiesServiceSuite))
}
