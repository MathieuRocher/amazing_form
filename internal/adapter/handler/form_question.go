package handler

import (
	"amazing_form/internal/adapter/application"
	"net/http"
	"strconv"

	domain "github.com/MathieuRocher/amazing_domain"
	"github.com/gin-gonic/gin"
)

type FormQuestionHandler struct {
	useCase application.FormQuestionUseCaseInterface
}

func NewFormQuestionHandler(uc application.FormQuestionUseCaseInterface) *FormQuestionHandler {
	return &FormQuestionHandler{useCase: uc}
}

func (h *FormQuestionHandler) GetFormQuestions(c *gin.Context) {
	formQuestions, _ := h.useCase.FindAll()

	c.JSON(http.StatusOK, gin.H{
		"message": formQuestions,
	})
}

func (h *FormQuestionHandler) CreateFormQuestion(c *gin.Context) {
	var payload domain.FormQuestion
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
	c.JSON(http.StatusOK, gin.H{"message": "formQuestion created"})
	return
}

func (h *FormQuestionHandler) GetFormQuestionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	formQuestion, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "formQuestion not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"formQuestion": formQuestion})
	return
}

func (h *FormQuestionHandler) UpdateFormQuestionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// est-ce qu'il existe
	formQuestion, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "formQuestion not found"})
		return
	}

	// je le remplace par le body
	err = c.Bind(&formQuestion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// je save
	err = h.useCase.Update(formQuestion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "formQuestion Updated"})
	return
}

func (h *FormQuestionHandler) DeleteFormQuestionByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.useCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "formQuestion not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "formQuestion deleted succesfully"})
	return

}

func (h *FormQuestionHandler) RegisterRoutes(rg *gin.RouterGroup) {
	group := rg.Group("/form-questions")
	{
		group.GET("/", h.GetFormQuestions)
		group.POST("/", h.CreateFormQuestion)
		group.GET("/:id", h.GetFormQuestionByID)
		group.PUT("/:id", h.UpdateFormQuestionByID)
		group.DELETE("/:id", h.DeleteFormQuestionByID)
	}
}
