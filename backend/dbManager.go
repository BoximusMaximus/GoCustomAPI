package main

import (
	"fmt"
	"context"

	"github.com/jackc/pgx/v5"
)

type character struct {
	ID    int16    `json:"id"`
	Name  string   `json:"name"`
	Power int    `json:"power"`
}

var characters = []character{
	{ID: 1, Name: "Link", Power: 56},
	{ID: 2, Name: "Zelda", Power: 27},
	{ID: 3, Name: "Ganon", Power: 48},
	{ID: 4, Name: "Midna", Power: 33},
}

func InitDb(connection *pgx.Conn){
	var sqlQuery string
	sqlQuery = "CREATE TABLE characters (" +
				"name varchar(80)," +
				"power int" +
				");"
	connection.Exec(context.Background(), sqlQuery)
	for i := 0; i < len(characters);i++{
		sqlQuery = fmt.Sprintf("INSERT INTO characters VALUES ('%s', '%d');", characters[i].Name, characters[i].Power)
		connection.Exec(context.Background(), sqlQuery)
	}

	

	
	
}

