package service

import (
	"testing"

	"calc/internal/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCalculationRepository struct {
	mock.Mock
}

func (m *MockCalculationRepository) CreateCalculation(calc entity.Calculation) error {
	args := m.Called(calc)
	return args.Error(0)
}

func (m *MockCalculationRepository) GetAllCalculations() ([]entity.Calculation, error) {
	args := m.Called()
	return args.Get(0).([]entity.Calculation), args.Error(1)
}

func (m *MockCalculationRepository) GetCalculationByID(id string) (entity.Calculation, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Calculation), args.Error(1)
}

func (m *MockCalculationRepository) UpdateCalculation(calc entity.Calculation) error {
	args := m.Called(calc)
	return args.Error(0)
}

func (m *MockCalculationRepository) DeleteCalculation(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestNewCalcService(t *testing.T) {
	mockRepo := new(MockCalculationRepository)
	service := NewCalcService(mockRepo)

	assert.NotNil(t, service)
}

func TestCalcService_CreateCalculation_Success(t *testing.T) {
	mockRepo := new(MockCalculationRepository)
	service := NewCalcService(mockRepo)

	expression := "2+2"
	expectedResult := "4"

	mockRepo.On("CreateCalculation", mock.MatchedBy(func(calc entity.Calculation) bool {
		return calc.Expression == expression && calc.Result == expectedResult
	})).Return(nil)

	calc, err := service.CreateCalculation(expression)

	assert.NoError(t, err)
	assert.Equal(t, expression, calc.Expression)
	assert.Equal(t, expectedResult, calc.Result)
	assert.NotEmpty(t, calc.ID)

	mockRepo.AssertExpectations(t)
}

func TestCalcService_GetAllCalculations_Success(t *testing.T) {
	mockRepo := new(MockCalculationRepository)
	service := NewCalcService(mockRepo)

	expectedCalculations := []entity.Calculation{
		{ID: "1", Expression: "2+2", Result: "4"},
		{ID: "2", Expression: "5*3", Result: "15"},
	}

	mockRepo.On("GetAllCalculations").Return(expectedCalculations, nil)

	calculations, err := service.GetAllCalculations()

	assert.NoError(t, err)
	assert.Len(t, calculations, 2)
	assert.Equal(t, expectedCalculations, calculations)

	mockRepo.AssertExpectations(t)
}

func TestCalcService_DeleteCalculation_Success(t *testing.T) {
	mockRepo := new(MockCalculationRepository)
	service := NewCalcService(mockRepo)

	id := "123"

	mockRepo.On("DeleteCalculation", id).Return(nil)

	err := service.DeleteCalculation(id)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCalcService_CreateCalculation_InvalidExpression(t *testing.T) {
	mockRepo := new(MockCalculationRepository)
	service := NewCalcService(mockRepo)

	expression := "2++2"

	calc, err := service.CreateCalculation(expression)

	assert.Error(t, err)
	assert.Equal(t, entity.Calculation{}, calc)
	assert.Contains(t, err.Error(), "Invalid token")
}
