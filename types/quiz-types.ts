/**
 * Shared TypeScript type definitions for ObservaQuiz
 * Generated from OpenAPI specification
 */

export interface HealthResponse {
  status: 'healthy';
  timestamp: string;
  version?: string;
}

export interface QuizQuestionsResponse {
  session_id: string;
  questions: Question[];
  current_question_index?: number;
}

export type QuestionType = 'multiple_choice' | 'text_input' | 'data_analysis';

export interface Question {
  id: string;
  type: QuestionType;
  text: string;
  options?: string[];
  data_context?: Record<string, any>;
}

export interface AnswerSubmission {
  session_id: string;
  question_id: string;
  answer: string | number | Record<string, any>;
  timestamp?: string;
}

export interface AnswerResponse {
  correct: boolean;
  explanation?: string;
  next_question?: Question;
  session_id: string;
  score?: number;
  quiz_complete?: boolean;
}

export interface ErrorResponse {
  error: string;
  message: string;
  details?: Record<string, any>;
}

// Client-side specific types
export interface QuizSession {
  id: string;
  questions: Question[];
  current_index: number;
  score: number;
  answers: AnswerSubmission[];
  completed: boolean;
  started_at: string;
}

export interface QuizState {
  session?: QuizSession;
  loading: boolean;
  error?: string;
  honeycomb_api_key?: string;
}

// Telemetry contracts
export interface TelemetryEvent {
  event_type: string;
  session_id: string;
  timestamp: string;
  properties: Record<string, any>;
}

export interface QuizStartedEvent extends TelemetryEvent {
  event_type: 'quiz_started';
  properties: {
    user_id?: string;
    honeycomb_team?: string;
  };
}

export interface QuestionAnsweredEvent extends TelemetryEvent {
  event_type: 'question_answered';
  properties: {
    question_id: string;
    question_type: QuestionType;
    answer: string | number | Record<string, any>;
    correct: boolean;
    time_taken_ms: number;
  };
}

export interface QuizCompletedEvent extends TelemetryEvent {
  event_type: 'quiz_completed';
  properties: {
    final_score: number;
    total_questions: number;
    completion_time_ms: number;
  };
}
