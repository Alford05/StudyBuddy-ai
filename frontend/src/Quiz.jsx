import React, { useState, useEffect } from "react";
import { getQuestion, checkAnswer } from "./api";

export default function Quiz({ words, firstQuestion, restartQuiz }) {
  const wordList = words.map((w) => w.trim());
  const [questionIndex, setQuestionIndex] = useState(0);
  const [question, setQuestion] = useState(firstQuestion);
  const [score, setScore] = useState(0);
  const [feedback, setFeedback] = useState("");

  const loadNext = async () => {
    if (questionIndex === 14) return;

    const res = await getQuestion(questionIndex + 1, wordList);
    setQuestion(res.data.question);
    setQuestionIndex(questionIndex + 1);
    setFeedback("");
  };

  const choose = async (i) => {
    const correctIndex = question.CorrectIndex;
    const res = await checkAnswer(i, correctIndex);
    const correct = res.data.correct;

    if (correct) {
      setScore((s) => s + 1);
      setFeedback("Correct!");
    } else {
      setFeedback("Wrong! Correct: " + res.data.correctAnswer);
    }

    setTimeout(loadNext, 1500);
  };

  if (!question)
    return <h3>Loading...</h3>;

  if (questionIndex >= 15)
    return (
      <div>
        <h2>Quiz Complete!</h2>
        <p>Score: {score} / 15</p>
        {/* Add Restart button */}
        <button onClick={restartQuiz}>Restart Quiz</button>
      </div>
    );

  return (
    <div style={{ padding: 30 }}>
      <h3>Question {questionIndex + 1}</h3>
      <p>{question.Sentence}</p>

      {question.Options.map((opt, i) => (
        <button key={i} onClick={() => choose(i)}>
          {opt}
        </button>
      ))}

      <p>{feedback}</p>
    </div>
  );
}
