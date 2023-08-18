package main

import (
	"context"

	"github.com/Nurquhart/learning-go/rest-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// init app
	err := initApp()
	if err != nil {
		panic(err)	
	}

	defer database.CloseMongoDB()
	
    app := fiber.New()

    app.Post("/", func(c *fiber.Ctx) error {
		// write a todo to the database
		sampleDoc := bson.M{"name": "sample todos"}
		collection := database.GetCollection("todos")

		nDoc, err := collection.InsertOne(context.TODO(), sampleDoc)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error inseting todo")
		}

		// send down info about todo
		return c.JSON(nDoc)
    })

    app.Listen(":3000")
}

func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}

	err = database.StartMongoDB()

	if err != nil {
		return err
	}
	return nil
}

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
