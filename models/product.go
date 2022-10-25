package models

import (
	"gorm.io/gorm"
)

// Struct Product
type Product struct {
	gorm.Model
	Id         int          `form:"id" json: "id" validate:"required"`
	Name       string       `form:"name" json: "name" validate:"required"`
	Image      string       `form:"image" json: "image" validate:"required"`
	Shop       string       `form:"shop" json: "shop" validate:"required"`
	Price      float32      `form:"price" json: "price" validate:"required"`
	Carts      []*Cart      `gorm:"many2many:cart_products;"`
	Transaksis []*Transaksi `gorm:"many2many:transaksi_products;"`
}

// Fungsi untuk membuat produk baru
func CreateProduct(db *gorm.DB, dataProduct *Product) (err error) {
	err = db.Create(dataProduct).Error
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk memanggil semua data produk
func GetProducts(db *gorm.DB, dataProduct *[]Product) (err error) {
	err = db.Find(dataProduct).Error
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk memanggil data produk berdasarkan idnya
func GetProductById(db *gorm.DB, dataProduct *Product, id int) (err error) {
	err = db.Where("id=?", id).First(dataProduct).Error
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk mengupdate/mengedit data produk berdasarkan id
func UpdateProduct(db *gorm.DB, dataProduct *Product) (err error) {
	db.Save(dataProduct)

	return nil
}

// Fungsi untuk mendestroy.menghapus produk berdasarkan id
func DeleteProduct(db *gorm.DB, dataProduct *Product, id int) (err error) {
	db.Where("id=?", id).Delete(dataProduct)

	return nil
}