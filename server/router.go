package server

import (
	"context"
	"go-hex-mongo/internal/adapters/db/mongodb"
	"go-hex-mongo/internal/adapters/handlers"
	"go-hex-mongo/internal/domains/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupRoutes(app *fiber.App) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve MongoDB URI and database name from environment variables
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")

	// Initialize MongoDB connection
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Select the database using the environment variable
	db := client.Database(dbName)

	// Initialize repository, service, and handler
	productRepo := mongodb.NewProductRepoImpl(db, "product")
	productService := services.NewProductServiceImpl(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Routes
	app.Post("/products", productHandler.CreateProduct)
	app.Get("/products/:id", productHandler.GetProductByID)
	app.Get("/products", productHandler.GetAllProducts)
	app.Put("/products/:id", productHandler.UpdateProduct)
	app.Delete("/products/:id", productHandler.DeleteProduct)
}
