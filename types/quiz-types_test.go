package types

import (
	"encoding/json"
	"testing"
	"time"
)

// TestQuestionTypeValidation tests the QuestionType validation
func TestQuestionTypeValidation(t *testing.T) {
	tests := []struct {
		name     string
		qType    QuestionType
		expected bool
	}{
		{"valid multiple choice", QuestionTypeMultipleChoice, true},
		{"valid text input", QuestionTypeTextInput, true},
		{"valid data analysis", QuestionTypeDataAnalysis, true},
		{"invalid empty", QuestionType(""), false},
		{"invalid random", QuestionType("random"), false},
		{"invalid case", QuestionType("Multiple_Choice"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.qType.IsValid(); got != tt.expected {
				t.Errorf("QuestionType.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestQuestionValidation tests the Question validation
func TestQuestionValidation(t *testing.T) {
	tests := []struct {
		name        string
		question    Question
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid text input question",
			question: Question{
				ID:   "test-id",
				Type: QuestionTypeTextInput,
				Text: "What is observability?",
			},
			expectError: false,
		},
		{
			name: "valid multiple choice question",
			question: Question{
				ID:      "test-id-2",
				Type:    QuestionTypeMultipleChoice,
				Text:    "Which tool is best for observability?",
				Options: []string{"Honeycomb", "DataDog", "New Relic"},
			},
			expectError: false,
		},
		{
			name: "invalid empty ID",
			question: Question{
				ID:   "",
				Type: QuestionTypeTextInput,
				Text: "What is observability?",
			},
			expectError: true,
			errorMsg:    "question ID cannot be empty",
		},
		{
			name: "invalid empty text",
			question: Question{
				ID:   "test-id",
				Type: QuestionTypeTextInput,
				Text: "",
			},
			expectError: true,
			errorMsg:    "question text cannot be empty",
		},
		{
			name: "invalid question type",
			question: Question{
				ID:   "test-id",
				Type: QuestionType("invalid"),
				Text: "What is observability?",
			},
			expectError: true,
			errorMsg:    "invalid question type: invalid",
		},
		{
			name: "multiple choice without options",
			question: Question{
				ID:   "test-id",
				Type: QuestionTypeMultipleChoice,
				Text: "Which tool is best?",
			},
			expectError: true,
			errorMsg:    "multiple choice questions must have options",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.question.Validate()
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

// TestQuestionJSONSerialization tests JSON marshaling/unmarshaling
func TestQuestionJSONSerialization(t *testing.T) {
	original := Question{
		ID:   "test-question-id",
		Type: QuestionTypeMultipleChoice,
		Text: "What is the best observability tool?",
		Options: []string{
			"Honeycomb",
			"DataDog", 
			"New Relic",
		},
		DataContext: map[string]interface{}{
			"difficulty": "beginner",
			"category":   "tools",
		},
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal question: %v", err)
	}

	// Unmarshal back
	var unmarshaled Question
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal question: %v", err)
	}

	// Compare
	if unmarshaled.ID != original.ID {
		t.Errorf("ID mismatch: got %s, want %s", unmarshaled.ID, original.ID)
	}
	if unmarshaled.Type != original.Type {
		t.Errorf("Type mismatch: got %s, want %s", unmarshaled.Type, original.Type)
	}
	if unmarshaled.Text != original.Text {
		t.Errorf("Text mismatch: got %s, want %s", unmarshaled.Text, original.Text)
	}
	if len(unmarshaled.Options) != len(original.Options) {
		t.Errorf("Options length mismatch: got %d, want %d", len(unmarshaled.Options), len(original.Options))
	}
}

// TestAnswerSubmissionValidation tests AnswerSubmission
func TestAnswerSubmissionValidation(t *testing.T) {
	now := time.Now()
	submission := AnswerSubmission{
		SessionID:  "session-123",
		QuestionID: "question-456",
		Answer:     "Honeycomb is the best!",
		Timestamp:  &now,
	}

	// Test JSON serialization
	jsonData, err := json.Marshal(submission)
	if err != nil {
		t.Fatalf("Failed to marshal answer submission: %v", err)
	}

	var unmarshaled AnswerSubmission
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal answer submission: %v", err)
	}

	if unmarshaled.SessionID != submission.SessionID {
		t.Errorf("SessionID mismatch: got %s, want %s", unmarshaled.SessionID, submission.SessionID)
	}
	if unmarshaled.QuestionID != submission.QuestionID {
		t.Errorf("QuestionID mismatch: got %s, want %s", unmarshaled.QuestionID, submission.QuestionID)
	}
}

// TestAnswerResponseValidation tests AnswerResponse
func TestAnswerResponseValidation(t *testing.T) {
	nextQuestion := Question{
		ID:   "next-question",
		Type: QuestionTypeTextInput,
		Text: "Follow-up question",
	}

	response := AnswerResponse{
		Correct:      true,
		Explanation:  stringPtr("Great answer! Honeycomb provides excellent observability."),
		NextQuestion: &nextQuestion,
		SessionID:    "session-123",
		Score:        intPtr(85),
		QuizComplete: boolPtr(false),
	}

	// Test JSON serialization
	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal answer response: %v", err)
	}

	var unmarshaled AnswerResponse
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal answer response: %v", err)
	}

	if unmarshaled.Correct != response.Correct {
		t.Errorf("Correct mismatch: got %v, want %v", unmarshaled.Correct, response.Correct)
	}
	if unmarshaled.SessionID != response.SessionID {
		t.Errorf("SessionID mismatch: got %s, want %s", unmarshaled.SessionID, response.SessionID)
	}
	if *unmarshaled.Score != *response.Score {
		t.Errorf("Score mismatch: got %d, want %d", *unmarshaled.Score, *response.Score)
	}
}

// Helper functions for pointer creation
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}

// TestQuizSessionValidation tests QuizSession
func TestQuizSessionValidation(t *testing.T) {
	questions := []Question{
		{
			ID:   "q1",
			Type: QuestionTypeTextInput,
			Text: "What is observability?",
		},
		{
			ID:   "q2", 
			Type: QuestionTypeMultipleChoice,
			Text: "Best tool?",
			Options: []string{"A", "B", "C"},
		},
	}

	session := QuizSession{
		ID:              "session-123",
		HoneycombAPIKey: "secret-key",
		Questions:       questions,
		CurrentIndex:    0,
		Score:           0,
		Answers:         []AnswerSubmission{},
		Completed:       false,
		StartedAt:       time.Now(),
		HoneycombTeam:   stringPtr("test-team"),
		UserID:          stringPtr("user-123"),
	}

	// Test that API key is not serialized
	jsonData, err := json.Marshal(session)
	if err != nil {
		t.Fatalf("Failed to marshal quiz session: %v", err)
	}

	jsonStr := string(jsonData)
	if containsSubstring(jsonStr, "secret-key") {
		t.Errorf("API key was serialized in JSON: %s", jsonStr)
	}

	// Test unmarshaling
	var unmarshaled QuizSession
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal quiz session: %v", err)
	}

	if unmarshaled.ID != session.ID {
		t.Errorf("ID mismatch: got %s, want %s", unmarshaled.ID, session.ID)
	}
	if len(unmarshaled.Questions) != len(session.Questions) {
		t.Errorf("Questions length mismatch: got %d, want %d", len(unmarshaled.Questions), len(session.Questions))
	}
}

// Helper function to check if string contains substring
func containsSubstring(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
		containsSubstringAt(s, substr, 1))))
}

func containsSubstringAt(s, substr string, start int) bool {
	if start >= len(s) {
		return false
	}
	if start+len(substr) > len(s) {
		return containsSubstringAt(s, substr, start+1)
	}
	if s[start:start+len(substr)] == substr {
		return true
	}
	return containsSubstringAt(s, substr, start+1)
}
