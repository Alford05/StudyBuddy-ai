import React, { useState } from "react";
import Quiz from "./Quiz";
import { startQuiz, restartQuizApi } from "./api";

export default function App() {
  const [words, setWords] = useState("");
  const [wordList, setWordList] = useState([]);
  const [quizStarted, setQuizStarted] = useState(false);
  const [initialQuestion, setInitialQuestion] = useState(null);

  const start = async () => {
    const list = words.split("\n").map((w) => w.trim()).filter((w) => w);
    if (list.length !== 10) {
      alert("Please enter exactly 10 words.");
      return;
    }
    setWordList(list);

    try {
      const res = await startQuiz(list);
      setInitialQuestion(res.data.question);
      setQuizStarted(true);
    } catch (err) {
        console.error("startQuiz error:", err.response || err);
        alert(
            "Error starting quiz: " + 
            (err.response?.data?.error || err.message || "check console")
        );
    }
  };

  const restartQuizWithSameWords = async () => {
    await restartQuizApi(null);
    setQuizStarted(false);
    setInitialQuestion(null);
  };

  const restartQuizWithNewWords = async () => {
    await restartQuizApi(null);
    setQuizStarted(false);
    setInitialQuestion(null);
    setWordList([]);
    setWords("");
  };

  if (!quizStarted)
    return (
      <div style={{ padding: 30 }}>
        <h2>Enter 10 vocabulary words</h2>
        <textarea
          rows={10}
          value={words}
          onChange={(e) => setWords(e.target.value)}
        />
        <br />
        <button onClick={start}>Start Quiz</button>
      </div>
    );

  return (
    <Quiz
      words={wordList} 
      firstQuestion={initialQuestion}
      restartQuiz={() => {
        // time to choose
        const restartOption = window.confirm(
            "Do you want to restart with the same words again?"
        );
        if (restartOption) {
            restartQuizWithSameWords();
        } else {
            restartQuizWithNewWords();
        }
     }} 
    />
  );
}
