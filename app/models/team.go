package models

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Teams struct {
	TeamID primitive.ObjectID `bson:"team_id,omitempty" json:"teamId"`
}

type Team struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string             `bson:"name,omitempty" json:"name"`
}

func (data Team) FindAll(db *mongo.Client) ([]Team, error) {
	teamCollection := db.Database(os.Getenv("DATABASE")).Collection("team")
	cursor, err := teamCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	var teams []Team
	if err = cursor.All(context.TODO(), &teams); err != nil {
		return nil, err
	}
	return teams, nil
}
