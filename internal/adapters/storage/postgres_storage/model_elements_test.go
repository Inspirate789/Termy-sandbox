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

type ModelElementsRepositorySuite struct {
	suite.Suite
	dbContainer testcontainers.Container
	db          *sqlx.DB
	migration   *migrations.Migration
}

func (s *ModelElementsRepositorySuite) BeforeAll(t provider.T) {
	t.Epic("ModelElements")
	t.WithNewStep("open DB connection", func(sCtx provider.StepCtx) {
		var err error
		s.dbContainer, s.db, err = test.SetupTestDatabase()
		sCtx.Require().NoError(err)
		s.migration = migrations.NewMigration(s.db)
	})
}

func (s *ModelElementsRepositorySuite) AfterAll(t provider.T) {
	t.WithNewStep("close DB connection", func(sCtx provider.StepCtx) {
		if s.dbContainer != nil {
			err := test.TeardownTestDatabase(s.dbContainer, s.db)
			sCtx.Require().NoError(err)
		} else {
			t.FailNow()
		}
	})
}

func (s *ModelElementsRepositorySuite) BeforeEach(t provider.T) {
	t.WithNewStep("migrate up", func(sCtx provider.StepCtx) {
		err := s.migration.UpTest()
		sCtx.Require().NoError(err)
	})
}

func (s *ModelElementsRepositorySuite) AfterEach(t provider.T) {
	t.WithNewStep("migrate down", func(sCtx provider.StepCtx) {
		err := s.migration.DownTest()
		sCtx.Require().NoError(err)
	})
}

func (s *ModelElementsRepositorySuite) TestSaveModelElements(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		id, err := store.SaveModelElements("test", []string{"noun", "verb", "adverb", "adjective", "participle"})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().Len(id, 5)
	}, allure.NewParameter("time", time.Now()))
}

func (s *ModelElementsRepositorySuite) TestGetAllModelElements(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		modelElements, err := store.GetAllModelElements("test")
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEmpty(modelElements)
		t.Log(modelElements)
	}, allure.NewParameter("time", time.Now()))
}

func TestModelElements(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(ModelElementsRepositorySuite))
}
