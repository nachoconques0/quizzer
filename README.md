# quizzer

### What I've done
A command-line application that allows users to participate in a fun and interactive quiz. The project demonstrates a simple quiz system using Go, with the following features:

- Users can answer multiple-choice questions through a CLI interface.
- Questions are fetched dynamically from an API.
- Once the quiz is completed, the user's answers are submitted to the API, which calculates the score.
- A leaderboard displays the user's performance compared to others.

### Dependencies Used
- Cobra & Promptui

### How to run it :scream_cat:
1. `git clone` this repo
2. Once you are inside the repo make sure to have the modules if some issues appears just run "go mod tidy"
3. In the terminal write "go run main.go" and the quiz should start.
   

### Directory

```
📦quizzer
 ┣ 📂cmd
 ┃ ┣ 📜root.go
 ┃ ┣ 📜startApi.go
 ┃ ┗ 📜startQuiz.go
 ┣ 📂internal
 ┃ ┣ 📂api
 ┃ ┃ ┣ 📜quiz.go
 ┃ ┃ ┗ 📜server.go
 ┃ ┣ 📂db
 ┃ ┃ ┗ 📜db.go
 ┃ ┣ 📂helpers
 ┃ ┃ ┣ 📜api.go
 ┃ ┃ ┣ 📜error.go
 ┃ ┃ ┗ 📜utils.go
 ┃ ┗ 📂model
 ┃ ┃ ┣ 📜item.go
 ┃ ┃ ┣ 📜quiz.go
 ┃ ┃ ┗ 📜submit_answer.go
 ┣ 📜.env
 ┣ 📜.gitignore
 ┣ 📜README.md
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┗ 📜main.go
 ```