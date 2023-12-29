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

type UnitsRepositorySuite struct {
	suite.Suite
	dbContainer testcontainers.Container
	db          *sqlx.DB
	migration   *migrations.Migration
}

func (s *UnitsRepositorySuite) BeforeAll(t provider.T) {
	t.Epic("Units")
	t.WithNewStep("open DB connection", func(sCtx provider.StepCtx) {
		var err error
		s.dbContainer, s.db, err = test.SetupTestDatabase()
		sCtx.Require().NoError(err)
		s.migration = migrations.NewMigration(s.db)
	})
}

func (s *UnitsRepositorySuite) AfterAll(t provider.T) {
	t.WithNewStep("close DB connection", func(sCtx provider.StepCtx) {
		if s.dbContainer != nil {
			err := test.TeardownTestDatabase(s.dbContainer, s.db)
			sCtx.Require().NoError(err)
		} else {
			t.FailNow()
		}
	})
}

func (s *UnitsRepositorySuite) BeforeEach(t provider.T) {
	t.WithNewStep("migrate up", func(sCtx provider.StepCtx) {
		err := s.migration.UpTest()
		sCtx.Require().NoError(err)
	})
}

func (s *UnitsRepositorySuite) AfterEach(t provider.T) {
	t.WithNewStep("migrate down", func(sCtx provider.StepCtx) {
		err := s.migration.DownTest()
		sCtx.Require().NoError(err)
	})
}

func (s *UnitsRepositorySuite) TestSaveUnits(t provider.T) {
	t.XSkip()

	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		err := store.SaveUnits("test", interfaces.SaveUnitsDTO{
			Units: []map[string]interfaces.SaveUnitDTO{
				{
					"ru": interfaces.SaveUnitDTO{ModelID: 2, Text: "ru1", PropertiesID: []int{5, 2}},
					"en": interfaces.SaveUnitDTO{ModelID: 1, Text: "en1", PropertiesID: []int{5, 2}},
				},
				{
					"ru": interfaces.SaveUnitDTO{ModelID: 4, Text: "ru2", PropertiesID: []int{1, 3}},
					"en": interfaces.SaveUnitDTO{ModelID: 3, Text: "en2", PropertiesID: []int{1, 4}},
				},
			},
			Contexts: map[string]string{
				"ru": "термин",
				"en": "term",
			},
		})
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsRepositorySuite) TestRenameUnit(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		err := store.RenameUnit("test", "en", "aboba1", "aboba11")
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsRepositorySuite) TestSetUnitProperties(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		err := store.SetUnitProperties("test", "en", "aboba1", []int{1, 2, 3, 4, 5})
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsRepositorySuite) TestDeleteUnit(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		err := store.DeleteUnit("test", "en", "aboba1")
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsRepositorySuite) TestGetAllUnits(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		units, err := store.GetAllUnits("test")
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEmpty(units)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsRepositorySuite) TestGetUnitsByModels(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		units, err := store.GetUnitsByModels("test", []int{2})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEmpty(units)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsRepositorySuite) TestGetUnitsByProperties(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		units, err := store.GetUnitsByProperties("test", []int{1, 2})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotEmpty(units)
	}, allure.NewParameter("time", time.Now()))
}

func TestUnits(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(UnitsRepositorySuite))
}
