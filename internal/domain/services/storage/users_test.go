package storage_test

import (
	"errors"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type UsersServiceSuite struct {
	suite.Suite
}

func (s *UsersServiceSuite) TestAddUser(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		err := service.AddUser(interfaces.UserDTO{
			Name:     "test",
			Password: "1234567",
			Role:     "educator",
		})
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		err := service.AddUser(interfaces.UserDTO{})
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
	}, allure.NewParameter("time", time.Now()))
}

func (s *UsersServiceSuite) TestGetUser(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		authRequest := entities.AuthRequest{Username: "Username", Password: "Password"}
		// act
		user, err := service.GetUser(authRequest)
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().EqualValues(user, &entities.User{
			ID:               0,
			Name:             authRequest.Username,
			Password:         authRequest.Password,
			Role:             entities.AdminRole,
			RegistrationDate: time.Now().Format(time.DateTime),
		})
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		authRequest := entities.AuthRequest{Username: "Username", Password: "Password"}
		// act
		user, err := service.GetUser(authRequest)
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
		sCtx.Assert().Nil(user)
	}, allure.NewParameter("time", time.Now()))
}

func TestUsers(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(UsersServiceSuite))
}
