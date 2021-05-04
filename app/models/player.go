package models

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Player struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string             `bson:"name,omitempty" json:"name"`
	Role string             `bson:"role,omitempty" json:"role"`
	Team primitive.ObjectID `bson:"team,omitempty" json:"team"`
}

type Players struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	TeamID   primitive.ObjectID `bson:"team_id,omitempty" json:"teamID"`
	MatchFee int                `bson:"match_fee,omitempty" json:"matchFee"`
}

func (data Player) FindAll(db *mongo.Client) ([]Player, error) {
	playerCollection := db.Database(os.Getenv("DATABASE")).Collection("player")
	cursor, err := playerCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	var players []Player
	if err = cursor.All(context.TODO(), &players); err != nil {
		return nil, err
	}
	return players, nil
}
