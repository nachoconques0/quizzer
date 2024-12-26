package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/nachoconques0/quizzer/internal/db"
	"github.com/nachoconques0/quizzer/internal/model"
)

type Service struct {
	db db.Data
}

func NewService(db db.Data) Service {
	return Service{db}
}

func (s Service) GetQuestion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s.db.GetQuestion()); err != nil {
		log.Printf("error encoding response: %s", err)
	}
}

func (s Service) SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	requestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading SubmitAnswers body: %s", err)
		return
	}

	var payload model.SubmitRequest
	err = json.Unmarshal(requestBytes, &payload)
	if err != nil {
	}

	quiz := s.db.GetQuiz()
	res := quiz.HandleAnswers(payload)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("error encoding response: %s", err)
	}
}
