package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Blog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *gorm.DB

func main() {
	app := fiber.New()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=pgbouncer user=postgres password=postgres dbname=django port=6432 sslmode=disable TimeZone=Asia/Bangkok",
		PreferSimpleProtocol: true, // disables prepared statements
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(20)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Second * 10)

	app.Get("/api/blogs", func(c *fiber.Ctx) error {
		var blogs []Blog
		db.Find(&blogs) // Use the connection for database operations
		return c.JSON(blogs)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// app.Get("/api/blogs", handlers.GetBlogs)

	app.Listen(":8000")
}
