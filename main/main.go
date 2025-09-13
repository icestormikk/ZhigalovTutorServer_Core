package main

import (
	"log"
	"net/http"
	"zhigalov_tutor_server_core/main/abstract/structs"
	"zhigalov_tutor_server_core/main/database"
	"zhigalov_tutor_server_core/main/server/controllers"
	"zhigalov_tutor_server_core/main/server/repos"
	"zhigalov_tutor_server_core/main/server/services"
)

func main() {
	router := http.NewServeMux()

	cfg := structs.NewEnvConfiguration(".env")
	db := database.NewPostgresDatabase(cfg)

	userRepo := repos.NewPostgresUserRepository(db)
	userService := services.NewDefaultUserService(userRepo)
	userController := controllers.NewDefaultUserController(userService)

	authService := services.NewDefaultAuthService(userRepo)
	authController := controllers.NewDefaultAuthController(authService)

	router.HandleFunc("GET /users", userController.GetUsers)
	//router.HandleFunc("GET /users/:id", userController.GetUser)
	router.HandleFunc("POST /users", userController.CreateUser)
	router.HandleFunc("POST /auth/login", authController.LoginUser)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panicln("Error starting server:", err)
	}
}
