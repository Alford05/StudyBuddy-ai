import axios from "axios";

export const startQuiz = (words) =>
    axios.post("/api/start", { words });

export const restartQuizApi = (words) =>
    axios.post("/api/restart", words ? { words } : {});

export const getNextQuestion = (index, words) =>
    axios.get(`/api/question/${index}`, {
        params: { words },
    });

export const checkAnswer = (questionIndex, selectedIndex) =>
    axios.post("/api/check", {
        questionIndex,
        selectedIndex,
    });


