# ObservaQuiz Contracts

This repository contains the shared contracts, interfaces, and coordination documentation for the ObservaQuiz project.

## Repository Structure

```
observaquiz-contracts/     # This repo - coordination & contracts
├── plan.md               # Master development plan
├── api-contracts/        # OpenAPI specifications
├── types/               # Shared type definitions (Go & TypeScript)
├── integration-tests/   # End-to-end contract validation
├── docs/               # Architecture & coordination docs
├── scripts/            # Cross-repo development tools
├── go.mod              # Go module definition
└── Makefile           # Build and test automation
```

## Related Repositories

- **Frontend**: [observaquiz-ui](https://github.com/honeycombio/observaquiz-ui) - React/TypeScript implementation
- **Backend**: [observaquiz-api](https://github.com/honeycombio/observaquiz-api) - Go AWS Lambda implementation

## Dependency Inversion Principle

This repository implements dependency inversion for the ObservaQuiz project:

- **Frontend** depends on contracts defined here (not on backend implementation)
- **Backend** implements contracts defined here (not coupled to frontend)
- **Integration tests** validate that implementations satisfy contracts

## Go Module Usage

### Installing as a Dependency

```bash
go get github.com/KentBeck/observaquiz-contracts
```

### Using in Go Code

```go
package main

import (
    "fmt"
    "github.com/KentBeck/observaquiz-contracts/types"
)

func main() {
    question := types.Question{
        ID:   "q1",
        Type: types.QuestionTypeTextInput,
        Text: "What is observability?",
    }
    
    if err := question.Validate(); err != nil {
        fmt.Printf("Invalid question: %v\n", err)
        return
    }
    
    fmt.Printf("Valid question: %s\n", question.Text)
}
```

### Available Types

- `types.Question` - Quiz question with validation
- `types.QuestionType` - Enum for question types
- `types.AnswerSubmission` - Answer submission data
- `types.AnswerResponse` - Response to answer submission
- `types.QuizQuestionsResponse` - Response when fetching questions
- `types.ErrorResponse` - Standard error response format
- `types.QuizSession` - Server-side session management
- `types.TelemetryEvent` - Base telemetry event type

## Development Workflow

1. **Contract-First Development**: Define APIs and types here first
2. **Independent Implementation**: Frontend and backend implement against stable contracts
3. **Contract Testing**: Validate implementations satisfy contracts
4. **Coordination**: Use `plan.md` to coordinate work across repositories

## Testing

### Run All Tests

```bash
make test
```

### Run Go Tests Only

```bash
make test-go
# or
go test -v ./types/...
```

### Run JavaScript Tests Only

```bash
make test-js
# or
npm test
```

### Validate OpenAPI Specification

```bash
make validate
# or
npm run validate:openapi
```

## Contract Validation

The contracts include comprehensive validation:

- **Type Safety**: All types have validation methods
- **JSON Compatibility**: Types serialize/deserialize correctly
- **OpenAPI Compliance**: Go types match OpenAPI specification exactly
- **Field Naming**: Consistent snake_case JSON field names

### Example Validation

```go
question := types.Question{
    ID:   "test-question",
    Type: types.QuestionTypeMultipleChoice,
    Text: "Which tool is best for observability?",
    Options: []string{"Honeycomb", "DataDog", "New Relic"},
}

// Validate the question
if err := question.Validate(); err != nil {
    log.Fatalf("Question validation failed: %v", err)
}

// Check question type
if !question.Type.IsValid() {
    log.Fatalf("Invalid question type: %s", question.Type)
}
```

## Getting Started

1. Clone this repository
2. Review `plan.md` for current development priorities
3. Check `api-contracts/` for latest API specifications
4. Run tests to validate current state: `make test`

## Contract-Driven TDD

1. **Red**: Write failing contract test
2. **Green**: Implement in frontend/backend to satisfy contract
3. **Refactor**: Improve implementation while maintaining contract compliance

This ensures both sides develop against stable, well-defined interfaces.

## CI/CD

The repository includes GitHub Actions workflows that:

- Validate OpenAPI specifications
- Run Go tests on multiple Go versions
- Run JavaScript/Jest tests
- Check code formatting and linting
- Run security scans
- Generate types from OpenAPI specs

## Contributing

1. All API changes must start with contract updates
2. Ensure all tests pass before submitting PRs
3. Follow the established JSON field naming conventions
4. Add tests for any new types or validation rules
5. Update documentation for any breaking changes

## Versioning

This repository follows semantic versioning:

- **Major**: Breaking changes to contracts
- **Minor**: New features, backward-compatible
- **Patch**: Bug fixes, no API changes

## License

ISC License - see LICENSE file for details.
