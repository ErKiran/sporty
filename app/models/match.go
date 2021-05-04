package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
