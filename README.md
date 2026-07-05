# StudyBuddy-ai

An AI-powered vocabulary learning application built with **Go**, **React**, and **Google Gemini**.

StudyBuddy helps students reinforce vocabulary through dynamically generated multiple-choice quizzes. Users provide a custom list of vocabulary words, and the application generates contextual sentences using Gemini AI, creating an interactive fill-in-the-blank learning experience.

---

## Features

* Generate vocabulary quizzes from a custom word bank
* AI-generated contextual sentences using Gemini
* Multiple-choice answer validation
* Session-based quiz tracking
* Single-word and two-word challenge questions
* React frontend with responsive UI
* Go backend REST API using Gin
* Dockerized deployment with Nginx reverse proxy

---

## Tech Stack

### Backend

* Go
* Gin Web Framework
* Cookie-based Session Management
* Google Gemini API

### Frontend

* React
* Vite
* Axios

### Infrastructure

* Docker
* Docker Compose
* Nginx

---

## Architecture

```mermaid
flowchart LR

    A[User] --> B[React Frontend]

    B -->|POST /api/start| C[Go API]
    B -->|GET Questions| C
    B -->|POST Answers| C

    C --> D[Session Store]

    C --> E[Gemini API]

    E -->|Generated Sentences| C

    C -->|Quiz Questions| B
```

---

## Quiz Flow

```mermaid
sequenceDiagram

    participant User
    participant Frontend
    participant Backend
    participant Gemini

    User->>Frontend: Enter 10 vocabulary words
    Frontend->>Backend: POST /api/start

    Backend->>Gemini: Generate sentence
    Gemini-->>Backend: Context sentence

    Backend-->>Frontend: Question + options

    User->>Frontend: Select answer
    Frontend->>Backend: POST /api/check

    Backend-->>Frontend: Correct / Incorrect

    loop Remaining Questions
        Frontend->>Backend: GET next question
        Backend->>Gemini: Generate sentence
        Gemini-->>Backend: Sentence
        Backend-->>Frontend: New question
    end
```

---

## Project Structure

```text
StudyBuddyv2/
│
├── backend/
│   ├── main.go
│   ├── handlers.go
│   ├── models.go
│   ├── question.go
│   ├── quiz.go
│   └── sentence_generator.go
│
├── frontend/
│   ├── src/
│   ├── package.json
│   └── vite.config.js
│
├── Dockerfile
├── docker-compose.yml
├── nginx.conf
└── README.md
```

---

## API Endpoints

### Start Quiz

```http
POST /api/start
```

Request:

```json
{
  "words": [
    "analyze",
    "concept",
    "infer",
    "justify",
    "context",
    "evaluate",
    "synthesize",
    "compare",
    "contrast",
    "evidence"
  ]
}
```

---

### Get Question

```http
GET /api/question/{index}
```

Returns a generated fill-in-the-blank question and answer choices.

---

### Check Answer

```http
POST /api/check
```

Request:

```json
{
  "questionIndex": 0,
  "selectedIndex": 2
}
```

---

### Restart Quiz

```http
POST /api/restart
```

Resets quiz progress while maintaining session state.

---

## Running Locally

### Prerequisites

* Go 1.23+
* Node.js
* Docker
* Google Gemini API Key

### Environment Variable

```bash
export SECRET_KEY=YOUR_GEMINI_API_KEY
```

### Run Backend

```bash
go run ./backend
```

### Run Frontend

```bash
cd frontend

npm install

npm run dev
```

---

## Docker Deployment

Build and start all services:

```bash
docker compose up --build
```

Application:

```text
http://localhost
```

---

## Learning Objectives

This project was built to practice:

* REST API design in Go
* Session management
* Integrating external AI services
* Frontend/backend communication
* Dockerized application deployment
* Dynamic content generation
* Application architecture and state management

---

## Future Improvements

* User authentication
* Persistent database storage
* Quiz history and analytics
* Difficulty levels
* Spaced repetition learning
* Teacher dashboards
* Progress tracking
* Cached AI responses to reduce API calls

---

## Author

Daniel Alford

Backend-focused developer building projects with Go, APIs, PostgreSQL, Docker, and cloud-native technologies.

