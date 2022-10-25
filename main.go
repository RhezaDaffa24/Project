package main

import (
	"learning/projectindividu/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	// // session
	store := session.New()

	app := fiber.New()

	// // static
	// app.Static("/", "./public", fiber.Static{
	// 	Index: "",
	// })

	// app.Static("/public", "./public")

	// Fungsi Check Login
	CheckLogin := func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		temp := sess.Get("username")
		if temp != nil {
			return c.Next()
		}

		return c.SendString("Anda Belum login")
	}

	// controllers
	authController := controllers.InitAuthAPIController(store)
	productsController := controllers.InitProductAPIController()
	cartController := controllers.InitCartAPIController(store)
	transaksiController := controllers.InitTransaksiAPIController(store)

	//API group
	auth := app.Group("/api")
	p := app.Group("/shop")
	c := app.Group("/cart")
	t := app.Group("/pay")

	//auth
	auth.Post("/login", authController.Login)
	auth.Get("/listakun", authController.ViewRegister)
	auth.Post("/register", authController.Register)
	auth.Post("/logout", authController.Logout)

	//products
	p.Get("/products", CheckLogin, productsController.GetProduct)
	p.Post("/products", CheckLogin, productsController.CreateProduct)
	p.Get("/products/detail/:id", CheckLogin, productsController.GetDetailProduct)
	p.Put("/products/:id", CheckLogin, productsController.UpdateProduct)
	p.Delete("/products/:id", CheckLogin, productsController.DeleteProduct)

	//cart
	c.Get("/addcart/:cartid/product/:productid", CheckLogin, cartController.AddCart)
	c.Get("/cart/:cartid", CheckLogin, cartController.GetCart)

	//transaksi
	t.Get("/listtransaksi/:userid", CheckLogin, transaksiController.GetTransaksi)
	t.Get("/detail/:transaksiid", CheckLogin, transaksiController.DetailTransaksi)
	t.Get("/checkout/:userid", CheckLogin, transaksiController.AddTransaksi)

	app.Listen(":3000")
}