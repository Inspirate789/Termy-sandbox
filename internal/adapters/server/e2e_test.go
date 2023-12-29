package server

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/migrations"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/test"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/services/storage"
	"github.com/gavv/httpexpect/v2"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/sirupsen/logrus"
	"github.com/testcontainers/testcontainers-go"
	"io"
	"net/http"
	"testing"
	"time"
)

type GinSuite struct {
	suite.Suite
	dbContainer testcontainers.Container
	db          *sqlx.DB
	migration   *migrations.Migration
}

func (s *GinSuite) BeforeAll(t provider.T) {
	t.Epic("Gin web framework")
	t.WithNewStep("open DB connection", func(sCtx provider.StepCtx) {
		var err error
		s.dbContainer, s.db, err = test.SetupTestDatabase()
		sCtx.Require().NoError(err)
		s.migration = migrations.NewMigration(s.db)
	})
}

func (s *GinSuite) AfterAll(t provider.T) {
	t.WithNewStep("close DB connection", func(sCtx provider.StepCtx) {
		if s.dbContainer != nil {
			err := test.TeardownTestDatabase(s.dbContainer, s.db)
			sCtx.Require().NoError(err)
		} else {
			t.FailNow()
		}
	})
}

func (s *GinSuite) BeforeEach(t provider.T) {
	t.WithNewStep("migrate up", func(sCtx provider.StepCtx) {
		err := s.migration.UpTest()
		sCtx.Require().NoError(err)
	})
}

func (s *GinSuite) AfterEach(t provider.T) {
	t.WithNewStep("migrate down", func(sCtx provider.StepCtx) {
		err := s.migration.DownTest()
		sCtx.Require().NoError(err)
	})
}

func (s *GinSuite) setupExpect(sCtx provider.StepCtx) (*httpexpect.Expect, error) {
	testStorage := postgres_storage.NewPostgresStorage(s.db)
	logger := logrus.StandardLogger()
	logger.SetOutput(io.Discard)
	service := storage.NewStorageService(testStorage, logger, 0)
	server, err := NewServer(8080, service, logger, nil)
	if err != nil {
		return nil, err
	}
	ginHandler := server.srv.Handler

	return httpexpect.WithConfig(httpexpect.Config{
		TestName: sCtx.Name(),
		Client: &http.Client{
			Transport: httpexpect.NewBinder(ginHandler),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(sCtx),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(sCtx, true),
		},
	}), nil
}

func (s *GinSuite) TestGinAuth(t provider.T) {
	var authToken string

	t.WithNewStep("Login", func(sCtx provider.StepCtx) {
		// arrange
		expect, err := s.setupExpect(sCtx)
		sCtx.Require().NoError(err)
		// act
		resp := expect.POST("/login").
			WithJSON(entities.AuthRequest{Username: "admin1", Password: "12345"}).
			Expect()
		// assert
		object := resp.Status(http.StatusOK).JSON().Object()
		authToken = object.Value("token").String().Raw()
		expire, err := time.Parse(time.RFC3339, object.Value("expire").String().Raw())
		sCtx.Require().NotEmpty(authToken)
		sCtx.Require().NoError(err)
		sCtx.Require().True(expire.After(time.Now()))
	}, allure.NewParameter("time", time.Now()))

	t.WithNewStep("Logout", func(sCtx provider.StepCtx) {
		// arrange
		expect, err := s.setupExpect(sCtx)
		sCtx.Require().NoError(err)
		// act
		resp := expect.DELETE("/logout").
			WithHeader("Authorization", "Bearer "+authToken).
			Expect()
		// assert
		resp.Status(http.StatusOK).NoContent()
	}, allure.NewParameter("time", time.Now()))
}

func TestGin(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(GinSuite))
}
