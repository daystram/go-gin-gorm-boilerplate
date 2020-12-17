package models

import (
	"time"

	"gorm.io/gorm"
)

type userOrm struct {
	db *gorm.DB
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Email     string `gorm:"unique"`
	Password  string
	Bio       string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type UserOrmer interface {
	GetOneByID(id uint) (user User, err error)
	GetOneByUsername(username string) (user User, err error)
	InsertUser(user User) (ID uint, err error)
	UpdateUser(user User) (err error)
}

func NewUserOrmer(db *gorm.DB) UserOrmer {
	//_ = db.AutoMigrate(&User{})		// builds table when enabled
	return &userOrm{db}
}


func (o *userOrm) GetOneByID(id uint) (user User, err error) {
	result := o.db.Model(&User{}).Where("id = ?", id).First(&user)
	return user, result.Error
}

func (o *userOrm) GetOneByUsername(username string) (user User, err error) {
	result := o.db.Model(&User{}).Where("username = ?", username).First(&user)
	return user, result.Error
}


func (o *userOrm) InsertUser(user User) (ID uint, err error) {
	result := o.db.Create(user)
	return user.ID, result.Error
}

func (o *userOrm) UpdateUser(user User) (err error) {
	result := o.db.Updates(user)
	return result.Error
}
