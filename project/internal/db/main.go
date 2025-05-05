package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Cliente de MongoDB global
var Client *mongo.Client

// Función que inicializa la conexión a MongoDB
func ConnectDB() (*mongo.Client, error) {
	// Obtener variables de entorno
	username := getEnv("MONGO_USERNAME", "admin")
	password := getEnv("MONGO_PASSWORD", "password")
	host := getEnv("MONGO_HOST", "localhost")
	port := getEnv("MONGO_PORT", "27017")

	// Construir string de conexión
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		username, password, host, port)

	// Opciones del cliente de MongoDB
	clientOptions := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(10 * time.Second)

	// Crear un contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conectar a MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a MongoDB: %v", err)
	}

	// Verificar la conexión
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("error al hacer ping a MongoDB: %v", err)
	}

	log.Println("Conexión exitosa a MongoDB")
	Client = client
	return client, nil
}

// Función para cerrar la conexión a MongoDB
func DisconnectDB(client *mongo.Client) {
	if client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Printf("Error al desconectar de MongoDB: %v", err)
	}
	log.Println("Conexión a MongoDB cerrada correctamente")
}

// Función auxiliar para obtener variables de entorno con valor por defecto
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
