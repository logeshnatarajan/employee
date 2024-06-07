package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/logeshnatarajan/employee/internal/model"
	"github.com/logeshnatarajan/employee/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEmployeeService struct {
	mock.Mock
}

// Ensure MockEmployeeService implements the EmployeeServiceInterface
var _ service.EmployeeServiceInterface = (*MockEmployeeService)(nil)

func (m *MockEmployeeService) CreateEmployee(emp model.Employee) (int, error) {
	args := m.Called(emp)
	return emp.ID, args.Error(1)
}

func (m *MockEmployeeService) GetEmployeeByID(id int) (model.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(model.Employee), args.Error(1)
}

func (m *MockEmployeeService) UpdateEmployee(emp model.Employee) error {
	args := m.Called(emp)
	return args.Error(0)
}

func (m *MockEmployeeService) DeleteEmployee(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockEmployeeService) ListEmployees(page, pageSize int) ([]model.Employee, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]model.Employee), args.Error(1)
}

func TestHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	emp := model.Employee{ID: 1, Name: "John Doe", Position: "Developer", Salary: 60000}
	empJSON, _ := json.Marshal(emp)

	t.Run("CreateEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		mockService.On("CreateEmployee", emp).Return(nil)

		router := gin.Default()
		RegisterRoutes(router, mockService)
		r := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(empJSON))

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusCreated, r.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("GetEmployeeByID", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		mockService.On("GetEmployeeByID", 1).Return(emp, nil)

		router := gin.Default()
		RegisterRoutes(router, mockService)
		r := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/employees/1", nil)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusOK, r.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("UpdateEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		emp.Name = "Jone Doe"
		emp.Position = "Developer"
		emp.Salary = 60000
		mockService.On("UpdateEmployee", emp).Return(nil)
		empJ, _ := json.Marshal(emp)

		router := gin.Default()
		RegisterRoutes(router, mockService)
		r := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(empJ))

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusOK, r.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("DeleteEmployee", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		mockService.On("DeleteEmployee", 1).Return(nil)

		router := gin.Default()
		RegisterRoutes(router, mockService)
		r := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/employees/1", nil)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusNoContent, r.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("ListEmployees", func(t *testing.T) {
		mockService := new(MockEmployeeService)
		mockService.On("ListEmployees", 1, 10).Return([]model.Employee{emp}, nil)

		router := gin.Default()
		RegisterRoutes(router, mockService)
		r := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/employees?page=1&pageSize=10", nil)

		router.ServeHTTP(r, req)
		assert.Equal(t, http.StatusOK, r.Code)
		mockService.AssertExpectations(t)
	})
}
