package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type fizzBuzzCount struct {
	Count uint `json:"count"`
}

func main() {
	router := gin.Default()
	router.Use(Options)
	router.POST("/fizzbuzz", postFizzBuzz)
	var port = goDotEnvVariable("SERVER_PORT")
	router.Run("localhost:" + port)
}
func Options(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-Type", "application/json")
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

// Get the count number and calculate the FizzBuzz message
func postFizzBuzz(c *gin.Context) {
	var params fizzBuzzCount

	// Call BindJSON to bind the received JSON
	if err := c.BindJSON(&params); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status:": 400, "message": err.Error()})
	}

	number := params.Count

	var message = ""

	if number%15 == 0 {
		message = goDotEnvVariable("THIRD_MESSAGE")
	} else if number%5 == 0 {
		message = goDotEnvVariable("SECOND_MESSAGE")
	} else if number%3 == 0 {
		message = goDotEnvVariable("FIRST_MESSAGE")
	} else {
		message = ""
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status:": 200, "message": message})
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {
	// load .env file
	godotenv.Load(".env")

	return os.Getenv(key)
}
