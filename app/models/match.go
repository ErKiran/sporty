package models

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Match struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Title           string             `bson:"title,omitempty"`
	Type            string             `bson:"type,omitempty"`
	Description     string             `bson:"description,omitempty"`
	Location        string             `bson:"location,omitempty"`
	StartDate       time.Time          `bson:"start_date,omitempty"`
	EndDate         time.Time          `bson:"end_date,omitempty"`
	MaximumAttendes int                `bson:"maximum_attendes,omitempty"`
	Players         []Players          `bson:"players,omitempty"`
	Teams           []Teams            `bson:"teams,omitempty"`
}

func (data *Match) Validate() error {

	if data.Title == "" {
		return errors.New("title is required")
	}

	if data.Type == "" {
		return errors.New("type is required")
	}

	if data.Description == "" {
		return errors.New("description is required")
	}

	if data.Location == "" {
		return errors.New("location is required")
	}

	if data.StartDate.IsZero() {
		return errors.New("start time is required")
	}

	if data.EndDate.IsZero() {
		return errors.New("end data is required")
	}
	return nil

}

func (data *Match) Create(db *mongo.Client) error {
	match := db.Database(os.Getenv("DATABASE")).Collection("match")

	newMatch, err := match.InsertOne(context.TODO(), data)

	if err != nil {
		return err
	}
	fmt.Println("newMatch", newMatch.InsertedID)
	return nil
}
