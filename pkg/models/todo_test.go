package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {

	assert.NoError(t, PrepareTestDatabase())

	todo, err := CreateTodo("明天下雨", "", 1)
	assert.NoError(t, err)
	assert.NotNil(t, todo)
	assert.Equal(t, "明天下雨", todo.Title)
}

func TestUpdateTodo(t *testing.T) {

	assert.NoError(t, PrepareTestDatabase())

	todo, _ := CreateTodo("明天下雨", "test", 1)
	todo, err := UpdateTodo(todo.ID, "后天下雨", "")
	assert.NoError(t, err)
	assert.NotNil(t, todo)
	assert.Equal(t, "后天下雨", todo.Title)
	assert.Equal(t, "", todo.Content)
}

func TestDeleteTodo(t *testing.T) {

	assert.NoError(t, PrepareTestDatabase())

	todo, _ := CreateTodo("明天下雨", "", 1)
	err := DeleteTodo(todo.ID)
	assert.NoError(t, err)
}

func TestGetTodo(t *testing.T) {

	assert.NoError(t, PrepareTestDatabase())

	todo, _ := CreateTodo("明天下雨", "", 1)
	todo, err := GetTodo(todo.ID)
	assert.NoError(t, err)
	assert.NotNil(t, todo)
	assert.Equal(t, "明天下雨", todo.Title)
}

func TestGetTodos(t *testing.T) {

	assert.NoError(t, PrepareTestDatabase())

	CreateTodo("明天下雨", "", 1)
	CreateTodo("明天下雨", "", 2)
	CreateTodo("明天下雨", "", 1)

	todos, err := GetTodos(1)
	assert.NoError(t, err)
	assert.NotNil(t, todos)
	assert.Equal(t, len(todos), 3)
}
