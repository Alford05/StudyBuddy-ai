package main

import (
	"log"
	"net/http"
	"sync"
)

type QuizState struct {
	WordBank        []string
	CurrentIndex    int
	CurrentQuestion Question
	Score           int
	QuestionBuffer  []Question
	mu              sync.Mutex
}

var quiz QuizState

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static"))) //serve frontend
	http.HandleFunc("/start", startQuizHandler)
	http.HandleFunc("/next", nextQuestionHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
