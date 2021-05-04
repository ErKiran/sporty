package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"play-it/app/models"
	"time"

	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Migrate() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := models.SetUp()

	faker := faker.New()

	database := client.Database(os.Getenv("DATABASE"))

	player := database.Collection("player")
	team := database.Collection("team")
	match := database.Collection("match")

	var teams []models.Team

	for i := 0; i < 5; i++ {
		teamData := models.Team{
			Name: faker.Gamer().Tag(),
		}
		teams = append(teams, teamData)
	}

	multipleTeams := make([]interface{}, 0)

	for _, v := range teams {
		multipleTeams = append(multipleTeams, v)
	}

	insertManyResult, err := team.InsertMany(context.TODO(), multipleTeams)
	if err != nil {
		log.Fatal(err)
	}

	teamsID := insertManyResult.InsertedIDs

	roles := []string{"player", "admin", "captain"}

	var players []models.Player

	for i := 0; i < 75; i++ {
		playerData := models.Player{
			Name: fmt.Sprintf("%s %s", faker.Person().FirstName(), faker.Person().LastName()),
			Role: roles[GetRandomNumber(len(roles))],
			Team: teamsID[GetRandomNumber(len(teamsID))].(primitive.ObjectID),
		}
		players = append(players, playerData)
	}

	multiplePlayers := make([]interface{}, 0)

	for _, v := range players {
		multiplePlayers = append(multiplePlayers, v)
	}

	_, err = player.InsertMany(context.TODO(), multiplePlayers)
	if err != nil {
		log.Fatal(err)
	}

	var matches []models.Match

	matchType := []string{"Match", "Pratice Session", "Event/Meeting"}

	for i := 0; i < 2; i++ {
		matchData := models.Match{
			Title:           faker.RandomLetter(),
			Description:     faker.Lorem().Sentence(15),
			Type:            matchType[GetRandomNumber(len(matchType))],
			Location:        faker.Address().City(),
			StartDate:       faker.Time().Time(time.Now()),
			EndDate:         faker.Time().Time(time.Now()),
			MaximumAttendes: faker.RandomDigit(),
			Teams: []models.Teams{
				{TeamID: teamsID[GetRandomNumber(len(teamsID))].(primitive.ObjectID)},
				{TeamID: teamsID[GetRandomNumber(len(teamsID))].(primitive.ObjectID)},
			},
		}
		matches = append(matches, matchData)
	}
	multipleMatches := make([]interface{}, 0)

	for _, v := range matches {
		multipleMatches = append(multipleMatches, v)
	}

	_, err = match.InsertMany(context.TODO(), multipleMatches)
	if err != nil {
		log.Fatal(err)
	}

}

func GetRandomNumber(length int) int {
	return rand.Intn(length)
}

func main() {
	Migrate()
}
