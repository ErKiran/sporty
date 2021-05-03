package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
	Role string             `bson:"role,omitempty"`
	Team primitive.ObjectID `bson:"team,omitempty"`
}

type Players struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	TeamID   primitive.ObjectID `bson:"team_id,omitempty"`
	MatchFee int                `bson:"match_fee,omitempty"`
}
