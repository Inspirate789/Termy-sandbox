package postgres_storage_test

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/migrations"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/test"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type ModelsRepositorySuite struct {
	suite.Suite
	dbContainer testcontainers.Container
	db          *sqlx.DB
	migration   *migrations.Migration
}

func (s *ModelsRepositorySuite) BeforeAll(t provider.T) {
	t.Epic("Models")
	t.WithNewStep("open DB connection", func(sCtx provider.StepCtx) {
		var err error
		s.dbContainer, s.db, err = test.SetupTestDatabase()
		sCtx.Require().NoError(err)
		s.migration = migrations.NewMigration(s.db)
	})
}

func (s *ModelsRepositorySuite) AfterAll(t provider.T) {
	t.WithNewStep("close DB connection", func(sCtx provider.StepCtx) {
		if s.dbContainer != nil {
			err := test.TeardownTestDatabase(s.dbContainer, s.db)
			sCtx.Require().NoError(err)
		} else {
			t.FailNow()
		}
	})
}

func (s *ModelsRepositorySuite) BeforeEach(t provider.T) {
	t.WithNewStep("migrate up", func(sCtx provider.StepCtx) {
		err := s.migration.UpTest()
		sCtx.Require().NoError(err)
	})
}

func (s *ModelsRepositorySuite) AfterEach(t provider.T) {
	t.WithNewStep("migrate down", func(sCtx provider.StepCtx) {
		err := s.migration.DownTest()
		sCtx.Require().NoError(err)
	})
}

func (s *ModelsRepositorySuite) TestSaveModels(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		id, err := store.SaveModels("test", []string{"a+a+a", "b+b+b", "c+c", "d", "e+e+e+e+e"})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().Len(id, 5)
	}, allure.NewParameter("time", time.Now()))
}

func (s *ModelsRepositorySuite) TestGetAllModels(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		models, err := store.GetAllModels("test")
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEmpty(models)
		t.Log(models)
	}, allure.NewParameter("time", time.Now()))
}

func TestModels(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(ModelsRepositorySuite))
}
