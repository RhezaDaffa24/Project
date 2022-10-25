package controllers

import (
	"learning/projectindividu/database"
	"learning/projectindividu/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginForm struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

// Struct AuthAPIController
type AuthAPIController struct {
	// Declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitAuthAPIController(s *session.Store) *AuthAPIController {
	db := database.InitDb()

	db.AutoMigrate(&models.User{})

	return &AuthAPIController{Db: db, store: s}
}

// Menampilkan Login
func (controller *AuthAPIController) Login(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)

	if err != nil {
		panic(err)
	}

	var dataUser models.User
	var dataForm LoginForm

	if err := c.BodyParser(&dataForm); err != nil {
		return c.JSON(fiber.Map{
			"msg": "Belum ada Username",
		})
	}

	//Mencari username
	errs := models.ViewUserByUsername(controller.Db, &dataUser, dataForm.Username)
	if errs != nil {
		return c.JSON(fiber.Map{
			"msg": "Tidak boleh kosong",
		})
	}

	//Hashing
	hash := bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(dataForm.Password))
	if hash == nil {
		sess.Set("username", dataUser.Username)
		sess.Set("userId", dataUser.ID)
		sess.Save()

		return c.JSON(fiber.Map{
			"msg": "Login berhasil",
		})
	}

	return c.JSON(fiber.Map{
		"msg": "Login error, Username atau Password yang dimasukkan tidak tepat",
	})
}

// Menampilkan akun apa saja yang telah register
func (controller *AuthAPIController) ViewRegister(c *fiber.Ctx) error {
	var dataUser []models.User
	err := models.ViewUser(controller.Db, &dataUser)
	if err != nil {
		return c.SendStatus(400)
	}
	return c.JSON(dataUser)
}

// Mendaftar akun baru
func (controller *AuthAPIController) Register(c *fiber.Ctx) error {
	var dataUser models.User
	var dataCart models.Cart

	if err := c.BodyParser(&dataUser); err != nil {
		return c.SendStatus(400)
	}

	
	bytes, _ := bcrypt.GenerateFromPassword([]byte(dataUser.Password), 10)
	hash := string(bytes)

	dataUser.Password = hash

	err := models.BuatAkun(controller.Db, &dataUser)
	if err != nil {
		return c.SendStatus(500)
	}

	errs := models.ViewUserByUsername(controller.Db, &dataUser, dataUser.Username)
	if errs != nil {
		return c.SendStatus(500)
	}

	errCart := models.BuatCart(controller.Db, &dataCart, dataUser.ID)
	if errCart != nil {
		return c.JSON(dataUser)
	}

	return c.JSON(dataUser)
}

// Keluar dari akun
func (controller *AuthAPIController) Logout(c *fiber.Ctx) error {

	sess, err := controller.store.Get(c)

	if err != nil {
		panic(err)
	}

	sess.Destroy()

	return c.JSON(fiber.Map{
		"msg": "Anda Berhasil Logout",
	})
}