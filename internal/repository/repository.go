package repository

import (
	"database/sql"
	"fmt"

	"github.com/logeshnatarajan/employee/internal/model"
)

type EmployeeRepository interface {
	CreateEmployee(emp model.Employee) (error, int)
	GetEmployeeByID(id int) (model.Employee, error)
	UpdateEmployee(emp model.Employee) error
	DeleteEmployee(id int) error
	ListEmployees(page, pageSize int) ([]model.Employee, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) CreateEmployee(emp model.Employee) (error, int) {
	query := "INSERT INTO employees (name, position, salary) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRow(query, emp.Name, emp.Position, emp.Salary).Scan(&emp.ID), emp.ID
}

func (r *PostgresRepository) GetEmployeeByID(id int) (model.Employee, error) {
	var emp model.Employee
	query := "SELECT id, name, position, salary FROM employees WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary)
	return emp, err
}

func (r *PostgresRepository) UpdateEmployee(emp model.Employee) error {
	fmt.Println("poduu")
	query := "UPDATE employees SET name = $1, position = $2, salary = $3 WHERE id = $4"
	_, err := r.db.Exec(query, emp.Name, emp.Position, emp.Salary, emp.ID)
	return err
}

func (r *PostgresRepository) DeleteEmployee(id int) error {
	query := "DELETE FROM employees WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) ListEmployees(page, pageSize int) ([]model.Employee, error) {
	offset := (page - 1) * pageSize
	query := "SELECT id, name, position, salary FROM employees LIMIT $1 OFFSET $2"
	rows, err := r.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var emp model.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}
