package api

import (
	"net/http"
	"strconv"

	"todo-api-gin-gorm/pkg/models"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) uint {

	session, err := models.SessionStore.Get(c.Request, "session")
	if err != nil {
		return 0
	}

	userID, ok := session.Values["user_id"].(uint)
	if !ok {
		return 0
	}

	return userID
}

type TodoInput struct {
	Title   string `json:"title,omitempty" binding:"required"`
	Content string `json:"content"`
}

func CreateTodo(c *gin.Context) {

	userID := Auth(c)
	if userID == 0 {
		errorJSON(c, http.StatusUnauthorized, errNotLogin)
		return
	}

	var data TodoInput
	if err := c.ShouldBindJSON(&data); err != nil {
		errorJSON(c, http.StatusBadRequest, errBadRequest)
		return
	}

	row, err := models.CreateTodo(data.Title, data.Content, userID)
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, errInternalServer)
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

func UpdateTodo(c *gin.Context) {

	userID := Auth(c)
	if userID == 0 {
		errorJSON(c, http.StatusUnauthorized, errNotLogin)
		return
	}

	id := c.Param("id")
	if id == "" {
		errorJSON(c, http.StatusBadRequest, errBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		errorJSON(c, http.StatusBadRequest, errBadRequest)
		return
	}

	var data TodoInput
	if err := c.ShouldBindJSON(&data); err != nil {
		errorJSON(c, http.StatusBadRequest, errBadRequest)
		return
	}

	row, err := models.UpdateTodo(uint(todoID), data.Title, data.Content)
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, errInternalServer)
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

func DeleteTodo(c *gin.Context) {

	userID := Auth(c)
	if userID == 0 {
		errorJSON(c, http.StatusUnauthorized, errNotLogin)
		return
	}

	id := c.Param("id")
	if id == "" {
		errorJSON(c, http.StatusBadRequest, errBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		errorJSON(c, http.StatusBadRequest, errBadRequest)
		return
	}

	if models.DeleteTodo(uint(todoID)); err != nil {
		errorJSON(c, http.StatusInternalServerError, errInternalServer)
		return
	}

	c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"id": todoID,
		},
	)
}

func GetTodo(c *gin.Context) {

	userID := Auth(c)
	if userID == 0 {
		errorJSON(c, http.StatusUnauthorized, errNotLogin)
		return
	}

	id := c.Param("id")
	if id == "" {
		errorJSON(c, http.StatusBadRequest, errBadRequest)
		return
	}

	todoID, err := strconv.Atoi(id)
	if err != nil {
		errorJSON(c, http.StatusBadRequest, errBadRequest)
		return
	}

	todo, err := models.GetTodo(uint(todoID))
	if err != nil {
		errorJSON(c, http.StatusForbidden, errNotExist)
		return
	}

	c.JSON(
		http.StatusOK,
		todo,
	)
}

func GetTodos(c *gin.Context) {

	userID := Auth(c)
	if userID == 0 {
		errorJSON(c, http.StatusUnauthorized, errNotLogin)
		return
	}

	todos, err := models.GetTodos(userID)
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, errInternalServer)
		return
	}

	c.JSON(
		http.StatusOK,
		todos,
	)
}
