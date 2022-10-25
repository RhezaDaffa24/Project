package controllers

import (
	"strconv"

	"learning/projectindividu/database"
	"learning/projectindividu/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type TransaksiAPIController struct {
	// Declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitTransaksiAPIController(s *session.Store) *TransaksiAPIController {
	db := database.InitDb()

	db.AutoMigrate(&models.Transaksi{})

	return &TransaksiAPIController{Db: db, store: s}
}

// Checkout
func (controller *TransaksiAPIController) AddTransaksi(c *fiber.Ctx) error {
	params := c.AllParams()
	intUserId, _ := strconv.Atoi(params["userid"])

	var dataTransaksi models.Transaksi
	var dataCart models.Cart

	err := models.ViewCart(controller.Db, &dataCart, intUserId)
	if err != nil {
		return c.SendStatus(400)
	}

	errs := models.BuatTransaksi(controller.Db, &dataTransaksi, uint(intUserId), dataCart.Products)
	if errs != nil {
		return c.SendStatus(400)
	}

	errss := models.UpdateCart(controller.Db, dataCart.Products, &dataCart, uint(intUserId))

	if errss != nil {
		return c.SendStatus(400)
	}

	if intUserId != 1 {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"Title": "History Transaksi",
		"Tid":   intUserId,
	})
}

// history transaksi
func (controller *TransaksiAPIController) GetTransaksi(c *fiber.Ctx) error {
	params := c.AllParams()

	intUserId, _ := strconv.Atoi(params["userid"])

	var dataTransaksi []models.Transaksi
	err := models.ViewTransaksiById(controller.Db, &dataTransaksi, intUserId)
	if err != nil {
		return c.SendStatus(500)
	}

	if intUserId != 1 {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"Title":      "History Transaksi",
		"Transaksis": dataTransaksi,
	})

}

// Menampilkan detail transaksi
func (controller *TransaksiAPIController) DetailTransaksi(c *fiber.Ctx) error {
	params := c.AllParams()

	intTransaksiId, _ := strconv.Atoi(params["transaksiid"])

	var dataTransaksi models.Transaksi
	err := models.ViewTransaksi(controller.Db, &dataTransaksi, intTransaksiId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{
		"Title":    "History Transaksi",
		"Products": dataTransaksi.Products,
	})
}