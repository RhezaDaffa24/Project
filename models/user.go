package models

import (
	"gorm.io/gorm"
)

// Struct User
type User struct {
	gorm.Model
	Name       string `form:"name" json: "name" validate:"required"`
	Username   string `form:"username" json: "username" validate:"required"`
	Email      string `form:"email" json: "email" validate:"required"`
	Password   string `form:"password" json: "password" validate:"required"`
	Cart       Cart
	Transaksis []Transaksi
}

// Membuat user baru
func BuatAkun(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Melihat semua data user
func ViewUser(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Menampilkan user berdasarkan username
func ViewUserByUsername(db *gorm.DB, User *User, username string) (err error) {
	err = db.Where("username=?", username).First(User).Error
	if err != nil {
		return err
	}
	return nil
}