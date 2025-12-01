package main

import "fmt"

// Question represents a single quiz question.
type Question struct {
	Sentence     string   `json:"sentence"`
	Options      []string `json:"options"`
	CorrectIndex int      `json:"correctIndex"`
	AnswerWords  []string `json:"answerWords"`
}

// QuizState holds the current state of the quiz.
type QuizState struct {
	WordBank     []string // List of words provided for the quiz
	CurrentIndex int      // Current question index
	Score        int      // Current score of the quiz
}

// Global variable to hold the state of the quiz.
var quizState QuizState

// StartQuiz initializes the quiz state when the quiz starts.
func StartQuiz(words []string) {
	quizState = QuizState{
		WordBank:     words,
		CurrentIndex: 0,
		Score:        0,
	}
}

// GetNextQuestion retrieves the next question for the quiz based on the current state.
func GetNextQuestion() (Question, error) {
	// Check if the quiz is finished
	if quizState.CurrentIndex >= len(quizState.WordBank) {
		return Question{}, fmt.Errorf("quiz completed")
	}

	// Generate the next question based on the current index
	q, err := buildSingleWordQuestion(quizState.WordBank, quizState.CurrentIndex, "", nil)
	if err != nil {
		return Question{}, err
	}

	// Increment the question index for the next round
	quizState.CurrentIndex++
	return q, nil
}

// CheckAnswer verifies the answer submitted by the user.
func CheckAnswer(selectedIndex int, correctIndex int) bool {
	// Check if the answer is correct
	if selectedIndex == correctIndex {
		quizState.Score++
		return true
	}
	return false
}

// GetScore returns the current score of the quiz.
func GetScore() int {
	return quizState.Score
}

// ResetQuiz resets the quiz state, starting it over from the beginning.
func ResetQuiz() {
	quizState = QuizState{}
}
