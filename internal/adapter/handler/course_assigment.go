package handler

import (
	"amazing_form/internal/adapter/application"
	"net/http"
	"strconv"

	domain "github.com/MathieuRocher/amazing_domain"
	"github.com/gin-gonic/gin"
)

type CourseAssignmentHandler struct {
	useCase application.CourseAssignmentUseCaseInterface
}

func NewCourseAssignmentHandler(uc application.CourseAssignmentUseCaseInterface) *CourseAssignmentHandler {
	return &CourseAssignmentHandler{useCase: uc}
}

func (h *CourseAssignmentHandler) GetCourseAssignments(c *gin.Context) {
	courseAssignments, _ := h.useCase.FindAll()

	c.JSON(http.StatusOK, gin.H{
		"message": courseAssignments,
	})
}

func (h *CourseAssignmentHandler) CreateCourseAssignment(c *gin.Context) {
	var payload domain.CourseAssignment
	err := c.Bind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	err = h.useCase.Create(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unknow sql error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "courseAssignment created"})
	return
}

func (h *CourseAssignmentHandler) GetCourseAssignmentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	courseAssignment, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "courseAssignment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"courseAssignment": courseAssignment})
	return
}

func (h *CourseAssignmentHandler) UpdateCourseAssignmentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// est-ce qu'il existe
	courseAssignment, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "courseAssignment not found"})
		return
	}

	// je le remplace par le body
	err = c.Bind(&courseAssignment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// je save
	err = h.useCase.Update(courseAssignment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "courseAssignment Updated"})
	return
}

func (h *CourseAssignmentHandler) DeleteCourseAssignmentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.useCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted succesfully"})
	return

}

func (h *CourseAssignmentHandler) RegisterRoutes(rg *gin.RouterGroup) {
	group := rg.Group("/course-assignments")
	{
		group.GET("", h.GetCourseAssignments)
		group.GET(":id", h.GetCourseAssignmentByID)
		group.POST("", h.CreateCourseAssignment)
		group.PUT(":id", h.UpdateCourseAssignmentByID)
		group.DELETE(":id", h.DeleteCourseAssignmentByID)
	}
}
