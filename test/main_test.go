package test

import (
	"testing"

	"todo-api-gin-gorm/pkg/models"
)

func TestMain(m *testing.M) {
	models.MainTest(m, "..")
}
