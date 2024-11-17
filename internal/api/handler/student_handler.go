package handler

import (
	"net/http"
	"strconv"
	"student-service/internal/model"
	"student-service/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StudentHandler struct {
	service service.StudentService
	logger  *zap.Logger
}

func NewStudentHandler(service service.StudentService, logger *zap.Logger) *StudentHandler {
	return &StudentHandler{
		service: service,
		logger:  logger,
	}
}

func (h *StudentHandler) Create(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateStudent(c.Request.Context(), &student); err != nil {
		h.logger.Error("Failed to create student", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.Error("Invalid student ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	student, err := h.service.GetStudent(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.Error("Failed to get student", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Update(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateStudent(c.Request.Context(), &student); err != nil {
		h.logger.Error("Failed to update student", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.Error("Invalid student ID", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	if err := h.service.DeleteStudent(c.Request.Context(), uint(id)); err != nil {
		h.logger.Error("Failed to delete student", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *StudentHandler) List(c *gin.Context) {
	students, err := h.service.ListStudents(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to list students", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}
