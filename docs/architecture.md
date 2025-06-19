# ObservaQuiz Architecture

## Overview

ObservaQuiz follows a **contract-first, dependency inversion** architecture where all components depend on shared contracts rather than each other.

```
┌─────────────────────────────────────────────────────────────┐
│                 observaquiz-contracts                       │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────┐ │
│  │   API Specs     │  │  Type Defs      │  │ Integration │ │
│  │  (OpenAPI)      │  │ (TS/Go)         │  │   Tests     │ │
│  └─────────────────┘  └─────────────────┘  └─────────────┘ │
└─────────────────────────────────────────────────────────────┘
           ▲                        ▲
           │                        │
           │ depends on             │ depends on
           │                        │
┌─────────────────┐        ┌─────────────────┐
│ observaquiz-ui  │        │ observaquiz-api │
│   (Frontend)    │        │   (Backend)     │
│                 │        │                 │
│ React/TypeScript│        │ Go/AWS Lambda   │
│ OpenTelemetry   │        │ OpenTelemetry   │
└─────────────────┘        └─────────────────┘
```

## Dependency Inversion Principle

### Traditional Coupling (❌)
```
Frontend ──depends on──> Backend
```

### Dependency Inversion (✅)
```
Frontend ──depends on──> Contracts <──implements── Backend
```

## Repository Structure

### observaquiz-contracts (This Repository)
- **Purpose**: Define contracts and coordinate development
- **Contents**: API specs, type definitions, integration tests, documentation
- **Dependencies**: None (pure abstractions)

### observaquiz-ui
- **Purpose**: React frontend implementation
- **Dependencies**: Contracts only (not backend)
- **Testing**: Against contract mocks

### observaquiz-api  
- **Purpose**: Go backend implementation
- **Dependencies**: Contracts only (not frontend)
- **Testing**: Contract compliance tests

## Development Workflow

### 1. Contract-First Development
```
1. Define API contract in OpenAPI spec
2. Generate types for TypeScript and Go
3. Write contract tests
4. Implement frontend against contract mocks
5. Implement backend to satisfy contracts
6. Run integration tests to validate compliance
```

### 2. TDD with Contracts
```
Red:    Write failing contract test
Green:  Implement in frontend/backend to pass contract
Refactor: Improve implementation while maintaining contract
```

### 3. Independent Development
- Frontend team works against stable contracts
- Backend team implements contracts independently
- Integration happens through contract validation
- No cross-team blocking on implementation details

## Communication Patterns

### API Communication
- REST API following OpenAPI specification
- JSON request/response payloads
- Standard HTTP status codes
- Consistent error response format

### Telemetry Integration
- OpenTelemetry for distributed tracing
- Structured logging with consistent fields
- Custom metrics for quiz-specific events
- Baggage for session context propagation

## Testing Strategy

### Contract Tests (Integration)
- Validate API endpoints match specification
- Test request/response schema compliance
- Verify error handling contracts
- End-to-end telemetry validation

### Unit Tests (Implementation)
- Frontend: Component tests with mocked contracts
- Backend: Handler tests with contract validation
- Both: Business logic tests independent of contracts

### Property-Based Testing
- Generate test data from OpenAPI schemas
- Validate serialization/deserialization
- Test edge cases and boundary conditions

## Deployment Architecture

### Frontend Deployment
- Static site generation
- CDN distribution
- Environment-specific configuration
- OpenTelemetry browser instrumentation

### Backend Deployment
- AWS Lambda functions
- API Gateway integration
- Environment-based configuration
- OpenTelemetry Lambda instrumentation

### Observability Stack
- Honeycomb for telemetry data
- Distributed tracing across services
- Custom dashboards for quiz metrics
- Alerting on error rates and performance

## Benefits of This Architecture

1. **Independent Development**: Teams can work in parallel
2. **Testability**: Each component can be tested in isolation
3. **Maintainability**: Changes to implementation don't break contracts
4. **Scalability**: Services can be scaled independently
5. **Reliability**: Contract tests catch integration issues early
6. **Documentation**: Contracts serve as living documentation
