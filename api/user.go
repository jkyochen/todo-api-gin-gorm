package api

import (
	"net/http"

	"todo-api-gin-gorm/pkg/models"

	"github.com/gin-gonic/gin"
)

type registerInput struct {
	Name     string `json:"name,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	Email    string `json:"email,omitempty" binding:"required"`
}

// Register regsiter user
func Register(c *gin.Context) {
	var data registerInput
	if err := c.ShouldBindJSON(&data); err != nil {
		errorJSON(c, http.StatusBadRequest, errBadRequest)
		return
	}

	row, err := models.CreateUser(data.Name, data.Password, data.Email)
	if models.IsErrUserAlreadyExist(err) {
		errorJSON(c, http.StatusConflict, errUserExist)
		return
	}
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

type loginInput struct {
	Email    string `json:"email,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

// Login user login
func Login(c *gin.Context) {
	var data loginInput
	if err := c.ShouldBindJSON(&data); err != nil {
		errorJSON(c, http.StatusBadRequest, errBadRequest)
		return
	}

	user, err := models.GetUser(data.Email, data.Password)
	if err != nil {
		errorJSON(c, http.StatusUnauthorized, errNotAuth)
		return
	}

	c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	)
}
