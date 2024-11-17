package api

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"student-service/internal/api/handler"
	"student-service/internal/config"
	"student-service/internal/repository"
	"student-service/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	cfg    *config.Config
	db     *gorm.DB
	logger *zap.Logger
	tracer opentracing.Tracer
	router *gin.Engine
}

func NewServer(cfg *config.Config, db *gorm.DB, logger *zap.Logger, tracer opentracing.Tracer) *Server {
	server := &Server{
		cfg:    cfg,
		db:     db,
		logger: logger,
		tracer: tracer,
		router: gin.Default(),
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	// Initialize dependencies
	studentRepo := repository.NewStudentRepository(s.db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService, s.logger)

	// API routes
	api := s.router.Group("/api/v1")
	{
		students := api.Group("/students")
		{
			students.POST("/", studentHandler.Create)
			students.GET("/:id", studentHandler.Get)
			students.PUT("/:id", studentHandler.Update)
			students.DELETE("/:id", studentHandler.Delete)
			students.GET("/", studentHandler.List)
		}
	}
}

func (s *Server) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.cfg.Server.Host, s.cfg.Server.Port),
		Handler:      s.router,
		ReadTimeout:  time.Duration(s.cfg.Server.Timeout) * time.Second,
		WriteTimeout: time.Duration(s.cfg.Server.Timeout) * time.Second,
	}

	// Start server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	s.logger.Info("Server started", zap.String("addr", server.Addr))

	<-ctx.Done()
	s.logger.Info("Server stopping")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server shutdown failed: %v", err)
	}

	s.logger.Info("Server stopped gracefully")
	return nil
}
