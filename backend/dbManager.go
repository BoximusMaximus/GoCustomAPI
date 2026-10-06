package main

import (
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
}

func InitDb(connection *pgx.Conn){

	sqlQuery := `CREATE TABLE characters (` +
				`name varchar(80),` +
				`power int` +
				`);`

	connection.Exec(context.Background(), sqlQuery)
	connection.Exec(context.Background(), `INSERT INTO characters VALUES ('Link', '56');`)
}

