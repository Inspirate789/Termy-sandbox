package postgres_storage_test

import (
	"database/sql"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/migrations"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/test"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type UsersRepositorySuite struct {
	suite.Suite
	dbContainer testcontainers.Container
	db          *sqlx.DB
	migration   *migrations.Migration
}

func (s *UsersRepositorySuite) BeforeAll(t provider.T) {
	t.Epic("Users")
	t.WithNewStep("open DB connection", func(sCtx provider.StepCtx) {
		var err error
		s.dbContainer, s.db, err = test.SetupTestDatabase()
		sCtx.Require().NoError(err)
		s.migration = migrations.NewMigration(s.db)
	})
}

func (s *UsersRepositorySuite) AfterAll(t provider.T) {
	t.WithNewStep("close DB connection", func(sCtx provider.StepCtx) {
		if s.dbContainer != nil {
			err := test.TeardownTestDatabase(s.dbContainer, s.db)
			sCtx.Require().NoError(err)
		} else {
			t.FailNow()
		}
	})
}

func (s *UsersRepositorySuite) BeforeEach(t provider.T) {
	t.WithNewStep("migrate up", func(sCtx provider.StepCtx) {
		err := s.migration.UpTest()
		sCtx.Require().NoError(err)
	})
}

func (s *UsersRepositorySuite) AfterEach(t provider.T) {
	t.WithNewStep("migrate down", func(sCtx provider.StepCtx) {
		err := s.migration.DownTest()
		sCtx.Require().NoError(err)
	})
}

func (s *UsersRepositorySuite) TestAddUser(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		err := store.AddUser(interfaces.UserDTO{
			Name:     "Test User",
			Password: "12345",
			Role:     "student",
		})
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UsersRepositorySuite) TestGetUser(t provider.T) {
	t.WithNewStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		user, err := store.GetUser(entities.AuthRequest{
			Username: "admin1",
			Password: "12345",
		})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal("admin1", user.Name)
		sCtx.Assert().Equal("12345", user.Password)
		sCtx.Assert().Equal("admin", user.Role)
	}, allure.NewParameter("time", time.Now()))

	t.WithNewStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		store := postgres_storage.NewPostgresStorage(s.db)
		// act
		_, err := store.GetUser(entities.AuthRequest{
			Username: "Who",
			Password: "12345",
		})
		// assert
		sCtx.Assert().True(errors.Is(err, sql.ErrNoRows))
	}, allure.NewParameter("time", time.Now()))
}

func TestUsers(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(UsersRepositorySuite))
}
