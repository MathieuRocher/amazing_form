package main

import (
	"amazing_form/internal/adapter/application"
	"amazing_form/internal/adapter/handler"
	"amazing_form/internal/adapter/repository"
	"amazing_form/internal/infrastructure/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()
	_ = database.DB.AutoMigrate(&repository.Form{}, &repository.FormQuestion{}, &repository.Course{}, &repository.CourseAssignment{})

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

	formPort := os.Getenv("FORM_PORT")
	if formPort == "" {
		log.Fatal("FORM_PORT must be set in environment")
	}

	err := router.Run("localhost:" + formPort)
	if err != nil {
		return
	}
}
