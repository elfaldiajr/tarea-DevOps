package main

import (
	"log"

	controller "github.com/elfaldiajr/tarea-DevOps/internal/controller"
	repository "github.com/elfaldiajr/tarea-DevOps/internal/repository"
	service "github.com/elfaldiajr/tarea-DevOps/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	userRepo, err := repository.NewUserRepository()
	if err != nil {
		log.Fatalf("Error al inicializar el repositorio: %v", err)
	}

	userService := service.NewUserService(userRepo)

	userController := controller.NewUserController(userService)

	router := gin.Default()

	userController.RegisterRoutes(router)

	// Iniciar el servidor
	log.Println("Servidor iniciado en http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
