package main

type StartRequest struct {
	Words []string `json:"words"`
}

type QuestionResponse struct {
	Question Question `json:"question"`
}

type CheckRequest struct {
	Words         []string `json:"words"`
	QuestionIndex int      `json:"questionIndex"`
	SelectedIndex int      `json:"selectedIndex"`
}

type CheckResponse struct {
	Correct       bool   `json:"correct"`
	CorrectAnswer string `json:"correctAnswer"`
}
