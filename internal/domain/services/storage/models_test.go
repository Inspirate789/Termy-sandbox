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

type ModelsServiceSuite struct {
	suite.Suite
}

func (s *ModelsServiceSuite) TestGetModels(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		models, err := service.GetModels("layer")
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(models.Models, []interfaces.OutputModelDTO{
			{ID: 0, Name: "0"},
			{ID: 1, Name: "1"},
			{ID: 2, Name: "2"},
		})
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		models, err := service.GetModels("layer")
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
		sCtx.Assert().Empty(models)
	}, allure.NewParameter("time", time.Now()))
}

func (s *ModelsServiceSuite) TestSaveModels(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		id, err := service.SaveModels("layer", interfaces.ModelNamesDTO{Models: []string{"a+b", "a+b", "a+b", "a+b", "a+b"}})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(id.Models, []int{0, 1, 2, 3, 4})
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		id, err := service.SaveModels("layer", interfaces.ModelNamesDTO{Models: []string{"0", "1", "2", "3", "4"}})
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
		sCtx.Assert().Empty(id)
	}, allure.NewParameter("time", time.Now()))
}

func TestModels(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(ModelsServiceSuite))
}
