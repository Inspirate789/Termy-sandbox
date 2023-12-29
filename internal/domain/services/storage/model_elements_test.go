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

type ModelElementsServiceSuite struct {
	suite.Suite
}

func (s *ModelElementsServiceSuite) TestGetModelElements(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		modelElements, err := service.GetModelElements("layer")
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(modelElements.Elements, []interfaces.OutputModelElementDTO{
			{ID: 0, Name: "0"},
			{ID: 1, Name: "1"},
			{ID: 2, Name: "2"},
		})
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		modelElements, err := service.GetModelElements("layer")
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
		sCtx.Assert().Empty(modelElements)
	}, allure.NewParameter("time", time.Now()))
}

func (s *ModelElementsServiceSuite) TestSaveModelElements(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		id, err := service.SaveModelElements("layer", interfaces.ModelElementNamesDTO{ModelElements: []string{"a", "b", "c", "d", "e"}})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().ElementsMatch(id.ModelElements, []int{0, 1, 2, 3, 4})
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		id, err := service.SaveModelElements("layer", interfaces.ModelElementNamesDTO{ModelElements: []string{"0", "1", "2", "3", "4"}})
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
		sCtx.Assert().Empty(id)
	}, allure.NewParameter("time", time.Now()))
}

func TestModelElements(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(ModelElementsServiceSuite))
}
