package main

import (
	"context"
	"log"
	"time"

	"github.com/elfaldiajr/tarea-DevOps/internal/db"
)

func main() {
	client, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	defer db.DisconnectDB(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error al verificar la conexión: %v", err)
	}

	log.Println("Conexión exitosa a MongoDB")
}
