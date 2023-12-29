package storage_integration_test

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/migrations"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/test"
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

type LayersServiceSuite struct {
	suite.Suite
	dbContainer testcontainers.Container
	db          *sqlx.DB
	migration   *migrations.Migration
}

func (s *LayersServiceSuite) BeforeAll(t provider.T) {
	t.Epic("Layers")
	t.WithNewStep("open DB connection", func(sCtx provider.StepCtx) {
		var err error
		s.dbContainer, s.db, err = test.SetupTestDatabase()
		sCtx.Require().NoError(err)
		s.migration = migrations.NewMigration(s.db)
	})
}

func (s *LayersServiceSuite) AfterAll(t provider.T) {
	t.WithNewStep("close DB connection", func(sCtx provider.StepCtx) {
		if s.dbContainer != nil {
			err := test.TeardownTestDatabase(s.dbContainer, s.db)
			sCtx.Require().NoError(err)
		} else {
			t.FailNow()
		}
	})
}

func (s *LayersServiceSuite) BeforeEach(t provider.T) {
	t.WithNewStep("migrate up", func(sCtx provider.StepCtx) {
		err := s.migration.UpTest()
		sCtx.Require().NoError(err)
	})
}

func (s *LayersServiceSuite) AfterEach(t provider.T) {
	t.WithNewStep("migrate down", func(sCtx provider.StepCtx) {
		err := s.migration.DownTest()
		sCtx.Require().NoError(err)
	})
}

func (s *LayersServiceSuite) TestSaveLayer(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		err := service.SaveLayer("new")
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))
}

func (s *LayersServiceSuite) TestGetLayers(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		layers, err := service.GetLayers()
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(layers.Layers, []string{"test"})
	}, allure.NewParameter("time", time.Now()))
}

func TestLayers(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(LayersServiceSuite))
}
