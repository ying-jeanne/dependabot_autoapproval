package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Define a route handler
	router.GET("/", func(c *gin.Context) {
		// Use the Redis client to get a value
		val, err := client.Get(c, "mykey").Result()
		if err != nil {
			c.String(500, "Error: "+err.Error())
			return
		}

		c.String(200, "Value: "+val)
	})

	// Start the server
	router.Run(":8080")
}
