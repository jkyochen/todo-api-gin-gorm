package models

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	// Needed for the MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// NewConn initializes a new gorm Engine
func NewConn() (err error) {

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME"))

	var db *gorm.DB

	if db, err = gorm.Open(os.Getenv("DATABASE_DRIVER"), connStr); err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(3 * time.Second)

	if err = db.DB().Ping(); err != nil {
		return err
	}

	db.AutoMigrate()

	return nil
}
