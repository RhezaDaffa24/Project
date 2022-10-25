package models

import (
	"gorm.io/gorm"
)

// Struct Transaksi
type Transaksi struct {
	gorm.Model
	Id       int `form:"id" json: "id" validate:"required"`
	UserID   uint
	Products []*Product `gorm:"many2many:transaksi_products;"`
}


// Menambahkan cart ke transaksi
func TambahTransaksi(db *gorm.DB, dataCart *Cart, dataProduct *Product) (err error) {
	dataCart.Products = append(dataCart.Products, dataProduct)
	err = db.Save(dataCart).Error
	if err != nil {
		return err
	}
	return nil
}

// Membuat transaksi baru
func BuatTransaksi(db *gorm.DB, dataTransaksi *Transaksi, userId uint, dataProduct []*Product) (err error) {
	dataTransaksi.UserID = userId
	dataTransaksi.Products = dataProduct

	err = db.Create(dataTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}

// Melihat semua transaksi
func ViewTransaksi(db *gorm.DB, dataTransaksi *Transaksi, id int) (err error) {
	err = db.Where("id=?", id).Preload("Products").Find(dataTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}

// Menampilkan transaksi berdasarkan id
func ViewTransaksiById(db *gorm.DB, dataTransaksi *[]Transaksi, id int) (err error) {
	err = db.Where("user_id=?", id).Find(dataTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}
