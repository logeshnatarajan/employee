# Employee Management System

This is a Go-based Employee Management System with a RESTful API. The application manages an employee database using PostgreSQL.

## Features

- Create, Read, Update, Delete (CRUD) operations for employees
- RESTful API with pagination
- Concurrency-safe implementation
- Unit testing for service and handler layers

## Technologies

- Go
- PostgreSQL
- Gin (HTTP Web Framework)
- Gorm (ORM Library)
- Testify (Testing Framework)

## Prerequisites

- Go 1.18+
- PostgreSQL
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

## Project Structure

employee/
│
├── internal/
│   ├── handler/
│   │   ├── employee_handler.go
│   │   ├── handler_test.go
│   │   ├── router.go
│   ├── model/
│   │   ├── employee.go
│   ├── repository/
│   │   ├── employee_repository.go
│   │   ├── repository.go
│   ├── service/
│   │   ├── service.go
│
├── migrations/
│   ├── 0001_create_employees_table.up.sql
│   ├── 0001_create_employees_table.down.sql
│   ├── 0002_seed_employees_table.up.sql
│   ├── 0002_seed_employees_table.down.sql
│
├── config/
│   ├── config.go
│
├── main.go
├── go.mod
├── go.sum
└── README.md

## Setup
### Step 1: Clone the Repository

```bash
git clone https://github.com/logeshnatarajan/employee.git
cd employee
```
### Step 2: Set Up PostgreSQL
#### 1: Install PostgreSQL if not already installed.
#### 2: Create a new database:
```bash
CREATE DATABASE employee_db;
```
### Step 3: Run Migrations
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
### Step 4: Install Dependencies
```bash
go mod tidy
```
### Step 5: Run the Application
```bash
go run main.go
```
### Step 6: Run Tests
```bash
go test ./internal/handler
```