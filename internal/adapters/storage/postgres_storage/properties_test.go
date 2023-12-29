package postgres_storage_test

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/migrations"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/test"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type PropertiesRepositorySuite struct {
	suite.Suite
	dbContainer testcontainers.Container
	db          *sqlx.DB
	migration   *migrations.Migration
}

func (s *PropertiesRepositorySuite) BeforeAll(t provider.T) {
	t.Epic("Properties")
	t.WithNewStep("open DB connection", func(sCtx provider.StepCtx) {
		var err error
		s.dbContainer, s.db, err = test.SetupTestDatabase()
		sCtx.Require().NoError(err)
		s.migration = migrations.NewMigration(s.db)
	})
}

func (s *PropertiesRepositorySuite) AfterAll(t provider.T) {
	t.WithNewStep("close DB connection", func(sCtx provider.StepCtx) {
		if s.dbContainer != nil {
			err := test.TeardownTestDatabase(s.dbContainer, s.db)
			sCtx.Require().NoError(err)
		} else {
			t.FailNow()
		}
	})
}

func (s *PropertiesRepositorySuite) BeforeEach(t provider.T) {
	t.WithNewStep("migrate up", func(sCtx provider.StepCtx) {
		err := s.migration.UpTest()
		sCtx.Require().NoError(err)
	})
}

func (s *PropertiesRepositorySuite) AfterEach(t provider.T) {
	t.WithNewStep("migrate down", func(sCtx provider.StepCtx) {
		err := s.migration.DownTest()
		sCtx.Require().NoError(err)
	})
}

func (s *PropertiesRepositorySuite) TestSaveProperties(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		id, err := store.SaveProperties([]string{"good", "bad", "awesome", "well", "brilliant"})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().Len(id, 5)
	}, allure.NewParameter("time", time.Now()))
}

func (s *PropertiesRepositorySuite) TestGetAllProperties(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		properties, err := store.GetAllProperties()
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEmpty(properties)
		t.Log(properties)
	}, allure.NewParameter("time", time.Now()))
}

func (s *PropertiesRepositorySuite) TestGetPropertiesByUnit(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		properties, err := store.GetPropertiesByUnit("test", interfaces.SearchUnitDTO{
			Lang: "en",
			Text: "aboba3",
		})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEmpty(properties)
		t.Log(properties)
	}, allure.NewParameter("time", time.Now()))
}

func TestProperties(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(PropertiesRepositorySuite))
}
