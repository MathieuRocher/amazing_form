package main

import (
	"amazing_form/internal/adapter/application"
	"amazing_form/internal/adapter/handler"
	"amazing_form/internal/adapter/repository"
	"amazing_form/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	router := gin.Default()

	database.InitDB()
	_ = database.DB.AutoMigrate(&repository.Course{}, &repository.CourseAssignment{}, &repository.Form{}, &repository.FormQuestion{})

	// c := cache.New(10*time.Minute, 10*time.Minute)
	// service.NewCacheService(c)
	// Importation des routes
	api := router.Group("/")

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

	reviewPort := os.Getenv("FORM_PORT")
	if reviewPort == "" {
		_ = godotenv.Load(".env") // charge localement si pas défini
		reviewPort = os.Getenv("FORM_PORT")
		if reviewPort == "" {
			reviewPort = "8081" // fallback si tout échoue
		}
	}

	err := router.Run("0.0.0.0:" + reviewPort)
	if err != nil {
		return
	}
}
