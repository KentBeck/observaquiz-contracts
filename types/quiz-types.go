// Package types contains shared Go type definitions for ObservaQuiz
// Generated from OpenAPI specification
package types

import (
	"fmt"
	"time"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   *string   `json:"version,omitempty"`
}

// QuizQuestionsResponse represents the response when fetching quiz questions
type QuizQuestionsResponse struct {
	SessionID            string     `json:"session_id"`
	Questions            []Question `json:"questions"`
	CurrentQuestionIndex *int       `json:"current_question_index,omitempty"`
}

// QuestionType represents the type of quiz question
type QuestionType string

const (
	QuestionTypeMultipleChoice QuestionType = "multiple_choice"
	QuestionTypeTextInput      QuestionType = "text_input"
	QuestionTypeDataAnalysis   QuestionType = "data_analysis"
)

// Question represents a quiz question
type Question struct {
	ID          string                 `json:"id"`
	Type        QuestionType           `json:"type"`
	Text        string                 `json:"text"`
	Options     []string               `json:"options,omitempty"`
	DataContext map[string]interface{} `json:"data_context,omitempty"`
}

// AnswerSubmission represents a submitted answer
type AnswerSubmission struct {
	SessionID  string      `json:"session_id"`
	QuestionID string      `json:"question_id"`
	Answer     interface{} `json:"answer"`
	Timestamp  *time.Time  `json:"timestamp,omitempty"`
}

// AnswerResponse represents the response to an answer submission
type AnswerResponse struct {
	Correct      bool      `json:"correct"`
	Explanation  *string   `json:"explanation,omitempty"`
	NextQuestion *Question `json:"next_question,omitempty"`
	SessionID    string    `json:"session_id"`
	Score        *int      `json:"score,omitempty"`
	QuizComplete *bool     `json:"quiz_complete,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string                 `json:"error"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// Server-side specific types

// QuizSession represents a quiz session in the backend
type QuizSession struct {
	ID              string             `json:"id"`
	HoneycombAPIKey string             `json:"-"` // Never serialize API key
	Questions       []Question         `json:"questions"`
	CurrentIndex    int                `json:"current_index"`
	Score           int                `json:"score"`
	Answers         []AnswerSubmission `json:"answers"`
	Completed       bool               `json:"completed"`
	StartedAt       time.Time          `json:"started_at"`
	CompletedAt     *time.Time         `json:"completed_at,omitempty"`
	HoneycombTeam   *string            `json:"honeycomb_team,omitempty"`
	UserID          *string            `json:"user_id,omitempty"`
}

// TelemetryEvent represents a telemetry event
type TelemetryEvent struct {
	EventType  string                 `json:"event_type"`
	SessionID  string                 `json:"session_id"`
	Timestamp  time.Time              `json:"timestamp"`
	Properties map[string]interface{} `json:"properties"`
}

// QuizStartedEvent represents a quiz started event
type QuizStartedEvent struct {
	TelemetryEvent
	UserID        *string `json:"user_id,omitempty"`
	HoneycombTeam *string `json:"honeycomb_team,omitempty"`
}

// QuestionAnsweredEvent represents a question answered event
type QuestionAnsweredEvent struct {
	TelemetryEvent
	QuestionID   string      `json:"question_id"`
	QuestionType string      `json:"question_type"`
	Answer       interface{} `json:"answer"`
	Correct      bool        `json:"correct"`
	TimeTakenMs  int64       `json:"time_taken_ms"`
}

// QuizCompletedEvent represents a quiz completed event
type QuizCompletedEvent struct {
	TelemetryEvent
	FinalScore       int   `json:"final_score"`
	TotalQuestions   int   `json:"total_questions"`
	CompletionTimeMs int64 `json:"completion_time_ms"`
}

// Validation methods

// IsValid checks if a QuestionType is valid
func (qt QuestionType) IsValid() bool {
	switch qt {
	case QuestionTypeMultipleChoice, QuestionTypeTextInput, QuestionTypeDataAnalysis:
		return true
	default:
		return false
	}
}

// Validate checks if a Question is valid
func (q *Question) Validate() error {
	if q.ID == "" {
		return fmt.Errorf("question ID cannot be empty")
	}
	if q.Text == "" {
		return fmt.Errorf("question text cannot be empty")
	}
	if !q.Type.IsValid() {
		return fmt.Errorf("invalid question type: %s", q.Type)
	}
	if q.Type == QuestionTypeMultipleChoice && len(q.Options) == 0 {
		return fmt.Errorf("multiple choice questions must have options")
	}
	return nil
}
