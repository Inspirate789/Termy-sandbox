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

type ModelsServiceSuite struct {
	suite.Suite
	dbContainer testcontainers.Container
	db          *sqlx.DB
	migration   *migrations.Migration
}

func (s *ModelsServiceSuite) BeforeAll(t provider.T) {
	t.Epic("Models")
	t.WithNewStep("open DB connection", func(sCtx provider.StepCtx) {
		var err error
		s.dbContainer, s.db, err = test.SetupTestDatabase()
		sCtx.Require().NoError(err)
		s.migration = migrations.NewMigration(s.db)
	})
}

func (s *ModelsServiceSuite) AfterAll(t provider.T) {
	t.WithNewStep("close DB connection", func(sCtx provider.StepCtx) {
		if s.dbContainer != nil {
			err := test.TeardownTestDatabase(s.dbContainer, s.db)
			sCtx.Require().NoError(err)
		} else {
			t.FailNow()
		}
	})
}

func (s *ModelsServiceSuite) BeforeEach(t provider.T) {
	t.WithNewStep("migrate up", func(sCtx provider.StepCtx) {
		err := s.migration.UpTest()
		sCtx.Require().NoError(err)
	})
}

func (s *ModelsServiceSuite) AfterEach(t provider.T) {
	t.WithNewStep("migrate down", func(sCtx provider.StepCtx) {
		err := s.migration.DownTest()
		sCtx.Require().NoError(err)
	})
}

func (s *ModelsServiceSuite) TestGetModels(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		models, err := service.GetModels("test")
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(models.Models, []interfaces.OutputModelDTO{
			{ID: 1, Name: "a+b"},
			{ID: 2, Name: "a+d+c"},
			{ID: 3, Name: "c+e"},
			{ID: 4, Name: "b+c+a"},
			{ID: 5, Name: "a"},
		})
	}, allure.NewParameter("time", time.Now()))
}

func (s *ModelsServiceSuite) TestSaveModels(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		testStorage := postgres_storage.NewPostgresStorage(s.db)
		logger := logrus.StandardLogger()
		logger.SetOutput(io.Discard)
		service := storage.NewStorageService(testStorage, logger, 0)
		// act
		id, err := service.SaveModels("test", interfaces.ModelNamesDTO{Models: []string{"a+a", "a+a+a", "a+a+a+a", "a+a+a+a+a", "a+a+a+a+a+a"}})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(id.Models, []int{6, 7, 8, 9, 10})
	}, allure.NewParameter("time", time.Now()))
}

func TestModels(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(ModelsServiceSuite))
}
