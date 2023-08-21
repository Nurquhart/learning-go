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

	// create a new playlist
    app.Post("/", func(c *fiber.Ctx) error {

		newPlaylist := new(models.PlaylistDTO)

		if err := c.BodyParser(newPlaylist); err != nil {
			return err
		}
		
		collection := database.GetCollection("Playlists") 
		nDoc, err := collection.InsertOne(context.TODO(), newPlaylist)
		if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error inseting todo")
			}

		return c.JSON(fiber.Map{"id": nDoc.InsertedID});
    })
  
	// get a list of your playlists
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
 
	// delete a playlist
	app.Delete("/:name", func(c *fiber.Ctx) error {

		name := c.Params("name")
 
		libCollection := database.GetCollection("Playlists") 
		libCollection.DeleteOne(context.TODO(), bson.M{"name": name})
		
		return c.JSON("delete success")

    })

	// change the name of a playlist
	app.Put("/:name", func(c *fiber.Ctx) error {

		name := c.Params("name")
		newName := new(models.NewName)

		if err := c.BodyParser(newName); err != nil {
			return err
		}
 
		// update := bson.D{{"$inc", bson.D{{"sizes.$", -2}}}}
		filter := bson.D{{Key: "name", Value: name}}
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: newName.Name}}}}

		libCollection := database.GetCollection("Playlists") 
		libCollection.UpdateOne(context.TODO(), filter, update) 
		
		return c.JSON("Name updated")

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
