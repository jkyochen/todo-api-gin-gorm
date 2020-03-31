package models

import (
	"fmt"
	"time"

	"todo-api-gin-gorm/pkg/config"

	"github.com/jinzhu/gorm"
	"github.com/wader/gormstore"

	// Needed for the MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SessionStore control session save and get
var SessionStore *gormstore.Store

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

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Name)

	if db, err = gorm.Open(config.Database.Driver, connStr); err != nil {
		return fmt.Errorf("Failed to connect to database: %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(3 * time.Second)

	if err = db.DB().Ping(); err != nil {
		return fmt.Errorf("Failed to ping to database: %v", err)
	}

	db.AutoMigrate(tables...)

	startDebug()
	newSession()
	return nil
}

func startDebug() {
	if config.Server.Debug {
		db.Debug()
	}
}

func newSession() {
	SessionStore = gormstore.New(db, []byte(config.Session.Key))
	go SessionStore.PeriodicCleanup(24*time.Hour, make(chan struct{}))
}
