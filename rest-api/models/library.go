package models


type Playlist struct {
	ID string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Songs []Song `json:"songs" bson:"songs"`
}

type Song struct {
	Title string `json:"title" bson:"title"`
	Artist string `json:"artist" bson:"artist"`
	Length string`json:"length" bson:"length"`
}

type PlaylistDTO struct {
	Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Songs []Song `json:"songs" bson:"songs"`
}