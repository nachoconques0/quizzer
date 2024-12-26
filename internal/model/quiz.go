package model

import (
	"fmt"
	"sort"

	"github.com/google/uuid"
)

type Question struct {
	ID                 uuid.UUID `json:"id"`
	Text               string    `json:"text"`
	Answers            []string  `json:"answer"`
	CorrectAnswerIndex int       `json:"-"`
}

type Quiz struct {
	Questions   []Question   `json:"questions"`
	Leaderboard *Leaderboard `json:"leaderboard"`
}

func NewQuiz() *Quiz {
	return &Quiz{
		Leaderboard: NewLeaderboard(),
		Questions:   []Question{},
	}
}

type User struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

type Leaderboard struct {
	Users  []User
	Result string
}

func NewLeaderboard() *Leaderboard {
	return &Leaderboard{
		Users: []User{
			{
				Name:   "pepe",
				Points: 1,
			},
			{
				Name:   "doge",
				Points: 3,
			},
			{
				Name:   "pumpndump",
				Points: 1,
			},
			{
				Name:   "memecoins<3",
				Points: 4,
			},
			{
				Name:   "musk",
				Points: 2,
			},
		},
	}
}

func (q *Quiz) AddQuestion(question Question) {
	q.Questions = append(q.Questions, question)
}

func (q *Quiz) HandleAnswers(req SubmitRequest) *Leaderboard {
	count := 0
	for _, answer := range req.Answers {
		for _, question := range q.Questions {
			if answer.QuestionID == question.ID {
				if answer.Value == question.CorrectAnswerIndex {
					count++
				}
			}
		}
	}
	q.updateLeaderBoard(count)
	return q.Leaderboard
}

func (q *Quiz) updateLeaderBoard(correctAnswers int) {
	q.Leaderboard = NewLeaderboard()
	q.Leaderboard.Users = append(q.Leaderboard.Users, User{
		Name:   "theonethatwillreadthecode",
		Points: correctAnswers,
	})
	sort.Sort(q.Leaderboard)
	q.Leaderboard.Result = calculateComparison(correctAnswers, *q.Leaderboard)
}

func calculateComparison(userPoints int, leaderboard Leaderboard) string {
	totalUsers := len(leaderboard.Users)
	betterThan := 0
	for _, user := range leaderboard.Users {
		if userPoints > user.Points {
			betterThan++
		}
	}

	percentage := float64(betterThan) / float64(totalUsers) * 100
	return fmt.Sprintf("You were better than %.2f%% of all quizzers", percentage)
}

func (lb Leaderboard) Len() int           { return len(lb.Users) }
func (lb Leaderboard) Less(i, j int) bool { return lb.Users[i].Points > lb.Users[j].Points }
func (lb Leaderboard) Swap(i, j int)      { lb.Users[i], lb.Users[j] = lb.Users[j], lb.Users[i] }
