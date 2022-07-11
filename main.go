package main

import (
	"log"
	"user-api/config"
	"user-api/internal/db"
	"user-api/internal/db/seeder"
	"user-api/internal/http/router"
	"user-api/internal/http/server"
	userRepo "user-api/internal/repository/user"
	authService "user-api/internal/service/auth"
	userService "user-api/internal/service/user"
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
	authService := authService.NewAuthService(config, userRepository)
	seeder.RunAdminSeeder(userService)

	e := server.New()
	router.NewUserRoute(e, config, userService)
	router.NewAuthRoute(e, config, authService, userService)

	e.Logger.Fatal(e.Start(":8080"))
}
