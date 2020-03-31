package api

import (
	"net/http"
	"strconv"

	"todo-api-gin-gorm/pkg/helper"
	"todo-api-gin-gorm/pkg/models"

	"github.com/gin-gonic/gin"
)

type todoInput struct {
	Title   string `json:"title,omitempty" binding:"required"`
	Content string `json:"content"`
}

// CreateTodo create todo
func CreateTodo(c *gin.Context) {

	userID := helper.GetUserDataSession(c.Request.Context())
	if userID == 0 {
		helper.ErrorJSON(c, http.StatusUnauthorized, helper.ErrNotLogin)
		return
	}

	var data todoInput
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.ErrorJSON(c, http.StatusBadRequest, helper.ErrBadRequest)
		return
	}

	row, err := models.CreateTodo(data.Title, data.Content, userID)
	if err != nil {
		helper.ErrorJSON(c, http.StatusInternalServerError, helper.ErrInternalServer)
		return
	}

	c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"id":         row.ID,
			"created_at": row.CreatedAt,
		},
	)
}

// UpdateTodo update todo
func UpdateTodo(c *gin.Context) {

	userID := helper.GetUserDataSession(c.Request.Context())
	if userID == 0 {
		helper.ErrorJSON(c, http.StatusUnauthorized, helper.ErrNotLogin)
		return
	}

	id := c.Param("id")
	if id == "" {
		helper.ErrorJSON(c, http.StatusBadRequest, helper.ErrBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorJSON(c, http.StatusBadRequest, helper.ErrBadRequest)
		return
	}

	var data todoInput
	if err := c.ShouldBindJSON(&data); err != nil {
		helper.ErrorJSON(c, http.StatusBadRequest, helper.ErrBadRequest)
		return
	}

	row, err := models.UpdateTodo(uint(todoID), data.Title, data.Content)
	if err != nil {
		helper.ErrorJSON(c, http.StatusInternalServerError, helper.ErrInternalServer)
		return
	}

	c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"id":        row.ID,
			"update_at": row.UpdatedAt,
		},
	)
}

// DeleteTodo delete todo
func DeleteTodo(c *gin.Context) {

	userID := helper.GetUserDataSession(c.Request.Context())
	if userID == 0 {
		helper.ErrorJSON(c, http.StatusUnauthorized, helper.ErrNotLogin)
		return
	}

	id := c.Param("id")
	if id == "" {
		helper.ErrorJSON(c, http.StatusBadRequest, helper.ErrBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorJSON(c, http.StatusBadRequest, helper.ErrBadRequest)
		return
	}

	if models.DeleteTodo(uint(todoID)); err != nil {
		helper.ErrorJSON(c, http.StatusInternalServerError, helper.ErrInternalServer)
		return
	}

	c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"id": todoID,
		},
	)
}

// GetTodo get single todo
func GetTodo(c *gin.Context) {

	userID := helper.GetUserDataSession(c.Request.Context())
	if userID == 0 {
		helper.ErrorJSON(c, http.StatusUnauthorized, helper.ErrNotLogin)
		return
	}

	id := c.Param("id")
	if id == "" {
		helper.ErrorJSON(c, http.StatusBadRequest, helper.ErrBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorJSON(c, http.StatusBadRequest, helper.ErrBadRequest)
		return
	}

	todo, err := models.GetTodo(uint(todoID))
	if err != nil {
		helper.ErrorJSON(c, http.StatusForbidden, helper.ErrNotExist)
		return
	}

	c.JSON(
		http.StatusOK,
		todo,
	)
}

// GetTodos get all todo about assign user
func GetTodos(c *gin.Context) {

	userID := helper.GetUserDataSession(c.Request.Context())
	if userID == 0 {
		helper.ErrorJSON(c, http.StatusUnauthorized, helper.ErrNotLogin)
		return
	}

	todos, err := models.GetTodos(userID)
	if err != nil {
		helper.ErrorJSON(c, http.StatusInternalServerError, helper.ErrInternalServer)
		return
	}

	c.JSON(
		http.StatusOK,
		todos,
	)
}
