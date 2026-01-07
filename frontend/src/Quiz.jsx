// frontend/src/Quiz.jsx
import React, { useState } from "react";
import { getNextQuestion, checkAnswer } from "./api";

export default function Quiz({ words, firstQuestion, restartQuiz }) {
  const wordList = words.map((w) => w.trim());

  // 0..14 = current question, 15 = finished
  const [questionIndex, setQuestionIndex] = useState(0);
  const [question, setQuestion] = useState(firstQuestion);
  const [score, setScore] = useState(0);
  const [feedback, setFeedback] = useState("");

  // Fetch a question by index (0..14)
  const loadQuestion = async (index) => {
    const res = await getNextQuestion(index, wordList);
    setQuestion(res.data.question);
    setQuestionIndex(index);
    setFeedback("");
  };

  const choose = async (selectedIndex) => {
    // Ask backend if this choice is correct for THIS question index
    const res = await checkAnswer(questionIndex, selectedIndex);
    const correct = res.data.correct;
    const correctIndex = question.correctIndex;

    if (correct) {
      setScore((s) => s + 1);
      setFeedback("Correct!");
    } else {
      setFeedback("Wrong! Correct: " + question.options[correctIndex]);
    }

    // After showing feedback, move to next question or finish
    setTimeout(async () => {
      // If this was question 14 (15th question total), end the quiz
      if (questionIndex === 14) {
        setQuestionIndex(15); // triggers "Quiz Complete" view
        setFeedback("");
        return;
      }

      // Otherwise, load the next question (index + 1)
      const nextIndex = questionIndex + 1;
      await loadQuestion(nextIndex);
    }, 1500);
  };

  if (!question) {
    return <h3>Loading...</h3>;
  }

  if (questionIndex >= 15) {
    return (
      <div>
        <h2>Quiz Complete!</h2>
        <p>Score: {score} / 15</p>
        <button onClick={restartQuiz}>Restart Quiz</button>
      </div>
    );
  }

  return (
    <div style={{ padding: 30 }}>
      <h3>Question {questionIndex + 1} of 15</h3>
      <p>{question.sentence}</p>

      {question.options.map((opt, i) => (
        <button
          key={i}
          onClick={() => choose(i)}
          style={{ marginRight: "8px" }}
        >
          {opt}
        </button>
      ))}

      <p>{feedback}</p>
    </div>
  );
}