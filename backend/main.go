package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	apiKey := os.Getenv("SecretKey")
	if apiKey == "" {
		log.Fatal("FATAL: SecretKey variable not set")
	}
	r.POST("/api/start", func(c *gin.Context) {
		handleStart(c, apiKey, rng)
	})
	r.GET("/api/question/:index", func(c *gin.Context) {
		handleQuestion(c, apiKey, rng)
	})
	r.POST("/api/check", handleCheck)

	r.Run(":8080")
}
