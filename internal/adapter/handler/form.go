package handler

import (
	"amazing_form/internal/adapter/application"
	"amazing_form/internal/adapter/handler/dto/form"
	"net/http"
	"strconv"

	domain "github.com/MathieuRocher/amazing_domain"

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
	// Query params
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	courseIDStr := c.Query("course_id")
	classIDStr := c.Query("class_id")

	var (
		forms    []domain.Form
		err      error
		page     *int
		limit    *int
		courseID *int
		classID  *int
	)

	// Parse optional filters
	if courseIDStr != "" {
		id, err := strconv.Atoi(courseIDStr)
		if err == nil {
			courseID = &id
		}
	}

	if classIDStr != "" {
		id, err := strconv.Atoi(classIDStr)
		if err == nil {
			classID = &id
		}
	}

	if pageStr != "" && limitStr != "" {
		p, err1 := strconv.Atoi(pageStr)
		l, err2 := strconv.Atoi(limitStr)
		if err1 == nil && err2 == nil {
			page = &p
			limit = &l
		}
	}

	if classID != nil || courseID != nil || page != nil || limit != nil {
		forms, err = h.useCase.FindAllFiltered(courseID, classID, page, limit)
	} else {
		forms, err = h.useCase.FindAll()
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": forms})
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
