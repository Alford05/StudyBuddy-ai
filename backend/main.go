package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("quiz_session", store))

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	apiKey := os.Getenv("SECRET_KEY")
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

	r.POST("/api/restart", func(c *gin.Context) {
		handleRestart(c)
	})

	r.Run(":8080")
}
