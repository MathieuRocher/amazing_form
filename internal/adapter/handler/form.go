package handler

import (
	"amazing_form/internal/adapter/application"
	"amazing_form/internal/adapter/handler/dto/form"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FormHandler struct {
	useCase application.FormUseCaseInterface
}

func NewFormHandler(uc application.FormUseCaseInterface) *FormHandler {
	return &FormHandler{useCase: uc}
}

func (h *FormHandler) GetForms(c *gin.Context) {
	forms, _ := h.useCase.FindAll()

	c.JSON(http.StatusOK, gin.H{
		"message": forms,
	})
}

func (h *FormHandler) CreateForm(c *gin.Context) {
	var payload form.CreateFormInput // TODO - Remplacer le domain.Form par handler.Form
	err := c.Bind(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed", "details": err.Error()})
		return
	}

	err = h.useCase.Create(payload.ToDomain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unknow sql error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "form created"})
	return
}

func (h *FormHandler) GetFormByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	form, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "form not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"form": form})
	return
}

func (h *FormHandler) UpdateFormByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// est-ce qu'il existe
	form, err := h.useCase.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "form not found"})
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
	c.JSON(http.StatusOK, gin.H{"message": "form Updated"})
	return
}

func (h *FormHandler) DeleteFormByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.useCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "form not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "form deleted succesfully"})
	return

}

func (h *FormHandler) RegisterRoutes(rg *gin.RouterGroup) {
	group := rg.Group("/forms")
	{
		group.GET("", h.GetForms)
		group.POST("", h.CreateForm)
		group.GET(":id", h.GetFormByID)
		group.PUT(":id", h.UpdateFormByID)
		group.DELETE(":id", h.DeleteFormByID)
	}
}
