package main

import (
	"log"
	"net/http"
	"user-api/config"
	"user-api/internal/db"
	"user-api/internal/http/router"
	"user-api/internal/http/server"
	userRepo "user-api/internal/repository/user"
	userService "user-api/internal/service/user"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := config.Load(".")
	if err != nil {
		log.Println(err)
	}
	db, err := db.NewDatabase(config)
	if err != nil {
		log.Println(err)
	}
	userRepository := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)

	e := server.New()
	router.NewUserRoute(e, userService)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
