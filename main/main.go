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
	repo := repos.NewPostgresUserRepository(db)
	service := services.NewDefaultUserService(repo)
	controller := controllers.NewDefaultUserController(service)

	router.HandleFunc("GET /users", controller.GetAllUsers)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panicln("Error starting server:", err)
	}
}
