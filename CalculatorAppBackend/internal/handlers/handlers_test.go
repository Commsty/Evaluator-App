package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"calc/internal/entity"
	"calc/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCalculationService struct {
	mock.Mock
}

func (m *MockCalculationService) CreateCalculation(expression string) (entity.Calculation, error) {
	args := m.Called(expression)
	return args.Get(0).(entity.Calculation), args.Error(1)
}

func (m *MockCalculationService) GetAllCalculations() ([]entity.Calculation, error) {
	args := m.Called()
	return args.Get(0).([]entity.Calculation), args.Error(1)
}

func (m *MockCalculationService) GetCalculationByID(id string) (entity.Calculation, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Calculation), args.Error(1)
}

func (m *MockCalculationService) UpdateCalculation(id, expression string) (entity.Calculation, error) {
	args := m.Called(id, expression)
	return args.Get(0).(entity.Calculation), args.Error(1)
}

func (m *MockCalculationService) DeleteCalculation(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupTestHandler(service service.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: service}
}

func TestCalculationHandler_GetCalculations_Success(t *testing.T) {
	mockService := new(MockCalculationService)
	handler := setupTestHandler(mockService)

	expectedCalculations := []entity.Calculation{
		{ID: "1", Expression: "2+2", Result: "4"},
		{ID: "2", Expression: "5*3", Result: "15"},
	}

	mockService.On("GetAllCalculations").Return(expectedCalculations, nil)

	req := httptest.NewRequest(http.MethodGet, "/calculations", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := handler.GetCalculations(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response []entity.Calculation
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, expectedCalculations[0].Expression, response[0].Expression)
	assert.Equal(t, expectedCalculations[0].Result, response[0].Result)

	mockService.AssertExpectations(t)
}

func TestCalculationHandler_PostCalculations_Success(t *testing.T) {
	mockService := new(MockCalculationService)
	handler := setupTestHandler(mockService)

	requestBody := entity.CalculationRequest{Expression: "10+5"}
	expectedCalc := entity.Calculation{ID: "123", Expression: "10+5", Result: "15"}

	mockService.On("CreateCalculation", "10+5").Return(expectedCalc, nil)

	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/calculations", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := handler.PostCalculations(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response entity.Calculation
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedCalc.ID, response.ID)
	assert.Equal(t, expectedCalc.Expression, response.Expression)
	assert.Equal(t, expectedCalc.Result, response.Result)

	mockService.AssertExpectations(t)
}

func TestCalculationHandler_DeleteCalculation_Success(t *testing.T) {
	mockService := new(MockCalculationService)
	handler := setupTestHandler(mockService)

	mockService.On("DeleteCalculation", "123").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/calculations/123", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("123")

	err := handler.DeleteCalculation(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.Empty(t, rec.Body.String())

	mockService.AssertExpectations(t)
}
