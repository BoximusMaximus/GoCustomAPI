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
	{ID: 1, Name: "Link", Power: 56},
	{ID: 2, Name: "Zelda", Power: 27},
	{ID: 3, Name: "Ganon", Power: 48},
}

func InitDb(){
	connection, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	defer connection.Close(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	testCharacters := psql.NewModel(character{}, connection)
		if err != nil {
			fmt.Println(err)
			return
		}

	for i := 0;i < len(characters);i++ {
		fmt.Println("Inserting Character")
		testCharacters.Insert(characters[i])
	}
	
}

