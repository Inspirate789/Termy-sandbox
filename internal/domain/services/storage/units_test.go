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

type UnitsServiceSuite struct {
	suite.Suite
}

func (s *UnitsServiceSuite) TestGetUnits(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		units, err := service.GetUnits("layer")
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(units, interfaces.OutputUnitsDTO{
			Units: []map[string]interfaces.OutputUnitDTO{
				{
					"ru": interfaces.OutputUnitDTO{ModelID: 0, RegDate: time.Now().Format(time.DateTime), Text: "ru1", PropertiesID: []int{0, 2}, ContextsID: []int{0, 1}},
					"en": interfaces.OutputUnitDTO{ModelID: 1, RegDate: time.Now().Format(time.DateTime), Text: "en1", PropertiesID: []int{0, 2}, ContextsID: []int{0, 2}},
				},
				{
					"ru": interfaces.OutputUnitDTO{ModelID: 0, RegDate: time.Now().Format(time.DateTime), Text: "ru2", PropertiesID: []int{1, 3}, ContextsID: []int{0, 1}},
					"en": interfaces.OutputUnitDTO{ModelID: 3, RegDate: time.Now().Format(time.DateTime), Text: "en2", PropertiesID: []int{1, 4}, ContextsID: []int{0, 2}},
				},
			},
			Contexts: []interfaces.ContextDTO{
				{ID: 0, Text: "0"},
				{ID: 1, Text: "1"},
				{ID: 2, Text: "2"},
			},
		})
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		units, err := service.GetUnits("layer")
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
		sCtx.Assert().Empty(units)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsServiceSuite) TestGetUnitsByModels(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		units, err := service.GetUnitsByModels("layer", interfaces.ModelsIdDTO{Models: []int{0, 1, 2, 3, 4}})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(units, interfaces.OutputUnitsDTO{
			Units: []map[string]interfaces.OutputUnitDTO{
				{
					"ru": interfaces.OutputUnitDTO{ModelID: 0, RegDate: time.Now().Format(time.DateTime), Text: "ru1", PropertiesID: []int{0, 2}, ContextsID: []int{0, 1}},
					"en": interfaces.OutputUnitDTO{ModelID: 1, RegDate: time.Now().Format(time.DateTime), Text: "en1", PropertiesID: []int{0, 2}, ContextsID: []int{0, 2}},
				},
				{
					"ru": interfaces.OutputUnitDTO{ModelID: 2, RegDate: time.Now().Format(time.DateTime), Text: "ru2", PropertiesID: []int{1, 3}, ContextsID: []int{0, 1}},
					"en": interfaces.OutputUnitDTO{ModelID: 3, RegDate: time.Now().Format(time.DateTime), Text: "en2", PropertiesID: []int{1, 4}, ContextsID: []int{0, 2}},
				},
			},
			Contexts: []interfaces.ContextDTO{
				{ID: 0, Text: "0"},
				{ID: 1, Text: "1"},
				{ID: 2, Text: "2"},
			},
		})
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		units, err := service.GetUnitsByModels("layer", interfaces.ModelsIdDTO{Models: []int{0, 1, 2, 3, 4}})
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
		sCtx.Assert().Empty(units)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsServiceSuite) TestGetUnitsByProperties(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		units, err := service.GetUnitsByProperties("layer", interfaces.PropertiesIdDTO{Properties: []int{0, 1, 2, 3, 4}})
		// assert
		sCtx.Assert().NoError(err)
		sCtx.Assert().Equal(units, interfaces.OutputUnitsDTO{
			Units: []map[string]interfaces.OutputUnitDTO{
				{
					"ru": interfaces.OutputUnitDTO{ModelID: 0, RegDate: time.Now().Format(time.DateTime), Text: "ru1", PropertiesID: []int{0, 1}, ContextsID: []int{0, 1}},
					"en": interfaces.OutputUnitDTO{ModelID: 1, RegDate: time.Now().Format(time.DateTime), Text: "en1", PropertiesID: []int{0, 1}, ContextsID: []int{0, 2}},
				},
				{
					"ru": interfaces.OutputUnitDTO{ModelID: 0, RegDate: time.Now().Format(time.DateTime), Text: "ru2", PropertiesID: []int{2, 3}, ContextsID: []int{0, 1}},
					"en": interfaces.OutputUnitDTO{ModelID: 3, RegDate: time.Now().Format(time.DateTime), Text: "en2", PropertiesID: []int{0, 3}, ContextsID: []int{0, 2}},
				},
			},
			Contexts: []interfaces.ContextDTO{
				{ID: 0, Text: "0"},
				{ID: 1, Text: "1"},
				{ID: 2, Text: "2"},
			},
		})
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		units, err := service.GetUnitsByProperties("layer", interfaces.PropertiesIdDTO{Properties: []int{0, 1, 2, 3, 4}})
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
		sCtx.Assert().Empty(units)
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsServiceSuite) TestSaveUnits(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		err := service.SaveUnits("layer", interfaces.SaveUnitsDTO{
			Units: []map[string]interfaces.SaveUnitDTO{
				{
					"ru": interfaces.SaveUnitDTO{ModelID: 0, Text: "ru1", PropertiesID: []int{0, 2}},
					"en": interfaces.SaveUnitDTO{ModelID: 1, Text: "en1", PropertiesID: []int{0, 2}},
				},
				{
					"ru": interfaces.SaveUnitDTO{ModelID: 0, Text: "ru2", PropertiesID: []int{1, 3}},
					"en": interfaces.SaveUnitDTO{ModelID: 3, Text: "en2", PropertiesID: []int{1, 4}},
				},
			},
			Contexts: map[string]string{
				"ru": "абоба",
				"en": "aboba",
			},
		})
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		err := service.SaveUnits("layer", interfaces.SaveUnitsDTO{
			Units: []map[string]interfaces.SaveUnitDTO{
				{
					"ru": interfaces.SaveUnitDTO{ModelID: 0, Text: "ru1", PropertiesID: []int{0, 2}},
					"en": interfaces.SaveUnitDTO{ModelID: 1, Text: "en1", PropertiesID: []int{0, 2}},
				},
				{
					"ru": interfaces.SaveUnitDTO{ModelID: 0, Text: "ru2", PropertiesID: []int{1, 3}},
					"en": interfaces.SaveUnitDTO{ModelID: 3, Text: "en2", PropertiesID: []int{1, 4}},
				},
			},
			Contexts: map[string]string{
				"ru": "абоба",
				"en": "aboba",
			},
		})
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsServiceSuite) TestUpdateUnits(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		newText := "abama"
		// act
		err := service.UpdateUnits("layer", interfaces.UpdateUnitsDTO{Units: []interfaces.UpdateUnitDTO{
			{Lang: "ru", OldText: "aboba", NewText: &newText, PropertiesID: []int{0, 1}},
			{Lang: "en", OldText: "ababa", NewText: &newText, PropertiesID: []int{1, 3}},
			{Lang: "ru", OldText: "abeba", NewText: &newText, PropertiesID: []int{2, 3}},
		}})
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		newText := "abama"
		// act
		err := service.UpdateUnits("layer", interfaces.UpdateUnitsDTO{Units: []interfaces.UpdateUnitDTO{
			{Lang: "ru", OldText: "aboba", NewText: &newText, PropertiesID: []int{0, 1}},
			{Lang: "en", OldText: "ababa", NewText: &newText, PropertiesID: []int{1, 3}},
			{Lang: "ru", OldText: "abeba", NewText: &newText, PropertiesID: []int{2, 3}},
		}})
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
	}, allure.NewParameter("time", time.Now()))
}

func (s *UnitsServiceSuite) TestDeleteUnit(t provider.T) {
	t.Parallel()

	t.WithNewAsyncStep("Positive test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServicePos()
		// act
		err := service.DeleteUnit("layer", interfaces.SearchUnitDTO{
			Lang: "ru",
			Text: "aboba",
		})
		// assert
		sCtx.Assert().NoError(err)
	}, allure.NewParameter("time", time.Now()))

	t.WithNewAsyncStep("Negative test", func(sCtx provider.StepCtx) {
		// arrange
		service := newServiceNeg(errors.New("newMockStorageNeg error"))
		// act
		err := service.DeleteUnit("layer", interfaces.SearchUnitDTO{
			Lang: "ru",
			Text: "aboba",
		})
		// assert
		sCtx.Assert().Error(err)
		sCtx.Assert().False(errors.Is(err, errors.New("newMockStorageNeg error")))
	}, allure.NewParameter("time", time.Now()))
}

func TestUnits(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}
	suite.RunSuite(t, new(UnitsServiceSuite))
}
