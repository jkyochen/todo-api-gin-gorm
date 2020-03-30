package models

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	// Needed for the MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

var tables []interface{}

func init() {
	tables = append(tables,
		new(User),
		new(Todo),
	)
}

// NewEngine initializes a new gorm Engine
func NewEngine() (err error) {

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME"))

	if db, err = gorm.Open(os.Getenv("DATABASE_DRIVER"), connStr); err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(3 * time.Second)

	if err = db.DB().Ping(); err != nil {
		return err
	}

	db.AutoMigrate(tables...)

	return nil
}
