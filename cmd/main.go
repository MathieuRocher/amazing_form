package main

import (
	"amazing_form/internal/adapter/application"
	"amazing_form/internal/adapter/handler"
	"amazing_form/internal/adapter/repository"
	"amazing_form/internal/infrastructure/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.InitDB()
	// c := cache.New(10*time.Minute, 10*time.Minute)
	// service.NewCacheService(c)
	// Importation des routes
	api := router.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	// FORM
	formRepository := repository.NewFormRepository()
	formUseCase := application.NewFormUseCase(formRepository)
	formHandler := handler.NewFormHandler(formUseCase)
	formHandler.RegisterRoutes(api)

	// FORM QUESTIONS
	formQuestionsRepository := repository.NewFormQuestionRepository()
	formQuestionsUseCase := application.NewFormQuestionUseCase(formQuestionsRepository)
	formQuestionsHandler := handler.NewFormQuestionHandler(formQuestionsUseCase)
	formQuestionsHandler.RegisterRoutes(api)

	// COURSE
	courseRepository := repository.NewCourseRepository()
	courseUseCase := application.NewCourseUsecase(courseRepository)
	courseHandler := handler.NewCourseHandler(courseUseCase)
	courseHandler.RegisterRoutes(api)

	// COURSE ASSIGNMENT
	courseAssignmentRepository := repository.NewCourseAssignmentRepository()
	courseAssignmentUseCase := application.NewCourseAssignmentUseCase(courseAssignmentRepository)
	courseAssignmentHandler := handler.NewCourseAssignmentHandler(courseAssignmentUseCase)
	courseAssignmentHandler.RegisterRoutes(api)
	router.Run("localhost:8081")
}
