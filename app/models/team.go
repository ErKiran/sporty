package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Teams struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	TeamID primitive.ObjectID `bson:"team_id,omitempty"`
}

type Team struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}
