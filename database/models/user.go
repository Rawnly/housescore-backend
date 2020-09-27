package models

import (
	"gorm.io/gorm"
	"housescore/database"
	"housescore/util"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name"`
	LastName string `gorm:"not null" json:"last_name"`
	Email string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Status bool `gorm:"default:false" json:"status"`
	Phone string `gorm:"unique" json:"phone"`
}

func (u *User) TableName() string {
	return "users"
}
func (u *User) BeforeSaves()(err error) {
	u.Password = util.EncryptPassword(u.Password)
	return
}
func (u *User) CheckPass(pswd string) bool  {
	return util.EncryptPassword(pswd) == u.Password

}

func (u *User) FindByEmail(email string) *User {
	db := database.Instance()

	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}

	return &user
}

func (u *User) Create() error {
	db := database.Instance()

	_ = u.BeforeSaves()
	if err := db.Table(u.TableName()).Save(u).Error; err!= nil {
		return err
	}

	return nil
}



