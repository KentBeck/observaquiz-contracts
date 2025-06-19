package types

import (
	"encoding/json"
	"testing"
)

// TestOpenAPICompliance tests that Go types match the OpenAPI specification
func TestOpenAPICompliance(t *testing.T) {
	t.Run("Question matches OpenAPI schema", func(t *testing.T) {
		// Test that Question can be serialized to match OpenAPI spec
		question := Question{
			ID:   "550e8400-e29b-41d4-a716-446655440000", // UUID format
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

		jsonData, err := json.Marshal(question)
		if err != nil {
			t.Fatalf("Failed to marshal question: %v", err)
		}

		// Verify JSON structure matches OpenAPI spec
		var jsonMap map[string]interface{}
		err = json.Unmarshal(jsonData, &jsonMap)
		if err != nil {
			t.Fatalf("Failed to unmarshal to map: %v", err)
		}

		// Check required fields exist
		requiredFields := []string{"id", "type", "text"}
		for _, field := range requiredFields {
			if _, exists := jsonMap[field]; !exists {
				t.Errorf("Required field '%s' missing from JSON", field)
			}
		}

		// Check type is valid enum value
		typeValue, ok := jsonMap["type"].(string)
		if !ok {
			t.Errorf("Type field is not a string")
		}
		validTypes := []string{"multiple_choice", "text_input", "data_analysis"}
		if !contains(validTypes, typeValue) {
			t.Errorf("Type '%s' is not a valid enum value", typeValue)
		}

		// Check options array exists for multiple choice
		if typeValue == "multiple_choice" {
			if _, exists := jsonMap["options"]; !exists {
				t.Errorf("Options field missing for multiple_choice question")
			}
		}
	})

	t.Run("QuizQuestionsResponse matches OpenAPI schema", func(t *testing.T) {
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

		response := QuizQuestionsResponse{
			SessionID:            "session-123",
			Questions:            questions,
			CurrentQuestionIndex: intPtr(0),
		}

		jsonData, err := json.Marshal(response)
		if err != nil {
			t.Fatalf("Failed to marshal quiz questions response: %v", err)
		}

		var jsonMap map[string]interface{}
		err = json.Unmarshal(jsonData, &jsonMap)
		if err != nil {
			t.Fatalf("Failed to unmarshal to map: %v", err)
		}

		// Check required fields
		requiredFields := []string{"session_id", "questions"}
		for _, field := range requiredFields {
			if _, exists := jsonMap[field]; !exists {
				t.Errorf("Required field '%s' missing from JSON", field)
			}
		}

		// Check questions is an array
		questionsField, ok := jsonMap["questions"].([]interface{})
		if !ok {
			t.Errorf("Questions field is not an array")
		}
		if len(questionsField) != 2 {
			t.Errorf("Expected 2 questions, got %d", len(questionsField))
		}
	})

	t.Run("AnswerSubmission matches OpenAPI schema", func(t *testing.T) {
		submission := AnswerSubmission{
			SessionID:  "session-123",
			QuestionID: "question-456",
			Answer:     "Honeycomb is the best!",
		}

		jsonData, err := json.Marshal(submission)
		if err != nil {
			t.Fatalf("Failed to marshal answer submission: %v", err)
		}

		var jsonMap map[string]interface{}
		err = json.Unmarshal(jsonData, &jsonMap)
		if err != nil {
			t.Fatalf("Failed to unmarshal to map: %v", err)
		}

		// Check required fields
		requiredFields := []string{"session_id", "question_id", "answer"}
		for _, field := range requiredFields {
			if _, exists := jsonMap[field]; !exists {
				t.Errorf("Required field '%s' missing from JSON", field)
			}
		}

		// Test different answer types (string, integer, object)
		testCases := []struct {
			name   string
			answer interface{}
		}{
			{"string answer", "text response"},
			{"integer answer", 42},
			{"object answer", map[string]interface{}{"selected": "option1", "confidence": 0.8}},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				submission.Answer = tc.answer
				jsonData, err := json.Marshal(submission)
				if err != nil {
					t.Fatalf("Failed to marshal submission with %s: %v", tc.name, err)
				}

				var unmarshaled AnswerSubmission
				err = json.Unmarshal(jsonData, &unmarshaled)
				if err != nil {
					t.Fatalf("Failed to unmarshal submission with %s: %v", tc.name, err)
				}
			})
		}
	})

	t.Run("AnswerResponse matches OpenAPI schema", func(t *testing.T) {
		nextQuestion := Question{
			ID:   "next-q",
			Type: QuestionTypeTextInput,
			Text: "Follow-up question",
		}

		response := AnswerResponse{
			Correct:      true,
			Explanation:  stringPtr("Excellent answer!"),
			NextQuestion: &nextQuestion,
			SessionID:    "session-123",
			Score:        intPtr(85),
			QuizComplete: boolPtr(false),
		}

		jsonData, err := json.Marshal(response)
		if err != nil {
			t.Fatalf("Failed to marshal answer response: %v", err)
		}

		var jsonMap map[string]interface{}
		err = json.Unmarshal(jsonData, &jsonMap)
		if err != nil {
			t.Fatalf("Failed to unmarshal to map: %v", err)
		}

		// Check required fields
		requiredFields := []string{"correct", "session_id"}
		for _, field := range requiredFields {
			if _, exists := jsonMap[field]; !exists {
				t.Errorf("Required field '%s' missing from JSON", field)
			}
		}

		// Check correct is boolean
		if _, ok := jsonMap["correct"].(bool); !ok {
			t.Errorf("Correct field is not a boolean")
		}

		// Check optional fields have correct types when present
		if explanation, exists := jsonMap["explanation"]; exists {
			if _, ok := explanation.(string); !ok {
				t.Errorf("Explanation field is not a string")
			}
		}

		if score, exists := jsonMap["score"]; exists {
			if _, ok := score.(float64); !ok { // JSON numbers are float64
				t.Errorf("Score field is not a number")
			}
		}

		if quizComplete, exists := jsonMap["quiz_complete"]; exists {
			if _, ok := quizComplete.(bool); !ok {
				t.Errorf("Quiz_complete field is not a boolean")
			}
		}
	})

	t.Run("ErrorResponse matches OpenAPI schema", func(t *testing.T) {
		errorResp := ErrorResponse{
			Error:   "validation_failed",
			Message: "The submitted answer is invalid",
			Details: map[string]interface{}{
				"field": "answer",
				"code":  "required",
			},
		}

		jsonData, err := json.Marshal(errorResp)
		if err != nil {
			t.Fatalf("Failed to marshal error response: %v", err)
		}

		var jsonMap map[string]interface{}
		err = json.Unmarshal(jsonData, &jsonMap)
		if err != nil {
			t.Fatalf("Failed to unmarshal to map: %v", err)
		}

		// Check required fields
		requiredFields := []string{"error", "message"}
		for _, field := range requiredFields {
			if _, exists := jsonMap[field]; !exists {
				t.Errorf("Required field '%s' missing from JSON", field)
			}
		}

		// Check field types
		if _, ok := jsonMap["error"].(string); !ok {
			t.Errorf("Error field is not a string")
		}
		if _, ok := jsonMap["message"].(string); !ok {
			t.Errorf("Message field is not a string")
		}
		if details, exists := jsonMap["details"]; exists {
			if _, ok := details.(map[string]interface{}); !ok {
				t.Errorf("Details field is not an object")
			}
		}
	})
}

