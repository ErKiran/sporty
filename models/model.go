package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetUp() *mongo.Client {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("nayaMatch")

	player := database.Collection("player")
	team := database.Collection("team")
	match := database.Collection("match")

	team1 := Team{
		Name: "Royal Fuckers India",
	}

	newTeam, err := team.InsertOne(context.TODO(), team1)
	if err != nil {
		log.Fatal(err)
	}

	player1 := Player{
		Name: "Kishan",
		Role: "Player",
		Team: newTeam.InsertedID.(primitive.ObjectID),
	}

	newPlayer, err := player.InsertOne(context.TODO(), player1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single player: ", newPlayer.InsertedID)

	fmt.Println("Inserted a single team ", newTeam.InsertedID)

	match1 := Match{
		Title:           "Nepal vs India",
		Type:            "International",
		Description:     "Khatra Match Hudai xa hai aaba",
		Location:        "Nepal",
		StartDate:       time.Now(),
		EndDate:         time.Now(),
		MaximumAttendes: 2000,
		Teams: []Teams{
			{
				TeamID: newTeam.InsertedID.(primitive.ObjectID),
			},
		},
		Players: []Players{
			{
				TeamID:   newTeam.InsertedID.(primitive.ObjectID),
				MatchFee: 1000,
			},
		},
	}

	newMatch, err := match.InsertOne(context.TODO(), match1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a match ", newMatch.InsertedID)

	fmt.Println("Connected to MongoDB!")
	return client
}
