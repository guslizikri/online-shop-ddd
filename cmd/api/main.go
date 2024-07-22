package main

import (
	"fmt"
	"log"
	"online-shop-ddd/external/database"
	"online-shop-ddd/internal/config"
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
}
