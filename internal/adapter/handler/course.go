package handler

import (
	"amazing_form/internal/adapter/application"
	"amazing_form/internal/adapter/handler/dto/course"
	"fmt"
	"net/http"
	"strconv"

	domain "github.com/MathieuRocher/amazing_domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CourseHandler struct {
	useCase application.CourseUsecaseInterface
}

func NewCourseHandler(uh application.CourseUsecaseInterface) *CourseHandler {
	return &CourseHandler{useCase: uh}
}

func (h *CourseHandler) RegisterRoutes(rg *gin.RouterGroup) {
	group := rg.Group("/courses")
	{
		group.GET("", h.GetCourses)
		group.POST("", h.CreateCourse)
		group.GET(":id", h.GetCourseByID)
		group.PUT(":id", h.UpdateCourseByID)
		group.DELETE(":id", h.DeleteCourseByID)

	}
}

func (h *CourseHandler) GetCourses(c *gin.Context) {
	forms, _ := h.useCase.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"message": forms,
	})
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {

	var payload course.CreateCourseInput
	err := c.Bind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	validate := validator.New()
	if err := validate.Struct(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}

	res := domain.Course{
		Title:       payload.Title,
		Description: payload.Description,
	}
	fmt.Println("payload : ", res)

	err = h.useCase.Create(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unknow sql error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "course created"})
}

func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	form, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"form": form})
}

func (h *CourseHandler) UpdateCourseByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// est-ce qu'il existe
	form, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "course not found"})
		return
	}

	// je le remplace par le body
	err = c.Bind(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	// je save
	err = h.useCase.Update(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "course Updated"})
}

func (h *CourseHandler) DeleteCourseByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.useCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "course not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "course deleted succesfully"})
}
