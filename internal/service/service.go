package service

import (
	"github.com/logeshnatarajan/employee/internal/model"
	"github.com/logeshnatarajan/employee/internal/repository"
)

type EmployeeServiceInterface interface {
	CreateEmployee(emp model.Employee) error
	GetEmployeeByID(id int) (model.Employee, error)
	UpdateEmployee(emp model.Employee) error
	DeleteEmployee(id int) error
	ListEmployees(page, pageSize int) ([]model.Employee, error)
}

type EmployeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(emp model.Employee) error {
	return s.repo.CreateEmployee(emp)
}

func (s *EmployeeService) GetEmployeeByID(id int) (model.Employee, error) {
	return s.repo.GetEmployeeByID(id)
}

func (s *EmployeeService) UpdateEmployee(emp model.Employee) error {
	return s.repo.UpdateEmployee(emp)
}

func (s *EmployeeService) DeleteEmployee(id int) error {
	return s.repo.DeleteEmployee(id)
}

func (s *EmployeeService) ListEmployees(page, pageSize int) ([]model.Employee, error) {
	return s.repo.ListEmployees(page, pageSize)
}
