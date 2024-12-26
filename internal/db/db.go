package db

import (
	"github.com/google/uuid"
	"github.com/nachoconques0/quizzer/internal/model"
)

type Data struct {
	Quiz *model.Quiz
}

func NewDB() Data {
	quiz := model.NewQuiz()
	return Data{
		Quiz: quiz,
	}
}

func (db Data) SeedQuiz() {
	db.Quiz.AddQuestion(
		model.Question{
			ID:   uuid.New(),
			Text: `What is the name of Link's horse in "The Legend of Zelda: Ocarina of Time"?`,
			Answers: []string{
				"Shadow",
				"Pepe",
				"Epona",
			},
			CorrectAnswerIndex: 2,
		},
	)
	db.Quiz.AddQuestion(
		model.Question{
			ID:   uuid.New(),
			Text: `Which instrument does Link play in "The Legend of Zelda: Majora's Mask"?`,
			Answers: []string{
				"Ocarina",
				"Harp",
				"Slash Electric Guitar",
			},
			CorrectAnswerIndex: 0,
		},
	)
	db.Quiz.AddQuestion(
		model.Question{
			ID:   uuid.New(),
			Text: `What is the name of the kingdom where most of the Zelda games take place?`,
			Answers: []string{
				"Hyrule",
				"Girona",
				"Kakariko Village",
			},
			CorrectAnswerIndex: 0,
		},
	)
	db.Quiz.AddQuestion(
		model.Question{
			ID:   uuid.New(),
			Text: `Who is the main antagonist in most of the Zelda games?`,
			Answers: []string{
				"Pedro Sanchez",
				"Ganondorf",
				"Bowser",
			},
			CorrectAnswerIndex: 1,
		},
	)
	db.Quiz.AddQuestion(
		model.Question{
			ID:   uuid.New(),
			Text: `What is the name of the legendary sword wielded by Link?`,
			Answers: []string{
				"Phantom Blade",
				"Skyfang",
				"Master Sword",
			},
			CorrectAnswerIndex: 2,
		},
	)
}

func (db Data) GetQuestion() []model.Question {
	return db.Quiz.Questions
}

func (db Data) GetQuiz() *model.Quiz {
	return db.Quiz
}

func (db Data) ResetLeaderboard() {
	db.Quiz.Leaderboard = model.NewLeaderboard()
}
