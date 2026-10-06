package main

import (
	"fmt"
	"net/http"
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/jackc/pgx/v5"


)



func GetCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, characters)
}

func PostCharacters(c *gin.Context) {
	var newCharacter character

	// Call BindJSON to bind the received JSON to
	// newCharacter.
	if err := c.BindJSON(&newCharacter); err != nil {
		return
	}

	// Add character to the slice
	characters = append(characters, newCharacter)
	c.IndentedJSON(http.StatusCreated, newCharacter)

}
func main() {
	connection, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close(context.Background())
	InitDb(connection)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// temp for testing
		// Start
		AllowAllOrigins: true,
		// End
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))
	router.GET("/characters", GetCharacters)
	router.POST("/characters", PostCharacters)

	router.Run(":8080")
}
