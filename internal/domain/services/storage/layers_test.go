package storage_test

import (
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/interfaces"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/pkg/errors"
	"testing"
	"time"
)

type LayersServiceSuite struct {
	suite.Suite
}

func (s *LayersServiceSuite) TestSaveLayer(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		err := service.SaveLayer("new")
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		err := service.SaveLayer("new")
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
	}, allure.NewParameter("time", time.Now()))
}

func (s *LayersServiceSuite) TestGetLayers(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		layers, err := service.GetLayers()
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(layers.Layers, []string{"1", "2", "3"})
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

func TestLayers(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(LayersServiceSuite))
}
