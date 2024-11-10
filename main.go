package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Magic 8-Ball responses
var responses = []string{
	"It is certain.",
	"It is decidedly so.",
	"Without a doubt.",
	"Yes – definitely.",
	"You may rely on it.",
	"As I see it, yes.",
	"Most likely.",
	"Outlook good.",
	"Yes.",
	"Signs point to yes.",
	"Reply hazy, try again.",
	"Ask again later.",
	"Better not tell you now.",
	"Cannot predict now.",
	"Concentrate and ask again.",
	"Don't count on it.",
	"My reply is no.",
	"My sources say no.",
	"Outlook not so good.",
	"Very doubtful.",
}

func main() {
	r := gin.Default()

	// Endpoint to respond to /ping with "ack"
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ack")
	})

	// Endpoint to respond to /8ball with a random Magic 8-Ball response
	r.GET("/8ball", func(c *gin.Context) {
		// Seed random generator and pick a random response
		rand.Seed(time.Now().UnixNano())
		response := responses[rand.Intn(len(responses))]
		c.JSON(http.StatusOK, gin.H{"response": response})
	})

	// Run the server on port 8080
	r.Run(":8080")
}