// Helper function to check if slice contains string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// TestJSONFieldNaming tests that JSON field names match OpenAPI spec exactly
func TestJSONFieldNaming(t *testing.T) {
	question := Question{
		ID:   "test",
		Type: QuestionTypeTextInput,
		Text: "test question",
	}

	jsonData, err := json.Marshal(question)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	jsonStr := string(jsonData)

	// Check that field names use snake_case as per OpenAPI spec
	expectedFields := []string{
		`"id"`,
		`"type"`,
		`"text"`,
	}

	for _, field := range expectedFields {
		if !containsString(jsonStr, field) {
			t.Errorf("Expected field %s not found in JSON: %s", field, jsonStr)
		}
	}

	// Check that we don't have camelCase versions
	unexpectedFields := []string{
		`"Id"`,
		`"Type"`,
		`"Text"`,
	}

	for _, field := range unexpectedFields {
		if containsString(jsonStr, field) {
			t.Errorf("Unexpected camelCase field %s found in JSON: %s", field, jsonStr)
		}
	}
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || 
		(len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || 
		containsStringAt(s, substr, 1))))
}

func containsStringAt(s, substr string, start int) bool {
	if start >= len(s) {
		return false
	}
	if start+len(substr) > len(s) {
		return containsStringAt(s, substr, start+1)
	}
	if s[start:start+len(substr)] == substr {
		return true
	}
	return containsStringAt(s, substr, start+1)
}
