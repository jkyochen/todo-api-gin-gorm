package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User represents the object of individual and member of organization.
type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index"`
	Name      string
	Password  string
	Lastlogin *time.Time
}

// CreateUser creates record of a new user.
func CreateUser(name, password, email string) (*User, error) {

	pdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	row := &User{
		Name:     name,
		Email:    email,
		Password: string(pdHash),
	}

	err = db.Transaction(func(tx *gorm.DB) error {

		if !tx.Where("email = ?", email).First(new(User)).RecordNotFound() {
			return ErrUserAlreadyExist{0, name}
		}

		return tx.Create(row).Error
	})

	return row, err
}
