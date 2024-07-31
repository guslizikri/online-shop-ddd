package main

import (
	"fmt"
	"log"
	"online-shop-ddd/apps/products"
	"online-shop-ddd/apps/transactions"
	"online-shop-ddd/apps/users"
	"online-shop-ddd/external/database"
	"online-shop-ddd/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "cmd/api/config.yaml"

	if err := config.LoadConfig(filename); err != nil {
		log.Fatal("ini error file config.yaml", err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		log.Fatal("ini error db start", err)
	}

	if db != nil {
		fmt.Println("DB Connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	users.Init(router, db)
	products.Init(router, db)
	transactions.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
