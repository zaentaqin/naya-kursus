package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	NoHp     string `json:"nohp"`
	ClassID  int    `json:"class_id"`
	Class    Class  `gorm:"foreignKey:ClassID" json:"class"`
}

type Class struct {
	gorm.Model
	ClassName string `json:"class_name"`
	Levels    string `json:"level"`
	Intructor string `json:"instructor"`
	Price     int    `json:"price"`
}

func main() {
	dsn := "host=localhost user=zaen password=zaen dbname=naya_kursus port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to Conecting to Databse", err)
	}
	fmt.Println("Database Conections Succussfull...")

	db.AutoMigrate(&User{}, &Class{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	app.Post("/class", func(c *fiber.Ctx) error {
		class := Class{}
		err := c.BodyParser(&class)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		db.Create(&class)
		if class.ClassName != "" && class.Price != 0 {
			return c.JSON(fiber.Map{
				"success": true,
				"message": "",
				"data":    class,
			})
		}
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintln("Class_name and price field is required"),
			"data":    nil,
		})
	})

	app.Get("/class", func(c *fiber.Ctx) error {
		name := c.Query("class_name")
		class := []Class{}

		if name != "" {
			result := db.Where("class_name LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&class)
			if result.RowsAffected != 0 {
				return c.JSON(fiber.Map{
					"success": true,
					"message": "",
					"data":    class,
				})
			}
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("class name %s not found", name),
				"data":    class,
			})
		}

		db.Find(&class)
		return c.JSON(fiber.Map{
			"success": true,
			"message": "",
			"data":    class,
		})
	})

	app.Get("/class/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		class := Class{}
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		result := db.First(&class, id)
		if result.RowsAffected != 0 {
			return c.JSON(fiber.Map{
				"success": true,
				"message": "",
				"data":    class,
			})
		}

		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("class with ID %d not found", id),
			"data":    nil,
		})
	})

	app.Put("/class/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		classRequest := Class{}
		if err := c.BodyParser(&classRequest); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		class := Class{}
		if err := db.First(&class, id).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		classRequest.ID = uint(id)
		if classRequest.ClassName == "" {
			classRequest.ClassName = class.ClassName
		}

		if classRequest.Levels == "" {
			classRequest.Levels = class.Levels
		}

		if classRequest.Intructor == "" {
			classRequest.Intructor = class.Intructor
		}

		if classRequest.Price == 0 {
			classRequest.Price = class.Price
		}

		if err := db.Updates(&classRequest).Error; err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"message": "",
			"data":    classRequest,
		})
	})

	app.Delete("/class/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		class := Class{}
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		result := db.Delete(&class, id)
		if result.RowsAffected != 0 {
			return c.JSON(fiber.Map{
				"success": true,
				"message": fmt.Sprintf("class with ID %d  successfully deleted ", id),
				"data":    class,
			})
		}

		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Class not found",
			"data":    nil,
		})

	})

	/*--------------------Router User-------------------------------------*/

	app.Post("/user", func(c *fiber.Ctx) error {
		user := User{}
		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		db.Create(&user)
		db.Preload("Class").Find(&user)
		return c.JSON(fiber.Map{
			"success": true,
			"message": "",
			"data":    user,
		})
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		user := []User{}

		db.Preload("Class").Find(&user)
		return c.JSON(fiber.Map{
			"success": true,
			"message": "",
			"data":    user,
		})
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadGateway).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		user := User{}
		result := db.Preload("Class").First(&user, id)
		if result.RowsAffected != 0 {
			return c.JSON(fiber.Map{
				"success": true,
				"message": "",
				"data":    user,
			})
		}

		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
			"data":    nil,
		})
	})

	app.Put("/user/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		userRequest := User{}
		if err := c.BodyParser(&userRequest); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		user := User{}
		if err := db.Preload("Class").First(&user, id).Error; err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		userRequest.ID = uint(id)
		userRequest.UserName = user.UserName

		if userRequest.FullName == "" {
			userRequest.FullName = user.FullName
		}

		if userRequest.Email == "" {
			userRequest.Email = user.Email
		}

		if userRequest.NoHp == "" {
			userRequest.NoHp = user.NoHp
		}

		if userRequest.ClassID == 0 {
			userRequest.ClassID = user.ClassID
		}

		if err := db.Updates(&userRequest).Error; err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Update Failed",
				"data":    nil,
			})
		}

		db.Preload("Class").First(&userRequest)
		return c.JSON(fiber.Map{
			"success": true,
			"message": "",
			"data":    userRequest,
		})
	})

	app.Delete("/user/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		result := db.Delete(&User{}, id)
		if result.RowsAffected != 0 {
			return c.JSON(fiber.Map{
				"success": "true",
				"message": fmt.Sprintf("user with ID %d succesfully deleted", id),
				"data":    User{},
			})
		}

		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	})

	type UserCount struct {
		ClassID   int
		UserCount int `json:"user_count"`
	}

	app.Get("/report", func(c *fiber.Ctx) error {
		UserCounts := []UserCount{}
		db.Table("users").
			Select("users.class_id, count(users.id) as user_count").
			Group("users.class_id").
			Order("user_count DESC").
			Scan(&UserCounts)

		classes := []Class{}

		for _, classUserCount := range UserCounts {
			if classUserCount.UserCount == UserCounts[0].UserCount {
				class := Class{}
				db.Where("id = ?", classUserCount.ClassID).First(&class)
				classes = append(classes, class)
			}
		}

		return c.JSON(fiber.Map{
			"Class": fiber.Map{
				"data": classes,
			},
			"UserCount": UserCounts[0].UserCount,
		})
	})

	app.Listen(":8080")
}
