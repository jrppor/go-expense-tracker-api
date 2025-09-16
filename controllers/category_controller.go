package controllers

import (
	"jrppor/go-expense-tracker-api/dto"
	"jrppor/go-expense-tracker-api/services"
	"net/http"
	"strconv"

	"strings"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{ svc services.CategoryService }

func NewCategoryController(svc services.CategoryService) *CategoryController {
	return &CategoryController{svc: svc}
}

func (ctrl *CategoryController) GetAllCategories(c *gin.Context) {
	list, err := ctrl.svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (ctrl *CategoryController) CreateCategory(c *gin.Context) {
	var req dto.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := ctrl.svc.Create(c.Request.Context(), req)
	if err != nil {
		// Translate duplicate key errors to 409 Conflict
		if isUniqueViolation(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "category with same name or slug already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (ctrl *CategoryController) UpdateCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var req dto.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := ctrl.svc.Update(c.Request.Context(), uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (ctrl *CategoryController) GetCategoryById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	e, err := ctrl.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, e)
}

func (ctrl *CategoryController) DeleteCategoryById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := ctrl.svc.Delete(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// isUniqueViolation returns true if error is a Postgres unique constraint violation (SQLSTATE 23505)
func isUniqueViolation(err error) bool {
	if err == nil {
		return false
	}
	// GORM wraps underlying driver errors; string match on code for simplicity across drivers
	// Avoid importing pg-specific packages; look for SQLSTATE 23505 or duplicate key phrase
	s := err.Error()
	return strings.Contains(s, "SQLSTATE 23505") || strings.Contains(strings.ToLower(s), "duplicate key value violates unique constraint")
}
