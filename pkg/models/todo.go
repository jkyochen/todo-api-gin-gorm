package models

import "github.com/jinzhu/gorm"

// Todo reminder thing
type Todo struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
}

func CreateTodo(title, content string, userID uint) (*Todo, error) {
	todo := Todo{Title: title, Content: content, UserID: userID}
	return &todo, db.Create(&todo).Error
}

func UpdateTodo(id uint, title, content string) (*Todo, error) {
	todo := new(Todo)
	if err := db.First(todo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	todo.Title = title
	todo.Content = content
	return todo, db.Save(todo).Error
}

func DeleteTodo(id uint) error {
	todo := new(Todo)
	return db.First(todo, "id = ?", id).Delete(todo).Error
}

func GetTodo(id uint) (*Todo, error) {
	todo := new(Todo)
	return todo, db.First(todo, "id = ?", id).Error
}

func GetTodos(userID uint) ([]*Todo, error) {
	var todos []*Todo
	return todos, db.Where("user_id = ?", userID).Find(&todos).Error
}
