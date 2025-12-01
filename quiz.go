package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

// ----- Types -----
type Question struct {
	Sentence     string   `json:"sentence"`
	Options      []string `json:"options"`
	CorrectIndex int      `json:"correct_index"`
}

type QuizState struct {
	mu              sync.Mutex
	WordBank        []string
	CurrentIndex    int
	CurrentQuestion Question
	Score           int
	QuestionBuffer  []Question
	RNG             *rand.Rand
	APIKey          string
}

var quiz QuizState

// ----- Main -----
func main() {
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	quiz.RNG = randSource

	apiKey := os.Getenv("SecretKey")
	if apiKey == "" {
		log.Println("SecretKey variable not set")
	}
	quiz.APIKey = apiKey // replace with your API key

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/start", startQuizHandler)
	http.HandleFunc("/next", nextQuestionHandler)
	http.HandleFunc("/answer", answerHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
