package model

import "github.com/google/uuid"

// SubmitRequest defines needed field to submit answers
type SubmitRequest struct {
	Answers []Answer `json:"answers,omitempty"`
}

// Answer defines the asset of answers in our service
type Answer struct {
	// ID Unique Identifier of the question
	QuestionID uuid.UUID `json:"question_id,omitempty"`
	// Answer
	Value int `json:"answer,omitempty"`
}
