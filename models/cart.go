package models

import (
	"gorm.io/gorm"
)

// Struct Cart
type Cart struct {
	gorm.Model
	UserID   uint
	Products []*Product `gorm:"many2many:cart_products;"`
}

// Membuat cart baru
func BuatCart(db *gorm.DB, dataCart *Cart, userId uint) (err error) {
	dataCart.UserID = userId
	err = db.Create(dataCart).Error
	if err != nil {
		return err
	}
	return nil
}

// Menambahkan isi cart
func TambahCart(db *gorm.DB, dataCart *Cart, product *Product) (err error) {
	dataCart.Products = append(dataCart.Products, product)
	err = db.Save(dataCart).Error
	if err != nil {
		return err
	}
	return nil
}

// Menampilkan semua isi cart
func ViewCart(db *gorm.DB, dataCart *Cart, id int) (err error) {
	err = db.Where("user_id=?", id).Preload("Products").Find(dataCart).Error
	if err != nil {
		return err
	}
	return nil
}

// Menampilkan cart sesuai dengan id
func ViewCartById(db *gorm.DB, dataCart *Cart, id int) (err error) {
	err = db.Where("user_id=?", id).First(dataCart).Error
	if err != nil {
		return err
	}
	return nil
}

// Menghapus cart saat sudah transaksi
func UpdateCart(db *gorm.DB, dataProduct []*Product, dataCart *Cart, userId uint) (err error) {
	db.Model(&dataCart).Association("Products").Delete(dataProduct)

	return nil
}