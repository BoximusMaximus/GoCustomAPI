package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gopsql/psql"
	"github.com/jackc/pgx/v5"
)

type character struct {
	ID    int16    `json:"id"`
	Name  string   `json:"name"`
	Games []string `json:"games"`
	Power uint8    `json:"power"`
}

var characters = []character{
	{ID: 1, Name: "Link", Games: []string{"Majoras Mask", "Twilight Princess", "Etc"}, Power: 56},
	{ID: 2, Name: "Zelda", Games: []string{"Wind Waker", "Spirit Tracks"}, Power: 27},
	{ID: 3, Name: "Ganon", Games: []string{"Ocarina Of Time", "Hyrule Warriors"}, Power: 48},
}

func InitDb(){
	connection, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println(err)
		return
	}

	testCharacters := psql.NewModel(character{}, connection)

	for i := 0;i < len(characters);i++ {
		testCharacters.Insert(characters[i])
	}
	
}

