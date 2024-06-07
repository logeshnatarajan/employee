package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/logeshnatarajan/employee/internal/service"
)

func RegisterRoutes(r *gin.Engine, svc service.EmployeeServiceInterface) {
	handler := NewEmployeeHandler(svc)
	r.POST("/employees", handler.CreateEmployee)
	r.GET("/employees/:id", handler.GetEmployeeByID)
	r.PUT("/employees/:id", handler.UpdateEmployee)
	r.DELETE("/employees/:id", handler.DeleteEmployee)
	r.GET("/employees", handler.ListEmployees)
}
