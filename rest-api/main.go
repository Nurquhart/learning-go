package main

import (
	"context"

	"github.com/Nurquhart/learning-go/rest-api/database"
	"github.com/Nurquhart/learning-go/rest-api/models"
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

		newPlaylist := new(models.PlaylistDTO)

		if err := c.BodyParser(newPlaylist); err != nil {
			return err
		}

		// mongodb is not great in this scenario, need this so songs is [] instead of null
		if newPlaylist.Songs == nil {
			newPlaylist.Songs = make([]models.Song, 0)
		}
		
		collection := database.GetCollection("Playlists") 
		nDoc, err := collection.InsertOne(context.TODO(), newPlaylist)
		if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error inseting todo")
			}

		return c.JSON(fiber.Map{"id": nDoc.InsertedID});
    })
  
	app.Get("/", func(c *fiber.Ctx) error {

		libCollection := database.GetCollection("Playlists") 
		cursor, err := libCollection.Find(context.TODO(), bson.M{})

		if err != nil {
			return err
		} 

		var playlists []models.Playlist
		if err = cursor.All(context.TODO(), &playlists); err != nil {
			return err
		}

		return c.JSON(playlists)

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
