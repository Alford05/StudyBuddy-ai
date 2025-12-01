package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleStart(c *gin.Context, apiKey string, rng *rand.Rand) {
	var req StartRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if len(req.Words) != 10 {
		c.JSON(400, gin.H{"error": "Exactly 10 words required"})
		return
	}

	//Start the quiz with provided words
	StartQuiz(req.Words)

	q, err := buildSingleWordQuestion(req.Words, 0, apiKey, rng)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, QuestionResponse{Question: q})
}

func handleQuestion(c *gin.Context, apiKey string, rng *rand.Rand) {
	indexStr := c.Param("index")
	index, _ := strconv.Atoi(indexStr)

	words := c.QueryArray("words")
	if len(words) != 10 {
		c.JSON(400, gin.H{"error": "Words must be provided"})
		return
	}

	var (
		q   Question
		err error
	)

	if index < 10 {
		q, err = buildSingleWordQuestion(words, index, apiKey, rng)
	} else {
		q, err = buildTwoWordQuestion(words, apiKey, rng)
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, QuestionResponse{Question: q})
}

func handleCheck(c *gin.Context) {
	var req CheckRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	// Rebuild the same question to verify correctness
	var correctPair Question
	var err error
	if req.QuestionIndex < 10 {
		correctPair, err = buildSingleWordQuestion(req.Words, req.QuestionIndex, "", nil)
	} else {
		correctPair, err = buildTwoWordQuestion(req.Words, "", nil)
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	correct := req.SelectedIndex == correctPair.CorrectIndex
	if correct {
		CheckAnswer(req.SelectedIndex, correctPair.CorrectIndex) //update score if correct
	}
	c.JSON(200, CheckResponse{
		Correct:       correct,
		CorrectAnswer: correctPair.Options[correctPair.CorrectIndex],
	})
}
